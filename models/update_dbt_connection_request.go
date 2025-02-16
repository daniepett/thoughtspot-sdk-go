package models

import (
	"os"
)

// UpdateDbtConnectionRequest struct for UpdateDbtConnectionRequest
type UpdateDbtConnectionRequest struct {
	// Unique ID of the DBT Connection.
	DbtConnectionIdentifier string `json:"dbt_connection_identifier"`
	// Name of the connection.
	ConnectionName string `json:"connection_name,omitempty"`
	// Name of the Database.
	DatabaseName string `json:"database_name,omitempty"`
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
	FileContent **os.File `json:"file_content,omitempty"`
}
