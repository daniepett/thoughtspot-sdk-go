package thoughtspot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/thoughtspot-sdk-go/models"
)

func (c *Client) CreateEmailCustomization(r models.CreateEmailCustomizationRequest) (*models.CreateEmailCustomizationResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/customization/email", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := models.CreateEmailCustomizationResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) SearchEmailCustomization(r models.CustomizationEmailSearchRequest) ([]models.CreateEmailCustomizationResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/customization/email/search", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	body, res, err := c.doRequest(req)
	if err != nil && res.StatusCode != 400 {
		return nil, err
	}

	conn := []models.CreateEmailCustomizationResponse{}

	if res.StatusCode == 400 {
		return nil, nil
	}

	err = json.Unmarshal(body, &conn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Client) UpdateEmailCustomization(r models.UpdateEmailCustomizationRequest) error {
	rb, err := json.Marshal(r)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/customization/email/update", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteOrgEmailCustomization(r models.CustomizationEmailDeleteRequest) error {
	rb, err := json.Marshal(r)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/customization/email/delete", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, res, err := c.doRequest(req)
	if err != nil && res.StatusCode != 400 {
		return err
	}

	return nil
}

func (c *Client) ValidateEmailCustomization() error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/customization/email/validate", c.HostURL), nil)
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
