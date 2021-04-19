package ssp

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
	sdk_ssp "github.com/docker/oscalkit/types/oscal/system_security_plan"
	information "github.com/infobeyondtech/oscal-processor/models/information"
	request_models "github.com/infobeyondtech/oscal-processor/models/requests"
)

func TestCreateSSP(t *testing.T) {
	fid, err := CreateFreshSSP()
	check(err)
	
	fmt.Printf("fid:" +fid)
}

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
	request := request_models.AddSystemCharacteristicReuqest{ 
		UUID : uid,
		SystemName : systemName,
		Description : desc,
		SecurityLevel : securitylevel}
	SetSystemCharacteristic(ssp, request)

	path := WriteToFile(ssp)		
	fmt.Printf("file path:" +path)

	ssp_cpy := &sdk_ssp.SystemSecurityPlan{}
	LoadFromFile(ssp_cpy, path)

	assert.Equal(string(ssp_cpy.SystemCharacteristics.SystemName), systemName)
	assert.Equal(string(ssp_cpy.SystemCharacteristics.Annotations[0].Value), deploymentModel)
	assert.Equal(string(ssp_cpy.SystemCharacteristics.SecuritySensitivityLevel), securitylevel)
}














