package profile

import (
	"encoding/xml"
	"errors"
	"fmt"
	
	
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
	. "github.com/ahmetb/go-linq/v3"
	sdk_profile "github.com/docker/oscalkit/types/oscal/profile"
	"github.com/google/uuid"
	"github.com/infobeyondtech/oscal-processor/context"
	
)


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

        version := "2015-01-22"
        oscalVersion := "1.0.0-milestone1"            // hardcoded oscal version
        sourceType := "application/oscal.catalog+xml" // hardcoded source type
        SetTitleVersion(p, version, oscalVersion, sourceType)

        guardImport(p)

        // examine Id
        Id := p.Id
        if Id != tt.expectId {
            t.Errorf("profile id, got: %s, expectId: %s", Id, tt.expectId)
        }

        // validate profile
        valid := validateProfile(p)
        if valid != true {
            t.Errorf("profile not valid = %v", valid)
        }
    }
}

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

            // examine title
            title := tt.args.profile.Metadata.Title
            if tt.expectTitle != string(title) {
                t.Errorf("got: %s, expectTitle: %s", string(title), tt.expectTitle)
            }

            // examine version
            version := tt.args.profile.Metadata.Version
            if tt.expectVersion != string(version) {
                t.Errorf("got: %s, expectVersion: %s", string(version), tt.expectVersion)
            }

            // examine oscal version
            oscalVersion := tt.args.profile.Metadata.OscalVersion
            if tt.expectOscalVersion != string(oscalVersion) {
                t.Errorf("got: %s, expectOscalVersion: %s", string(oscalVersion), tt.expectOscalVersion)
            }

        })

        // set id
        fid := uuid.New().String()
        SetID(p, "_"+fid)
        guardImport(p)

        // validate profile
        valid := validateProfile(p)
        if valid != true {
            t.Errorf("profile not valid = %v", valid)
        }
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

            // examine role id, org name and title
            if roleId != tt.expectRoleID {
                t.Errorf("got: %s, expectroleId: %s", roleId, tt.expectRoleID)
            }
            if string(orgName) != tt.expectOrgName {
                t.Errorf("got: %s, expectOrgName: %s", orgName, tt.expectOrgName)
            }
            if string(title) != tt.expectTitle {
                t.Errorf("got: %s, expectTitle: %s", title, tt.expectTitle)
            }

            // examine party id
            if string(partyId) != tt.expectPartyID {
                t.Errorf("got: %s, expectPartyID: %s", partyId, tt.expectPartyID)
            }

            // examine party and role relationship
            if firstRPRoleId != tt.expectRoleID {
                t.Errorf("got: %s, expectroleId: %s", firstRPRoleId, tt.expectRoleID)
            }

            // examine first responsible party id
            if string(firstRPPartyId) != tt.expectPartyID {
                t.Errorf("got: %s, expectPartyID: %s", firstRPPartyId, tt.expectPartyID)
            }
        })

        version := "2015-01-22"
        oscalVersion := "1.0.0-milestone1"            // hardcoded oscal version
        sourceType := "application/oscal.catalog+xml" // hardcoded source type
        SetTitleVersion(p, version, oscalVersion, sourceType)

        // set id
        fid := uuid.New().String()
        SetID(p, "_"+fid)
        guardImport(p)

        // validate profile
        valid := validateProfile(p)
        if valid != true {
            t.Errorf("profile not valid = %v", valid)
        }
    }
}

func TestAddControls(t *testing.T) {
    p := &sdk_profile.Profile{}
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
                profile:   p,
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

        })

        version := "2015-01-22"
        oscalVersion := "1.0.0-milestone1"            // hardcoded oscal version
        sourceType := "application/oscal.catalog+xml" // hardcoded source type
        SetTitleVersion(p, version, oscalVersion, sourceType)

        // set id
        fid := uuid.New().String()
        SetID(p, "_"+fid)
        guardImport(p)

        // validate profile
        valid := validateProfile(p)
        if valid != true {
            t.Errorf("profile not valid = %v", valid)
        }
    }
}

func TestAddBackMatter(t *testing.T) {
    p := &sdk_profile.Profile{}
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
                profile: p,
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

        })

        version := "2015-01-22"
        oscalVersion := "1.0.0-milestone1"            // hardcoded oscal version
        sourceType := "application/oscal.catalog+xml" // hardcoded source type
        SetTitleVersion(p, version, oscalVersion, sourceType)

        // set id
        fid := uuid.New().String()
        SetID(p, "_"+fid)
        guardImport(p)

        // validate profile
        valid := validateProfile(p)
        if valid != true {
            t.Errorf("profile not valid = %v", valid)
        }
    }
}

func TestCreateProfile(t *testing.T) {
    orgName := "Infobeyondtech"
    orgEmail := "info@infobeyondtech.com"
    title := "Infobeyond BASELINE_2020"
    baseline := "NIST Special Publication 800-53 Revision 4: Security and Privacy Controls for Federal Information Systems and Organizations"
    controls := []string{"cp-1", "cp-10", "cp-2", "cp-3", "cp-4", "ir-1", "ir-2", "ir-3", "ir-4", "ir-5", "ir-6"}
    source := []string{"NIST_SP-800-53_rev4_catalog.xml"}
    filePath, err := CreateProfile(controls, baseline, source, title, "be3f5ab3-dbe0-4293-a2e0-8182c7fddc24", orgName, orgEmail)
    check(err)

    // create profile will generate a file under download folder
    // read from file and load it into a profile structure
    p := &sdk_profile.Profile{}
    LoadFromFile(p, filePath)

    // validate profile
    valid := validateProfile(p)
    if valid != true {
        t.Errorf("CreateProfile() result not valid = %v", valid)
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

// helper function: track error
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// helper function: validate profile
func validateProfile(p *Profile) bool {

    // validate profile
    parent := context.DownloadDir
    targetFile := parent + "/" + p.Id
    targetFile = context.ExpandPath(targetFile)
    xmlFile := targetFile + ".xml"

    out, e := xml.MarshalIndent(p, "  ", "    ")
    check(e)
    ioErr := ioutil.WriteFile(xmlFile, out, 0644)

    // validate the xml file
    valid, ioErr := Validate(xmlFile)

    check(ioErr)

    // os.Remove(xmlFile)

    return valid
}


