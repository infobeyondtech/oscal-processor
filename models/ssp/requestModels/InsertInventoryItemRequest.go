package requestmodels

// An InsertInventoryItem contains references to components and responsible parties
type InsertInventoryItem struct {
	Description string `json:"description,omitempty"`

	AssetID string `json:"assetID,omitempty"`

	ResponsibleParties []string `json:"responsibleParties,omitempty"`

	ImplementComponents []string `json:"ImplementComponents,omitempty"`
}
