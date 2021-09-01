package exporter

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/go-kit/log/level"
)

type ProjectTimeDuration struct {
	Pid         int
	ProjectName string
	Duration    time.Duration
	YearMonth   string
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
	yearMonth := fmt.Sprintf("%v/%d", now.Year(), now.Month())
	ptj := make([]ProjectTimeDuration, 0, len(m))
	for pid, duration := range m {
		project, err := e.GetProject(ctx, strconv.Itoa(pid))
		if err != nil {
			level.Info(e.Logger).Log("msg", err.Error(), "pid", pid)
			continue
		}
		p := ProjectTimeDuration{
			Pid:         pid,
			ProjectName: project.Name,
			Duration:    duration,
			YearMonth:   yearMonth,
		}
		ptj = append(ptj, p)
	}
	return ptj, nil
}
