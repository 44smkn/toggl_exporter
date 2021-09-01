package exporter

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/44smkn/toggl_exporter/pkg/model"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/prometheus/exporter-toolkit/web"
)

const (
	namespace = "toggl"
)

var (
	projectDuration = prometheus.NewDesc(prometheus.BuildFQName(namespace, "project_duration", "seconds"), "total time for time entiries by project", []string{"project_name", "month"}, nil)
)

type Exporter struct {
	WebConfig     string
	ListenAddress string
	MetricsPath   string

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

	ctx := context.Background()
	pds, err := e.GetProjectDurations(ctx)
	if err != nil {
		level.Error(e.Logger).Log("msg", fmt.Sprintf("failed to get project durations: %v", err))
		return
	}
	for _, pd := range pds {
		ch <- prometheus.MustNewConstMetric(projectDuration, prometheus.CounterValue, pd.Duration.Seconds(), pd.ProjectName, pd.YearMonth)
	}
}

func (e *Exporter) ListenAndServe() error {
	prometheus.MustRegister(e)
	prometheus.MustRegister(version.NewCollector("toggl_exporter"))

	level.Info(e.Logger).Log("msg", "Listening on address", "address", e.ListenAddress)
	http.Handle(e.MetricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
<head><title>Toggl Exporter</title></head>
<body>
<h1>Toggl Exporter</h1>
<p><a href='` + e.MetricsPath + `'>Metrics</a></p>
</body>
</html>`))
	})

	srv := &http.Server{Addr: e.ListenAddress}
	if err := web.ListenAndServe(srv, e.WebConfig, e.Logger); err != nil {
		return err
	}
	return nil
}
