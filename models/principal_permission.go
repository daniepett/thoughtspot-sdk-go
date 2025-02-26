package models

type PrincipalPermission struct {
	PrincipalId      string   `json:"principal_id"`
	PrincipalName    string   `json:"principal_name"`
	Permission       string   `json:"permission"`
	SharedPermission string   `json:"shared_permission"`
	GroupPermission  []string `json:"group_permission"`
}
