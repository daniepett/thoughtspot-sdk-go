package models

type SearchCustomCalendarsRequest struct {
	ConnectionIdentifier string      `json:"connection_identifier,omitempty"`
	NamePattern          string      `json:"name_pattern,omitempty"`
	RecordOffset         string      `json:"record_offset,omitempty"`
	RecordSize           string      `json:"record_size,omitempty"`
	SortOptions          SortOptions `json:"sort_options,omitempty"`
}
