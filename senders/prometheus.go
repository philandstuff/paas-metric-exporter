package senders

import (
    "time"

    "github.com/fatih/structs"

	"github.com/alphagov/paas-metric-exporter/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusSender struct {
    labelNames []string

    counterVecs map[string]prometheus.CounterVec
    gaugeVecs map[string]prometheus.GaugeVec
    histogramVecs map[string]prometheus.HistogramVec
}

var _ metrics.Sender = &PrometheusSender{}

func NewPrometheusSender() *PrometheusSender {
    labelNames := []string{
        "App",
        "CellId",
        "GUID",
        "Instance",
        "Job",
        "Organisation",
        "Space",
    }

    counterVecs := make(map[string]prometheus.CounterVec);
    gaugeVecs := make(map[string]prometheus.GaugeVec);
    histogramVecs := make(map[string]prometheus.HistogramVec);

    return &PrometheusSender{
        labelNames,
        counterVecs,
        gaugeVecs,
        histogramVecs,
    }
}

func (s *PrometheusSender) Gauge(metric metrics.GaugeMetric) error {
    gaugeVec, present := s.gaugeVecs[metric.Name()]

    if !present {
        options := prometheus.GaugeOpts{ Name: metric.Name(), Help: " " }
        gaugeVec = *prometheus.NewGaugeVec(options, s.labelNames)

        prometheus.MustRegister(gaugeVec)
        s.gaugeVecs[metric.Name()] = gaugeVec
    }

    labels := s.labels(metric)
    value := float64(metric.Value)

    gaugeVec.With(labels).Set(value)

    return nil
}

func (s *PrometheusSender) Incr(metric metrics.CounterMetric) error {
    counterVec, present := s.counterVecs[metric.Name()]

    if !present {
        options := prometheus.CounterOpts{ Name: metric.Name(), Help: " " }
        counterVec = *prometheus.NewCounterVec(options, s.labelNames)

        prometheus.MustRegister(counterVec)
        s.counterVecs[metric.Name()] = counterVec
    }

    labels := s.labels(metric)
    value := float64(metric.Value)

    counterVec.With(labels).Add(value)

    return nil
}

func (s *PrometheusSender) PrecisionTiming(metric metrics.PrecisionTimingMetric) error {
    histogramVec, present := s.histogramVecs[metric.Name()]

    if !present {
        options := prometheus.HistogramOpts{ Name: metric.Name(), Help: " " }
        histogramVec = *prometheus.NewHistogramVec(options, s.labelNames)

        prometheus.MustRegister(histogramVec)
        s.histogramVecs[metric.Name()] = histogramVec
    }

    labels := s.labels(metric)
    value := float64(metric.Value) / float64(time.Second)

    histogramVec.With(labels).Observe(value)

    return nil
}

func (s *PrometheusSender) labels(metric metrics.Metric) prometheus.Labels {
    labels := make(prometheus.Labels)
    fields := structs.Map(metric)

    for k, v := range fields {
        for _, n := range s.labelNames {
            if k == n {
                labels[k] = v.(string)
            }
        }
    }

    return labels
}
