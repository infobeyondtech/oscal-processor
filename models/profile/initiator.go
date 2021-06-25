package profile

import (
    "encoding/xml"
    "errors"
    "fmt"
    "io/ioutil"
    "time"

    "github.com/google/uuid"

    "github.com/docker/oscalkit/pkg/oscal_source"
    sdk_profile "github.com/docker/oscalkit/types/oscal/profile"
    "github.com/docker/oscalkit/types/oscal/validation_root"

    "github.com/infobeyondtech/oscal-processor/context"
    profile_models "github.com/infobeyondtech/oscal-processor/models/data_models/profile_model"
)

// Given a set of controls, a set of catalogs, and a baseline,
// generate a unique ID, which can be used for the following operations.
func CreateProfile(ctrls []string, baseline string, ctlgs []string, title string, orgUuid string, orgName string, orgEmail string) (string, error) {
    // A unique file
    version := "2015-01-22"
    oscalVersion := "1.0.0-milestone1"            // hardcoded oscal version
    sourceType := "application/oscal.catalog+xml" // hardcoded source type

    fid := uuid.New().String()
    parent := context.DownloadDir
    targetFile := parent + "/" + fid
    targetFile = context.ExpandPath(targetFile)
    xmlFile := targetFile + ".xml"

    // generate profile and write to file
    p := &sdk_profile.Profile{}

    SetID(p, "_"+fid)
    SetTitleVersion(p, version, oscalVersion, title)
    partyID := orgUuid

    AddRoleParty(p, "creator", "Document Creator", partyID, orgName, orgEmail)
    sourceID := "catalog"
    controls := ctrls
    AddControls(p, controls, "#"+sourceID)

    // check ctlgs and baseline
    if len(ctlgs) == 0 {
        return fid, errors.New("ctlgs cannot be empty")
    }
    if len(baseline) == 0 {
        return fid, errors.New("baseline cannot be empty")
    }

    // set ctlgs and baseline
    description := baseline
    source := ctlgs[0]

    // back matter
    AddBackMatter(p, sourceID, description, source, sourceType)

    // marshal
    out, e := xml.MarshalIndent(p, "  ", "    ")
    if e != nil {
        return fid, e
    }

    // target file has no file type, but content is in xml format
    err := ioutil.WriteFile(xmlFile, out, 0644)

    // Returns the unique file id, if everything is correct
    return xmlFile, err
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

func WriteToFile(p *sdk_profile.Profile) string {
    parent := context.DownloadDir
    targetFile := parent + "/" + p.Id
    targetFile = context.ExpandPath(targetFile)
    xmlFile := targetFile + ".xml"

    out, e := xml.MarshalIndent(p, "  ", "    ")
    check(e)

    // set modification date
    dt := time.Now()
    guardMetadata(p)
    p.Metadata.LastModified = validation_root.LastModified(dt.String())

    ioErr := ioutil.WriteFile(xmlFile, out, 0644)
    check(ioErr)

    return xmlFile
}

func SetID(profile *sdk_profile.Profile, id string) {
    profile.Id = id
}

// SetTitleVersion : set profile title, version, oscal version, modify date in metadata
func SetTitleVersion(profile *sdk_profile.Profile, version string, oscalVersion string, title string) error {
    // metadata
    guardMetadata(profile)
    profile.Metadata.Title = validation_root.Title(title)
    profile.Metadata.LastModified = LastModified((time.Now().Format(time.RFC3339)))
    profile.Metadata.Version = validation_root.Version(version)
    profile.Metadata.OscalVersion = validation_root.OscalVersion(oscalVersion)

    return nil
}

// AddRoleParty : append a role to profile metadata. Note: does not check duplicate roles
func AddRoleParty(profile *sdk_profile.Profile, roleID string, title string, partyID string, orgName string, email string) error {
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

    return nil
}

// AddAddress : append an address to profile metadata. Note: does not check duplicate addresses
func AddParty(profile *sdk_profile.Profile, orgName string, addressLines []string, city string, state string, postalCode string, roleId string, partyId string) error {
    guardMetadata(profile)

    // check if the party exist
    //addressExist := false
    newtParty := &Party{}
    /*for _, p := range profile.Metadata.Parties {
        if p.Id == partyID {
            addressExist = true
            existParty = &p
            break
        }
    }
    if !addressExist {
        return errors.New("address not exist")
    }*/

    var addressArr []Address

    address := &Address{
        City:       validation_root.City(city),
        PostalCode: validation_root.PostalCode(postalCode),
        State:      validation_root.State(state)}

    // address lines
    for _, line := range addressLines {
        addressLine := validation_root.AddrLine(line)
        address.PostalAddress = append(address.PostalAddress, addressLine)
    }

    newtParty.Org = &Org{
        OrgName:   validation_root.OrgName(orgName),
        Addresses: addressArr,
    }

    newtParty.Org.Addresses = append(newtParty.Org.Addresses, *address)
    newtParty.Id = partyId
    profile.Metadata.Parties = append(profile.Metadata.Parties, *newtParty)

    // finding the responsible role-id
    var notFound = true

    for index, rparty := range profile.Metadata.ResponsibleParties {
        if rparty.RoleId == roleId {
            rparty.PartyIds = append(rparty.PartyIds, (validation_root.PartyId(partyId)))
            profile.Metadata.ResponsibleParties[index].PartyIds = rparty.PartyIds
            notFound = false
        }
    }

    if notFound == true {
        var partyId_arr []PartyId
        partyId_arr = append(partyId_arr, (validation_root.PartyId(partyId)))

        responsibleParty := ResponsibleParty{
            RoleId:   roleId,
            PartyIds: partyId_arr,
        }
        profile.Metadata.ResponsibleParties = append(profile.Metadata.ResponsibleParties, responsibleParty)
    }

    return nil
}

// AddControls : append a list of controls to profile. Note: does not check duplicate controls
func AddControls(profile *sdk_profile.Profile, controls []string, reference string) error {

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
        var found = false
        for _, h := range existImport.Include.IdSelectors {
            fmt.Print("h", h)
            fmt.Print("*control", *control)
            if h == *control {
                found = true
                break
            }
        }
        if found == false {
            existImport.Include.IdSelectors = append(existImport.Include.IdSelectors, *control)
        }
    }

    if !importExist {
        // this is a newly created import
        profile.Imports = append(profile.Imports, *existImport)
    }
    return nil
}

