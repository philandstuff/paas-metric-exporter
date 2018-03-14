package metrics

import (
	"time"

	"github.com/alphagov/paas-metric-exporter/senders"
)

var _ Metric = CounterMetric{}
var _ Metric = GaugeMetric{}
var _ Metric = FGaugeMetric{}
var _ Metric = TimingMetric{}
var _ Metric = PrecisionTimingMetric{}

//go:generate counterfeiter -o mocks/metric.go . Metric
type Metric interface {
	Send(sender senders.Sender, template string) error
	Name() string
}

type CounterMetric struct {
	App          string
	CellId       string
	GUID         string
	Instance     string
	Job          string
	Metric       string
	Organisation string
	Space        string

	Value int64
}

func (m CounterMetric) Name() string {
	return m.Metric
}

func (m CounterMetric) Send(sender senders.Sender, tmpl string) error {
	tmplName, err := render(tmpl, m)
	if err != nil {
		return err
	}

	return sender.Incr(tmplName, m.Value)
}

type GaugeMetric struct {
	App          string
	CellId       string
	GUID         string
	Instance     string
	Job          string
	Metric       string
	Organisation string
	Space        string

	Value int64
}

func (m GaugeMetric) Name() string {
	return m.Metric
}

func (m GaugeMetric) Send(sender senders.Sender, tmpl string) error {
	tmplName, err := render(tmpl, m)
	if err != nil {
		return err
	}

	return sender.Gauge(tmplName, m.Value)
}

type FGaugeMetric struct {
	App          string
	CellId       string
	GUID         string
	Instance     string
	Job          string
	Metric       string
	Organisation string
	Space        string

	Value float64
}

func (m FGaugeMetric) Name() string {
	return m.Metric
}

func (m FGaugeMetric) Send(sender senders.Sender, tmpl string) error {
	tmplName, err := render(tmpl, m)
	if err != nil {
		return err
	}

	return sender.FGauge(tmplName, m.Value)
}

type TimingMetric struct {
	App          string
	CellId       string
	GUID         string
	Instance     string
	Job          string
	Metric       string
	Organisation string
	Space        string

	Value int64
}

func (m TimingMetric) Name() string {
	return m.Metric
}

func (m TimingMetric) Send(sender senders.Sender, tmpl string) error {
	tmplName, err := render(tmpl, m)
	if err != nil {
		return err
	}

	return sender.Timing(tmplName, m.Value)
}

type PrecisionTimingMetric struct {
	App          string
	CellId       string
	GUID         string
	Instance     string
	Job          string
	Metric       string
	Organisation string
	Space        string

	Value time.Duration
}

func (m PrecisionTimingMetric) Name() string {
	return m.Metric
}

func (m PrecisionTimingMetric) Send(sender senders.Sender, tmpl string) error {
	tmplName, err := render(tmpl, m)
	if err != nil {
		return err
	}

	return sender.PrecisionTiming(tmplName, m.Value)
}
