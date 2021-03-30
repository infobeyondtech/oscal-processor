package requestmodels

// An InsertImplementedRequirementRequest contains all information to insert a record in to ssp
type InsertImplementedRequirementRequest struct {
	ControlID string `json:"controlID,omitempty"`

	Statements []Statement `json:"statements,omitempty"`
}

// A Statement is how a control and its sub-parts are implemented
type Statement struct {
	StatementID string `json:"statementID,omitempty"`

	ByComponents []ByComponent `json:"bycomponents,omitempty"`
}

// A ByComponent includes one component and a few parameters
type ByComponent struct {
	ComponentID string `json:"componentID,omitempty"`

	Description string `json:"description,omitempty"`

	SetParameters []SetParameter `json:"setParameters,omitempty"`
}

// SetParameter describes a pair of key and values
type SetParameter struct {
	ParamID string `json:"paramID,omitempty"`

	Value string `json:"value,omitempty"`
}

// ResponsibleRole refers to a party
type ResponsibleRole struct {
	RoleID string `json:"RoleID,omitempty"`

	Party *Party `json:"party,omitempty"`
}

type ResponsibleParty struct {
	Role *Role

	Party *Party
}

type Role struct {
	Title string

	Annotation string

	RoleId string
}

type Party struct {
	Type string

	Name string
}
