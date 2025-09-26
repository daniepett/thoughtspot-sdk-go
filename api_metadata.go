package thoughtspot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/thoughtspot-sdk-go/models"
)

func (c *Client) ImportMetadataTML(r models.ImportMetadataTMLRequest) ([]models.ImportMetadataTMLResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/metadata/tml/import", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var m []models.ImportMetadataTMLResponse
	err = json.Unmarshal(body, &m)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *Client) ExportMetadataTML(r models.ExportMetadataTMLRequest) ([]models.ExportMetadataTMLResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/metadata/tml/export", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, res, err := c.doRequest(req)
	if err != nil && res.StatusCode != 400 {
		return nil, err
	}

	if res.StatusCode == 400 {
		return nil, nil
	}

	var m []models.ExportMetadataTMLResponse
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *Client) DeleteMetadata(r models.DeleteMetadataRequest) error {
	rb, err := json.Marshal(r)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/metadata/delete", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
