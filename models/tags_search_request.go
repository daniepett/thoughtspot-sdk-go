package models

type TagsSearchRequest struct {
	TagIdentifier string `json:"tag_identifier,omitempty"`
	NamePattern   string `json:"name_pattern,omitempty"`
	Color         string `json:"color,omitempty"`
}
