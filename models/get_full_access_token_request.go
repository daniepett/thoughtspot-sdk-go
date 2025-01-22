package models


// GetFullAccessTokenRequest struct for GetFullAccessTokenRequest
type GetFullAccessTokenRequest struct {
	// Username of the ThoughtSpot user. The username is stored in the `name` attribute of the user object.
	Username string `json:"username"`
	// Password of the user account
	Password *string `json:"password,omitempty"`
	// The secret key string provided by the ThoughtSpot application server. ThoughtSpot generates a secret key when Trusted authentication is enabled.
	SecretKey *string `json:"secret_key,omitempty"`
	// Token validity duration in seconds
	ValidityTimeInSec *int32 `json:"validity_time_in_sec,omitempty"`
	// ID of the Org context to log in to. If the Org ID is not specified and secret key is provided then user will be logged into the org corresponding to the secret key, and if secret key is not provided then user will be logged in to the Org context of their previous login session.
	OrgId *int32 `json:"org_id,omitempty"`
	// Email address of the user. Specify this attribute when creating a new user (just-in-time (JIT) provisioning).
	Email *string `json:"email,omitempty"`
	// Indicates display name of the user. Use this parameter to provision a user just-in-time (JIT).
	DisplayName *string `json:"display_name,omitempty"`
	//    Creates a new user if the specified username does not already exist in ThoughtSpot. To provision a user just-in-time (JIT), set this attribute to true.      Note: For JIT provisioning of a user, the secret_key is required. 
	AutoCreate bool `json:"auto_create,omitempty"`
	// ID or name of the groups to which the newly created user belongs. Use this parameter to provision a user just-in-time (JIT).
	GroupIdentifiers []string `json:"group_identifiers,omitempty"`
	// <div><span class=\"since-beta-tag\">Version: 9.10.5.cl or later</span><span class=\"since-beta-tag\">Deprecated: 10.4.0.cl and later</span></div>  Define attributes such as Runtime filters and Runtime parameters to send security entitlements to a user session. For more information, see [Documentation](https://developers.thoughtspot.com/docs/abac-user-parameters).
	UserParameters *UserParameterOptions `json:"user_parameters,omitempty"`
}
