package toggl

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/44smkn/toggl_exporter/pkg/model"
)

const (
	timeEntriesURI = "/time_entries"
	ISO8601        = "2006-01-02T15:04:05-07:00"
)

// TimeEntry is the object representing the time entry in toggl world.
// It is bound with Get time_entries API Response of toggl.
// See: https://github.com/toggl/toggl_api_docs/blob/master/chapters/time_entries.md#get-time-entry-details
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

// GetProject returns array of objects representing time entries in toggl_exporter.
// It retrive time entries data bound with api key and create array of peculiar time entry object.
func (r *TimeEntryRepository) GetTimeEntries(ctx context.Context) ([]model.TimeEntry, error) {
	query := fmt.Sprintf("start_date=%s", getBeginningOfMonthQueryParam(time.Now().UTC()))
	req, err := r.newRequest(ctx, http.MethodGet, timeEntriesURI, &query, nil)
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
	case http.StatusNotFound:
		return nil, errors.New(fmt.Sprintf("url parameter may be not valid. status is %v", res.Status))
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

func getBeginningOfMonthQueryParam(now time.Time) string {
	return url.QueryEscape(time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC).Format(ISO8601))
}
