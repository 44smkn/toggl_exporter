package model

import (
	"context"
	"time"
)

type TimeEntry struct {
	Pid      int
	Duration time.Duration
}

type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TimeEntryRepository interface {
	GetTimeEntries(ctx context.Context) ([]TimeEntry, error)
}

type ProjectRepository interface {
	GetProject(ctx context.Context, pid string) (*Project, error)
}
