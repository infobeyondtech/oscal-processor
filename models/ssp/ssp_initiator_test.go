package ssp

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	sdk_ssp "github.com/docker/oscalkit/types/oscal/system_security_plan"
	information "github.com/infobeyondtech/oscal-processor/models/information"
	request_models "github.com/infobeyondtech/oscal-processor/models/data_models"
)

func TestAddUser(t *testing.T){
	assert := assert.New(t)
	ssp := &sdk_ssp.SystemSecurityPlan{}
	roleId := "asset-owner"
	title := "Audit Team"
	db_user := &information.User{
		Id: "2",
		UUID: "ae8de94c-835d-4303-83b1-114b6a117a07",
		RoleId: roleId ,
		Type: "internal",
		Title: title,
	}
	AddUser(ssp, *db_user)	
	path := WriteToFile(ssp)		// this function will initiate file id if not set
	fmt.Printf("file path:" +path)

	// load from file and check its field
	ssp_cpy := &sdk_ssp.SystemSecurityPlan{}
	LoadFromFile(ssp_cpy, path)
	firstUser := ssp_cpy.SystemImplementation.Users[0]	

	assert.Equal(string(firstUser.Title), title)
	assert.Equal(string(firstUser.RoleIds[0]), roleId)
}

func TestAddParty(t *testing.T){
	assert := assert.New(t)
	partyId := "3b2a5599-cc37-403f-ae36-5708fa804b27"
	ssp := &sdk_ssp.SystemSecurityPlan{}
	AddParty(ssp, partyId)

	path := WriteToFile(ssp)		
	fmt.Printf("file path:" +path)
	ssp_cpy := &sdk_ssp.SystemSecurityPlan{}
	LoadFromFile(ssp_cpy, path)
	firstParty := ssp_cpy.Metadata.Parties[0]

	assert.Equal(firstParty.Id, partyId)
}

func TestAddComponent(t *testing.T){
	assert := assert.New(t)
	componentUUID := "795533ab-9427-4abe-820f-0b571bacfe6d"
	componentType := "policy"
	componentTitle := "Enterprise Logging, Monitoring, and Alerting Policy"
	compnentState := "operational"

	ssp := &sdk_ssp.SystemSecurityPlan{}
	AddComponent(ssp, componentUUID, []sdk_ssp.ResponsibleRole{})	// add component without roles

	path := WriteToFile(ssp)		
	fmt.Printf("file path:" +path)
	ssp_cpy := &sdk_ssp.SystemSecurityPlan{}
	LoadFromFile(ssp_cpy, path)
	firstComponent := ssp_cpy.SystemImplementation.Components[0]

	assert.Equal(firstComponent.Id, componentUUID)
	assert.Equal(firstComponent.ComponentType, componentType)
	assert.Equal(firstComponent.Status.State, compnentState)
	assert.Equal(string(firstComponent.Title), componentTitle)
}

func TestSetTitileVersion(t *testing.T){
	assert := assert.New(t)
	ssp := &sdk_ssp.SystemSecurityPlan{}
	title := "Enterprise Logging and Auditing System Security Plan"
	version := "1.0"
	oscal_version := "1.0.0-rc1"
	request := request_models.SetTitleVersionRequest{ Title: title , Version: version, OscalVersion: oscal_version}

	SetTitleVersion(ssp, request)
	path := WriteToFile(ssp)		
	fmt.Printf("file path:" +path)

	ssp_cpy := &sdk_ssp.SystemSecurityPlan{}
	LoadFromFile(ssp_cpy, path)
		
	assert.Equal(string(ssp_cpy.Metadata.Version), version)
	assert.Equal(string(ssp_cpy.Metadata.Title), title)
	assert.Equal(string(ssp_cpy.Metadata.OscalVersion), oscal_version)
}

