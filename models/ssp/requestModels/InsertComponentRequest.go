package requestmodels

// An InsertComponentRequest contains basic information about a components and references to responsible roles
type InsertComponentRequest struct {
	ResponsibleRoles []string

	// todo: other fields
}
