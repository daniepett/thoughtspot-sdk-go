package models

// ExportMetadataTMLRequest struct for ExportMetadataTMLRequest
type ExportMetadataTMLRequest struct {
	// Metadata objects.
	Metadata []ExportMetadataTypeInput `json:"metadata"`
	// Indicates whether to export associated metadata objects of specified metadata objects.
	ExportAssociated bool `json:"export_associated,omitempty"`
	// Adds FQNs of the referenced objects. For example, if you are exporting a Liveboard and its associated objects, the API returns the Liveboard TML data with the FQNs of the referenced worksheet. If the exported TML data includes FQNs, you don't need to manually add FQNs of the referenced objects during TML import.
	ExportFqn bool `json:"export_fqn,omitempty"`
	// TML EDOC content format.  **Note: exporting in YAML format currently requires manual formatting of the output. For more details on the workaround, please click [here](https://developers.thoughtspot.com/docs/known-issues#_version_9_12_0_cl)**
	EdocFormat string `json:"edoc_format,omitempty"`
	// Indicates whether to export worksheet TML in DEFAULT or V1 or V2 version.
	ExportSchemaVersion *string `json:"export_schema_version,omitempty"`
	// Indicates whether to export table while exporting connection.
	ExportDependent bool `json:"export_dependent,omitempty"`
	// Indicates whether to export connection as dependent while exporting table/worksheet/answer/liveboard. This will only be active when export_associated is true.
	ExportConnectionAsDependent bool `json:"export_connection_as_dependent,omitempty"`
	// Indicates whether to export is happening from all orgs context.
	AllOrgsOverride bool `json:"all_orgs_override,omitempty"`
	// Flags to specify additional options for export. <br/>  <span class=\"since-beta-tag\">Version: 10.5.0.cl or later</span>
	ExportOptions ExportOptions `json:"export_options,omitempty"`
}
