package models

// CreateConfigRequest struct for CreateConfigRequest
type CreateConfigRequest struct {
	// URL for connecting to remote repository
	RepositoryUrl string `json:"repository_url"`
	// Username to authenticate connection to remote repository
	Username string `json:"username"`
	// Access token corresponding to the user to authenticate connection to remote repository
	AccessToken string `json:"access_token"`
	//    Applicable when Orgs is enabled in the cluster      List of Org ids or name. Provide value -1 for cluster level. Example : [\"OrgID1-or-Name1\", \"OrgID2-or-Name2\"]         Note: If no value is specified, then the configurations will be returned for all orgs the user has access to  <br/>  <span class=\"since-beta-tag\">Version: 9.5.0.cl or later</span>
	OrgIdentifier *string `json:"org_identifier,omitempty"`
	// List the remote branches to configure. Example:[development, production]
	BranchNames []string `json:"branch_names,omitempty"`
	// Name of the remote branch where objects from this Thoughtspot instance will be versioned. <br/>  <span class=\"since-beta-tag\">Version: 9.7.0.cl or later</span>
	CommitBranchName string `json:"commit_branch_name,omitempty"`
	// Maintain mapping of guid for the deployment to an instance <br/>  <span class=\"since-beta-tag\">Version: 9.4.0.cl or later</span>
	EnableGuidMapping bool `json:"enable_guid_mapping,omitempty"`
	//    Name of the branch where the configuration files related to operations between Thoughtspot and version control repo should be maintained.      Note: If no branch name is specified, then by default, ts_config_files branch is considered. Ensure this branch exists before configuration. <br/>  <span class=\"since-beta-tag\">Version: 9.7.0.cl or later</span>
	ConfigurationBranchName *string `json:"configuration_branch_name,omitempty"`
}