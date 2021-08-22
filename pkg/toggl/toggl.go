package toggl

import (
	"net/http"
	"time"
)

const (
	togglBaseURL = "api.track.toggl.com"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string, timeout int) *Client {
	httpClient := &http.Client{
		Timeout: time.Duration(timeout),
	}
	return &Client{
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}
