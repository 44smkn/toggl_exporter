package exporter

import (
	"sync"

	"github.com/44smkn/toggl_exporter/pkg/model"
	"github.com/go-kit/kit/log"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace = "toggl"
)

var (
	timeEntries = prometheus.NewDesc(prometheus.BuildFQName(namespace, "time_entries", "seconds"), "total time for time entiries", []string{"project_name"}, nil)
)

type Exporter struct {
	mutex  sync.RWMutex
	Logger log.Logger

	model.TimeEntryRepository
	model.ProjectRepository
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- timeEntries
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	// call
	ch <- prometheus.MustNewConstMetric(timeEntries, prometheus.CounterValue, 0)
}
