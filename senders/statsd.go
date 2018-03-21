package senders

import (
	quipo_statsd "github.com/quipo/statsd"
	"github.com/alphagov/paas-metric-exporter/metrics"
	"github.com/alphagov/paas-metric-exporter/presenters"
)

type StatsdSender struct {
	Client quipo_statsd.Statsd
    presenter presenters.PathPresenter
}

var _ metrics.Sender = StatsdSender{}

const DefaultTemplate =
    "{{.Space}}.{{.App}}.{{.Instance}}.{{.Metric}}{{if .Metadata.statusRange}}.{{.Metadata.statusRange}}{{end}}"

func NewStatsdSender(client quipo_statsd.Statsd, template string) (StatsdSender, error) {
    presenter, err := presenters.NewPathPresenter(template)
    sender := StatsdSender { Client: client, presenter: presenter }

    return sender, err
}

func (s StatsdSender) Gauge(metric metrics.GaugeMetric) error {
	stat, err := s.presenter.Present(metric)
    if err != nil {
        return err
    }

	return s.Client.Gauge(stat, metric.Value)
}

func (s StatsdSender) Incr(metric metrics.CounterMetric) error {
	stat, err := s.presenter.Present(metric)
    if err != nil {
        return err
    }

	return s.Client.Incr(stat, metric.Value)
}

func (s StatsdSender) PrecisionTiming(metric metrics.PrecisionTimingMetric) error {
	stat, err := s.presenter.Present(metric)
    if err != nil {
        return err
    }

	return s.Client.PrecisionTiming(stat, metric.Value)
}
