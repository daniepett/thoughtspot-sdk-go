package models

// PermissionsMetadataTypeInput MetadataType InputType used in Permission API's
type PermissionsMetadataTypeInput struct {
	//   Type of metadata object.     Required if the name of the object is set as the identifier. This attribute is optional when the object GUID is specified as the identifier.
	Type string `json:"type,omitempty"`
	// Unique ID or name of the metadata object.
	Identifier string `json:"identifier"`
}
