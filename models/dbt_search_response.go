package models

// DbtSearchResponse struct for DbtSearchResponse
type DbtSearchResponse struct {
	DbtConnectionIdentifier string `json:"dbt_connection_identifier,omitempty"`
	ProjectName             string `json:"project_name,omitempty"`
	ConnectionId            string `json:"connection_id,omitempty"`
	ConnectionName          string `json:"connection_name,omitempty"`
	CdwDatabase             string `json:"cdw_database,omitempty"`
	ImportType              string `json:"import_type,omitempty"`
	AuthorName              string `json:"author_name,omitempty"`
}
