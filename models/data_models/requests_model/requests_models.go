package data_models

type CreatProfileRequest struct {
    Baseline string   `json:"baseline" binding:"required"`
    Controls []string `json:"controls" binding:"required"`
    Catalogs []string `json:"catalogs" binding:"required"`

    Title    string `json:"title" binding:"required"`
    OrgUuid  string `json:"orgUuid" binding:"required"`
    OrgName  string `json:"orgName" binding:"required"`
    OrgEmail string `json:"orgEmail" binding:"required"`
}

type AddPartyRequest struct {
    OrgName    string   `json:"orgName" binding:"required"`
    UUID       string   `json:"uuid" binding:"required"`
    Addresses  []string `json:"addresses" binding:"required"`
    City       string   `json:"city" binding:"required"`
    State      string   `json:"state" binding:"required"`
    PostalCode string   `json:"postalCode" binding:"required"`
    RoleId     string   `json:"roleId" binding:"required"`
    PartyId    string   `json:"partyId" binding:"required"`
}

type SetTitleVersionRequest struct {
    //UUID         string `json:"uuid" binding:"required"`
    Title        string `json:"title" binding:"required"`
    ProfileId    string `json:"profileId" binding:"required"`
    Version      string `json:"version"`
    OscalVersion string `json:"oscalVersion"`
}

type AddRolePartyRequest struct {
    UUID    string `json:"uuid" binding:"required"`
    RoleID  string `json:"roleID" binding:"required"`
    Title   string `json:"title" binding:"required"`
    PartyID string `json:"partyId" binding:"required"`
    OrgName string `json:"orgName" binding:"required"`
    Email   string `json:"email" binding:"required"`
}

type AddControlRequest struct {
    UUID       string   `json:"uuid" binding:"required"`
    ControlIDs []string `json:"controlIDs" binding:"required"`
}

// below are requests related to ssp
type AddSystemCharacteristicReuqest struct {
    FileID          string `json:"fileID" binding:"required"`
    UUID            string `json:"uuid" binding:"required"`
    SystemName      string `json:"systemName"`
    Description     string `json:"description"`
    DeploymentModel string `json:"deploymentModel"`
    SecurityLevel   string `json:"securityLevel"`

    SystemInformationTitle       string `json:"systemInformationTitle"`
    SystemInformationDescription string `json:"systemInformationDescription"`
    ConfidentialityImpact        string `json:"confidentialityImpact"`
    IntegrityImpact              string `json:"integrityImpact"`
    AvailabilityImpact           string `json:"availabilityImpact"`
}

type InsertInventoryItemRequest struct {
    FileID              string         `json:"fileID" binding:"required"`
    InventoryItemID     string         `json:"inventoryItemID" binding:"required"`
    ImplementComponents []string       `json:"implementComponents" binding:"required"`
    ResponsibleParties  []RolePartyMap `json:"responsibleParties"`
}

type InsertImplementedRequirementRequest struct {
    FileID     string      `json:"fileID" binding:"required"`
    UUID       string      `json:"uuid" binding:"required"`
    ControlID  string      `json:"controlID" binding:"required"`
    Statements []Statement `json:"statements"`
}

type SetParameter struct {
    ParamID string `json:"paramID" binding:"required"`
    Value   string `json:"value" binding:"required"`
}

type ByComponent struct {
    ComponentID        string         `json:"componentID" binding:"required"`
    Description        string         `json:"description"`
    SetParameters      []SetParameter `json:"setParameters"`
    ResponsibleParties []RolePartyMap `json:"responsibleParties"`
}

type Statement struct {
    StatementID  string        `json:"statementID" binding:"required"`
    ByComponents []ByComponent `json:"bycomponents" binding:"required"`
}

type RolePartyMap struct {
    UserUUID string `json:"userUUID" binding:"required"`
    PartyUUIDs []string `json:"partyUUIDs" binding:"required"`
} 


type RemoveElementRequest struct {
    FileID    string `json:"fileID" binding:"required"`
    ElementID string `json:"elementID" binding:"required"`
}

type CompareDiffRequest struct {
    UUID string `xml:"uuid,attr,omitempty" json:"uuid,omitempty"`
    Baseline string `xml:"baseline,attr,omitempty" json:"baseline,omitempty"`
}
