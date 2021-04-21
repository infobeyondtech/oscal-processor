package ssp

import (
	"encoding/xml"
	"io/ioutil"
	"fmt"

	"github.com/google/uuid"

	sdk_profile "github.com/docker/oscalkit/types/oscal/profile"
	sdk_ssp "github.com/docker/oscalkit/types/oscal/system_security_plan"
	request_models "github.com/infobeyondtech/oscal-processor/models/requests"
	information "github.com/infobeyondtech/oscal-processor/models/information"
	"github.com/docker/oscalkit/types/oscal/validation_root"
	"github.com/infobeyondtech/oscal-processor/context"
)

// Create an empty SSP
func CreateFreshSSP() (string, error) {

	fid := uuid.New().String()
	ssp := &sdk_ssp.SystemSecurityPlan{}

	// set id
	ssp.Id = fid

	out, err1 := xml.MarshalIndent(ssp, "  ", "    ")
	check(err1)

	err := ioutil.WriteFile(fid, out, 0644)
	check(err)

	return fid, nil
}

func SetTitleVersion(ssp *sdk_ssp.SystemSecurityPlan, request request_models.SetTitleVersionRequest){
	GuardMetaData(ssp)

	ssp.Metadata.Title = sdk_ssp.Title(request.Title)
	ssp.Metadata.Version = Version(request.Version)
	ssp.Metadata.OscalVersion = OscalVersion(request.OscalVersion)
}

func SetSystemCharacteristic(ssp *sdk_ssp.SystemSecurityPlan, request request_models.AddSystemCharacteristicReuqest ){
	GuardSystemCharacteristics(ssp)

	ssp.SystemCharacteristics.SystemName = sdk_ssp.SystemName(request.SystemName)
	ssp.SystemCharacteristics.Description = &sdk_ssp.Markup{Raw:request.Description}
	ssp.SystemCharacteristics.SecuritySensitivityLevel = sdk_ssp.SecuritySensitivityLevel(request.SecurityLevel)

	annotation := &Annotation {Name: "deployment-model", Value:request.DeploymentModel }
	ssp.SystemCharacteristics.Annotations = append(ssp.SystemCharacteristics.Annotations, *annotation)

	systemId := &sdk_ssp.SystemId{ IdentifierType: "https://ietf.org/rfc/rfc4122", Value:request.UUID}
	ssp.SystemCharacteristics.SystemIds = append(ssp.SystemCharacteristics.SystemIds, *systemId)
}

// initiate a ssp instance from an existing xml file
func LoadFromFile(ssp *sdk_ssp.SystemSecurityPlan, path string){
	dat, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("error: %v", e)
		return
	}

	// unmarshal into data structure
	marshalError := xml.Unmarshal([]byte(dat), &ssp)
	if marshalError != nil {
		fmt.Printf("error: %v", marshalError)
		return
	}
}

func WriteToFile(ssp *sdk_ssp.SystemSecurityPlan) string{
	parent := context.DownloadDir
	if(ssp.Id == ""){
		ssp.Id = uuid.New().String()
	}	
	targetFile := parent + "/" + ssp.Id
	targetFile = context.ExpandPath(targetFile)
	xmlFile := targetFile + ".xml"

	out, e := xml.MarshalIndent(ssp, "  ", "    ")
	check(e)
	
	ioErr := ioutil.WriteFile(xmlFile, out, 0644)
	check(ioErr)

	return xmlFile
}

// insert an inventory item
func AddInventoryItem(ssp *sdk_ssp.SystemSecurityPlan, request request_models.InsertInventoryItemRequest){
	sdk_itm := &sdk_ssp.InventoryItem{}	

	// fetch inventory information for item
	db_itm := information.GetInventoryItem(request.InventoryItemID)
	sdk_itm.AssetId = db_itm.AssetId
	sdk_itm.Description = &sdk_ssp.Markup{Raw:db_itm.Description}
	sdk_itm.Id = db_itm.UUID

	// fetch parties for item
	for _, partyRoleMap := range request.ResponsibleParties {

		sdk_party := &sdk_ssp.ResponsibleParty{}
		
		// insert user role detail in the header
		db_user := information.GetUser(partyRoleMap.UserUUID)
		AddUser(ssp, db_user)
		sdk_party.RoleId = db_user.RoleId
		
		for _, partyId := range partyRoleMap.PartyUUIDs{
			sdk_party.PartyIds = append(sdk_party.PartyIds, PartyId(partyId))

			// insert party detail in the header
			AddParty(ssp, partyId)
		}

		sdk_itm.ResponsibleParties = append(sdk_itm.ResponsibleParties, *sdk_party)
	}
	
	// fetch component for item
	for _, component_id := range request.ImplementComponents{
		implement_component := &sdk_ssp.ImplementedComponent{ ComponentId : component_id }
		sdk_itm.ImplementedComponents = append(sdk_itm.ImplementedComponents, *implement_component)
		
		// insert component detail in the header
		AddComponent(ssp, component_id, []sdk_ssp.ResponsibleRole{})	
	}

	// insert inventory item into ssp systemimplementation section
	GuardSystemImplementation(ssp)
	ssp.SystemImplementation.SystemInventory.InventoryItems = append(ssp.SystemImplementation.SystemInventory.InventoryItems, *sdk_itm)
}

