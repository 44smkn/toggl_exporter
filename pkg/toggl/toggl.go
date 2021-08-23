package toggl

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	togglAPIBaseURL               = "https://api.track.toggl.com/api/v8"
	currentUserWithRelatedDataURI = "/me?with_related_data=true"
)

type Client struct {
	apiKey     string
	URL        *url.URL
	httpClient *http.Client
}

func NewClient(apiKey string, timeout int) *Client {
	parsedURL, _ := url.ParseRequestURI(togglAPIBaseURL)
	httpClient := &http.Client{
		Timeout: time.Duration(timeout),
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
	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func (c *Client) GetTimeEntriesGroupByProject(ctx context.Context) (*User, error) {
	req, err := c.newRequest(ctx, http.MethodGet, currentUserWithRelatedDataURI, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// TODO: check status code

	var user User
	if err := decodeBody(res, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
