package models

// ExportMetadataTMLRequest struct for ExportMetadataTMLRequest
type ExportMetadataTMLResponse struct {
	Edoc string                        `json:"edoc"`
	Info ExportMetadataTMLResponseInfo `json:"info"`
}

type ExportMetadataTMLResponseInfo struct {
	Filename string                              `json:"filename"`
	Id       string                              `json:"id"`
	Name     string                              `json:"name"`
	Type     string                              `json:"type"`
	Status   ExportMetadataTMLResponseInfoStatus `json:"status"`
}

type ExportMetadataTMLResponseInfoStatus struct {
	ErrorMessage string `json:"error_message"`
	StatusCode   string `json:"status_code"`
	ErrorCode    int32  `json:"error_code"`
}
