package models

type ImportMetadataTMLResponse struct {
	Response     ImportMetadataTMLResponseResponse `json:"response"`
	RequestIndex int32                             `json:"request_index"`
}

type ImportMetadataTMLResponseResponse struct {
	Header ImportMetadataTMLResponseResponseHeader `json:"header"`
	Status ImportMetadataTMLResponseStatus         `json:"status"`
	Action string                                  `json:"action"`
}

type ImportMetadataTMLResponseResponseHeader struct {
	AuthorName          string `json:"author_name,omitempty"`
	SchemaStripe        string `json:"schema_stripe,omitempty"`
	AuthorGuid          string `json:"author_guid,omitempty"`
	Created             int64  `json:"created,omitempty"`
	MetadataType        string `json:"metadata_type,omitempty"`
	OwnerGuid           string `json:"owner_guid,omitempty"`
	Type                string `json:"type,omitempty"`
	IdGuid              string `json:"id_guid,omitempty"`
	IsVersioningEnabled bool   `json:"is_versioning_enabled,omitempty"`
	DatabaseStripe      string `json:"database_stripe,omitempty"`
	Name                string `json:"name,omitempty"`
	ModifiedBy          string `json:"modified_by,omitempty"`
	ObjId               string `json:"objId,omitempty"`
	Modified            int64  `json:"modified,omitempty"`
	AuthorDisplayName   string `json:"author_display_name,omitempty"`
}

type ImportMetadataTMLResponseStatus struct {
	ErrorMessage string `json:"error_message"`
	StatusCode   string `json:"status_code"`
	ErrorCode    int32  `json:"error_code"`
}
