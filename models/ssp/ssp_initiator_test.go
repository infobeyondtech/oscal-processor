package ssp

import (
	"testing"
	"fmt"
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
	ssp := &sdk_ssp.SystemSecurityPlan{}
	db_user := &information.User{
		Id: "2",
		UUID: "ae8de94c-835d-4303-83b1-114b6a117a07",
		RoleId: "asset-owner",
		Type: "internal",
		Title: "Audit Team",
	}
	AddUser(ssp, *db_user)

	// todo: check the field in ssp

	WriteToFile(ssp)	// this function will initiate file id if not set

	fmt.Printf("file name:" +ssp.Id)

}

