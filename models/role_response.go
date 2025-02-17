package models

// RoleResponse struct for RoleResponse
type RoleResponse struct {
	// Unique Id of the role.
	Id string `json:"id"`
	// Name of the role
	Name string `json:"name"`
	// Description of the role
	Description string `json:"description"`
	// number of groups assigned with this role
	GroupsAssignedCount int32 `json:"groups_assigned_count,omitempty"`
	// Orgs in which role exists.
	Orgs []GenericInfo `json:"orgs,omitempty"`
	// Details of groups assigned with this role
	Groups []GenericInfo `json:"groups,omitempty"`
	// Privileges granted to the role.
	Privileges []string `json:"privileges"`
	// Permission details of the Role
	Permission string `json:"permission,omitempty"`
	// Unique identifier of author of the role.
	AuthorId string `json:"author_id,omitempty"`
	// Unique identifier of modifier of the role.
	ModifierId string `json:"modifier_id,omitempty"`
	// Creation time of the role in milliseconds.
	CreationTimeInMillis int64 `json:"creation_time_in_millis,omitempty"`
	// Last modified time of the role in milliseconds.
	ModificationTimeInMillis int64`json:"modification_time_in_millis,omitempty"`
	// Indicates whether the role is deleted.
	Deleted bool `json:"deleted,omitempty"`
	// Indicates whether the role is deprecated.
	Deprecated bool `json:"deprecated,omitempty"`
	// Indicates whether the role is external.
	External bool `json:"external,omitempty"`
	// Indicates whether the role is hidden.
	Hidden bool `json:"hidden,omitempty"`
	// Indicates whether the role is shared via connection
	SharedViaConnection bool `json:"shared_via_connection,omitempty"`
}
