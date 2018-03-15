package senders_test

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_model/go"

	. "github.com/alphagov/paas-metric-exporter/metrics"
	. "github.com/alphagov/paas-metric-exporter/senders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

)

var _ = Describe("PrometheusSender", func() {
	Describe("#Incr", func() {
        It("sends a counter metric to prometheus", func() {
            sender := NewPrometheusSender()

            families := captureMetrics(func () {
                sender.Incr(CounterMetric{
                    Metric: "my_counter",
                    Value: 1,
                    App: "some_value",
                })
            })

            family := families[0]
            metrics := family.GetMetric()
            metric := metrics[0]
            labels := metric.GetLabel()

            Expect(len(families)).To(Equal(1))
            Expect(len(metrics)).To(Equal(1))

            Expect(family.GetName()).To(Equal("my_counter"))
            Expect(metric.Counter.GetValue()).To(Equal(float64(1)))

            Expect(labels[0].GetName()).To(Equal("App"))
            Expect(labels[0].GetValue()).To(Equal("some_value"))
        })
    })
})

type m = []*io_prometheus_client.MetricFamily

func captureMetrics(callback func()) m {
    gatherer := prometheus.DefaultGatherer

    before, _ := gatherer.Gather()
    callback()
    after, _ := gatherer.Gather()

    return subtract(after, before)
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
