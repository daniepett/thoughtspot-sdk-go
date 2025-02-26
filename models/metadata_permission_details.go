package models

// PrincipalsInput struct for PrincipalsInput
type MetadataPermissionDetails struct {
	MetadataId              string                    `json:"metadata_id"`
	MetadataName            string                    `json:"metadata_name"`
	MetadataType            string                    `json:"metadata_type"`
	MetadataOwner           ObjectIDAndName           `json:"metadata_owner"`
	MetadataAuthor          ObjectIDAndName           `json:"metadata_author"`
	PrincipalPermissionInfo []PrincipalPermissionInfo `json:"principal_permission_info"`
}
