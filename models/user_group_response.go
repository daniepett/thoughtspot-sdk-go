package models

// UserGroupResponse struct for UserGroupResponse
type UserGroupResponse struct {
	// The unique identifier of the object
	AuthorId string `json:"author_id,omitempty"`
	// Indicates whether the response has complete detail of the group.
	CompleteDetail bool `json:"complete_detail,omitempty"`
	// Content details of the group
	Content map[string]interface{} `json:"content,omitempty"`
	// Creation time of the group in milliseconds
	CreationTimeInMillis float32 `json:"creation_time_in_millis,omitempty"`
	// Liveboards that are assigned as default Liveboards to the group.
	DefaultLiveboards []UserGroup `json:"default_liveboards,omitempty"`
	// Indicates whether the group is deleted
	Deleted bool `json:"deleted,omitempty"`
	// Indicates whether the group is deprecated
	Deprecated bool `json:"deprecated,omitempty"`
	// Description of the group
	Description string `json:"description,omitempty"`
	// Display name of the group.
	DisplayName string `json:"display_name"`
	// Indicates whether the group is external
	External bool `json:"external,omitempty"`
	// Generation number of the group
	GenerationNumber int32 `json:"generation_number,omitempty"`
	// Indicates whether the group is hidden
	Hidden bool `json:"hidden,omitempty"`
	// The unique identifier of the object
	Id string `json:"id"`
	// Index number of the group
	Index int32 `json:"index,omitempty"`
	// Index version number of the group
	IndexVersion int32 `json:"index_version,omitempty"`
	// Metadata version number of the group
	MetadataVersion int32 `json:"metadata_version,omitempty"`
	// Last modified time of the group in milliseconds.
	ModificationTimeInMillis float32 `json:"modification_time_in_millis,omitempty"`
	// The unique identifier of the object
	ModifierId string `json:"modifier_id,omitempty"`
	// Name of the group.
	Name string `json:"name"`
	// Orgs in which group exists.
	Orgs []UserGroup `json:"orgs,omitempty"`
	// The unique identifier of the object
	OwnerId string `json:"owner_id,omitempty"`
	// Parent type of the group.
	ParentType string `json:"parent_type,omitempty"`
	// Privileges which are assigned to the group
	Privileges []string `json:"privileges,omitempty"`
	// Groups who are part of the group
	SubGroups []UserGroup `json:"sub_groups,omitempty"`
	// Indicates whether the group is a system group.
	SystemGroup bool `json:"system_group,omitempty"`
	// Tags associated with the group.
	Tags []UserGroup `json:"tags,omitempty"`
	// Type of the group.
	Type string `json:"type,omitempty"`
	// Users who are part of the group.
	Users []UserGroup `json:"users,omitempty"`
	// Visibility of the group. The SHARABLE makes a group visible to other users and groups, and thus allows them to share objects.
	Visibility string `json:"visibility"`
	// List of roles assgined to the user
	Roles []Role `json:"roles,omitempty"`
}
