package models

// ExportMetadataTMLRequest struct for ExportMetadataTMLRequest
type ExportMetadataTMLResponse struct {
	Edoc string                        `json:"edoc"`
	Info ExportMetadataTMLResponseInfo `json:"info"`
}

type ExportMetadataTMLResponseInfo struct {
	Filename string `json:"filename"`
	Id       string `json:"id"`
	Name     string `json:"name"`
}