func TestSetCharacteristics(t *testing.T){
	assert := assert.New(t)
	ssp := &sdk_ssp.SystemSecurityPlan{}

	uid := "d7456980-9277-4dcb-83cf-f8ff0442623b"
	systemName := "Enterprise Logging and Auditing System"
	desc := "This is an example of a system that provides enterprise logging and log auditing capabilities."
	deploymentModel := "private"
	securitylevel := "moderate"
	systemInformationTitle := "System and Network Monitoring"
	systemInformationDescription := "This system maintains historical logging and auditing information for all client devices connected to this system."
	confidentialityImpact := "fips-199-moderate"
	integrityImpact := "fips-199-moderate"
	availabilityImpact := "fips-199-low"
	request := request_models.AddSystemCharacteristicReuqest{ 
		UUID : uid,
		SystemName : systemName,
		Description : desc,
		SecurityLevel : securitylevel,
		DeploymentModel : deploymentModel,
		SystemInformationTitle : systemInformationTitle,
		SystemInformationDescription : systemInformationDescription,
		ConfidentialityImpact: confidentialityImpact,
		IntegrityImpact : integrityImpact ,
		AvailabilityImpact : availabilityImpact,
	}
	SetSystemCharacteristic(ssp, request)

	path := WriteToFile(ssp)		
	fmt.Printf("file path:" +path)

	ssp_cpy := &sdk_ssp.SystemSecurityPlan{}
	LoadFromFile(ssp_cpy, path)

	assert.Equal(string(ssp_cpy.SystemCharacteristics.SystemName), systemName)
	assert.Equal(string(ssp_cpy.SystemCharacteristics.Annotations[0].Value), deploymentModel)
	assert.Equal(string(ssp_cpy.SystemCharacteristics.SecuritySensitivityLevel), securitylevel)
}

func TestAddInventoryItem(t *testing.T){
	assert := assert.New(t)
	ssp := &sdk_ssp.SystemSecurityPlan{}
	item_uid := "c9c32657-a0eb-4cf2-b5c1-20928983063c"
	implementComponent_id := "795533ab-9427-4abe-820f-0b571bacfe6d"
	responsibleParties := []RolePartyMap{}

	// two responsible parties
	party1 := RolePartyMap{ UserUUID : "9824089b-322c-456f-86c4-4111c4200f69", PartyUUIDs:[]string{"833ac398-5c9a-4e6b-acba-2a9c11399da0"}}
	party2 := RolePartyMap{ UserUUID : "ae8de94c-835d-4303-83b1-114b6a117a07", PartyUUIDs:[]string{"3b2a5599-cc37-403f-ae36-5708fa804b27"}}
	responsibleParties = append(responsibleParties, party1)
	responsibleParties = append(responsibleParties, party2)

	request := request_models.InsertInventoryItemRequest{
		InventoryItemID : item_uid,
		ImplementComponents : []string{implementComponent_id},
		ResponsibleParties : responsibleParties,
	}

	AddInventoryItem(ssp, request)
	path := WriteToFile(ssp)		
	fmt.Printf("file path:" +path)

	ssp_cpy := &sdk_ssp.SystemSecurityPlan{}
	LoadFromFile(ssp_cpy, path)

	first_itm := ssp_cpy.SystemImplementation.SystemInventory.InventoryItems[0]
	assert.Equal(first_itm.Id, item_uid)
	assert.Equal(first_itm.ImplementedComponents[0].ComponentId, implementComponent_id)
	assert.Equal(len(first_itm.ResponsibleParties),2)
}

func TestAddImplementedRequirement(t *testing.T){
	assert := assert.New(t)
	ssp := &sdk_ssp.SystemSecurityPlan{}
	statements := []request_models.Statement{}

	uuid := "aaadb3ff-6ae8-4332-92db-211468c52af2"
	control_id := "au-1"

	// set parameters
	parameter := request_models.SetParameter{ ParamID : "au-1_prm_1", Value: "all staff and contractors within the organization"}
	parameters := []request_models.SetParameter{parameter}
	
	// users roles
	parties := []string{"ec485dcf-2519-43f5-8e7d-014cc315332d"}
	role := request_models.RolePartyMap{UserUUID:"46ee87f8-724d-42de-907a-670fcb8bd0e3", PartyUUIDs:parties}
	roles := []request_models.RolePartyMap{role}
	
	by_component := request_models.ByComponent{
		ComponentID : "795533ab-9427-4abe-820f-0b571bacfe6d",
		Description : "The legal department develops, documents, and disseminates this policy to all staff and contractors within the organization.",		
		SetParameters : parameters,
		ResponsibleParties : roles,
	}
	by_components := []request_models.ByComponent{by_component}
	smt1 := request_models.Statement{ StatementID : "f3887a91-9ed3-425c-b305-21e4634a1c34",
		 ByComponents:by_components,}
	statements = append(statements, smt1)

	request := request_models.InsertImplementedRequirementRequest{
		UUID: uuid,
		ControlID: control_id,
		Statements: statements,
	}

	AddImplementedRequirement(ssp, request)
	path := WriteToFile(ssp)		
	fmt.Printf("file path:" +path)

	ssp_cpy := &sdk_ssp.SystemSecurityPlan{}
	LoadFromFile(ssp_cpy, path)

	assert.Equal(ssp_cpy.ControlImplementation.ImplementedRequirements[0].ControlId, control_id)
	assert.Equal(ssp_cpy.ControlImplementation.ImplementedRequirements[0].Id, uuid)

}

