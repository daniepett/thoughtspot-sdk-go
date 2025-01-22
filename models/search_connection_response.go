package models


// SearchConnectionResponse struct for SearchConnectionResponse
type SearchConnectionResponse struct {
	// Unique ID of the connection.
	Id string `json:"id"`
	// Name of the connection.
	Name string `json:"name"`
	// Description of the connection.
	Description string `json:"description,omitempty"`
	// Type of data warehouse.
	DataWarehouseType string `json:"data_warehouse_type"`
	DataWarehouseObjects *DataWarehouseObjects `json:"data_warehouse_objects,omitempty"`
	// Details of the connection.
	Details map[string]interface{} `json:"details,omitempty"`
}
