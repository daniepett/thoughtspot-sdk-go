package models

type DataWarehouseConfig struct {
	Configuration     map[string]interface{}                `json:"configuration"`
	ExternalDatabases []DataWarehouseConfigExternalDatabase `json:"externalDatabases,omitempty"`
}

type DataWarehouseConfigConfiguration struct {
	AccountName       string `json:"account_name,omitempty"`
	User              string `json:"user,omitempty"`
	Password          string `json:"password,omitempty"`
	PrivateKey        string `json:"private_key,omitempty"`
	Role              string `json:"role,omitempty"`
	Warehouse         string `json:"warehouse,omitempty"`
	Database          string `json:"database,omitempty"`
	OauthClientId     string `json:"oauth_client_id,omitempty"`
	OauthClientSecret string `json:"oauth_client_secret,omitempty"`
	Scope             string `json:"scope,omitempty"`
	AuthUrl           string `json:"auth_url,omitempty"`
	AccessTokenUrl    string `json:"access_token_url,omitempty"`
}

type DataWarehouseConfigExternalDatabase struct {
	Name string `json:"name"`
}
