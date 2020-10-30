package ssp

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/google/uuid"

	sdk_profile "github.com/docker/oscalkit/types/oscal/profile"
	sdk_ssp "github.com/docker/oscalkit/types/oscal/system_security_plan"
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

// handle error
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
