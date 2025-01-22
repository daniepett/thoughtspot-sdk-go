package models

// UpdateUserGroupRequest struct for UpdateUserGroupRequest
type UpdateUserGroupRequest struct {
	// Name of the group to modify.
	Name string `json:"name,omitempty"`
	// ID of the Liveboards to be assigned as default Liveboards to the users in the group.
	DefaultLiveboardIdentifiers []string `json:"default_liveboard_identifiers,omitempty"`
	// Description for the group.
	Description string `json:"description,omitempty"`
	// Display name of the group.
	DisplayName string `json:"display_name,omitempty"`
	// Privileges to assign to the group.
	Privileges []string `json:"privileges,omitempty"`
	// GUID or name of the sub groups. A subgroup is a group assigned to a parent group.
	SubGroupIdentifiers []string `json:"sub_group_identifiers,omitempty"`
	// Type of the group
	Type string `json:"type,omitempty"`
	// GUID or name of the users to assign to the group.
	UserIdentifiers []string `json:"user_identifiers"`
	// Visibility of the group. To make a group visible to other users and groups, set the visibility to SHAREABLE.
	Visibility string `json:"visibility,omitempty"`
	// Role identifiers of the Roles that should be assigned to the group.
	RoleIdentifiers []string `json:"role_identifiers,omitempty"`
	// Type of update operation. Default operation type is REPLACE
	Operation string `json:"operation,omitempty"`
}
