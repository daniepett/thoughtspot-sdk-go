package models

// SearchRolesRequest struct for SearchRolesRequest
type SearchRolesRequest struct {
	// unique ID or name of the Roles
	RoleIdentifiers []string `json:"role_identifiers,omitempty"`
	// Unique Id or name of the Organisation
	OrgIdentifiers []string `json:"org_identifiers,omitempty"`
	// Unique Id or name of the User Group
	GroupIdentifiers []string `json:"group_identifiers,omitempty"`
	// Privileges assigned to the Role. See [Documentation](https://developers.thoughtspot.com/docs/rbac#_role_categories_and_privileges)for supported roles privileges.
	Privileges []string `json:"privileges,omitempty"`
	// Indicates whether the Role is deprecated.
	// Deprecated
	Deprecated bool `json:"deprecated,omitempty"`
	// Indicates whether the Role is external
	// Deprecated
	External bool `json:"external,omitempty"`
	// Indicates whether the Role is shared via connection
	// Deprecated
	SharedViaConnection bool `json:"shared_via_connection,omitempty"`
	// Permission details of the Role
	// Deprecated
	Permissions []string `json:"permissions,omitempty"`
}
