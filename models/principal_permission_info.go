package models

type PrincipalPermissionInfo struct {
	PrincipalType        string                `json:"principal_type"`
	PrincipalSubType     string                `json:"principal_sub_type"`
	PrincipalPermissions []PrincipalPermission `json:"principal_permissions"`
}
