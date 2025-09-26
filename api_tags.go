package thoughtspot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/thoughtspot-sdk-go/models"
)

func (c *Client) CreateTag(r models.TagsCreateRequest) (*models.Tag, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/tags/create", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	m := models.Tag{}
	err = json.Unmarshal(body, &m)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (c *Client) SearchTags(r models.TagsSearchRequest) ([]models.Tag, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/tags/search", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var m []models.Tag
	err = json.Unmarshal(body, &m)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (c *Client) UpdateTag(id string, r models.TagsUpdateRequest) error {
	rb, err := json.Marshal(r)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/tags/%s/update", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteTag(id string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/tags/%s/delete", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) AssignTag(r models.TagsAssignRequest) error {
	rb, err := json.Marshal(r)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/tags/assign", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) UnassignTag(r models.TagsUnassignRequest) error {
	rb, err := json.Marshal(r)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/tags/unassign", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
