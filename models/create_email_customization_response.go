package models

type CreateEmailCustomizationResponse struct {
	TenantID           string                        `json:"tenant_id"`
	Org                Org                           `json:"org"`
	Name               string                        `json:"name"`
	TemplateProperties TemplatePropertiesInputCreate `json:"template_properties"`
}
