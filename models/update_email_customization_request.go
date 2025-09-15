package models

type UpdateEmailCustomizationRequest struct {
	OrgIdentifier      string                        `json:"org_identifier"`
	TemplateProperties TemplatePropertiesInputCreate `json:"template_properties"`
}
