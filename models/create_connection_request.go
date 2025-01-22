package models

// CreateConnectionRequest struct for CreateConnectionRequest
type CreateConnectionRequest struct {
	// Unique name for the connection.
	Name string `json:"name"`
	// Description of the connection.
	Description string `json:"description,omitempty"`
	// Type of the data warehouse.
	DataWarehouseType string `json:"data_warehouse_type"`
	// Connection configuration attributes in JSON format. To create a connection with tables, include table attributes. See the documentation above for sample JSON.
	DataWarehouseConfig map[string]interface{} `json:"data_warehouse_config"`
	// Validates the connection metadata if tables are included. If you are creating a connection without tables, specify `false`.
	Validate bool `json:"validate"`
}
