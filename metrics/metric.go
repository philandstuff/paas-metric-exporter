package metrics

import (
	"time"

	"github.com/alphagov/paas-metric-exporter/presenters"
)

func render(template string, m Metric) (string, error) {
    presenter := presenters.PathPresenter { Template: template }
    return presenter.Present(m)
}

var _ Metric = CounterMetric{}
var _ Metric = GaugeMetric{}
var _ Metric = FGaugeMetric{}
var _ Metric = TimingMetric{}
var _ Metric = PrecisionTimingMetric{}

//go:generate counterfeiter -o mocks/sender.go . Sender
type Sender interface {
	Gauge2(metric GaugeMetric) error
	FGauge2(metric FGaugeMetric) error
	Incr2(metric CounterMetric) error
	Timing2(metric TimingMetric) error
	PrecisionTiming2(metric PrecisionTimingMetric) error
}

//go:generate counterfeiter -o mocks/metric.go . Metric
type Metric interface {
	Send(sender Sender, template string) error
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

func (m CounterMetric) Send(sender Sender, tmpl string) error {
	//tmplName, err := render(tmpl, m)
	//if err != nil {
	//	return err
	//}

	return sender.Incr2(m)
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

func (m GaugeMetric) Send(sender Sender, tmpl string) error {
	//tmplName, err := render(tmpl, m)
	//if err != nil {
	//	return err
	//}

	return sender.Gauge2(m)
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

func (m FGaugeMetric) Send(sender Sender, tmpl string) error {
	//tmplName, err := render(tmpl, m)
	//if err != nil {
	//	return err
	//}

	return sender.FGauge2(m)
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

func (m TimingMetric) Send(sender Sender, tmpl string) error {
	//tmplName, err := render(tmpl, m)
	//if err != nil {
	//	return err
	//}

	return sender.Timing2(m)
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

func (m PrecisionTimingMetric) Send(sender Sender, tmpl string) error {
	//tmplName, err := render(tmpl, m)
	//if err != nil {
	//	return err
	//}

	return sender.PrecisionTiming2(m)
}
