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
	togglAPIBaseURL = "https://api.track.toggl.com/api/v8"
)

// Client facilitates making HTTP requests to the Toggl API
type Client struct {
	apiKey     string
	URL        *url.URL
	httpClient *http.Client
}

// NewClient initializes a Client
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

func (c *Client) newRequest(ctx context.Context, method, spath string, rawQuery *string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)
	if rawQuery != nil {
		u.RawQuery = *rawQuery
	}
	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.apiKey, "api_token")
	return req, nil
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}
