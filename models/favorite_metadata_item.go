package models

// FavoriteMetadataItem struct for FavoriteMetadataItem
type FavoriteMetadataItem struct {
	// Unique ID of the metadata object.
	Id string `json:"id"`
	// name of the metadata object.
	Name string `json:"name"`
	//   Type of metadata object.     Required if the name of the object is set as the identifier. This attribute is optional when the object GUID is specified as the identifier.
	Type string `json:"type"`
}