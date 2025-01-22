package models

// Table struct for Table
type Table struct {
	// Name of the table.
	Name string `json:"name"`
	// Columns of the table.
	Columns []Column `json:"columns,omitempty"`
	// Type of table. Either view or table
	Type string `json:"type,omitempty"`
	// Description of the table
	Description string `json:"description,omitempty"`
	// Determines if the table is selected
	Selected bool `json:"selected,omitempty"`
	// Determines if the table is linked
	Linked bool `json:"linked,omitempty"`
	// List of relationships for the table
	Relationships []map[string]interface{} `json:"relationships,omitempty"`
}
