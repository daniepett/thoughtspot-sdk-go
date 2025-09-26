package models

type TagsUpdateRequest struct {
	Name  string `json:"name"`
	Color string `json:"color,omitempty"`
}
