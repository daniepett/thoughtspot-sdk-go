package models

// UserParameterOptions Define attributes such as Runtime filters and Runtime parameters to send security entitlements to a user session. For more information, see [Documentation](https://developers.thoughtspot.com/docs/abac-user-parameters). 
type UserParameterOptions struct {
	Objects []UserObject `json:"objects,omitempty"`
	// Objects to apply the User_Runtime_Filters.  Examples to set the `runtime_filters` : ```json { \"column_name\": \"Color\", \"operator\": \"EQ\", \"values\": [\"red\"], \"persist\": false } ```
	RuntimeFilters []RuntimeFilters `json:"runtime_filters,omitempty"`
	// Objects to apply the User_Runtime_Sorts.  Examples to set the `runtime_sorts` : ```json { \"column_name\": \"Color\", \"order\": \"ASC\", \"persist\": false } ```
	RuntimeSorts []RuntimeSorts `json:"runtime_sorts,omitempty"`
	// Objects to apply the Runtime_Parameters.  Examples to set the `parameters` : ```json { \"name\": \"Color\", \"values\": [\"Blue\"], \"persist\": false } ```
	Parameters []RuntimeParameters `json:"parameters,omitempty"`
}
