package senders

import (
	"log"

	"github.com/alphagov/paas-metric-exporter/metrics"
	"github.com/alphagov/paas-metric-exporter/presenters"
)

type DebugClient struct {
	Prefix string
    Template string
}

var _ metrics.Sender = DebugClient{}

func (d DebugClient) Gauge(metric metrics.GaugeMetric) error {
    presenter := presenters.PathPresenter{ Template: d.Template }
	stat, _ := presenter.Present(metric)

	log.Printf("gauge %s %d\n", d.Prefix+stat, metric.Value)
	return nil
}

func (d DebugClient) FGauge(metric metrics.FGaugeMetric) error {
    presenter := presenters.PathPresenter{ Template: d.Template }
	stat, _ := presenter.Present(metric)

	log.Printf("gauge %s %d\n", d.Prefix+stat, metric.Value)
	return nil
}

func (d DebugClient) Incr(metric metrics.CounterMetric) error {
    presenter := presenters.PathPresenter{ Template: d.Template }
	stat, _ := presenter.Present(metric)

	log.Printf("incr %s %d\n", d.Prefix+stat, metric.Value)
	return nil
}

func (d DebugClient) Timing(metric metrics.TimingMetric) error {
    presenter := presenters.PathPresenter{ Template: d.Template }
	stat, _ := presenter.Present(metric)

	log.Printf("timing %s %d\n", d.Prefix+stat, metric.Value)
	return nil
}

func (d DebugClient) PrecisionTiming(metric metrics.PrecisionTimingMetric) error {
    presenter := presenters.PathPresenter{ Template: d.Template }
	stat, _ := presenter.Present(metric)

	log.Printf("timing %s %d\n", d.Prefix+stat, metric.Value)
	return nil
}
