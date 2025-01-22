package models

// RuntimeSorts Objects to apply the Runtime_Sorts.
type RuntimeSorts struct {
	// The column name to apply filter.
	ColumnName string `json:"column_name,omitempty"`
	// Order for the sort.
	Order string `json:"order,omitempty"`
	// Flag to persist the runtime sorts. <br/>  <span class=\"since-beta-tag\">Version: 9.12.0.cl or later</span>
	Persist bool `json:"persist,omitempty"`
	// Object to apply the runtime sort.
	Objects []UserObject `json:"objects,omitempty"`
}