func TestRemoveInventoryItem(t *testing.T){
	
	// add an inventory item
	// assert := assert.New(t)
	ssp := &sdk_ssp.SystemSecurityPlan{}
	item_uid := "c9c32657-a0eb-4cf2-b5c1-20928983063c"
	implementComponent_id := "795533ab-9427-4abe-820f-0b571bacfe6d"
	responsibleParties := []RolePartyMap{}

	// responsible parties
	party1 := RolePartyMap{ UserUUID : "9824089b-322c-456f-86c4-4111c4200f69", PartyUUIDs:[]string{"833ac398-5c9a-4e6b-acba-2a9c11399da0"}}
	party2 := RolePartyMap{ UserUUID : "ae8de94c-835d-4303-83b1-114b6a117a07", PartyUUIDs:[]string{"3b2a5599-cc37-403f-ae36-5708fa804b27"}}
	responsibleParties = append(responsibleParties, party1)
	responsibleParties = append(responsibleParties, party2)

	request := request_models.InsertInventoryItemRequest{
		InventoryItemID : item_uid,
		ImplementComponents : []string{implementComponent_id},
		ResponsibleParties : responsibleParties,
	}

	// add inventory item
	AddInventoryItem(ssp, request)

	// test remove the only inventory item
	RemoveInventoryItemAt(ssp, item_uid)

	path := WriteToFile(ssp)		
	fmt.Printf("file path:" +path)
}

func TestRemoveImplemented(t *testing.T){
	ssp := &sdk_ssp.SystemSecurityPlan{}
	statements := []request_models.Statement{}

	uuid := "aaadb3ff-6ae8-4332-92db-211468c52af2"
	control_id := "au-1"

	// set parameters
	parameter := request_models.SetParameter{ ParamID : "au-1_prm_1", Value: "all staff and contractors within the organization"}
	parameters := []request_models.SetParameter{parameter}
	
	// users roles
	parties := []string{"ec485dcf-2519-43f5-8e7d-014cc315332d"}
	role := request_models.RolePartyMap{UserUUID:"46ee87f8-724d-42de-907a-670fcb8bd0e3", PartyUUIDs:parties}
	roles := []request_models.RolePartyMap{role}
	
	by_component := request_models.ByComponent{
		ComponentID : "795533ab-9427-4abe-820f-0b571bacfe6d",
		Description : "The legal department develops, documents, and disseminates this policy to all staff and contractors within the organization.",		
		SetParameters : parameters,
		ResponsibleParties : roles,
	}
	by_components := []request_models.ByComponent{by_component}
	smt1 := request_models.Statement{ StatementID : "f3887a91-9ed3-425c-b305-21e4634a1c34",
		 ByComponents:by_components,}
	statements = append(statements, smt1)

	request := request_models.InsertImplementedRequirementRequest{
		UUID: uuid,
		ControlID: control_id,
		Statements: statements,
	}

	// add implemented requirement
	AddImplementedRequirement(ssp, request)

	RemoveImplementedRequirementAt(ssp, uuid)

	path := WriteToFile(ssp)		
	fmt.Printf("file path:" +path)
}

