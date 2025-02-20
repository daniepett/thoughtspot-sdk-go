package thoughtspot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/thoughtspot-sdk-go/models"
)

func (c *Client) CreateUserGroup(r models.CreateUserGroupRequest) (*models.UserGroupResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/groups/create", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	m := models.UserGroupResponse{}
	err = json.Unmarshal(body, &m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (c *Client) SearchUserGroups(r models.SearchUserGroupsRequest) ([]models.UserGroupResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/groups/search", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var m []models.UserGroupResponse
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *Client) UpdateUserGroup(id string, r models.UpdateUserGroupRequest) error {
	rb, err := json.Marshal(r)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/groups/%s/update", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteUserGroup(id string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/groups/%s/delete", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
