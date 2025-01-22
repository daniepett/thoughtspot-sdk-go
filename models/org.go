package models

// Org The current Org context of the user.
type Org struct {
	// The ID of the object.
	Id int32 `json:"id"`
	// Name of the object.
	Name string `json:"name"`
}