// SetMerge : set the merge value in profile
func SetMerge(profile *sdk_profile.Profile, asIs string) error {
    guardMerge(profile)
    profile.Merge.AsIs = AsIs(asIs)
    return nil
}

// AddModification : append a modification to profile. Note: does not check duplicate modifications
func AddModification(profile *sdk_profile.Profile, controlID string, position string, name string, value string) error {

    guardModify(profile)
    // todo: only support alter for the time being
    alter := &Alter{ControlId: controlID}
    prop := &Prop{Name: name, Value: value}
    addition := &Addition{Position: position}
    addition.Properties = append(addition.Properties, *prop)
    alter.Additions = append(alter.Additions, *addition)
    profile.Modify.Alterations = append(profile.Modify.Alterations, *alter)

    return nil
}

// AddBackMatter : append a back matter source to profile backmatter. Note: does not check duplicate back matter source
func AddBackMatter(profile *sdk_profile.Profile, id string, desc string, link string, media string) error {
    guardBackMatter(profile)
    resource := &Resource{Id: id, Desc: Desc(desc)}
    rlink := &RLink{Href: link, MediaType: media}
    resource.Rlinks = append(resource.Rlinks, *rlink)
    profile.BackMatter.Resources = append(profile.BackMatter.Resources, *resource)
    return nil
}

// underlaying functions
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

// Validate : examine if this profile is a valid
func Validate(path string) (bool, error) {
    os, err := oscal_source.Open(path)
    if err != nil {
        return false, err
    }
    defer os.Close()

    err = os.Validate()
    if err != nil {
        return false, err
    }

    return true, nil
}