func TestIntegration(t *testing.T){
	assert := assert.New(t)
	// set title version in metadata
	ssp := &sdk_ssp.SystemSecurityPlan{}
	title := "Enterprise Logging and Auditing System Security Plan"
	version := "1.0"
	oscal_version := "1.0.0-milestone1"
	request := request_models.SetTitleVersionRequest{ Title: title , Version: version, OscalVersion: oscal_version}
	SetTitleVersion(ssp, request)

	// add system characteristics
	uid := "d7456980-9277-4dcb-83cf-f8ff0442623b"
	systemName := "Enterprise Logging and Auditing System"
	desc := "This is an example of a system that provides enterprise logging and log auditing capabilities."
	deploymentModel := "private"
	securitylevel := "moderate"
	systemInformationTitle := "System and Network Monitoring"
	systemInformationDescription := "This system maintains historical logging and auditing information for all client devices connected to this system."
	confidentialityImpact := "fips-199-moderate"
	integrityImpact := "fips-199-moderate"
	availabilityImpact := "fips-199-low"
	request1 := request_models.AddSystemCharacteristicReuqest{ 
		UUID : uid,
		SystemName : systemName,
		Description : desc,
		SecurityLevel : securitylevel,
		DeploymentModel : deploymentModel,
		SystemInformationTitle : systemInformationTitle,
		SystemInformationDescription : systemInformationDescription,
		ConfidentialityImpact: confidentialityImpact,
		IntegrityImpact : integrityImpact ,
		AvailabilityImpact : availabilityImpact,
	}
	SetSystemCharacteristic(ssp, request1)

	// add system implementation
	uuid := "aaadb3ff-6ae8-4332-92db-211468c52af2"
	control_id := "au-1"
	statements := []request_models.Statement{}
	parameter := request_models.SetParameter{ ParamID : "au-1_prm_1", Value: "all staff and contractors within the organization"}
	parameters := []request_models.SetParameter{parameter}
	parties := []string{"ec485dcf-2519-43f5-8e7d-014cc315332d"}
	role := request_models.RolePartyMap{UserUUID:"46ee87f8-724d-42de-907a-670fcb8bd0e3", PartyUUIDs:parties}
	roles := []request_models.RolePartyMap{role}
	
	by_component := request_models.ByComponent{
		ComponentID : "795533ab-9427-4abe-820f-0b571bacfe6d",
		Description : "The legal department develops, documents, and disseminates this policy to all staff and contractors within the organization.",		
		SetParameters : parameters,
		ResponsibleParties : roles,
	}
	by_components := []request_models.ByComponent{by_component}
	smt1 := request_models.Statement{ StatementID : "au-1smt.a",
		 ByComponents:by_components,}
	statements = append(statements, smt1)

	request2 := request_models.InsertImplementedRequirementRequest{
		UUID: uuid,
		ControlID: control_id,
		Statements: statements,
	}
	AddImplementedRequirement(ssp, request2)

	// add inventory item
	item_uid := "c9c32657-a0eb-4cf2-b5c1-20928983063c"
	implementComponent_id := "795533ab-9427-4abe-820f-0b571bacfe6d"
	responsibleParties := []RolePartyMap{}
	party1 := RolePartyMap{ UserUUID : "9824089b-322c-456f-86c4-4111c4200f69", PartyUUIDs:[]string{"833ac398-5c9a-4e6b-acba-2a9c11399da0"}}
	party2 := RolePartyMap{ UserUUID : "ae8de94c-835d-4303-83b1-114b6a117a07", PartyUUIDs:[]string{"3b2a5599-cc37-403f-ae36-5708fa804b27"}}
	responsibleParties = append(responsibleParties, party1)
	responsibleParties = append(responsibleParties, party2)

	request3 := request_models.InsertInventoryItemRequest{
		InventoryItemID : item_uid,
		ImplementComponents : []string{implementComponent_id},
		ResponsibleParties : responsibleParties,
	}
	AddInventoryItem(ssp, request3)

	// write to file
	path := WriteToFile(ssp)		
	fmt.Printf("file path: " +path)

	// load ssp model from file
	profileName := "NIST_SP-800-53_rev4_MODERATE-baseline_profile.xml"
	sspModel:= MakeSystemSecurityPlanModel(path, profileName)

	assert.Equal(sspModel.MetaDataModel.Title, ssp.Metadata.Title)

	// todo: check more fields
}














