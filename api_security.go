package thoughtspot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/thoughtspot-sdk-go/models"
)

func (c *Client) ShareMetadata(r models.ShareMetadataRequest) error {
	rb, err := json.Marshal(r)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/security/metadata/share", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) FetchPermissionsOnMetadata(r models.FetchPermissionsOnMetadataRequest) (*models.PermissionOfMetadataResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/security/metadata/fetch-permissions", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	m := models.PermissionOfMetadataResponse{}
	err = json.Unmarshal(body, &m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}
