package models

// ImportMetadataTMLRequest struct for ImportMetadataTMLRequest
type ImportMetadataTMLRequest struct {
	// Details of TML objects.  **Note: importing TML in YAML format, when coming directly from our Playground, is currently requires manual formatting. For more details on the workaround, please click [here](https://developers.thoughtspot.com/docs/known-issues#_version_9_12_0_cl)**
	MetadataTmls []string `json:"metadata_tmls"`
	// Specifies the import policy for the TML import.
	ImportPolicy *string `json:"import_policy,omitempty"`
	// If selected, creates TML objects with new GUIDs.
	CreateNew bool `json:"create_new,omitempty"`
	// If import is happening from all orgs context.
	AllOrgsContext bool `json:"all_orgs_context,omitempty"`
	// Boolean to indicate if the CDW validation for table imports should be skipped.
	SkipCdwValidationForTables bool `json:"skip_cdw_validation_for_tables,omitempty"`
	// <div><span class=\"since-beta-tag\">Version: 10.5.0.cl or later</span></div>  Boolean to indicate if the large metadata validation should be enabled.
	EnableLargeMetadataValidation bool `json:"enable_large_metadata_validation,omitempty"`
}
