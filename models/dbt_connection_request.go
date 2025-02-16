package models

// DbtConnectionRequest struct for DbtConnectionRequest
type DbtConnectionRequest struct {
	// Name of the connection.
	ConnectionName string `json:"connection_name"`
	// Name of the Database.
	DatabaseName string `json:"database_name"`
	// Mention type of Import
	ImportType string `json:"import_type,omitempty"`
	// Access token is mandatory when Import_Type is DBT_CLOUD.
	AccessToken string `json:"access_token,omitempty"`
	// DBT URL is mandatory when Import_Type is DBT_CLOUD.
	DbtUrl string `json:"dbt_url,omitempty"`
	// Account ID is mandatory when Import_Type is DBT_CLOUD
	AccountId string `json:"account_id,omitempty"`
	// Project ID is mandatory when Import_Type is DBT_CLOUD
	ProjectId string `json:"project_id,omitempty"`
	// DBT Environment ID\"
	DbtEnvId string `json:"dbt_env_id,omitempty"`
	// Name of the project
	ProjectName string `json:"project_name,omitempty"`
	// Upload DBT Manifest and Catalog artifact files as a ZIP file. This field is Mandatory when Import Type is 'ZIP_FILE'
	FileContent string `json:"file_content,omitempty"`
}
