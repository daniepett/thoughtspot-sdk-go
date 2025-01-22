package models


// GetTokenResponse struct for GetTokenResponse
type GetTokenResponse struct {
	// Bearer auth token.
	Token string `json:"token"`
	// Token creation time in milliseconds.
	CreationTimeInMillis float32 `json:"creation_time_in_millis"`
	// Token expiration time in milliseconds.
	ExpirationTimeInMillis float32 `json:"expiration_time_in_millis"`
	// Username to whom the token is issued.
	ValidForUserId string `json:"valid_for_user_id"`
	// Unique identifier of the user to whom the token is issued.
	ValidForUsername string `json:"valid_for_username"`
}