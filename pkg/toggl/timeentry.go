package toggl

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/44smkn/toggl_exporter/pkg/model"
)

const (
	timeEntriesURI = "/time_entries"
)

type TimeEntry struct {
	ID          int           `json:"id"`
	Wid         int           `json:"wid"`
	Pid         int           `json:"pid,omitempty"`
	Billable    bool          `json:"billable"`
	Start       time.Time     `json:"start"`
	Stop        time.Time     `json:"stop"`
	Duration    time.Duration `json:"duration"`
	Description string        `json:"description"`
	Tags        []string      `json:"tags"`
	At          time.Time     `json:"at"`
}

type TimeEntryRepository struct {
	*Client
}

func (r *TimeEntryRepository) GetTimeEntries(ctx context.Context) ([]model.TimeEntry, error) {
	req, err := r.newRequest(ctx, http.MethodGet, timeEntriesURI, nil)
	if err != nil {
		return nil, err
	}

	res, err := r.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusForbidden:
		return nil, errors.New(fmt.Sprintf("APIKey may be not valid. status is %v", res.Status))
	}

	var rawTimeEntries []TimeEntry
	if err := decodeBody(res, &rawTimeEntries); err != nil {
		return nil, err
	}

	timeEntries := make([]model.TimeEntry, 0, len(rawTimeEntries))
	for _, re := range rawTimeEntries {
		duration := re.Duration * time.Second
		if duration < 0 {
			duration = time.Now().Sub(re.Start)
		}

		e := model.TimeEntry{
			Pid:      re.Pid,
			Duration: duration,
		}
		timeEntries = append(timeEntries, e)
	}

	return timeEntries, nil
}
