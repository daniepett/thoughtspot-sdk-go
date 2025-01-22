package models

// ExportOptions Flags to specify additional options for export. This will only be active when UserDefinedId in TML is enabled.
type ExportOptions struct {
	// Boolean Flag to whether to export user_defined_id of referenced object. This will only be respected when UserDefinedId in TML is enabled.
	ExportWithReferenceObjectId bool `json:"export_with_reference_object_id,omitempty"`
	// Boolean flag to whether to export guid of the object. This will only be respected when UserDefinedId in TML is enabled.
	IncludeGuid bool `json:"include_guid"`
	// Boolean flag to whether to export user_defined_id of the object. This will only be respected when UserDefinedId in TML is enabled.
	IncludeObjectId bool `json:"include_object_id,omitempty"`
}
