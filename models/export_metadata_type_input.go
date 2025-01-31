package models

// ExportMetadataTypeInput MetadataType InputType used in Export MetadataType API
type ExportMetadataTypeInput struct {
	//   Type of metadata.     Required if the name of the object is set as the identifier. This attribute is optional when the object GUID is specified as the identifier.
	Type string `json:"type,omitempty"`
	// Unique ID or name of the metadata object.
	Identifier string `json:"identifier"`
}
