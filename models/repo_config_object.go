package models

// RepoConfigObject struct for RepoConfigObject
type RepoConfigObject struct {
	// Remote repository URL configured
	RepositoryUrl string `json:"repository_url,omitempty"`
	// Username to authenticate connection to the version control system
	Username string `json:"username,omitempty"`
	// Name of the remote branch where objects from this Thoughtspot instance will be versioned.
	CommitBranchName string `json:"commit_branch_name,omitempty"`
	// Branches that have been pulled in local repository
	Branches []string `json:"branches,omitempty"`
	// Maintain mapping of guid for the deployment to an instance
	EnableGuidMapping bool `json:"enable_guid_mapping,omitempty"`
	// Name of the branch where the configuration files related to operations between Thoughtspot and version control repo should be maintained.
	ConfigurationBranchName string `json:"configuration_branch_name,omitempty"`
	Org *Org `json:"org,omitempty"`
}