// helper function: track error
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func MakeProfileModel(path string) ProfileModel {
    // load from file
    profile := &sdk_profile.Profile{}
    profile_Model := profile_models.ProfileModel{}
    dat, e := ioutil.ReadFile(path)

    if e != nil {
        fmt.Printf("error: %v", e)
        return profile_Model
    }

    // unmarshal into data structure
    marshalError := xml.Unmarshal([]byte(dat), &profile)
    if marshalError != nil {
        fmt.Printf("error: %v", marshalError)
        return profile_Model
    }

    profile_Model.Metadata = profile_models.Metadata{}
    profile_Model.Imports = profile_models.Import{}
    profile_Model.BackMatter = profile_models.BackMatter{}

    // metadata
    if profile.Metadata != nil {
        profile_Model.Metadata.Title = string(profile.Metadata.Title)
        profile_Model.Metadata.Version = string(profile.Metadata.Version)
        profile_Model.Metadata.OscalVersion = string(profile.Metadata.OscalVersion)
        profile_Model.Metadata.LastModified = string(profile.Metadata.LastModified)
        var party_arr []profile_models.Party
        for _, party := range profile.Metadata.Parties {
            var AddressArr []profile_models.Address
            for _, address := range party.Org.Addresses {
                var PostalAddressArr []string
                for _, PostalAddress := range address.PostalAddress {
                    PostalAddressArr = append(PostalAddressArr, string(PostalAddress))
                }
                new_address := profile_models.Address{
                    Type:          string(address.Type),
                    PostalAddress: PostalAddressArr,
                    City:          string(address.City),
                    State:         string(address.State),
                    PostalCode:    string(address.PostalCode),
                    Country:       string(address.Country),
                }
                AddressArr = append(AddressArr, new_address)
            }

            new_org := profile_models.Org{
                OrgName:   string(party.Org.OrgName),
                Addresses: AddressArr,
            }

            new_party := profile_models.Party{
                Id:  party.Id,
                Org: new_org,
            }
            party_arr = append(party_arr, new_party)
        }
        profile_Model.Metadata.Parties = party_arr

        var responsibleParties_arr []profile_models.ResponsibleParty
        for _, responsibleParty := range profile.Metadata.ResponsibleParties {
            var partyIdArr []string
            for _, partyId := range responsibleParty.PartyIds {
                partyIdArr = append(partyIdArr, string(partyId))
            }
            new_responsibleParty := profile_models.ResponsibleParty{
                RoleId:   responsibleParty.RoleId,
                PartyIds: partyIdArr,
            }
            responsibleParties_arr = append(responsibleParties_arr, new_responsibleParty)
        }
        profile_Model.Metadata.ResponsibleParties = responsibleParties_arr
    }

    profile_Model.Imports.Href = string(profile.Imports[0].Href)
    for _, include := range profile.Imports[0].Include.IdSelectors {
        profile_Model.Imports.Include = append(profile_Model.Imports.Include, include.ControlId)
    }

    for _, resource := range profile.BackMatter.Resources {
        var rlink_arr []profile_models.Rlink
        for _, rlinks := range resource.Rlinks {
            rlink := profile_models.Rlink{
                Href:      rlinks.Href,
                MediaType: rlinks.MediaType,
            }
            rlink_arr = append(rlink_arr, rlink)
        }

        resources := profile_models.Resource{
            Id:     resource.Id,
            Desc:   string(resource.Desc),
            Rlinks: rlink_arr,
        }
        profile_Model.BackMatter.Resources = append(profile_Model.BackMatter.Resources, resources)
    }
    return profile_Model

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

type ProfileModel = profile_models.ProfileModel

// find the value in a property array
func findPropValue(properties []Prop, key string) string {
    if properties == nil {
        return ""
    }
    for _, prop := range properties {
        if prop.Name == key {
            return prop.Value
        }
    }
    return ""
}

type Rlink = validation_root.Rlink
