package models

// PrincipalsInput struct for PrincipalsInput
type PrincipalsInput struct {
	// Unique ID or name of the principal object such as a user or group.
	Identifier string `json:"identifier"`
	// Principal type.
	Type string `json:"type,omitempty"`
}
