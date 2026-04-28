package models

type MetadataListItemInput struct {
	Identifier    string   `json:"identifier,omitempty"`
	ObjIdentifier string   `json:"obj_identifier,omitempty"`
	NamePattern   string   `json:"name_pattern,omitempty"`
	Type          string   `json:"type,omitempty"`
	Subtypes      []string `json:"subtypes,omitempty"`
}
