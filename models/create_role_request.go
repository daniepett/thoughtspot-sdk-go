package models

// CreateRoleRequest struct for CreateRoleRequest
type CreateRoleRequest struct {
	// Unique name of the Role.
	Name string `json:"name"`
	// Description of the Role.
	Description string `json:"description,omitempty"`
	// Privileges granted to the Role. See [Documentation](https://developers.thoughtspot.com/docs/rbac#_role_categories_and_privileges)for supported roles privileges.
	Privileges []string `json:"privileges,omitempty"`
}
