package thoughtspot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/thoughtspot-sdk-go/models"
)

func (c *Client) CreateRole(r models.CreateRoleRequest) (*models.RoleResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/roles/create", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	m := models.RoleResponse{}
	err = json.Unmarshal(body, &m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (c *Client) SearchRoles(r models.SearchRolesRequest) ([]models.SearchRoleResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/roles/search", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var m []models.SearchRoleResponse
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *Client) UpdateRole(id string, r models.UpdateRoleRequest) (*models.RoleResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/roles/%s/update", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	m := models.RoleResponse{}
	err = json.Unmarshal(body, &m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}
func (c *Client) DeleteRole(id string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/roles/%s/delete", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
