package models

type MetadataSearchSortOptions struct {
	FieldName string `json:"field_name,omitempty"`
	Order     string `json:"order,omitempty"`
}
