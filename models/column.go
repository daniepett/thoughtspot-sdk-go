package models


// Column struct for Column
type Column struct {
	// Name of the column
	Name string `json:"name"`
	// Data type of the column
	DataType string `json:"data_type"`
	// Determines if the column schema is an aggregate
	IsAggregate string `json:"is_aggregate,omitempty"`
	// Determines if the column schema can be imported
	CanImport bool `json:"can_import,omitempty"`
	// Determines if the table is selected
	Selected bool `json:"selected,omitempty"`
	// Determines if the table is linked
	IsLinkedActive bool `json:"is_linked_active,omitempty"`
}
