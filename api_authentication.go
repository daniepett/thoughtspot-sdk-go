package thoughtspot

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/daniepett/thoughtspot-sdk-go/models"
)

func (c *Client) GetCurrentUserInfo() (*models.User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/auth/session/user", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()

	req.URL.RawQuery = q.Encode()

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
