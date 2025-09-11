package thoughtspot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/daniepett/thoughtspot-sdk-go/models"
)

func (c *Client) CreateCalendar(r models.CreateCustomCalendarRequest) (*models.CalendarResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/calendars/create", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, _, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	res := models.CalendarResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) SearchCalendars(r models.SearchCustomCalendarsRequest) ([]models.CalendarResponse, error) {
	rb, err := json.Marshal(r)

	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/calendars/search", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	body, res, err := c.doRequest(req)
	if err != nil && res.StatusCode != 400 {
		return nil, err
	}

	conn := []models.CalendarResponse{}

	if res.StatusCode == 400 {
		return nil, nil
	}

	err = json.Unmarshal(body, &conn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Client) UpdateCalendar(id string, r models.UpdateCustomCalendarRequest) error {
	rb, err := json.Marshal(r)

	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/calendars/%s/update", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteCalendar(id string) error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/calendars/%s/delete", c.HostURL, id), nil)
	if err != nil {
		return err
	}

	_, _, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
