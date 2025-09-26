package thoughtspot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/thoughtspot-sdk-go/models"
)

// GetToken - Get a new token for user
func (c *Client) GetToken() (*models.GetTokenResponse, error) {
	// if c.Auth.Username == "" || c.Auth.Password == "" {
	// 	return nil, fmt.Errorf("Missing ClientID and ClientSecret")
	// }

	// if c.Auth.GrantType != "client_credentials" {
	// 	return nil, fmt.Errorf("Grant type not supported")
	// }

	rb, err := json.Marshal(c.Auth)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/token/full", c.HostURL), strings.NewReader(string(rb)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	// if http.DetectContentType(body) == "text/html; charset=utf-8" {
	// 	return nil, fmt.Errorf("The cluster appears to be in maintenance or startup mode")
	// }

	ar := models.GetTokenResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
