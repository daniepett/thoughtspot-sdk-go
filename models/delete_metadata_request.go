package models

// DeleteMetadataRequest struct for DeleteMetadataRequest
type DeleteMetadataRequest struct {
	// Metadata objects.
	Metadata []DeleteMetadataTypeInput `json:"metadata"`
	// Indicates whether to delete disabled metadata objects.
	DeleteDisabledObjects bool `json:"delete_disabled_objects,omitempty"`
}
