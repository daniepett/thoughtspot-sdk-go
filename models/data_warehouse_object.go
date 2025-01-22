package models


// DataWarehouseObjects struct for DataWarehouseObjects
type DataWarehouseObjects struct {
	// Databases of the connection.
	Databases []Database `json:"databases"`
}