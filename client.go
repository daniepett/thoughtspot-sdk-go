package thoughtspot

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/daniepett/thoughtspot-sdk-go/models"
)

// HostURL - Default QlikCloud URL

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       models.GetObjectAccessTokenRequest
}

// NewClient -
func NewClient(host, username, password, orgId *string) (*Client, error) {
	print(host)
	c := Client{
		HTTPClient: &http.Client{Timeout: 20 * time.Minute},
		HostURL:    fmt.Sprintf("https://%s.thoughtspot.cloud/api/rest/2.0", *host),
	}

	// If username or password not provided, return empty client
	// if *username == nil || *password == nil {
	// 	return &c, nil
	// }

	c.Auth = models.GetObjectAccessTokenRequest{
		Username: *username,
		Password: password,
		OrgId:    orgId,
	}

	// c.Auth.GrantType = "client_credentials"

	ar, err := c.GetToken()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, *http.Response, error) {
	token := c.Token

	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, res, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, res, err
	}

	statusOK := res.StatusCode >= 200 && res.StatusCode < 300

	if !statusOK {
		return nil, res, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, res, err
}
