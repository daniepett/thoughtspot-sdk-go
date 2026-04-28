package models

type PermissionsInput struct {
	Principal PrincipalsInput `json:"principal"`
	ShareMode string          `json:"share_mode"`
}
