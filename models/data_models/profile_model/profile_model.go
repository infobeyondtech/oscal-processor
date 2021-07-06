package profile_models

import (
    "strings"

    "github.com/docker/oscalkit/types/oscal/nominal_catalog"
    "github.com/docker/oscalkit/types/oscal/validation_root"
)

type ProfileModel struct {
    Metadata   Metadata   `xml:"metadata,omitempty" json:"metadata,omitempty"`
    Imports    Import     `xml:"import,omitempty" json:"imports,omitempty"`
    BackMatter BackMatter `xml:"back-matter,omitempty" json:"backMatter,omitempty"`
}

type Metadata struct {
    // Role, Responsible party
    Title              string             `json:"title" binding:"required"`
    Version            string             `json:"version" binding:"required"`
    OscalVersion       string             `json:"oscalversion" binding:"required"`
    LastModified       string             `json:"lastModified" binding:"required"`
    Roles              []Role             `xml:"role,omitempty" json:"roles,omitempty"`
    Parties            []Party            `json:"parties" binding:"required"`
    ResponsibleParties []ResponsibleParty `xml:"responsible-party,omitempty" json:"responsible-parties,omitempty"`
}

type Role struct {

    // Unique identifier of the containing object
    Id string `xml:"id,attr,omitempty" json:"id,omitempty"`

    // A title for display and navigation
    Title string `xml:"title,omitempty" json:"title,omitempty"`
}

type Prop struct {
    // Identifying the purpose and intended use of the property, part or other object.
    Name string `xml:"name,attr,omitempty" json:"name,omitempty"`

    // Unique identifier of the containing object
    Id string `xml:"id,attr,omitempty" json:"id,omitempty"`

    // A namespace qualifying the name.
    Ns string `xml:"ns,attr,omitempty" json:"ns,omitempty"`

    // Indicating the type or classification of the containing object
    Class string `xml:"class,attr,omitempty" json:"class,omitempty"`
    Value string `xml:",chardata" json:"value,omitempty"`
}

// A reference to a local or remote resource
type Link struct {
    // A link to a document or document fragment (actual, nominal or projected)
    Href string `xml:"href,attr,omitempty" json:"href,omitempty"`

    // Describes the type of relationship provided by the link. This can be an indicator of the link's purpose.
    Rel string `xml:"rel,attr,omitempty" json:"rel,omitempty"`

    // Describes the media type of the linked resource
    MediaType string `xml:"media-type,attr,omitempty" json:"mediaType,omitempty"`
    Value     string `xml:",chardata" json:"value,omitempty"`
}

// Markup ...
type Markup struct {
    Raw string `xml:",innerxml" json:"raw,omitempty" yaml:"raw,omitempty"`
}

func MarkupFromPlain(plain string) *Markup {
    plain = strings.ReplaceAll(plain, "&", "&amp;")
    plain = strings.ReplaceAll(plain, "<", "&lt;")
    plain = strings.ReplaceAll(plain, "<", "&gt;")
    return &Markup{
        Raw: "<p>" + plain + "</p>",
    }
}

// A name/value pair with optional explanatory remarks.
type Annotation struct {

    // Identifying the purpose and intended use of the property, part or other object.
    Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
    // Unique identifier of the containing object
    Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
    // A namespace qualifying the name.
    Ns string `xml:"ns,attr,omitempty" json:"ns,omitempty"`
    // Indicates the value of the characteristic.
    Value string `xml:"value,attr,omitempty" json:"value,omitempty"`

    // Additional commentary on the parent item.
    Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
}

// A reference to a set of organizations or persons that have responsibility for performing a referenced role relative to the parent context.
type ResponsibleParty struct {

    // The role that the party is responsible for.
    RoleId string `xml:"role-id,attr,omitempty" json:"roleId,omitempty"`

    // References a  defined in .
    PartyIds []string `xml:"party-id,omitempty" json:"party-ids,omitempty"`
}

// A responsible entity, either singular (an organization or person) or collective (multiple persons)
type Party struct {
    Org Org    `xml:"org,omitempty" json:"org,omitempty"`
    Id  string `xml:"id,attr,omitempty" json:"id,omitempty"`
}

// A person, with contact information
type Person struct {

    // Full (legal) name of an individual
    PersonName PersonName `xml:"person-name,omitempty" json:"personName,omitempty"`
    // A common name, short name or acronym
    ShortName ShortName `xml:"short-name,omitempty" json:"shortName,omitempty"`
    // Affiliated organization
    OrgName OrgName `xml:"org-name,omitempty" json:"orgName,omitempty"`
    // An identifier for a person (such as an ORCID) using a designated scheme.
    PersonIds []PersonId `xml:"person-id,omitempty" json:"person-ids,omitempty"`
    // An identifier for an organization using a designated scheme.
    OrganizationIds []OrgId `xml:"org-id,omitempty" json:"organization-ids,omitempty"`
    // References a  defined in .
    LocationIds []LocationId `xml:"location-id,omitempty" json:"location-ids,omitempty"`
    // Email address
    EmailAddresses []Email `xml:"email,omitempty" json:"email-addresses,omitempty"`
    // Contact number by telephone
    TelephoneNumbers []Phone `xml:"phone,omitempty" json:"telephone-numbers,omitempty"`
    // URL for web site or Internet presence
    URLs []Url `xml:"url,omitempty" json:"URLs,omitempty"`
    // A value with a name, attributed to the containing control, part, or group.
    Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
    // A reference to a local or remote resource
    Links []Link `xml:"link,omitempty" json:"links,omitempty"`
    // Additional commentary on the parent item.
    Remarks *Remarks `xml:"remarks,omitempty" json:"remarks,omitempty"`
    // A postal address.
    Addresses []Address `xml:"address,omitempty" json:"addresses,omitempty"`
    // A name/value pair with optional explanatory remarks.
    Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
}

