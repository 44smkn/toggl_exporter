package exporter

import (
	"context"
	"strconv"
	"time"

	"github.com/go-kit/log/level"
)

type ProjectTimeDuration struct {
	Pid         int
	ProjectName string
	Duration    time.Duration
	Year        string
	Month       string
}

func (e *Exporter) GetProjectDurations(ctx context.Context) ([]ProjectTimeDuration, error) {
	timeEntries, err := e.GetTimeEntries(ctx)
	if err != nil {
		return nil, err
	}

	m := make(map[int]time.Duration)
	for _, e := range timeEntries {
		m[e.Pid] += e.Duration
	}

	now := time.Now().UTC()
	year := strconv.Itoa(now.Year())
	month := now.Month().String()

	ptj := make([]ProjectTimeDuration, 0, len(m))
	for pid, duration := range m {
		if pid == 0 {
			continue
		}
		project, err := e.GetProject(ctx, strconv.Itoa(pid))
		if err != nil {
			level.Info(e.Logger).Log("msg", err.Error(), "pid", pid)
			continue
		}
		p := ProjectTimeDuration{
			Pid:         pid,
			ProjectName: project.Name,
			Duration:    duration,
			Year:        year,
			Month:       month,
		}
		ptj = append(ptj, p)
	}
	return ptj, nil
}
