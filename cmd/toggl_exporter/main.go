package main

import (
	"os"
	"time"

	"github.com/44smkn/toggl_exporter/pkg/exporter"
	"github.com/44smkn/toggl_exporter/pkg/toggl"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/common/promlog/flag"
	"github.com/prometheus/common/version"
	webflag "github.com/prometheus/exporter-toolkit/web/kingpinflag"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		webConfig     = webflag.AddFlags(kingpin.CommandLine)
		listenAddress = kingpin.Flag("web.listen-address", "Address to listen on for web interface and telemetry.").Default(":9981").Envar("WEB_LISTEN_ADDRESS").String()
		metricsPath   = kingpin.Flag("web.telemetry-path", "Path under which to expose metrics.").Default("/metrics").Envar("WEB_TELEMETRY_PATH").String()
		togglAPIKey   = kingpin.Flag("toggl.api-key", "write later...").Envar("TOGGL_API_KEY").String()
		togglTimeout  = kingpin.Flag("toggl.req-timeout-seconds", "Timeout for trying").Envar("TOGGL_REQ_TIMEOUT_SECONDS").Duration()
	)

	promlogConfig := &promlog.Config{}
	flag.AddFlags(kingpin.CommandLine, promlogConfig)
	kingpin.Version(version.Print("toggl_exporter"))
	kingpin.HelpFlag.Short('h')
	kingpin.Parse()
	logger := promlog.New(promlogConfig)

	level.Info(logger).Log("msg", "Starting toggl_exporter", "version", version.Info())
	level.Info(logger).Log("msg", "Build context", "context", version.BuildContext())

	exporter := NewExporter(*togglAPIKey, *togglTimeout, logger)
	if err := exporter.ListenAndServe(*listenAddress, *webConfig, *metricsPath); err != nil {
		level.Error(logger).Log("msg", "Error starting HTTP server", "err", err)
		os.Exit(1)
	}
}

func NewExporter(togglAPIKey string, togglTimeout time.Duration, logger log.Logger) *exporter.Exporter {
	client := toggl.NewClient(togglAPIKey, togglTimeout)
	timeEntryRepository := &toggl.TimeEntryRepository{Client: client}
	projectRepository := &toggl.ProjectRepository{Client: client}

	return &exporter.Exporter{
		Logger: logger,

		TimeEntryRepository: timeEntryRepository,
		ProjectRepository:   projectRepository,
	}
}
