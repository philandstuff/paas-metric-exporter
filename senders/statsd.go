package senders

import (
	quipo_statsd "github.com/quipo/statsd"
	"github.com/alphagov/paas-metric-exporter/metrics"
	"github.com/alphagov/paas-metric-exporter/presenters"
)

type StatsdSender struct {
	Client *quipo_statsd.StatsdClient
    Template string
}

var _ metrics.Sender = StatsdSender{}

func NewStatsdSender(statsdEndpoint string, statsdPrefix string, template string) StatsdSender {
	return StatsdSender {
        Client: quipo_statsd.NewStatsdClient(statsdEndpoint, statsdPrefix),
        Template: template,
    }
}

func (s StatsdSender) Start() {
	s.Client.CreateSocket()
}

func (s StatsdSender) Gauge(metric metrics.GaugeMetric) error {
    presenter := presenters.PathPresenter{ Template: s.Template }
	stat, _ := presenter.Present(metric)

	return s.Client.Gauge(stat, metric.Value)
}

func (s StatsdSender) FGauge(metric metrics.FGaugeMetric) error {
    presenter := presenters.PathPresenter{ Template: s.Template }
	stat, _ := presenter.Present(metric)

	return s.Client.FGauge(stat, metric.Value)
}

func (s StatsdSender) Incr(metric metrics.CounterMetric) error {
    presenter := presenters.PathPresenter{ Template: s.Template }
	stat, _ := presenter.Present(metric)

	return s.Client.Incr(stat, metric.Value)
}

func (s StatsdSender) Timing(metric metrics.TimingMetric) error {
    presenter := presenters.PathPresenter{ Template: s.Template }
	stat, _ := presenter.Present(metric)

	return s.Client.Timing(stat, metric.Value)
}

func (s StatsdSender) PrecisionTiming(metric metrics.PrecisionTimingMetric) error {
    presenter := presenters.PathPresenter{ Template: s.Template }
	stat, _ := presenter.Present(metric)

	return s.Client.PrecisionTiming(stat, metric.Value)
}
