package models

// DeleteMetadataTypeInput MetadataType InputType used in Delete MetadataType API
type DeleteMetadataTypeInput struct {
	//   Type of metadata.     Required if the name of the object is set as the identifier. This attribute is optional when the object GUID is specified as the identifier.
	Type string `json:"type,omitempty"`
	// Unique ID or name of the metadata object.
	Identifier string `json:"identifier"`
}
