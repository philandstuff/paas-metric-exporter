package senders

import (
	"time"

	"github.com/fatih/structs"

	"github.com/alphagov/paas-metric-exporter/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

type PrometheusSender struct {
	counterVecs   map[string]prometheus.CounterVec
	gaugeVecs     map[string]prometheus.GaugeVec
	histogramVecs map[string]prometheus.HistogramVec
}

var _ metrics.Sender = &PrometheusSender{}

func NewPrometheusSender() *PrometheusSender {
	counterVecs := make(map[string]prometheus.CounterVec)
	gaugeVecs := make(map[string]prometheus.GaugeVec)
	histogramVecs := make(map[string]prometheus.HistogramVec)

	return &PrometheusSender{
		counterVecs,
		gaugeVecs,
		histogramVecs,
	}
}

func (s *PrometheusSender) Gauge(metric metrics.GaugeMetric) error {
	gaugeVec, present := s.gaugeVecs[metric.Name()]
	labelNames := buildLabelsFromMetric(metric)

	if !present {
		options := prometheus.GaugeOpts{Name: metric.Name(), Help: " "}
		gaugeVec = *prometheus.NewGaugeVec(options, labelNames)

		prometheus.MustRegister(gaugeVec)
		s.gaugeVecs[metric.Name()] = gaugeVec
	}

	labels := s.labels(metric, labelNames)
	value := float64(metric.Value)

	gaugeVec.With(labels).Set(value)

	return nil
}

func (s *PrometheusSender) Incr(metric metrics.CounterMetric) error {
	counterVec, present := s.counterVecs[metric.Name()]
	labelNames := buildLabelsFromMetric(metric)

	if !present {
		options := prometheus.CounterOpts{Name: metric.Name(), Help: " "}
		counterVec = *prometheus.NewCounterVec(options, labelNames)

		prometheus.MustRegister(counterVec)
		s.counterVecs[metric.Name()] = counterVec
	}

	labels := s.labels(metric, labelNames)
	value := float64(metric.Value)

	counterVec.With(labels).Add(value)

	return nil
}

func (s *PrometheusSender) PrecisionTiming(metric metrics.PrecisionTimingMetric) error {
	histogramVec, present := s.histogramVecs[metric.Name()]
	labelNames := buildLabelsFromMetric(metric)

	if !present {
		options := prometheus.HistogramOpts{Name: metric.Name(), Help: " "}
		histogramVec = *prometheus.NewHistogramVec(options, labelNames)

		prometheus.MustRegister(histogramVec)
		s.histogramVecs[metric.Name()] = histogramVec
	}

	labels := s.labels(metric, labelNames)
	value := float64(metric.Value) / float64(time.Second)

	histogramVec.With(labels).Observe(value)

	return nil
}

func (s *PrometheusSender) labels(metric metrics.Metric, labelNames []string) prometheus.Labels {
	labels := make(prometheus.Labels)
	fields := structs.Map(metric)

	for mk, mv := range metric.GetMetadata() {
		fields[mk] = mv
	}

	for k, v := range fields {
		for _, n := range labelNames {
			if k == n {
				labels[k] = v.(string)
			}
		}
	}

	return labels
}
func buildLabelsFromMetric(metric metrics.Metric) (labelNames []string) {
	labelNames = append(labelNames,
		"App",
		"CellId",
		"GUID",
		"Instance",
		"Job",
		"Organisation",
		"Space",
	)

	for k := range metric.GetMetadata() {
		labelNames = append(labelNames, k)
	}

	return labelNames
}
