// Provides the functionality of creating a profile
package profile

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/google/uuid"

	sdk_profile "github.com/docker/oscalkit/types/oscal/profile"
	"github.com/docker/oscalkit/types/oscal/validation_root"

	"github.com/infobeyondtech/oscal-processor/context"
)

// Given a set of controls, a set of catalogs, and a baseline,
// generate a unique ID, which can be used for the following operations.
func CreateProfile(ctrls []string, baseline string, ctlgs []string) (string, error) {
	// A unique file
	fid := uuid.New().String()
	parent := context.DownloadDir
	targetFile := parent + "/" + fid
	targetFile = context.ExpandPath(targetFile)

	// generate profile and write to file
	p := &sdk_profile.Profile{}

	// todo: many fields here are hardcoded
	SetID(p, "uuid-be3f5ab3-dbe0-4293-a2e0-8182c7fddc23")
	SetTitleVersion(p, "2015-01-22", "1.0.0-milestone1", "Infobeyond BASELINE")
	partyID := "IT-JTF"
	orgName := "Infobeyondtech"
	orgEmail := "info@infobeyondtech.com"
	AddRoleParty(p, "creator", "Document Creator", partyID, orgName, orgEmail)

	addressLines := []string{"InfoBeyond Technology LLC", "320 Whittington PKWY, STE 117", "Louisville, KY, USA 40222-4917"}
	AddAddress(p, partyID, addressLines, "Louvisville", "KY", "40222-4917")

	SetMerge(p, "true")

	sourceID := "catalog"
	controls := ctrls
	AddControls(p, controls, "#"+sourceID)

	AddModification(p, "cp-1", "starting", "priority", "P1")

	// check ctlgs and baseline
	if len(ctlgs) == 0 {
		return fid, errors.New("ctlgs cannot be empty")
	}
	if len(baseline) == 0 {
		return fid, errors.New("baseline cannot be empty")
	}

	// todo: give ctlgs and baseline to the correct field
	description := baseline
	source := ctlgs[0]

	sourceType := "application/oscal.catalog+xml"
	AddBackMatter(p, sourceID, description, source, sourceType)

	// marshal
	out, e := xml.MarshalIndent(p, "  ", "    ")
	if e != nil {
		return fid, e
	}

	err := ioutil.WriteFile(targetFile, out, 0644)

	// Returns the unique file id, if everything is correct
	return fid, err
}

// LoadFromFile : initiate a profile using a xml file
func LoadFromFile(profile *sdk_profile.Profile, path string) {
	dat, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Printf("error: %v", e)
		return
	}

	// unmarshal into data structure
	marshalError := xml.Unmarshal([]byte(dat), &profile)
	if marshalError != nil {
		fmt.Printf("error: %v", marshalError)
		return
	}
}

func SetID(profile *sdk_profile.Profile, id string) {
	profile.Id = id
}

// SetTitleVersion : set profile title, version, oscal version, modify date in metadata
func SetTitleVersion(profile *sdk_profile.Profile, version string, oscalVersion string, title string) {
	// metadata
	guardMetadata(profile)
	profile.Metadata.Title = validation_root.Title(title)
	profile.Metadata.LastModified = LastModified((time.Now().Format(time.RFC3339)))
	profile.Metadata.Version = validation_root.Version(version)
	profile.Metadata.OscalVersion = validation_root.OscalVersion(oscalVersion)
}

// AddRoleParty : append a role to profile metadata. Note: does not check duplicate roles
func AddRoleParty(profile *sdk_profile.Profile, roleID string, title string, partyID string, orgName string, email string) {
	// metadata
	guardMetadata(profile)

	// add role
	role := &Role{Id: roleID}
	role.Title = validation_root.Title(title)
	profile.Metadata.Roles = append(profile.Metadata.Roles, *role)

	// insert party
	pid := partyID
	party := &Party{Id: pid}
	org := &Org{OrgName: validation_root.OrgName(orgName)}
	org.Addresses = []Address{}

	// WARNING: email address location is inconsistent with oscalkit validation
	// oscalkit expect email address to appear after address while the profile schema defined email address before address
	// however, oscalkit validation says email is an optional field, therefore I omit it for the time-being
	// org.EmailAddresses = append(org.EmailAddresses, validation_root.Email(email))
	party.Org = org
	profile.Metadata.Parties = append(profile.Metadata.Parties, *party)

	// insert link to role
	relation := &ResponsibleParty{RoleId: roleID}
	relation.PartyIds = append(relation.PartyIds, PartyId(pid))
	profile.Metadata.ResponsibleParties = append(profile.Metadata.ResponsibleParties, *relation)
}

