package toggl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	togglAPIBaseURL = "https://api.track.toggl.com/api/v8"
	timeEntriesURI  = "/time_entries"
)

type Client struct {
	apiKey     string
	URL        *url.URL
	httpClient *http.Client
}

func NewClient(apiKey string, timeout time.Duration) *Client {
	parsedURL, _ := url.ParseRequestURI(togglAPIBaseURL)
	httpClient := &http.Client{
		Timeout: timeout,
	}
	return &Client{
		apiKey:     apiKey,
		URL:        parsedURL,
		httpClient: httpClient,
	}
}

func (c *Client) newRequest(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)
	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.SetBasicAuth(c.apiKey, "api_token")
	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func (c *Client) GetTimeEntries(ctx context.Context) ([]TimeEntry, error) {
	req, err := c.newRequest(ctx, http.MethodGet, timeEntriesURI, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusForbidden:
		return nil, errors.New(fmt.Sprintf("APIKey may be not valid. status is %v", res.Status))
	}

	var timeEntries []TimeEntry
	if err := decodeBody(res, &timeEntries); err != nil {
		return nil, err
	}

	return timeEntries, nil
}
