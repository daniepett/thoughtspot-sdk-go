package models

// FetchPermissionsOnMetadataRequest struct for FetchPermissionsOnMetadataRequest
type FetchPermissionsOnMetadataRequest struct {
	// GUID or name of the metadata object.
	Metadata []PermissionsMetadataTypeInput `json:"metadata"`
	// User or group objects for which you want to fetch permissions. If not specified, the API returns all users and groups that can access the specified metadata objects.
	Principals []PrincipalsInput `json:"principals,omitempty"`
	// Indicates whether to fetch permissions of dependent metadata objects.
	IncludeDependentObjects bool `json:"include_dependent_objects,omitempty"`
	// The starting record number from where the records should be included for each metadata type.
	RecordOffset int32 `json:"record_offset,omitempty"`
	// The number of records that should be included for each metadata type.
	RecordSize int32 `json:"record_size,omitempty"`
	// <div><span class=\"since-beta-tag\">Version: 10.3.0.cl or later</span></div>  Specifies the type of permission. Valid values are:     EFFECTIVE - If the user permission to the metadata objects is granted by the privileges assigned to the groups to which they belong.     DEFINED - If a user or user group received access to metadata objects via object sharing by another user.
	PermissionType string `json:"permission_type,omitempty"`
}
