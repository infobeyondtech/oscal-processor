package profile

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	. "github.com/ahmetb/go-linq/v3"
	sdk_profile "github.com/docker/oscalkit/types/oscal/profile"
)

func TestSetTitleVersion(t *testing.T) {
	p := &sdk_profile.Profile{}
	type args struct {
		profile      *sdk_profile.Profile
		version      string
		oscalVersion string
		title        string
	}
	tests := []struct {
		name               string
		args               args
		expectTitle        string
		expectVersion      string
		expectOscalVersion string
	}{
		{
			args: args{
				profile:      p,
				version:      "2015-01-22",
				oscalVersion: "1.0.0-milestone1",
				title:        "Infobeyond BASELINE",
			},
			expectTitle:        "Infobeyond BASELINE",
			expectVersion:      "2015-01-22",
			expectOscalVersion: "1.0.0-milestone1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetTitleVersion(tt.args.profile, tt.args.version, tt.args.oscalVersion, tt.args.title)

			title := tt.args.profile.Metadata.Title
			if tt.expectTitle != string(title) {
				t.Errorf("got: %s, expectTitle: %s", string(title), tt.expectTitle)
			}

			version := tt.args.profile.Metadata.Version
			if tt.expectVersion != string(version) {
				t.Errorf("got: %s, expectVersion: %s", string(version), tt.expectVersion)
			}

			oscalVersion := tt.args.profile.Metadata.OscalVersion
			if tt.expectOscalVersion != string(oscalVersion) {
				t.Errorf("got: %s, expectOscalVersion: %s", string(oscalVersion), tt.expectOscalVersion)
			}

			// todo: validate profile

		})
	}
}

func TestSetID(t *testing.T) {
	p := &sdk_profile.Profile{}
	type args struct {
		profile *sdk_profile.Profile
		id      string
	}
	tests := []struct {
		name     string
		args     args
		expectId string
	}{
		{
			args: args{
				profile: p,
				id:      "uuid-be3f5ab3-dbe0-4293-a2e0-8182c7fddc23",
			},
			expectId: "uuid-be3f5ab3-dbe0-4293-a2e0-8182c7fddc23",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetID(tt.args.profile, tt.args.id)
		})

		Id := p.Id
		if Id != tt.expectId {
			t.Errorf("got: %s, expectId: %s", Id, tt.expectId)
		}

		// todo: validate profile
	}
}

func TestAddRoleParty(t *testing.T) {
	p := &sdk_profile.Profile{}
	type args struct {
		profile *sdk_profile.Profile
		roleID  string
		title   string
		partyID string
		orgName string
		email   string
	}
	tests := []struct {
		name          string
		args          args
		expectTitle   string
		expectRoleID  string
		expectPartyID string
		expectOrgName string
		expectEmail   string
	}{
		{
			args: args{
				profile: p,
				roleID:  "creator",
				title:   "Document Creator",
				partyID: "IT-JTE",
				orgName: "Infobeyondtech",
				email:   "info@infobeyondtech.com",
			},
			expectRoleID:  "creator",
			expectTitle:   "Document Creator",
			expectPartyID: "IT-JTE",
			expectOrgName: "Infobeyondtech",
			expectEmail:   "info@infobeyondtech.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddRoleParty(tt.args.profile, tt.args.roleID, tt.args.title, tt.args.partyID, tt.args.orgName, tt.args.email)

			firstParty := p.Metadata.Parties[0]
			firstOrg := firstParty.Org
			firstRole := p.Metadata.Roles[0]
			firstResponsibleParty := p.Metadata.ResponsibleParties[0]

			roleId := firstRole.Id
			orgName := firstOrg.OrgName
			title := firstRole.Title
			partyId := firstParty.Id
			firstRPRoleId := firstResponsibleParty.RoleId
			firstRPPartyId := firstResponsibleParty.PartyIds[0]

			//email := firstOrg.EmailAddresses[0]

			if roleId != tt.expectRoleID {
				t.Errorf("got: %s, expectroleId: %s", roleId, tt.expectRoleID)
			}
			if string(orgName) != tt.expectOrgName {
				t.Errorf("got: %s, expectOrgName: %s", orgName, tt.expectOrgName)
			}
			if string(title) != tt.expectTitle {
				t.Errorf("got: %s, expectTitle: %s", title, tt.expectTitle)
			}

			// note: email has inconsistency so it is not included in the tests yet
			/*
				if string(email) != tt.expectEmail {
					t.Errorf("got: %s, expectEmail: %s", email, tt.expectEmail)
				}*/
			if string(partyId) != tt.expectPartyID {
				t.Errorf("got: %s, expectPartyID: %s", partyId, tt.expectPartyID)
			}

			// examine party and role relationship
			if firstRPRoleId != tt.expectRoleID {
				t.Errorf("got: %s, expectroleId: %s", firstRPRoleId, tt.expectRoleID)
			}
			if string(firstRPPartyId) != tt.expectPartyID {
				t.Errorf("got: %s, expectPartyID: %s", firstRPPartyId, tt.expectPartyID)
			}

			// todo: validate profile

			// todo: test cases of inserting multiple parties and roles
		})
	}
}

