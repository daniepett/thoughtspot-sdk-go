package models

// SearchUserGroupsRequest struct for SearchUserGroupsRequest
type SearchUserGroupsRequest struct {
	// GUID of Liveboards that are assigned as default Liveboards to the users in the group.
	DefaultLiveboardIdentifiers []string `json:"default_liveboard_identifiers,omitempty"`
	// Description of the group
	Description string `json:"description,omitempty"`
	// Display name of the group
	DisplayName string `json:"display_name,omitempty"`
	// A pattern to match case-insensitive name of the Group object.
	NamePattern string `json:"name_pattern,omitempty"`
	// GUID or name of the group
	GroupIdentifier string `json:"group_identifier,omitempty"`
	// ID or name of the Org to which the group belongs
	OrgIdentifiers []string `json:"org_identifiers,omitempty"`
	// Privileges assigned to the group.
	Privileges []string `json:"privileges,omitempty"`
	// GUID or name of the sub groups. A subgroup is a group assigned to a parent group.
	SubGroupIdentifiers []string `json:"sub_group_identifiers,omitempty"`
	// Group type.
	Type *string `json:"type,omitempty"`
	// GUID or name of the users assigned to the group.
	UserIdentifiers []string `json:"user_identifiers,omitempty"`
	// Visibility of the group. To make a group visible to other users and groups, set the visibility to SHAREABLE.
	Visibility string `json:"visibility,omitempty"`
	// Filter groups with a list of Roles assigned to a group
	RoleIdentifiers []string `json:"role_identifiers,omitempty"`
	// The starting record number from where the records should be included.
	RecordOffset int32 `json:"record_offset,omitempty"`
	// The number of records that should be included.
	RecordSize int32 `json:"record_size,omitempty"`
	// Sort options to filter group details.
	SortOptions SortOptions `json:"sort_options,omitempty"`
}
