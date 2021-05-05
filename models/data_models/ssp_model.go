package data_models

type SystemSecurityPlanModel struct{
	ImportProfile string `json:"importProfile" binding:"required"`
	MetaDataModel MetaData `json:"metaDataModel" binding:"required"`
	SystemCharacteristicModel SystemCharacteristic `json:"SystemCharacteristicModel" binding:"required"`
	SystemImplementationModel SystemImplementation `json:"systemImplementationModel" binding:"required"`
	ControlImplementationModel ControlImplementation `json:"controlImplementationModel" binding:"required"`
}

type SystemCharacteristic struct{
	//UUID string `json:"uuid" binding:"required"`
	SystemName string `json:"systemName" binding:"required"`
	Description string `json:"description" binding:"required"`
	//DeploymentModel string `json:"deploymentModel" binding:"required"`
	SecurityLevel string `json:"securityLevel" binding:"required"`

	SystemInformationTitle string `json:"systemInformationTitle" binding:"required"`
	SystemInformationDescription string `json:"systemInformationDescription" binding:"required"`
	ConfidentialityImpact string `json:"confidentialityImpact" binding:"required"`
	IntegrityImpact string `json:"integrityImpact" binding:"required"`
	AvailabilityImpact string `json:"availabilityImpact" binding:"required"`
}

type MetaData struct{
	Title        string `json:"title" binding:"required"`
	Version      string `json:"version" binding:"required"`
	OscalVersion string `json:"oscalversion" binding:"required"`
	LastModified string `json:"lastModified" binding:"required"`

	Parties [] Party `json:"parties" binding:"required"`
}

type SystemImplementation struct{
	Users []User `json:"users" binding:"required"`

	Components []Component `json:"components" binding:"required"`

	InventoryItems []InventoryItem `json:"inventoryItems" binding:"required"`
}

type ControlImplementation struct{
	ImplementedRequirements []ImplementedRequirement `json:"implementedRequirements" binding:"required"`
}

type Party struct{
	Name string `json:"name" binding:"required"`
	Uuid string `json:"uuid" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type User struct{
	Uuid string `json:"uuid" binding:"required"`
	Title string `json:"title" binding:"required"`
	Type string `json:"type" binding:"required"`
	RoleId string  `json:"roleId" binding:"required"`
}

type Component struct {
	Uuid string `json:"uuid" binding:"required"`
	Type string `json:"type" binding:"required"`
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status string `json:"status" binding:"required"`
	ResponsibleRoles [] ResponsibleRole `json:"responsibleRoles" binding:"required"`
}

type ResponsibleRole struct{
	RoleId string `json:"roleId" binding:"required"`
	PartyUuid string `json:"partyUuid" binding:"required"`
}

type InventoryItem struct{
	Uuid string `json:"uuid" binding:"required"`
	Description string `json:"description" binding:"required"`
	AssetId string `json:"assetId" binding:"required"`

	ResponsibleParties []ResponsibleParty `json:"responsibleParties" binding:"required"`
}

type ResponsibleParty struct{
	RoleId string `json:"roleId" binding:"required"`
	PartyUuid string `json:"partyUuid" binding:"required"`
}

type ImplementedRequirement struct{
	ControlId string `json:"controlId" binding:"required"`
	Uuid string `json:"uuid" binding:"required"`

	Statements []StatementModel `json:"statements" binding:"required"`
}

type StatementModel struct {
	StatementId string 	 `json:"statementId" binding:"required"`

	ByComponents []ByComponentModel `json:"byComponents" binding:"required"`
}

type ByComponentModel struct{
	Description string  `json:"description" binding:"required"`
	ComponentUuid string  `json:"componentUuid" binding:"required"`

	Parameters []Parameter `json:"parameters" binding:"required"`
}

type Parameter struct {
	ParamId string `json:"paramId" binding:"required"`
	Value string `json:"value" binding:"required"`
}