// insert an implemented requirement
func AddImplementedRequirement(ssp *sdk_ssp.SystemSecurityPlan, requirement request_models.InsertImplementedRequirementRequest){
	sdk_requirement := &sdk_ssp.ImplementedRequirement{}
	sdk_requirement.ControlId = requirement.ControlID
	sdk_requirement.Id =  requirement.UUID

	for _, statement := range requirement.Statements{			
		// from request statement to sdk statement
		sdk_statement := &sdk_ssp.Statement{}
		sdk_statement.StatementId = statement.StatementID

		for _, byComponent := range statement.ByComponents{
			// from byComponent to sdk byComponent
			sdk_byComponent := &sdk_ssp.ByComponent{}
			sdk_byComponent.ComponentId = byComponent.ComponentID
			sdk_byComponent.Description = &sdk_ssp.Markup{Raw:byComponent.Description}
			responsibleRoles := []sdk_ssp.ResponsibleRole{}
			
			// component parameters
			for _, param := range byComponent.SetParameters{
				sdk_param := &sdk_ssp.SetParameter{}
				// from setParams to sdk params
				sdk_param.ParamId = param.ParamID
				sdk_param.Value = sdk_ssp.Value(param.Value)

				sdk_byComponent.ParameterSettings = append(sdk_byComponent.ParameterSettings, *sdk_param)
			}

			// responsible roles and users for a component
			for _, partyRoleMap := range byComponent.ResponsibleParties {

				sdk_role:= &sdk_ssp.ResponsibleRole{}
				
				// insert user role detail in the header
				db_user := information.GetUser(partyRoleMap.UserUUID)
				AddUser(ssp, db_user)
				sdk_role.RoleId = db_user.RoleId
				
				for _, partyId := range partyRoleMap.PartyUUIDs{
					sdk_role.PartyIds = append(sdk_role.PartyIds, PartyId(partyId))
		
					// insert party detail in the header
					AddParty(ssp, partyId)
				}
		
				//sdk_byComponent.ResponsibleRoles = append(sdk_byComponent.ResponsibleRoles, *sdk_role)
				responsibleRoles = append(responsibleRoles, *sdk_role)
			}			

			AddComponent(ssp, byComponent.ComponentID, responsibleRoles)
			sdk_statement.ByComponents = append(sdk_statement.ByComponents, *sdk_byComponent)
		}
		sdk_requirement.Statements = append(sdk_requirement.Statements, *sdk_statement)
	}

	GuardControlImplementation(ssp)
	ssp.ControlImplementation.ImplementedRequirements = append(ssp.ControlImplementation.ImplementedRequirements, *sdk_requirement)
}

// private func to add a component in system-implementation, check duplicates
func AddComponent(ssp *sdk_ssp.SystemSecurityPlan, componentId string, responsibleRoles []sdk_ssp.ResponsibleRole){
	db_component := information.GetComponent(componentId)
	sdk_component := &sdk_ssp.Component{}
	sdk_component.Id = db_component.UUID
	sdk_component.Description =  &sdk_ssp.Markup{Raw:db_component.Description}
	sdk_component.Status = &sdk_ssp.Status{State:db_component.State}
	sdk_component.Title = Title(db_component.Title)
	sdk_component.ComponentType = db_component.Type

	// version and last-modified property
	versionProperty := &Prop{Name:"version", Value:db_component.Version}
	lastModifiedProperty := &Prop{Name:"last-modified-date", Value:db_component.LastModified}
	sdk_component.Properties = append(sdk_component.Properties, *versionProperty)
	sdk_component.Properties = append(sdk_component.Properties, *lastModifiedProperty )

	// insert the responsible role
	copy(sdk_component.ResponsibleRoles, responsibleRoles)

	// insert into ssp component collection
	GuardSystemImplementation(ssp)
	ssp.SystemImplementation.Components = append(ssp.SystemImplementation.Components, *sdk_component)
}

