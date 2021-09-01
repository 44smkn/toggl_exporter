package main

import (
	"os"

	"github.com/44smkn/toggl_exporter/pkg/config"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/version"
)

func main() {
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	exporter := config.InitExporter(promlogConfig, logger)
	prometheus.MustRegister(exporter)
	prometheus.MustRegister(version.NewCollector("toggl_exporter"))
	if err := exporter.ListenAndServe(); err != nil {
		level.Error(logger).Log("msg", "Error starting HTTP server", "err", err)
		os.Exit(1)
	}
}
