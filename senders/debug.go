package senders

import (
	"log"

	"github.com/alphagov/paas-metric-exporter/metrics"
	"github.com/alphagov/paas-metric-exporter/presenters"
)

type DebugSender struct {
	Prefix string
    presenter presenters.PathPresenter
}

var _ metrics.Sender = DebugSender{}

func NewDebugSender(statsdPrefix string, template string) (DebugSender, error) {
    presenter, err := presenters.NewPathPresenter(template)
    sender := DebugSender { Prefix: statsdPrefix, presenter: presenter }

	return sender, err
}

func (d DebugSender) Gauge(metric metrics.GaugeMetric) error {
	stat, _ := d.presenter.Present(metric)

	log.Printf("gauge %s %d\n", d.Prefix+stat, metric.Value)
	return nil
}

func (d DebugSender) FGauge(metric metrics.FGaugeMetric) error {
	stat, _ := d.presenter.Present(metric)

	log.Printf("gauge %s %d\n", d.Prefix+stat, metric.Value)
	return nil
}

func (d DebugSender) Incr(metric metrics.CounterMetric) error {
	stat, _ := d.presenter.Present(metric)

	log.Printf("incr %s %d\n", d.Prefix+stat, metric.Value)
	return nil
}

func (d DebugSender) Timing(metric metrics.TimingMetric) error {
	stat, _ := d.presenter.Present(metric)

	log.Printf("timing %s %d\n", d.Prefix+stat, metric.Value)
	return nil
}

func (d DebugSender) PrecisionTiming(metric metrics.PrecisionTimingMetric) error {
	stat, _ := d.presenter.Present(metric)

	log.Printf("timing %s %d\n", d.Prefix+stat, metric.Value)
	return nil
}
