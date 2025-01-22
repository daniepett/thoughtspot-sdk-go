package models


// Database struct for Database
type Database struct {
	// Name of the database.
	Name string `json:"name"`
	// Schemas of the database.
	Schemas []SchemaObject `json:"schemas,omitempty"`
	// Determines if the object is auto created.
	AutoCreated bool `json:"auto_created,omitempty"`
}