package models


// DataWarehouseObjectInput struct for DataWarehouseObjectInput
type DataWarehouseObjectInput struct {
	// Name of the database.
	Database string `json:"database,omitempty"`
	// Name of the schema within the database.
	Schema string `json:"schema,omitempty"`
	// Name of the table within the schema.
	Table string `json:"table,omitempty"`
	// Name of the column within the table.
	Column string `json:"column,omitempty"`
}
