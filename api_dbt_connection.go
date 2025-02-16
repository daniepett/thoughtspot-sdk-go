package thoughtspot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/thoughtspot-sdk-go/models"
)

func (c *Client) CreateDbtConnection(r models.DbtConnectionRequest) (*map[string]interface{}, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	// fmt.Print("This is the request", string(rb))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dbt/dbt-connection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		fmt.Print("This is the create error", string(err.Error()))
		return nil, err
	}

	fmt.Print("This is the create response", string(body))

	var connTemplate interface{}
	err = json.Unmarshal(body, &connTemplate)
	if err != nil {
		return nil, err
	}

	connMap := connTemplate.(map[string]interface{})
	return &connMap, nil
}

func (c *Client) SearchDbtConnection() ([]models.DbtSearchResponse, error) {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dbt/search", c.HostURL), nil)
	if err != nil {
		return nil, err
	}
	body, res, err := c.doRequest(req)
	if err != nil && res.StatusCode != 400 {
		return nil, err
	}

	fmt.Print("This is the search_response", string(body))

	conn := []models.DbtSearchResponse{}

	if res.StatusCode == 400 {
		return conn, nil
	}

	err = json.Unmarshal(body, &conn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Client) UpdateDbtConnection(r models.UpdateDbtConnectionRequest) (*map[string]interface{}, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dbt/update-dbt-connection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	var connTemplate interface{}
	err = json.Unmarshal(body, &connTemplate)
	if err != nil {
		return nil, err
	}

	connMap := connTemplate.(map[string]interface{})
	return &connMap, nil
}

func (c *Client) DeleteDbtConnection(id string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/dbt/%s/delete", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
