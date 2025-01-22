package models


// ConnectionInput struct for ConnectionInput
type ConnectionInput struct {
	// Unique ID or name of the connection.
	Identifier string `json:"identifier,omitempty"`
	// A pattern to match case-insensitive name of the connection object. User `%` for a wildcard match.
	NamePattern string `json:"name_pattern,omitempty"`
	// Filter options for databases, schemas, tables and columns.
	DataWarehouseObjects []DataWarehouseObjectInput `json:"data_warehouse_objects,omitempty"`
}