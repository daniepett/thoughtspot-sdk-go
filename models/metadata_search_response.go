package models

type MetadataSearchResponse struct {
	MetadataId    string `json:"metadata_id"`
	MetadataName  string `json:"metadata_name"`
	MetadataType  string `json:"metadata_type"`
	MetadataObjId string `json:"metadata_obj_id"`
}

// dependent_objects
// incomplete_objects
// metadata_detail
// metadata_header
// visualization_headers
// stats
