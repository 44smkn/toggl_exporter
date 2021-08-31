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

	ptj := make([]ProjectTimeDuration, 0, len(m))
	for pid, duration := range m {
		project, err := e.GetProject(ctx, strconv.Itoa(pid))
		if err != nil {
			level.Info(e.Logger).Log("err", err.Error())
			continue
		}
		p := ProjectTimeDuration{
			Pid:         pid,
			ProjectName: project.Name,
			Duration:    duration,
		}
		ptj = append(ptj, p)
	}
	return ptj, nil
}