// private func to add a user in system-implementation, check duplicates
func AddUser(ssp *sdk_ssp.SystemSecurityPlan, db_user information.User){
	
	sdk_user :=  &sdk_ssp.User{}
	sdk_user.Title = Title(db_user.Title)
	sdk_user.Id = db_user.UUID
	sdk_user.RoleIds = []sdk_ssp.RoleId{sdk_ssp.RoleId(db_user.RoleId)}

	annotation := &Annotation {Name: "type", Value:db_user.Type }
	sdk_user.Annotations = []sdk_ssp.Annotation{ *annotation }

	// insert into ssp header
	GuardSystemImplementation(ssp)
	ssp.SystemImplementation.Users = append(ssp.SystemImplementation.Users, *sdk_user)
}

// private func to add a party in meta data, check duplicates
func AddParty(ssp *sdk_ssp.SystemSecurityPlan, partyId string){
	
	db_party := information.GetParty(partyId);
	sdk_party := &Party{}
	sdk_party.Id = db_party.UUID

	// todo: party name and type are missing in this sdk

	// insert into ssp header
	GuardMetaData(ssp)
	ssp.Metadata.Parties = append(ssp.Metadata.Parties, *sdk_party)
} 

func GuardMetaData(ssp *sdk_ssp.SystemSecurityPlan){
	if(ssp.Metadata == nil){
		ssp.Metadata = &sdk_ssp.Metadata{}
	}
}

func GuardSystemImplementation(ssp *sdk_ssp.SystemSecurityPlan){
	if(ssp.SystemImplementation == nil){
		ssp.SystemImplementation = &sdk_ssp.SystemImplementation{}
	}
}

func GuardControlImplementation(ssp *sdk_ssp.SystemSecurityPlan){
	if(ssp.ControlImplementation == nil){
		ssp.ControlImplementation = &sdk_ssp.ControlImplementation{}
	}
}

func GuardSystemCharacteristics(ssp *sdk_ssp.SystemSecurityPlan){
	if(ssp.SystemCharacteristics == nil){
		ssp.SystemCharacteristics = &sdk_ssp.SystemCharacteristics{}
	}
}


// Handle error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ImplementedComponent
type ImplementedComponent = sdk_ssp.ImplementedComponent

// Metadata : field in profile
type Metadata = validation_root.Metadata

// Profile : basic profile entity
type Profile = sdk_profile.Profile

// Role : field in Metadata
type Role = validation_root.Role

// Merge : field in profile
type Merge = sdk_profile.Merge

// Import : field in Metadata
type Import = sdk_profile.Import

// Modify : field in Profile
type Modify = sdk_profile.Modify

// BackMatter : field in Profile
type BackMatter = validation_root.BackMatter

// Party : field in Metadata
type Party = validation_root.Party

// Org : field in Metadata
type Org = validation_root.Org

// Version : field in Metadata
type Version = validation_root.Version

// ResponsibleParty : field in Metadata
type ResponsibleParty = validation_root.ResponsibleParty

// PartyId : field in Metadata
type PartyId = validation_root.PartyId

// Include : field in Profile
type Include = sdk_profile.Include

// Call : field in Include
type Call = sdk_profile.Call

// LastModified : field in Metadata
type LastModified = validation_root.LastModified

// Resource : field in BackMatter
type Resource = validation_root.Resource

// RLink : field in BackMatter
type RLink = validation_root.Rlink

// OscalVersion: field in Metadata
type OscalVersion = validation_root.OscalVersion

// Desc : field in BackMatter
type Desc = validation_root.Desc

// Alter : field in Modify
type Alter = sdk_profile.Alter

// Addition : field in Imports
type Addition = sdk_profile.Add

// Prop : field in Modify
type Prop = sdk_profile.Prop

// Address : field in Metadata
type Address = validation_root.Address

// AsIs : field in Modify
type AsIs = sdk_profile.AsIs

//
type RolePartyMap = request_models.RolePartyMap

//
type Title = validation_root.Title

//
type Annotation  =validation_root.Annotation