type Remarks = Markup

type PartyId string

// Full (legal) name of an individual

type PersonName string

// Full (legal) name of an organization

type OrgName string

// A common name, short name or acronym

type ShortName string

// A single line of an address.

type AddrLine string

// City, town or geographical region for mailing address

type City string

// State, province or analogous geographical region for mailing address

type State string

// Postal or ZIP code for mailing address

type PostalCode string

// Country for mailing address

type Country string

// Email address

type Email string

// Contact number by telephone
type Phone struct {
    // Indicates the type of phone number.
    Type  string `xml:"type,attr,omitempty" json:"type,omitempty"`
    Value string `xml:",chardata" json:"value,omitempty"`
}

// An identifier for a person (such as an ORCID) using a designated scheme.
type PersonId struct {
    // Indicating the type of identifier, address, email or other data item.
    Type  string `xml:"type,attr,omitempty" json:"type,omitempty"`
    Value string `xml:",chardata" json:"value,omitempty"`
}

// An identifier for an organization using a designated scheme.
type OrgId struct {
    // Indicating the type of identifier, address, email or other data item.
    Type  string `xml:"type,attr,omitempty" json:"type,omitempty"`
    Value string `xml:",chardata" json:"value,omitempty"`
}

// References a  defined in .

type LocationId string

// URL for web site or Internet presence

type Url string

// An organization or legal entity (not a person), with contact information
type Org struct {
    // Full (legal) name of an organization
    OrgName string `xml:"org-name,omitempty" json:"orgName,omitempty"`
    // A postal address.
    Addresses []Address `xml:"address,omitempty" json:"addresses,omitempty"`
}

// A postal address.
type Address struct {

    // Indicates the type of address.
    Type string `xml:"type,attr,omitempty" json:"type,omitempty"`

    // A single line of an address.
    PostalAddress []string `xml:"addr-line,omitempty" json:"postal-address,omitempty"`
    // City, town or geographical region for mailing address
    City string `xml:"city,omitempty" json:"city,omitempty"`
    // State, province or analogous geographical region for mailing address
    State string `xml:"state,omitempty" json:"state,omitempty"`
    // Postal or ZIP code for mailing address
    PostalCode string `xml:"postal-code,omitempty" json:"postalCode,omitempty"`
    // Country for mailing address
    Country string `xml:"country,omitempty" json:"country,omitempty"`
}

// An Import element designates a catalog, profile, or other resource to be
//          included (referenced and potentially modified) by this profile.
type Import struct {

    // A link to a document or document fragment (actual, nominal or projected)
    Href string `xml:"href,attr,omitempty" json:"href,omitempty"`

    // Specifies which controls to include from the resource (source catalog) being
    //           imported
    Include []string `xml:"include,omitempty" json:"include,omitempty"`
}

// A structured information object representing a security or privacy control. Each security or privacy control within the Catalog is defined by a distinct control instance.
type Control struct {

    // Unique identifier of the containing object
    Id string `xml:"id,attr,omitempty" json:"id,omitempty"`
    // Indicating the type or classification of the containing object
    Class string `xml:"class,attr,omitempty" json:"class,omitempty"`

    // A title for display and navigation
    Title Title `xml:"title,omitempty" json:"title,omitempty"`
    // A value with a name, attributed to the containing control, part, or group.
    Properties []Prop `xml:"prop,omitempty" json:"properties,omitempty"`
    // A reference to a local or remote resource
    Links []Link `xml:"link,omitempty" json:"links,omitempty"`
    // Parameters provide a mechanism for the dynamic assignment of value(s) in a control.
    Parameters []Param `xml:"param,omitempty" json:"parameters,omitempty"`
    // A name/value pair with optional explanatory remarks.
    Annotations []Annotation `xml:"annotation,omitempty" json:"annotations,omitempty"`
    // A partition or component of a control or part
    Parts []Part `xml:"part,omitempty" json:"parts,omitempty"`
    // A structured information object representing a security or privacy control. Each security or privacy control within the Catalog is defined by a distinct control instance.
    Controls []Control `xml:"control,omitempty" json:"controls,omitempty"`
}

type Param = nominal_catalog.Param

type Part = nominal_catalog.Part

type Title string

// A collection of citations and resource references.
type BackMatter struct {

    // A resource associated with the present document, which may be a pointer to other data or a citation.
    Resources []Resource `xml:"resource,omitempty" json:"resources,omitempty"`
}

// A resource associated with the present document, which may be a pointer to other data or a citation.
type Resource struct {

    // Unique identifier of the containing object
    Id string `xml:"id,attr,omitempty" json:"id,omitempty"`

    // A short textual description
    Desc string `xml:"desc,omitempty" json:"desc,omitempty"`

    // A pointer to an external copy of a document with optional hash for verification
    Rlinks []Rlink `xml:"rlink,omitempty" json:"rlinks,omitempty"`
}

type Desc validation_root.Desc

// A pointer to an external copy of a document with optional hash for verification
type Rlink struct {

    // A link to a document or document fragment (actual, nominal or projected)
    Href string `xml:"href,attr,omitempty" json:"href,omitempty"`
    // Describes the media type of the linked resource
    MediaType string `xml:"media-type,attr,omitempty" json:"mediaType,omitempty"`
}


