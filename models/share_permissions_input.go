package models

// SharePermissionsInput struct for SharePermissionsInput
type SharePermissionsInput struct {
	Principal PrincipalsInput `json:"principal"`
	// Type of access to the shared object
	ShareMode string `json:"share_mode"`
}
