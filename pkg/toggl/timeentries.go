package toggl

import "time"

type TimeEntries []struct {
	ID          int       `json:"id"`
	Wid         int       `json:"wid"`
	Pid         int       `json:"pid,omitempty"`
	Billable    bool      `json:"billable"`
	Start       time.Time `json:"start"`
	Stop        time.Time `json:"stop"`
	Duration    int       `json:"duration"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	At          time.Time `json:"at"`
}
