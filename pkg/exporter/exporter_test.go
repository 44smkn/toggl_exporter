package exporter_test

import (
	"context"
	"testing"

	"github.com/44smkn/toggl_exporter/pkg/exporter"
	"github.com/44smkn/toggl_exporter/pkg/model"
	"github.com/go-kit/kit/log"
	"github.com/prometheus/common/promlog"
)

type FakeExporter struct {
	Logger log.Logger

	model.TimeEntryRepository
	model.ProjectRepository
}

func NewExporter(t *testing.T, tmr model.TimeEntryRepository, pr model.ProjectRepository) *exporter.Exporter {
	t.Helper()
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)

	return &exporter.Exporter{
		Logger:              logger,
		TimeEntryRepository: tmr,
		ProjectRepository:   pr,
	}
}

type FakeTimeEntryRepository struct {
	entriesFn func() []model.TimeEntry
}

type FakeProjectRepository struct {
	projectFn func(pid string) *model.Project
}

func NewFakeTimeEntryRepository(fn func() []model.TimeEntry) *FakeTimeEntryRepository {
	return &FakeTimeEntryRepository{
		entriesFn: fn,
	}
}

func NewFakeProjectRepository(fn func(pid string) *model.Project) *FakeProjectRepository {
	return &FakeProjectRepository{
		projectFn: fn,
	}
}

func (r FakeTimeEntryRepository) GetTimeEntries(ctx context.Context) ([]model.TimeEntry, error) {
	return r.entriesFn(), nil
}

func (r FakeProjectRepository) GetProject(ctx context.Context, pid string) (*model.Project, error) {
	return r.projectFn(pid), nil
}