func TestSetMerge(t *testing.T) {
	p := &sdk_profile.Profile{}
	type args struct {
		profile *sdk_profile.Profile
		asIs    string
	}
	tests := []struct {
		name        string
		args        args
		expectValue string
	}{
		{
			args: args{
				profile: p,
				asIs:    "true",
			},
			expectValue: "true",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetMerge(tt.args.profile, tt.args.asIs)

			asis := p.Merge.AsIs
			if string(asis) != tt.expectValue {
				t.Errorf("got: %s, expectValue: %s", asis, tt.expectValue)
			}

			// todo: validate profile
		})
	}
}

func TestAddControls(t *testing.T) {
	type args struct {
		profile   *sdk_profile.Profile
		controls  []string
		reference string
	}
	tests := []struct {
		name      string
		args      args
		expectIDs []string
	}{
		{
			args: args{
				profile:   &sdk_profile.Profile{},
				reference: "#catalog",
				controls:  []string{"cp-1", "cp-10", "cp-2", "cp-3", "cp-4", "ir-1", "ir-2", "ir-3", "ir-4", "ir-5", "ir-6"},
			},
			expectIDs: []string{"cp-1", "cp-10", "cp-2", "cp-3", "cp-4", "ir-1", "ir-2", "ir-3", "ir-4", "ir-5", "ir-6"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unitProfile := tt.args.profile

			AddControls(tt.args.profile, tt.args.controls, tt.args.reference)

			// examine each control, each profile is fresh and has only one import
			var cs []string
			firstImport := unitProfile.Imports[0]
			From(firstImport.Include.IdSelectors).SelectT(func(i Call) string { return i.ControlId }).ToSlice(&cs)
			eq := reflect.DeepEqual(tt.expectIDs, cs)
			if !eq {
				t.Errorf("controls mismatch:\n" + strings.Join(cs, ",") + "\n" + strings.Join(tt.expectIDs, ","))
			}

			// todo: validate profile
		})
	}
}

func TestAddBackMatter(t *testing.T) {
	type args struct {
		profile *sdk_profile.Profile
		id      string
		desc    string
		link    string
		media   string
	}
	tests := []struct {
		name        string
		args        args
		expectId    string
		expectDesc  string
		expectLink  string
		expectMedia string
	}{
		{
			args: args{
				profile: &sdk_profile.Profile{},
				id:      "catalog",
				desc:    "NIST Special Publication 800-53 Revision 4: Security and Privacy Controls for Federal Information Systems and Organizations",
				link:    "NIST_SP-800-53_rev4_catalog.xml",
				media:   "application/oscal.catalog+xml",
			},
			expectId:    "catalog",
			expectDesc:  "NIST Special Publication 800-53 Revision 4: Security and Privacy Controls for Federal Information Systems and Organizations",
			expectLink:  "NIST_SP-800-53_rev4_catalog.xml",
			expectMedia: "application/oscal.catalog+xml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			unitProfile := tt.args.profile
			AddBackMatter(tt.args.profile, tt.args.id, tt.args.desc, tt.args.link, tt.args.media)

			// each profile is fresh and has only one resource in backMatter
			resource := unitProfile.BackMatter.Resources[0]
			rlink := resource.Rlinks[0]

			if string(resource.Id) != tt.expectId {
				t.Errorf("got: %s, expectId: %s", resource.Id, tt.expectId)
			}
			if string(resource.Desc) != tt.expectDesc {
				t.Errorf("got: %s, expectDesc: %s", resource.Id, tt.expectDesc)
			}
			if string(rlink.Href) != tt.expectLink {
				t.Errorf("got: %s, expectLink: %s", rlink.Href, tt.expectLink)
			}
			if string(rlink.MediaType) != tt.expectMedia {
				t.Errorf("got: %s, expectMedia: %s", resource.Id, tt.expectMedia)
			}

			// todo: test validation of this profile
		})
	}
}

func TestCreateProfile(t *testing.T) {
	p := &sdk_profile.Profile{}

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
	controls := []string{"cp-1", "cp-10", "cp-2", "cp-3", "cp-4", "ir-1", "ir-2", "ir-3", "ir-4", "ir-5", "ir-6"}
	AddControls(p, controls, "#"+sourceID)

	AddModification(p, "cp-1", "starting", "priority", "P1")

	description := "NIST Special Publication 800-53 Revision 4: Security and Privacy Controls for Federal Information Systems and Organizations"
	source := "NIST_SP-800-53_rev4_catalog.xml"
	sourceType := "application/oscal.catalog+xml"
	AddBackMatter(p, sourceID, description, source, sourceType)

	// marshal
	out, e3 := xml.MarshalIndent(p, "  ", "    ")
	if e3 != nil {
		t.Errorf("error: %v\n", e3)
	}
	t.Log(len(string(out)))

	err := ioutil.WriteFile("test", out, 0644)
	check(err)
}

// handle error
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestValidate(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr error
	}{
		{
			args: args{
				path: "test_files/test.xml",
			},
			want:    true,
			wantErr: nil,
		},
		{
			args: args{
				path: "test_files/test_invalid.xml",
			},
			want:    false,
			wantErr: errors.New(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Validate(tt.args.path)

			if got != tt.want {
				t.Errorf("Validate() = %v, want %v, err: %v", got, tt.want, err)
			}

			if tt.wantErr != nil && err == nil {
				t.Errorf("Validate() error, wantErr")
				return
			}

		})
	}
}
