package models

// RuntimeFilters Objects to apply the Runtime_Filters.
type RuntimeFilters struct {
	// The column name to apply filter.
	ColumnName string `json:"column_name"`
	// Value of the filters.
	Values []string `json:"values"`
	// Operator value. Example: EQ
	Operator string `json:"operator"`
	// Flag to persist the runtime filters. <br/>  <span class=\"since-beta-tag\">Version: 9.12.0.cl or later</span>
	Persist bool `json:"persist,omitempty"`
	// Object to apply the runtime filter.
	Objects []UserObject `json:"objects,omitempty"`
}
