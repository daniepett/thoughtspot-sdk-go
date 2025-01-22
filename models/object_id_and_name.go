package models

// ObjectIDAndName The object representation with ID and Name.
type ObjectIDAndName struct {
	// The unique identifier of the object.
	Id string `json:"id,omitempty"`
	// Name of the object.
	Name string `json:"name,omitempty"`
}
