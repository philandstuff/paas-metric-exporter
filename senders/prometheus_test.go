package senders_test

import (
    "time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_model/go"

	. "github.com/alphagov/paas-metric-exporter/metrics"
	. "github.com/alphagov/paas-metric-exporter/senders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)

var _ = Describe("PrometheusSender", func() {
    sender := NewPrometheusSender()

	Describe("#Incr", func() {
        It("sends a counter metric to prometheus", func() {
            families := captureMetrics(func () {
                sender.Incr(CounterMetric{
                    Metric: "counter_incremented_once",
                    Value: 1,
                    App: "some_value",
                })
            })

            family := families[0]
            metrics := family.GetMetric()
            metric := metrics[0].Counter
            labels := metrics[0].GetLabel()

            Expect(len(families)).To(Equal(1))
            Expect(len(metrics)).To(Equal(1))

            Expect(family.GetName()).To(Equal("counter_incremented_once"))
            Expect(metric.GetValue()).To(Equal(float64(1)))

            Expect(labels[0].GetName()).To(Equal("App"))
            Expect(labels[0].GetValue()).To(Equal("some_value"))
        })

        It("does not error when called multiple times", func() {
            counterMetric := CounterMetric{
                Metric: "counter_incremented_multiple_times",
                Value: 1,
                App: "some_value",
            }

            families := captureMetrics(func () {
                sender.Incr(counterMetric)
                sender.Incr(counterMetric)
                sender.Incr(counterMetric)
            })

            metrics := families[0].GetMetric()
            metric := metrics[0].Counter

            Expect(len(families)).To(Equal(1))
            Expect(len(metrics)).To(Equal(1))

            Expect(metric.GetValue()).To(Equal(float64(3)))
        })

        It("includes Metadata as additional labels", func() {
            families := captureMetrics(func () {
                sender.Incr(CounterMetric{
                    Metric: "response",
                    Metadata: map[string]string{ "statusRange": "2xx" },
                    Value: 1,
                })
            })

            metrics := families[0].GetMetric()
            labels := metrics[0].GetLabel()
            metadata := labels[len(labels) - 1]

            Expect(metadata.GetName()).To(Equal("statusRange"))
            Expect(metadata.GetValue()).To(Equal("2xx"))
        })
    })

    Describe("#Gauge", func() {
        It("sends a floating point gauge metric to prometheus", func() {
            families := captureMetrics(func () {
                sender.Gauge(GaugeMetric{
                    Metric: "my_gauge",
                    Value: 3,
                })
            })

            family := families[0]
            metrics := family.GetMetric()
            metric := metrics[0].Gauge

            Expect(family.GetName()).To(Equal("my_gauge"))
            Expect(metric.GetValue()).To(Equal(3.0))
        })
    })

    Describe("#PrecisionTiming", func() {
        It("sends a histogram metric into a sensible bucket in prometheus", func() {
            families := captureMetrics(func () {
                sender.PrecisionTiming(PrecisionTimingMetric{
                    Metric: "my_precise_time",
                    Value: time.Duration(3142) * time.Millisecond,
                })
            })

            family := families[0]
            metrics := family.GetMetric()
            metric := metrics[0].Histogram
            buckets := metric.GetBucket()

            Expect(family.GetName()).To(Equal("my_precise_time"))
            Expect(metric.GetSampleCount()).To(Equal(uint64(1)))
            Expect(metric.GetSampleSum()).To(Equal(3.142))

            last_buckets := buckets[len(buckets) - 3:]

            Expect(last_buckets[0].GetUpperBound()).To(Equal(2.5))
            Expect(last_buckets[0].GetCumulativeCount()).To(Equal(uint64(0)))

            Expect(last_buckets[1].GetUpperBound()).To(Equal(5.0))
            Expect(last_buckets[1].GetCumulativeCount()).To(Equal(uint64(1)))

            Expect(last_buckets[2].GetUpperBound()).To(Equal(10.0))
            Expect(last_buckets[2].GetCumulativeCount()).To(Equal(uint64(1)))
        })
    })
})

type m = []*io_prometheus_client.MetricFamily

func captureMetrics(callback func()) m {
    gatherer := prometheus.DefaultGatherer

    before, _ := gatherer.Gather()
    callback()
    after, _ := gatherer.Gather()

    subtracted := subtract(after, before)
    Expect(len(subtracted)).To(BeNumerically(">", 0),
        "expected to capture some new metrics",
    )

    return subtracted
}

func subtract(aSlice m, bSlice m) m {
    var subtracted m

    Outer:
    for _, a := range aSlice {
        for _, b := range bSlice {
            if a.GetName() == b.GetName() {
                continue Outer
            }
        }
        subtracted = append(subtracted, a)
    }

    return subtracted
}
