package models

// User struct for User
type User struct {
	// Unique identifier of the user.
	Id string `json:"id"`
	// Name of the user.
	Name string `json:"name"`
	// Display name of the user.
	DisplayName string `json:"display_name"`
	// Visibility of the users. The `SHARABLE` property makes a user visible to other users and group, who can share objects with the user.
	Visibility string `json:"visibility"`
	// Unique identifier of author of the user.
	AuthorId string `json:"author_id,omitempty"`
	// Defines whether the user can change their password.
	CanChangePassword bool `json:"can_change_password,omitempty"`
	// Defines whether the response has complete detail of the user.
	CompleteDetail bool `json:"complete_detail,omitempty"`
	// Creation time of the user in milliseconds.
	CreationTimeInMillis float32 `json:"creation_time_in_millis,omitempty"`
	// CurrentOrg *Org `json:"current_org,omitempty"`
	// // Indicates whether the user is deleted.
	// Deleted bool `json:"deleted,omitempty"`
	// // Indicates whether the user is deprecated.
	// Deprecated bool `json:"deprecated,omitempty"`
	// // Type of the user account.
	// AccountType string `json:"account_type,omitempty"`
	// // Status of the user account.
	// AccountStatus string `json:"account_status,omitempty"`
	// // Email of the user.
	// Email string `json:"email,omitempty"`
	// // Expiration time of the user in milliseconds.
	// ExpirationTimeInMillis float32 `json:"expiration_time_in_millis,omitempty"`
	// // Indicates whether the user is external.
	// External bool `json:"external,omitempty"`
	// // Metadata objects to add to the users' favorites list.
	// FavoriteMetadata []FavoriteMetadataItem `json:"favorite_metadata,omitempty"`
	// // Timestamp of the first login session of the user in milliseconds.
	// FirstLoginTimeInMillis float32 `json:"first_login_time_in_millis,omitempty"`
	// // Group mask of the user.
	// GroupMask int32 `json:"group_mask,omitempty"`
	// // Indicates whether the user is hidden.
	// Hidden bool `json:"hidden,omitempty"`
	// HomeLiveboard *ObjectIDAndName `json:"home_liveboard,omitempty"`
	// // Incomplete details of user if any present.
	// IncompleteDetails map[string]interface{} `json:"incomplete_details,omitempty"`
	// // Indicates whether it is first login of the user.
	// IsFirstLogin bool `json:"is_first_login,omitempty"`
	// // Last modified time of the user in milliseconds.
	// ModificationTimeInMillis float32 `json:"modification_time_in_millis,omitempty"`
	// // Unique identifier of modifier of the user.
	// ModifierId string `json:"modifier_id,omitempty"`
	// // User preference for receiving email notifications on shared Answers or Liveboard.
	// NotifyOnShare bool `json:"notify_on_share,omitempty"`
	// // The user preference for turning off the onboarding experience.
	// OnboardingExperienceCompleted bool `json:"onboarding_experience_completed,omitempty"`
	// // Orgs to which the user belongs.
	// Orgs []Org `json:"orgs,omitempty"`
	// // Unique identifier of owner of the user.
	// OwnerId string `json:"owner_id,omitempty"`
	// // Parent type of the user.
	// ParentType string `json:"parent_type,omitempty"`
	// // Privileges which are assigned to the user.
	// Privileges []string `json:"privileges,omitempty"`
	// // User's preference to revisit the new user onboarding experience.
	// ShowOnboardingExperience bool `json:"show_onboarding_experience,omitempty"`
	// // Indicates whether the user is a super user.
	// SuperUser bool `json:"super_user,omitempty"`
	// // Indicates whether the user is a system user.
	// SystemUser bool `json:"system_user,omitempty"`
	// // Tags associated with the user.
	// Tags []ObjectIDAndName `json:"tags,omitempty"`
	// // Unique identifier of tenant of the user.
	// TenantId string `json:"tenant_id,omitempty"`
	// // Groups to which the user is assigned.
	// UserGroups []ObjectIDAndName `json:"user_groups,omitempty"`
	// // Inherited User Groups which the user is part of.
	// UserInheritedGroups []ObjectIDAndName `json:"user_inherited_groups,omitempty"`
	// // Indicates whether welcome email is sent for the user.
	// WelcomeEmailSent bool `json:"welcome_email_sent,omitempty"`
	// // Privileges which are assigned to the user with org.
	// OrgPrivileges map[string]interface{} `json:"org_privileges,omitempty"`
	// // Locale for the user.
	// PreferredLocale string `json:"preferred_locale,omitempty"`
	// // Properties for the user
	// ExtendedProperties map[string]interface{} `json:"extended_properties,omitempty"`
	// // Preferences for the user
	// ExtendedPreferences map[string]interface{} `json:"extended_preferences,omitempty"`
	// // User Parameters which are specified for the user via JWToken
	// UserParameters map[string]interface{} `json:"user_parameters,omitempty"`
	// // Access Control Properties which are specified for the user via JWToken
	// AccessControlProperties map[string]interface{} `json:"access_control_properties,omitempty"`
}