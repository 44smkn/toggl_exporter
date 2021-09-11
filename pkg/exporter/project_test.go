package exporter_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/44smkn/toggl_exporter/pkg/exporter"
	"github.com/44smkn/toggl_exporter/pkg/model"
	"github.com/google/go-cmp/cmp"
)

func TestGetProjectDurations(t *testing.T) {
	t.Parallel()

	timeEntriesFn := func() []model.TimeEntry {
		return []model.TimeEntry{
			{
				Pid:      99,
				Duration: 50,
			},
			{
				Pid:      199,
				Duration: 20,
			},
			{
				Pid:      99,
				Duration: 50,
			},
		}
	}

	ProjectFn := func(pid string) *model.Project {
		switch pid {
		case "99":
			return &model.Project{
				ID:   99,
				Name: "Reading Books",
			}
		case "199":
			return &model.Project{
				ID:   199,
				Name: "Sleep",
			}
		}
		return nil
	}

	now := time.Now().UTC()
	want := []exporter.ProjectTimeDuration{
		{
			Pid:         99,
			ProjectName: "Reading Books",
			Duration:    100,
			Year:        strconv.Itoa(now.Year()),
			Month:       now.Month().String(),
		},
		{
			Pid:         199,
			ProjectName: "Sleep",
			Duration:    20,
			Year:        strconv.Itoa(now.Year()),
			Month:       now.Month().String(),
		},
	}

	exporter := NewExporter(t, NewFakeTimeEntryRepository(timeEntriesFn), NewFakeProjectRepository(ProjectFn))

	ctx := context.Background()
	got, err := exporter.GetProjectDurations(ctx)
	if err != nil {
		t.Errorf(err.Error())
	}

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("GetProjectDurations mismatch (-want +got):\n%s", diff)
	}
}
