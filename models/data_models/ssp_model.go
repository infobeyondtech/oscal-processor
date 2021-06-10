package data_models

type SystemSecurityPlanModel struct{
	ImportProfile string `json:"importProfile" binding:"required"`
	MetaDataModel MetaData `json:"metaDataModel" binding:"required"`
	SystemCharacteristicModel SystemCharacteristic `json:"SystemCharacteristicModel"`
	SystemImplementationModel SystemImplementation `json:"systemImplementationModel"`
	ControlImplementationModel ControlImplementation `json:"controlImplementationModel"`
}

type SystemCharacteristic struct{
	SystemName string `json:"systemName" binding:"required"`
	Description string `json:"description"`
	SecurityLevel string `json:"securityLevel"`

	SystemInformationTitle string `json:"systemInformationTitle"`
	SystemInformationDescription string `json:"systemInformationDescription"`
	ConfidentialityImpact string `json:"confidentialityImpact"`
	IntegrityImpact string `json:"integrityImpact" `
	AvailabilityImpact string `json:"availabilityImpact"`
}

type MetaData struct{
	Title        string `json:"title"`
	Version      string `json:"version"`
	OscalVersion string `json:"oscalversion"`
	LastModified string `json:"lastModified"`

	Parties []Party `json:"parties"`
}

type SystemImplementation struct{
	Users []User `json:"users"`

	Components []Component `json:"components"`

	InventoryItems []InventoryItem `json:"inventoryItems"`
}

type ControlImplementation struct{
	ImplementedRequirements []ImplementedRequirement `json:"implementedRequirements" binding:"required"`
}

type Party struct{
	Name string `json:"name" binding:"required"`
	Uuid string `json:"uuid" binding:"required"`
	Type string `json:"type"`
}

type User struct{
	Uuid string `json:"uuid" binding:"required"`
	Title string `json:"title"`
	Type string `json:"type"`
	RoleId string  `json:"roleId" `
}

type Component struct {
	Uuid string `json:"uuid" binding:"required"`
	Type string `json:"type"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`
	ResponsibleRoles [] ResponsibleRole `json:"responsibleRoles"`
}

type ResponsibleRole struct{
	RoleId string `json:"roleId" binding:"required"`
	PartyIds []string `json:"partyIds" binding:"required"`
}

type InventoryItem struct{
	Uuid string `json:"uuid" binding:"required"`
	Description string `json:"description"`
	AssetId string `json:"assetId"`
	ImplementComponentIds []string `json:"implementComponentIds" binding:"required"`

	ResponsibleParties []ResponsibleParty `json:"responsibleParties"`
}

type ResponsibleParty struct{
	RoleId string `json:"roleId" binding:"required"`
	PartyUuid string `json:"partyUuid" binding:"required"`
}

type ImplementedRequirement struct{
	ControlId string `json:"controlId" binding:"required"`
	Uuid string `json:"uuid" binding:"required"`

	Statements []StatementModel `json:"statements"`
}

type StatementModel struct {
	StatementId string `json:"statementId" binding:"required"`

	ByComponents []ByComponentModel `json:"byComponents" binding:"required"`
}

type ByComponentModel struct{
	Description string  `json:"description"`
	ComponentUuid string  `json:"componentUuid" binding:"required"`

	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	ParamId string `json:"paramId" binding:"required"`
	Value string `json:"value" binding:"required"`
}


