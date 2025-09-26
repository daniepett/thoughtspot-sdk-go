package models

type TagsUnassignRequest struct {
	Metadata       []TagMetadataTypeInput `json:"metadata"`
	TagIdentifiers []string               `json:"tag_identifiers"`
}
