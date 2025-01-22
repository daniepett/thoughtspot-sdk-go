package models

// UpdateRoleRequest struct for UpdateRoleRequest
type UpdateRoleRequest struct {
	// Name of the Role.
	Name string `json:"name"`
	// Description of the Role.
	Description string `json:"description,omitempty"`
	// Privileges granted to the role. See [Documentation](https://developers.thoughtspot.com/docs/rbac#_role_categories_and_privileges)for supported roles privileges.
	Privileges []string `json:"privileges,omitempty"`
}