// AddAddress : append an address to profile metadata. Note: does not check duplicate addresses
func AddAddress(profile *sdk_profile.Profile, partyID string, addressLines []string, city string, state string, postalCode string) {
	guardMetadata(profile)

	// check if the party exist
	addressExist := false
	existParty := &Party{}
	for _, p := range profile.Metadata.Parties {
		if p.Id == partyID {
			addressExist = true
			existParty = &p
			break
		}
	}
	if !addressExist {
		return
	}

	address := &Address{
		City:       validation_root.City(city),
		PostalCode: validation_root.PostalCode(postalCode),
		State:      validation_root.State(state)}

	// address lines
	for _, line := range addressLines {
		addressLine := validation_root.AddrLine(line)
		address.PostalAddress = append(address.PostalAddress, addressLine)
	}

	// append address
	if existParty.Org == nil {
		existParty.Org = &Org{}
	}
	existParty.Org.Addresses = append(existParty.Org.Addresses, *address)
	//existParty.Org.EmailAddresses = append(existParty.Org.EmailAddresses, validation_root.Email(email))
}

// AddControls : append a list of controls to profile. Note: does not check duplicate controls
func AddControls(profile *sdk_profile.Profile, controls []string, reference string) {

	guardImport(profile)

	// check if the same reference exist
	importExist := false
	existImport := &Import{Href: reference}
	for _, i := range profile.Imports {
		if i.Href == reference {
			importExist = true
			existImport = &i
			break
		}
	}

	if existImport.Include == nil {
		existImport.Include = &Include{}
	}

	// iterate over controls
	for _, c := range controls {
		control := &Call{ControlId: c}
		existImport.Include.IdSelectors = append(existImport.Include.IdSelectors, *control)
	}

	if !importExist {
		// this is a newly created import
		profile.Imports = append(profile.Imports, *existImport)
	}

}

// SetMerge : set the merge value in profile
func SetMerge(profile *sdk_profile.Profile, asIs string) {
	guardMerge(profile)
	profile.Merge.AsIs = AsIs(asIs)
}

// AddModification : append a modification to profile. Note: does not check duplicate modifications
func AddModification(profile *sdk_profile.Profile, controlID string, position string, name string, value string) {

	guardModify(profile)
	// todo: only support alter for the time being
	alter := &Alter{ControlId: controlID}
	prop := &Prop{Name: name, Value: value}
	addition := &Addition{Position: position}
	addition.Properties = append(addition.Properties, *prop)
	alter.Additions = append(alter.Additions, *addition)
	profile.Modify.Alterations = append(profile.Modify.Alterations, *alter)
}

// AddBackMatter : append a back matter source to profile backmatter. Note: does not check duplicate back matter source
func AddBackMatter(profile *sdk_profile.Profile, id string, desc string, link string, media string) {
	guardBackMatter(profile)
	resource := &Resource{Id: id, Desc: Desc(desc)}
	rlink := &RLink{Href: link, MediaType: media}
	resource.Rlinks = append(resource.Rlinks, *rlink)
	profile.BackMatter.Resources = append(profile.BackMatter.Resources, *resource)
}

//
func guardMetadata(profile *sdk_profile.Profile) {
	if profile.Metadata == nil {
		profile.Metadata = &Metadata{}
	}
}

func guardImport(profile *sdk_profile.Profile) {
	if profile.Imports == nil {
		profile.Imports = []Import{}
	}
}

func guardMerge(profile *sdk_profile.Profile) {
	if profile.Merge == nil {
		profile.Merge = &Merge{}
	}
}

func guardModify(profile *sdk_profile.Profile) {
	if profile.Modify == nil {
		profile.Modify = &Modify{}
	}
}

func guardBackMatter(profile *sdk_profile.Profile) {
	if profile.BackMatter == nil {
		profile.BackMatter = &BackMatter{}
	}
}

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
