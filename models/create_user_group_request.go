package models

// CreateUserGroupRequest struct for CreateUserGroupRequest
type CreateUserGroupRequest struct {
	// Name of the group. The group name must be unique.
	Name string `json:"name"`
	// Display name for the group.
	DisplayName string `json:"display_name"`
	// GUID of the Liveboards to assign as default Liveboards to the users in the group.
	DefaultLiveboardIdentifiers []string `json:"default_liveboard_identifiers,omitempty"`
	// Description of the group
	Description string `json:"description,omitempty"`
	// Privileges to assign to the group
	Privileges []string `json:"privileges,omitempty"`
	// GUID or name of the sub groups. A subgroup is a group assigned to a parent group.
	SubGroupIdentifiers []string `json:"sub_group_identifiers,omitempty"`
	// Group type.
	Type string `json:"type,omitempty"`
	// GUID or name of the users to assign to the group.
	UserIdentifiers []string `json:"user_identifiers,omitempty"`
	// Visibility of the group. To make a group visible to other users and groups, set the visibility to SHAREABLE.
	Visibility string `json:"visibility,omitempty"`
	// Role identifiers of the roles that should be assigned to the group.
	RoleIdentifiers []string `json:"role_identifiers,omitempty"`
}
