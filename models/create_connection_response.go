package models


// CreateConnectionResponse struct for CreateConnectionResponse
type CreateConnectionResponse struct {
	// ID of the connection created.
	Id string `json:"id"`
	// Name of the connection.
	Name string `json:"name"`
	// Type of data warehouse.
	DataWarehouseType string `json:"data_warehouse_type"`
	// Details of the connection.
	Details map[string]interface{} `json:"details,omitempty"`
}