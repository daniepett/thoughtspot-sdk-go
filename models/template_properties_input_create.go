package models

type TemplatePropertiesInputCreate struct {
	CtaButtonBgColor             string `json:"cta_button_bg_color,omitempty"`
	CtaTextFontColor             string `json:"cta_text_font_color,omitempty"`
	PrimaryBgColor               string `json:"primary_bg_color,omitempty"`
	HomeURL                      string `json:"home_url,omitempty"`
	LogoURL                      string `json:"logo_url,omitempty"`
	FontFamily                   string `json:"font_family,omitempty"`
	ProductName                  string `json:"product_name,omitempty"`
	FooterAddress                string `json:"footer_address,omitempty"`
	FooterPhone                  string `json:"footer_phone,omitempty"`
	ReplacementValueForLiveboard string `json:"replacement_value_for_liveboard,omitempty"`
	ReplacementValueForAnswer    string `json:"replacement_value_for_answer,omitempty"`
	ReplacementValueForSpotIQ    string `json:"replacement_value_for_spot_iq,omitempty"`
	HideFooterAddress            bool   `json:"hide_footer_address,omitempty"`
	HideFooterPhone              bool   `json:"hide_footer_phone,omitempty"`
	HideManageNotification       bool   `json:"hide_manage_notification,omitempty"`
	HideMobileAppNudge           bool   `json:"hide_mobile_app_nudge,omitempty"`
	HidePrivacyPolicy            bool   `json:"hide_privacy_policy,omitempty"`
	HideProductName              bool   `json:"hide_product_name,omitempty"`
	HideTSVocabularyDefinitions  bool   `json:"hide_ts_vocabulary_definitions,omitempty"`
	HideNotificationStatus       bool   `json:"hide_notification_status,omitempty"`
	HideErrorMessage             bool   `json:"hide_error_message,omitempty"`
	HideUnsubscribeLink          bool   `json:"hide_unsubscribe_link,omitempty"`
	HideModifyAlert              bool   `json:"hide_modify_alert,omitempty"`
}
