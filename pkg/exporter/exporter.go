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
	projectDuration = prometheus.NewDesc(prometheus.BuildFQName(namespace, "project_duration", "seconds"), "total time for time entiries", []string{"project_name"}, nil)
)

type Exporter struct {
	mutex  sync.RWMutex
	Logger log.Logger

	model.TimeEntryRepository
	model.ProjectRepository
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- projectDuration
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	// call
	ch <- prometheus.MustNewConstMetric(projectDuration, prometheus.CounterValue, 0)
}
