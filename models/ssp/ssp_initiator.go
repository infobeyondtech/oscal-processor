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
)

// Create an empty SSP
func CreateFreshSSP() (string, error) {

	fid := uuid.New().String()
	ssp := &sdk_ssp.SystemSecurityPlan{}

	// set id
	ssp.Id = fid

	out, err1 := xml.MarshalIndent(ssp, "  ", "    ")
	check(err1)

	err := ioutil.WriteFile("ssp_test", out, 0644)
	check(err)

	return fid, nil
}

func SetTitleVersion(ssp *sdk_ssp.SystemSecurityPlan, request request_models.SetTitleVersionRequest){

}

func SetSystemCharacteristic(ssp *sdk_ssp.SystemSecurityPlan, request request_models.AddSystemCharacteristicReuqest ){


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
		sdk_party.RoleId = partyRoleMap.RoleID
		
		// insert user role detail in the header
		AddUser(ssp, partyRoleMap.RoleID)
		
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
		AddComponent(ssp, component_id)
	}
	
	// todo: insert inventory item into ssp
}

// insert an implemented requirement
func AddImplementedRequirement(ssp *sdk_ssp.SystemSecurityPlan, requirement request_models.InsertImplementedRequirementRequest){

}

// private func to add a component in system-implementation, check duplicates
func AddComponent(ssp *sdk_ssp.SystemSecurityPlan, componentId string){
	db_component := information.GetComponent(componentId)
	sdk_component := &sdk_ssp.Component{}
	sdk_component.Id = db_component.UUID

	// todo: convert to sdk_component

	// todo: insert into ssp component collection
}

// private func to add a user in system-implementation, check duplicates
func AddUser(ssp *sdk_ssp.SystemSecurityPlan, userId string){
	db_user :=	information.GetUser(userId)

	// todo: convert to sdk_user

	// todo: insert into ssp header
}

// private func to add a party in meta data, check duplicates
func AddParty(ssp *sdk_ssp.SystemSecurityPlan, partyId string){
	db_party := information.GetParty(partyId);

	// todo: convert to sdk_party

	// todo: insert into ssp header
} 


func GuardMetaData(ssp *sdk_ssp.SystemSecurityPlan){

}

func GuardSystemImplementation(ssp *sdk_ssp.SystemSecurityPlan){

}

func GuardControlImplementation(ssp *sdk_ssp.SystemSecurityPlan){

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
