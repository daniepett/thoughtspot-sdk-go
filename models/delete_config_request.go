package models

// DeleteConfigRequest struct for DeleteConfigRequest
type DeleteConfigRequest struct {
	//    Applicable when Orgs is enabled in the cluster      Indicator to consider cluster level or org level config. Set it to false to delete configuration from current org. If set to true, then the configuration at cluster level and orgs that inherited the configuration from cluster level will be deleted.  <br/>  <span class=\"since-beta-tag\">Version: 9.5.0.cl or later</span>
	ClusterLevel bool `json:"cluster_level,omitempty"`
}
