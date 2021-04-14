package ssp

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	sdk_ssp "github.com/docker/oscalkit/types/oscal/system_security_plan"
	//request_models "github.com/infobeyondtech/oscal-processor/models/requests"
	information "github.com/infobeyondtech/oscal-processor/models/information"
)

func TestCreateSSP(t *testing.T) {
	fid, err := CreateFreshSSP()
	check(err)
	
	fmt.Printf("file name:" +fid)
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

	firstUser := ssp.SystemImplementation.Users[0]
	
	WriteToFile(ssp)	// this function will initiate file id if not set
	fmt.Printf("file name:" +ssp.Id)

	assert.Equal(string(firstUser.Title), title)
	assert.Equal(string(firstUser.RoleIds[0]), roleId)
}


