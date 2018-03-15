package senders

import (
	quipo_statsd "github.com/quipo/statsd"
	"github.com/alphagov/paas-metric-exporter/metrics"
	"github.com/alphagov/paas-metric-exporter/presenters"
)

type StatsdSender struct {
	Client *quipo_statsd.StatsdClient
    presenter presenters.PathPresenter
}

var _ metrics.Sender = StatsdSender{}

func NewStatsdSender(statsdEndpoint string, statsdPrefix string, template string) StatsdSender {
	return StatsdSender {
        Client: quipo_statsd.NewStatsdClient(statsdEndpoint, statsdPrefix),
        presenter: presenters.PathPresenter{ Template: template },
    }
}

func (s StatsdSender) Start() {
	s.Client.CreateSocket()
}

func (s StatsdSender) Gauge(metric metrics.GaugeMetric) error {
	stat, _ := s.presenter.Present(metric)

	return s.Client.Gauge(stat, metric.Value)
}

func (s StatsdSender) FGauge(metric metrics.FGaugeMetric) error {
	stat, _ := s.presenter.Present(metric)

	return s.Client.FGauge(stat, metric.Value)
}

func (s StatsdSender) Incr(metric metrics.CounterMetric) error {
	stat, _ := s.presenter.Present(metric)

	return s.Client.Incr(stat, metric.Value)
}

func (s StatsdSender) Timing(metric metrics.TimingMetric) error {
	stat, _ := s.presenter.Present(metric)

	return s.Client.Timing(stat, metric.Value)
}

func (s StatsdSender) PrecisionTiming(metric metrics.PrecisionTimingMetric) error {
	stat, _ := s.presenter.Present(metric)

	return s.Client.PrecisionTiming(stat, metric.Value)
}
