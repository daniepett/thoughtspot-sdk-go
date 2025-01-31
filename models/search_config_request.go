package models

// SearchConfigRequest struct for SearchConfigRequest
type SearchConfigRequest struct {
	//    Applicable when Orgs is enabled in the cluster      List of Org ids or name. Provide value -1 for cluster level. Example : [\"OrgID1-or-Name1\", \"OrgID2-or-Name2\"]         Note: If no value is specified, then the configurations will be returned for all orgs the user has access to  <br/>  <span class=\"since-beta-tag\">Version: 9.5.0.cl or later</span>
	OrgIdentifiers []string `json:"org_identifiers,omitempty"`
}