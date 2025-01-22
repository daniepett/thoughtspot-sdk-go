package models


// SortOptionInput struct for SortOptionInput
type SortOptionInput struct {
	// Name of the field to apply the sort on.
	FieldName string `json:"field_name,omitempty"`
	// Sort order : ASC(Ascending) or DESC(Descending).
	Order string `json:"order,omitempty"`
}
