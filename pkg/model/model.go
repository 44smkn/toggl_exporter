package model

import (
	"context"
	"time"
)

// TimeEntry represents individual units of time in Toggl Track.
// Unused properties are removed and some of needed properties are made use easily.
type TimeEntry struct {
	Pid      int
	Duration time.Duration
}

// Project represents project in Toggl Track, which is behave as attribute of the time entry.
// Unused properties are removed and some of needed properties are made use easily.
type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TimeEntryRepository interface is the implementation required of time entry.
type TimeEntryRepository interface {
	GetTimeEntries(ctx context.Context) ([]TimeEntry, error)
}

// TimeEntryRepository interface is the implementation required of project.
type ProjectRepository interface {
	GetProject(ctx context.Context, pid string) (*Project, error)
}
