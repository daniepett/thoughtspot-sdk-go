package models

// ImportMetadataTMLRequest struct for ImportMetadataTMLRequest
type ImportMetadataTMLResponse struct {
	// Details of TML objects.  **Note: importing TML in YAML format, when coming directly from our Playground, is currently requires manual formatting. For more details on the workaround, please click [here](https://developers.thoughtspot.com/docs/known-issues#_version_9_12_0_cl)**
	Response ImportMetadataTMLResponseResponse `json:"response"`
}

type ImportMetadataTMLResponseResponse struct {
	// Details of TML objects.  **Note: importing TML in YAML format, when coming directly from our Playground, is currently requires manual formatting. For more details on the workaround, please click [here](https://developers.thoughtspot.com/docs/known-issues#_version_9_12_0_cl)**
	Header map[string]interface{} `json:"header"`
}
