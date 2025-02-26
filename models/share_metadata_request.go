package models

// ShareMetadataRequest struct for ShareMetadataRequest
type ShareMetadataRequest struct {
	// Type of metadata. Required if identifier in metadata_identifies is a name. 1. Liveboard 2. Answers 3. LOGICAL_TABLE for any data object such as table, worksheet or view. 4. LOGICAL_COLUMN
	MetadataType string `json:"metadata_type,omitempty"`
	// Unique ID or name of metadata objects. Note: All the names should belong to same metadata_type
	MetadataIdentifiers []string `json:"metadata_identifiers,omitempty"`
	// Metadata details for sharing objects.
	Metadata []ShareMetadataTypeInput `json:"metadata,omitempty"`
	// Permission details for sharing the objects.
	Permissions []SharePermissionsInput `json:"permissions"`
	// Options to specify details of Liveboard. First Liveboard encountered in payload is considered to be the corresponding Liveboard.
	VisualizationIdentifiers []string `json:"visualization_identifiers,omitempty"`
	// Email IDs to which notifications will be sent.
	Emails []string `json:"emails,omitempty"`
	// Message to be included in notification.
	Message string `json:"message"`
	// Sends object URLs in the customized format in email notifications.
	EnableCustomUrl bool `json:"enable_custom_url,omitempty"`
	// Flag to notify user when any object is shared.
	NotifyOnShare bool `json:"notify_on_share,omitempty"`
	// Flag to make the object discoverable.
	HasLenientDiscoverability bool `json:"has_lenient_discoverability,omitempty"`
}
