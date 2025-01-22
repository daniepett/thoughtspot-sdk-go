package models

// RuntimeParameters Objects to apply the Runtime_Parameters.
type RuntimeParameters struct {
	// The name of the parameter.
	Name string `json:"name"`
	// The array of values.
	Values []string `json:"values"`
	// Flag to persist the parameters. <br/>  <span class=\"since-beta-tag\">Version: 9.12.0.cl or later</span>
	Persist bool `json:"persist,omitempty"`
	// Object to apply the runtime parameter.
	Objects []UserObject `json:"objects,omitempty"`
}
