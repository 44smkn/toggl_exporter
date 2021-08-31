package toggl_test

import (
	"context"
	"testing"
	"time"

	"github.com/44smkn/toggl_exporter/pkg/toggl"
)

func TestGetTimeEntriesGroupByProject(t *testing.T) {
	// apiKey := os.Getenv("TOGGL_API_KEY")
	apiKey := "f147b98528581098b2e6832d5ee1358f"
	c := toggl.NewClient(apiKey, 30*time.Second)
	ctx := context.Background()
	timeEntries, err := c.GetTimeEntries(ctx)
	if err != nil {
		t.Errorf("failed to request: %v", err)
	}
	t.Errorf("%+v", timeEntries)
}

func TestGetProject(t *testing.T) {
	// apiKey := os.Getenv("TOGGL_API_KEY")
	apiKey := "f147b98528581098b2e6832d5ee1358f"
	c := toggl.NewClient(apiKey, 30*time.Second)
	ctx := context.Background()
	project, err := c.GetProject(ctx, "150667633")
	if err != nil {
		t.Errorf("failed to request: %v", err)
	}
	t.Errorf("%+v", project)
}
