package models

type CalendarResponse struct {
	CalendarName             string `json:"calendar_name,omitempty"`
	ConnectionName           string `json:"connection_name,omitempty"`
	DataWarehouseType        string `json:"data_warehouse_type,omitempty"`
	ModificationTimeInMillis string `json:"modification_time_in_millis,omitempty"`
	AuthorName               string `json:"author_name,omitempty"`
	ConnectionId             string `json:"connection_id,omitempty"`
	CalendarId               string `json:"calendar_id,omitempty"`
}
