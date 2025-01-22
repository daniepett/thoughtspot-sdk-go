package models

type UpdateConnectionRequest struct {
	// Unique name for the connection.
	Name string `json:"name"`
	// Description of the connection.
	Description string `json:"description,omitempty"`
	// Type of the data warehouse.
	DataWarehouseConfig map[string]interface{} `json:"data_warehouse_config"`
	// Validates the connection metadata if tables are included. If you are creating a connection without tables, specify `false`.
	Validate bool `json:"validate"`
}
