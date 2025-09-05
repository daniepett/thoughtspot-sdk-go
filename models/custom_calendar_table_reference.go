package models

type CustomCalendarTableReference struct {
	ConnectionIdentifier string `json:"connection_identifier"`
	DatabaseName         string `json:"database_name"`
	SchemaName           string `json:"schema_name"`
	TableName            string `json:"table_name"`
}
