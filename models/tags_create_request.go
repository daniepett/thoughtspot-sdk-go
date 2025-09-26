package models

type TagsCreateRequest struct {
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}
