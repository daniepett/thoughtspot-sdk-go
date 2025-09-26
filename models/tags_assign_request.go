package models

type TagsAssignRequest struct {
	Metadata       []TagMetadataTypeInput `json:"metadata"`
	TagIdentifiers []string               `json:"tag_identifiers"`
}
