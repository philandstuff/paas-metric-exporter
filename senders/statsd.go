package senders

import (
	"time"

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

func (s StatsdSender) Gauge2(metric metrics.GaugeMetric) error {
    presenter := presenters.PathPresenter{ Template: s.Template }
	stat, _ := presenter.Present(metric)

	return s.Client.Gauge(stat, metric.Value)
}

func (s StatsdSender) Gauge(stat string, value int64) error {
	return s.Client.Gauge(stat, value)
}

func (s StatsdSender) FGauge2(metric metrics.FGaugeMetric) error {
    presenter := presenters.PathPresenter{ Template: s.Template }
	stat, _ := presenter.Present(metric)

	return s.Client.FGauge(stat, metric.Value)
}

func (s StatsdSender) FGauge(stat string, value float64) error {
	return s.Client.FGauge(stat, value)
}

func (s StatsdSender) Incr2(metric metrics.CounterMetric) error {
    presenter := presenters.PathPresenter{ Template: s.Template }
	stat, _ := presenter.Present(metric)

	return s.Client.Incr(stat, metric.Value)
}

func (s StatsdSender) Incr(stat string, count int64) error {
	return s.Client.Incr(stat, count)
}

func (s StatsdSender) Timing2(metric metrics.TimingMetric) error {
    presenter := presenters.PathPresenter{ Template: s.Template }
	stat, _ := presenter.Present(metric)

	return s.Client.Timing(stat, metric.Value)
}

func (s StatsdSender) Timing(stat string, delta int64) error {
	return s.Client.Timing(stat, delta)
}

func (s StatsdSender) PrecisionTiming2(metric metrics.PrecisionTimingMetric) error {
    presenter := presenters.PathPresenter{ Template: s.Template }
	stat, _ := presenter.Present(metric)

	return s.Client.PrecisionTiming(stat, metric.Value)
}

func (s StatsdSender) PrecisionTiming(stat string, delta time.Duration) error {
	return s.Client.PrecisionTiming(stat, delta)
}
