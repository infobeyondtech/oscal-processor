package request_models

type CreatProfileRequest struct {
	Baseline string   `json:"baseline" binding:"required"`
	Controls []string `json:"controls" binding:"required"`
	Catalogs []string `json:"catalogs" binding:"required"`

	Title    string `json:"title" binding:"required"`
	OrgUuid  string `json:"orgUuid" binding:"required"`
	OrgName  string `json:"orgName" binding:"required"`
	OrgEmail string `json:"orgEmail" binding:"required"`
}

type AddAddressRequest struct {
	UUID         string `json:"uuid" binding:"required"`
	Addresses  []string `json:"addresses" binding:"required"`
	City       string   `json:"city" binding:"required"`
	State      string   `json:"state" binding:"required"`
	PostalCode string   `json:"postalCode" binding:"required"`
}

type SetTitleVersionRequest struct {
	UUID         string `json:"uuid" binding:"required"`
	Title        string `json:"title" binding:"required"`
	Version      string `json:"version" binding:"required"`
	OscalVersion string `json:"oscalversion" binding:"required"`
}

type AddRolePartyRequest struct {
	UUID    string `json:"uuid" binding:"required"`
	RoleID  string `json:"roleID" binding:"required"`
	Title   string `json:"title" binding:"required"`
	PartyID string `json:"partyId" binding:"required"`
	OrgName string `json:"orgName" binding:"required"`
	Email   string `json:"email" binding:"required"`
}



// below are requests related to ssp
type AddSystemCharacteristicReuqest struct{
	UUID string `json:"uuid" binding:"required"`
	SystemName string `json:"systemName" binding:"required"`
	Description string `json:"description" binding:"required"`
	DeploymentModel string `json:"deploymentModel" binding:"required"`
	SecurityLevel string `json:"securityLevel" binding:"required"`
}

type InsertInventoryItemRequest struct {
	UUID         string `json:"uuid" binding:"required"`
	InventoryItemID string `json:"inventoryItemID" binding:"required"`
	ImplementComponents []string `json:"implementComponents" binding:"required"`
	ResponsibleParties []RolePartyMap `json:"responsibleParties" binding:"required"`
}

type InsertImplementedRequirementRequest struct {
	UUID         string `json:"uuid" binding:"required"`
	ControlID string `json:"controlID" binding:"required"`
	Statements []Statement `json:"statements" binding:"required"`
}

type SetParameter struct {
	ParamID string `json:"paramID" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type ByComponent struct {
	ComponentID string `json:"componentID" binding:"required"`
	Description string `json:"description" binding:"required"`
	SetParameters []SetParameter `json:"setParameters" binding:"required"`
}

type Statement struct {
	StatementID string `json:"statementID" binding:"required"`
	ByComponents []ByComponent `json:"bycomponents" binding:"required"`
}

type RolePartyMap struct {
	UserUUID string `json:"UserUUID" binding:"required"`
	PartyUUIDs []string `json:"PartyUUIDs" binding:"required"`
}