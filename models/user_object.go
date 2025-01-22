package models

// UserObject Objects to apply the User_Object.
type UserObject struct {
	//   Type of object.     Required if the name of the object is set as the identifier. This attribute is optional when the object GUID is specified as the identifier.
	Type string `json:"type,omitempty"`
	// Unique name/id of the object.
	Identifier string `json:"identifier"`
}