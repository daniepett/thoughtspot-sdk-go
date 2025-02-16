package thoughtspot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/thoughtspot-sdk-go/models"
)

func (c *Client) CreateConnection(r models.CreateConnectionRequest) (*models.CreateConnectionResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	fmt.Print("This is the request", string(rb))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/connection/create", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	conn := models.CreateConnectionResponse{}
	err = json.Unmarshal(body, &conn)
	if err != nil {
		return nil, err
	}

	return &conn, nil
}

func (c *Client) SearchConnection(r models.SearchConnectionRequest) ([]models.SearchConnectionResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/connection/search", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	body, res, err := c.doRequest(req)
	if err != nil && res.StatusCode != 400 {
		return nil, err
	}

	conn := []models.SearchConnectionResponse{}

	if res.StatusCode == 400 {
		return conn, nil
	}

	err = json.Unmarshal(body, &conn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Client) UpdateConnection(id string, r models.UpdateConnectionRequest) error {
	rb, err := json.Marshal(r)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/connections/%s/update", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteConnection(id string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/connections/%s/delete", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
