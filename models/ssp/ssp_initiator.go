package ssp

import (
	"encoding/xml"
	"io/ioutil"
	"fmt"
	"time"

	"github.com/google/uuid"
	sdk_profile "github.com/docker/oscalkit/types/oscal/profile"
	sdk_ssp "github.com/docker/oscalkit/types/oscal/system_security_plan"
	data_models "github.com/infobeyondtech/oscal-processor/models/data_models"
	information "github.com/infobeyondtech/oscal-processor/models/information"
	"github.com/docker/oscalkit/types/oscal/validation_root"
	"github.com/infobeyondtech/oscal-processor/context"
)

func SetTitleVersion(ssp *sdk_ssp.SystemSecurityPlan, request data_models.SetTitleVersionRequest){
	GuardMetaData(ssp)

	ssp.Metadata.Title = sdk_ssp.Title(request.Title)
	ssp.Metadata.Version = Version(request.Version)
	ssp.Metadata.OscalVersion = OscalVersion(request.OscalVersion)
}

func SetSystemCharacteristic(ssp *sdk_ssp.SystemSecurityPlan, request data_models.AddSystemCharacteristicReuqest){
	GuardSystemCharacteristics(ssp)

	ssp.SystemCharacteristics.SystemName = sdk_ssp.SystemName(request.SystemName)
	ssp.SystemCharacteristics.Description = &sdk_ssp.Markup{Raw:request.Description}
	ssp.SystemCharacteristics.SecuritySensitivityLevel = sdk_ssp.SecuritySensitivityLevel(request.SecurityLevel)

	annotation := &Annotation {Name: "deployment-model", Value:request.DeploymentModel }
	ssp.SystemCharacteristics.Annotations = append(ssp.SystemCharacteristics.Annotations, *annotation)

	systemId := &sdk_ssp.SystemId{ IdentifierType: "https://ietf.org/rfc/rfc4122", Value:request.UUID}
	ssp.SystemCharacteristics.SystemIds = append(ssp.SystemCharacteristics.SystemIds, *systemId)

	GuardSystemInformation(ssp)

	// initiate fields in impact info
	impactInfo := &sdk_ssp.InformationType{}
	impactInfo.Title = sdk_ssp.Title(request.SystemInformationTitle)
	impactInfo.Description = &sdk_ssp.Markup{Raw:request.SystemInformationDescription}
	impactInfo.ConfidentialityImpact = &sdk_ssp.ConfidentialityImpact{Base:sdk_ssp.Base(request.ConfidentialityImpact)}
	impactInfo.AvailabilityImpact = &sdk_ssp.AvailabilityImpact{Base: sdk_ssp.Base(request.AvailabilityImpact)}
	impactInfo.IntegrityImpact = &sdk_ssp.IntegrityImpact{Base: sdk_ssp.Base(request.IntegrityImpact)}

	ssp.SystemCharacteristics.SystemInformation.InformationTypes = append(ssp.SystemCharacteristics.SystemInformation.InformationTypes, *impactInfo)
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

// initiate a ssp instance for the given file id
func LoadFromFileById(ssp *sdk_ssp.SystemSecurityPlan, fileId string){
	parent := context.DownloadDir
	targetFile := parent + "/" + fileId
	targetFile = context.ExpandPath(targetFile)
	xmlFile := targetFile + ".xml"

	dat, e := ioutil.ReadFile(xmlFile)
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

// marshal a ssp into a xml file, returns the xml path
func WriteToFile(ssp *sdk_ssp.SystemSecurityPlan) string{
	parent := context.DownloadDir
	if(ssp.Id == ""){
		ssp.Id = uuid.New().String()
	}	
	targetFile := parent + "/" + ssp.Id
	targetFile = context.ExpandPath(targetFile)
	xmlFile := targetFile + ".xml"

	// set modification date
	dt := time.Now()
	GuardMetaData(ssp)
	ssp.Metadata.LastModified = validation_root.LastModified(dt.String())

	// marshal to xml
	out, e := xml.MarshalIndent(ssp, "  ", "    ")
	check(e)
	
	ioErr := ioutil.WriteFile(xmlFile, out, 0644)
	check(ioErr)

	return xmlFile
}

// insert an inventory item
func AddInventoryItem(ssp *sdk_ssp.SystemSecurityPlan, request data_models.InsertInventoryItemRequest){
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
	if(ssp.SystemImplementation.SystemInventory==nil){
		ssp.SystemImplementation.SystemInventory = &sdk_ssp.SystemInventory{}
	}
	ssp.SystemImplementation.SystemInventory.InventoryItems = append(ssp.SystemImplementation.SystemInventory.InventoryItems, *sdk_itm)
}

// insert an implemented requirement
func AddImplementedRequirement(ssp *sdk_ssp.SystemSecurityPlan, requirement data_models.InsertImplementedRequirementRequest){
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
	GuardSystemImplementation(ssp)

	// checkt existing components
	for _, component := range ssp.SystemImplementation.Components{
		if(component.Id == componentId){
			// merge roles within a component and finish		
			toAppend := []sdk_ssp.ResponsibleRole {}
			
			for _, newRole := range responsibleRoles{
				duplicate := false
				for _, existRole := range component.ResponsibleRoles{
					if(newRole.RoleId == existRole.RoleId){
						duplicate = true
						break
					}
				}
				if(!duplicate){
					toAppend = append(toAppend, newRole)
				}
			}
			component.ResponsibleRoles = append(component.ResponsibleRoles, toAppend...)
			return
		}
	}

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
	sdk_component.ResponsibleRoles = responsibleRoles

	// insert into ssp component collection	
	ssp.SystemImplementation.Components = append(ssp.SystemImplementation.Components, *sdk_component)
}

// private func to add a user in system-implementation, check duplicates
func AddUser(ssp *sdk_ssp.SystemSecurityPlan, db_user information.User){
	GuardSystemImplementation(ssp)

	// no duplicate users
	for _, user := range ssp.SystemImplementation.Users{
		if(user.Id == db_user.UUID){
			return
		}		
	}
	
	// get user info from DB
	sdk_user :=  &sdk_ssp.User{}
	sdk_user.Title = Title(db_user.Title)
	sdk_user.Id = db_user.UUID
	sdk_user.RoleIds = []sdk_ssp.RoleId{sdk_ssp.RoleId(db_user.RoleId)}

	annotation := &Annotation {Name: "type", Value:db_user.Type }
	sdk_user.Annotations = []sdk_ssp.Annotation{ *annotation }

	// insert into ssp header
	ssp.SystemImplementation.Users = append(ssp.SystemImplementation.Users, *sdk_user)
}

// private func to add a party in meta data, check duplicates
func AddParty(ssp *sdk_ssp.SystemSecurityPlan, partyId string){
	GuardMetaData(ssp)

	// no duplicate parties
	for _, party := range ssp.Metadata.Parties{
		if(party.Id == partyId){
			return
		}		
	}
	
	// get party info from DB
	db_party := information.GetParty(partyId);
	sdk_party := &Party{}
	sdk_party.Id = db_party.UUID

	// party name and type are missing in this sdk
	// using properties field in sdk_party for party name and type
	nameProperty := &Prop{Name:"name", Value:db_party.RoleId}
	typeProperty := &Prop{Name:"type", Value:db_party.Type}
	sdk_party.Properties = append(sdk_party.Properties, *nameProperty)
	sdk_party.Properties = append(sdk_party.Properties, *typeProperty)

	// insert into ssp header
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

func GuardSystemInformation(ssp *sdk_ssp.SystemSecurityPlan){
	GuardSystemCharacteristics(ssp)
	if(ssp.SystemCharacteristics.SystemInformation == nil){
		ssp.SystemCharacteristics.SystemInformation = &sdk_ssp.SystemInformation{}
	}
}

func MakeSystemSecurityPlanModel(path string, profileName string) SystemSecurityPlanModel{
	// load from file
	ssp := &sdk_ssp.SystemSecurityPlan{}
	sspModel := SystemSecurityPlanModel{}
	dat, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("error: %v", e)
		return sspModel
	}

	// unmarshal into data structure
	marshalError := xml.Unmarshal([]byte(dat), &ssp)
	if marshalError != nil {
		fmt.Printf("error: %v", marshalError)
		return sspModel
	}

	// turn into data model
	sspModel.ImportProfile = profileName
	sspModel.MetaDataModel = data_models.MetaData{}
	sspModel.SystemCharacteristicModel = data_models.SystemCharacteristic{}
	sspModel.SystemImplementationModel = data_models.SystemImplementation{}
	sspModel.ControlImplementationModel = data_models.ControlImplementation{}

	// metadata
	if(ssp.Metadata!=nil){
		sspModel.MetaDataModel.Version = string(ssp.Metadata.Version)
		sspModel.MetaDataModel.OscalVersion = string(ssp.Metadata.OscalVersion)
		sspModel.MetaDataModel.Title = string(ssp.Metadata.Title)
		sspModel.MetaDataModel.LastModified = string(ssp.Metadata.LastModified)
		for _, party := range ssp.Metadata.Parties{
			// find type and name property from properties
			partyName := findPropValue(party.Properties, "name")
			partyType := findPropValue(party.Properties, "type")
			partyModel := data_models.Party{
				Uuid: party.Id,	
				Name: partyName,
				Type: partyType,
			}
			sspModel.MetaDataModel.Parties = append(sspModel.MetaDataModel.Parties, partyModel)
		}	
	}

	// System Characteristic
	if(ssp.SystemCharacteristics!= nil && ssp.SystemCharacteristics.SystemInformation != nil){

		informationValue := ssp.SystemCharacteristics.SystemInformation.InformationTypes[0]
		sspModel.SystemCharacteristicModel.SystemName = string(ssp.SystemCharacteristics.SystemName)
		sspModel.SystemCharacteristicModel.Description = string(ssp.SystemCharacteristics.Description.Raw)
    	sspModel.SystemCharacteristicModel.SecurityLevel = string(ssp.SystemCharacteristics.SecuritySensitivityLevel)

		sspModel.SystemCharacteristicModel.SystemInformationTitle = string(informationValue.Title)
		sspModel.SystemCharacteristicModel.SystemInformationDescription = string(informationValue.Description.Raw)
		sspModel.SystemCharacteristicModel.IntegrityImpact = string(informationValue.IntegrityImpact.Base)
		sspModel.SystemCharacteristicModel.ConfidentialityImpact = string(informationValue.ConfidentialityImpact.Base)
		sspModel.SystemCharacteristicModel.AvailabilityImpact = string(informationValue.AvailabilityImpact.Base)
	}

	// System Implementation
	if(ssp.SystemImplementation!=nil){		
	for _, user := range ssp.SystemImplementation.Users{
	userModel := data_models.User{
		Uuid: user.Id,
		Title: string(user.Title),
		Type: string(user.Annotations[0].Value),
		RoleId: string(user.RoleIds[0]),			
	}
	sspModel.SystemImplementationModel.Users = append(sspModel.SystemImplementationModel.Users, userModel)
	}	
	
	// sspModel.SystemImplementationModel.Components
	for _, component := range ssp.SystemImplementation.Components{
		componentModel := data_models.Component{
			Uuid : component.Id,
			Type : component.ComponentType,
			Title : string(component.Title),
			Description : string(component.Description.Raw),
			Status : component.Status.State,
		}
		// responsible roles
		for _, role := range component.ResponsibleRoles{
			ResponsibleRole := data_models.ResponsibleRole{
				RoleId: role.RoleId,
			}
			for _, partyId := range role.PartyIds{
				ResponsibleRole.PartyIds = append(ResponsibleRole.PartyIds, string(partyId))
			}
			componentModel.ResponsibleRoles = append(componentModel.ResponsibleRoles, ResponsibleRole)
		}
		sspModel.SystemImplementationModel.Components = append(sspModel.SystemImplementationModel.Components, componentModel)
	}

	// sspModel.SystemImplementationModel.InventoryItems
	for _, item := range ssp.SystemImplementation.SystemInventory.InventoryItems{
		itemModel := data_models.InventoryItem{
			Uuid : item.Id,
			Description : item.Description.Raw,
			AssetId : item.AssetId,		
		}
		for _,impl := range item.ImplementedComponents{
			itemModel.ImplementComponentIds = append(itemModel.ImplementComponentIds, impl.ComponentId)
		}
		sspModel.SystemImplementationModel.InventoryItems = append(sspModel.SystemImplementationModel.InventoryItems, itemModel)
	}
	}

	// Control Implementation
	// sspModel.ControlImplementationModel.ImplementedRequirements
	if(ssp.ControlImplementation!=nil){
	for _, req := range ssp.ControlImplementation.ImplementedRequirements{
		reqModel := data_models.ImplementedRequirement{
			Uuid: req.Id,
			ControlId: req.ControlId,
		}

		// statements
		for _, statement := range req.Statements{
			statementModel := data_models.StatementModel{
				StatementId : statement.StatementId,
			}
			// byComponents array
			for _, bycomponent := range statement.ByComponents{
				bycomponentModel := data_models.ByComponentModel{
					ComponentUuid : bycomponent.ComponentId,
					Description : bycomponent.Description.Raw,					
				}
				// setParameter arary
				for _, setParam:= range bycomponent.ParameterSettings{
					paramModel := data_models.Parameter{
						ParamId: setParam.ParamId,
						Value: string(setParam.Value),
					}
					bycomponentModel.Parameters = append(bycomponentModel.Parameters, paramModel)
				}
				statementModel.ByComponents = append(statementModel.ByComponents, bycomponentModel)
			}
			reqModel.Statements = append(reqModel.Statements, statementModel)
		}

		sspModel.ControlImplementationModel.ImplementedRequirements = append(sspModel.ControlImplementationModel.ImplementedRequirements, reqModel)
	}
	}

	return sspModel
}

// Handle error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// find the value in a property array
func findPropValue(properties []Prop, key string) string{
	if(properties==nil){
		return ""
	}
	for _, prop := range properties{
		if(prop.Name == key){
			return prop.Value
		}
	}
	return ""
}

// remove an implemented requirement
func RemoveImplementedRequirementAt(ssp *sdk_ssp.SystemSecurityPlan, reqId string){
	// check if element container in the xml exist
	if(ssp.ControlImplementation==nil){
		return
	}

	// find the index of the element
	index := -1
	for i:=0;i<len(ssp.ControlImplementation.ImplementedRequirements);i++{
		if(ssp.ControlImplementation.ImplementedRequirements[i].Id == reqId){
			index = i
			break
		}
	}
	if(index == -1){
		return	// didn't find the element
	}

	// handle the case where the target element is the only element left
	if(index==0 || len(ssp.ControlImplementation.ImplementedRequirements)==1){
		ssp.ControlImplementation = nil
		return
	}else{
		// remove that slice at index
		ssp.ControlImplementation.ImplementedRequirements[index] = ssp.ControlImplementation.ImplementedRequirements[len(ssp.ControlImplementation.ImplementedRequirements)-1]
		ssp.ControlImplementation.ImplementedRequirements = ssp.ControlImplementation.ImplementedRequirements[:len(ssp.ControlImplementation.ImplementedRequirements)-1]
		return
	}
}

// remove an inventory item 
func RemoveInventoryItemAt(ssp *sdk_ssp.SystemSecurityPlan, itemId string){
	// check if element container in the xml exist
	if(ssp.SystemImplementation==nil){
		return
	}
	if(ssp.SystemImplementation.SystemInventory==nil){
		return
	}

	// find the index of the element
	index := -1
	for i:=0;i<len(ssp.SystemImplementation.SystemInventory.InventoryItems);i++{
		if(ssp.SystemImplementation.SystemInventory.InventoryItems[i].Id == itemId){
			index = i
			break
		}
	}
	if(index == -1){
		return // didn't find the element
	}

	// handle the case where the target element is the only element left
	if(index==0 || len(ssp.SystemImplementation.SystemInventory.InventoryItems)==1){
		ssp.SystemImplementation.SystemInventory = nil
		return
	}else{
		// remove that slice at index
		ssp.SystemImplementation.SystemInventory.InventoryItems[index] = ssp.SystemImplementation.SystemInventory.InventoryItems[len(ssp.SystemImplementation.SystemInventory.InventoryItems)-1]
		ssp.SystemImplementation.SystemInventory.InventoryItems = ssp.SystemImplementation.SystemInventory.InventoryItems[:len(ssp.SystemImplementation.SystemInventory.InventoryItems)-1]
		return
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
type RolePartyMap = data_models.RolePartyMap

//
type Title = validation_root.Title

//
type Annotation  =validation_root.Annotation

// 
type SystemSecurityPlanModel = data_models.SystemSecurityPlanModel

