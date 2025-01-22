package models


// SchemaObject struct for SchemaObject
type SchemaObject struct {
	// Name of the schema.
	Name string `json:"name"`
	// Tables in the schema.
	Tables []Table `json:"tables,omitempty"`
}
