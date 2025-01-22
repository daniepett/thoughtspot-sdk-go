package models


// SearchConnectionRequest struct for SearchConnectionRequest
type SearchConnectionRequest struct {
	// List of connections and name pattern
	Connections []ConnectionInput `json:"connections,omitempty"`
	// Array of types of data warehouse defined for the connection.
	DataWarehouseTypes []string `json:"data_warehouse_types,omitempty"`
	// The starting record number from where the records should be included.
	RecordOffset *int32 `json:"record_offset,omitempty"`
	// The number of records that should be included.
	RecordSize *int32 `json:"record_size,omitempty"`
	// Unique ID or name of tags.
	TagIdentifiers []string `json:"tag_identifiers,omitempty"`
	// Data warehouse object type.
	DataWarehouseObjectType *string `json:"data_warehouse_object_type,omitempty"`
	// Sort options.
	SortOptions *SortOptionInput `json:"sort_options,omitempty"`
	// Indicates whether to include complete details of the connection objects.
	IncludeDetails bool `json:"include_details,omitempty"`
	// Configuration values. If empty we are fetching configuration from datasource based on given connection id. If required you can provide config details to fetch specific details. Example input: {}, {\"warehouse\":\"SMALL_WH\",\"database\":\"DEVELOPMENT\"}. This is only applicable when data_warehouse_object_type is selected.
	Configuration map[string]interface{} `json:"configuration,omitempty"`
	// List of authentication types to fetch data_ware_house_objects from external Data warehouse. This is only applicable when data_warehouse_object_type is selected.
	AuthenticationType *string `json:"authentication_type,omitempty"`
}