package metrics_test

import (
	. "github.com/alphagov/paas-metric-exporter/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"time"
)

type FakeStatsdClient struct {
	timingCalled          bool
	precisionTimingCalled bool
	incrCalled            bool
	gaugeCalled           bool
	fGaugeCalled          bool
	stat                  string
	value                 int64
	fValue                float64
	precisionTimingValue  time.Duration
}

func (f *FakeStatsdClient) PrecisionTiming(metric PrecisionTimingMetric) error {
	f.precisionTimingCalled = true
	f.stat = metric.Metric // should this be the presented name?
	f.precisionTimingValue = metric.Value
	return nil
}

func (f *FakeStatsdClient) Incr(metric CounterMetric) error {
	f.incrCalled = true
	f.stat = metric.Metric // should this be the presented name?
	f.value = metric.Value
	return nil
}

func (f *FakeStatsdClient) Gauge(metric GaugeMetric) error {
	f.gaugeCalled = true
	f.stat = metric.Metric // should this be the presented name?
	f.value = metric.Value
	return nil
}

var _ = Describe("Metric", func() {
	var (
		fakeStatsdClient *FakeStatsdClient
	)

	Describe("#NewCounterMetric", func() {
		It("creates a new CounterMetric", func() {
			metric := CounterMetric{Metric: "my.counter.metric", Value: 1}

			Expect(metric.Name()).To(Equal("my.counter.metric"))
			Expect(metric.Value).To(Equal(int64(1)))
		})
	})

	Describe("#NewGaugeMetric", func() {
		It("creates a new GaugeMetric", func() {
			metric := GaugeMetric{Metric: "my.gauge.metric", Value: 20}

			Expect(metric.Name()).To(Equal("my.gauge.metric"))
			Expect(metric.Value).To(Equal(int64(20)))
		})
	})

	Describe("#NewPrecisionTimingMetric", func() {
		It("creates a new PrecisionTimingMetric", func() {
			metric := PrecisionTimingMetric{Metric: "my.precision.timing.metric", Value: 100 * time.Millisecond}

			Expect(metric.Name()).To(Equal("my.precision.timing.metric"))
			Expect(metric.Value).To(Equal(100 * time.Millisecond))
		})
	})

	Describe("#Send", func() {
		BeforeEach(func() {
			fakeStatsdClient = new(FakeStatsdClient)
		})

		Context("with a PrecisionTimingMetric", func() {
			Context("without prefix", func() {
				It("sends the Metric to StatsD with time.Duration precision", func() {
					metric := PrecisionTimingMetric{Metric: "http.responsetimes.api_10_244_0_34_xip_io", Value: 50 * time.Millisecond}
					metric.Send(fakeStatsdClient)

					Expect(fakeStatsdClient.precisionTimingCalled).To(BeTrue())
					Expect(fakeStatsdClient.stat).To(Equal("http.responsetimes.api_10_244_0_34_xip_io"))
					Expect(fakeStatsdClient.precisionTimingValue).To(Equal(50 * time.Millisecond))
				})
			})

			Context("with prefix", func() {
				It("sends the Metric to StatsD with time.Duration precision", func() {
					metric := PrecisionTimingMetric{Metric: "http.responsetimes.api_10_244_0_34_xip_io", Value: 50 * time.Millisecond}
					metric.Send(fakeStatsdClient)

					Expect(fakeStatsdClient.precisionTimingCalled).To(BeTrue())
					Expect(fakeStatsdClient.stat).To(Equal("http.responsetimes.api_10_244_0_34_xip_io"))
					Expect(fakeStatsdClient.precisionTimingValue).To(Equal(50 * time.Millisecond))
				})
			})
		})

		Context("with a CounterMetric", func() {
			Context("without prefix", func() {
				It("sends the Metric to StatsD with int64 precision", func() {
					metric := CounterMetric{Metric: "http.statuscodes.api_10_244_0_34_xip_io.200", Value: 1}
					metric.Send(fakeStatsdClient)

					Expect(fakeStatsdClient.incrCalled).To(BeTrue())
					Expect(fakeStatsdClient.stat).To(Equal("http.statuscodes.api_10_244_0_34_xip_io.200"))
					Expect(fakeStatsdClient.value).To(Equal(int64(1)))
				})
			})

			Context("with prefix", func() {
				It("sends the Metric to StatsD with int64 precision", func() {
					metric := CounterMetric{Metric: "http.statuscodes.api_10_244_0_34_xip_io.200", Value: 1}
					metric.Send(fakeStatsdClient)

					Expect(fakeStatsdClient.incrCalled).To(BeTrue())
					Expect(fakeStatsdClient.stat).To(Equal("http.statuscodes.api_10_244_0_34_xip_io.200"))
					Expect(fakeStatsdClient.value).To(Equal(int64(1)))
				})
			})
		})

		Context("with a GaugeMetric", func() {
			Context("without prefix", func() {
				It("sends the Metric to StatsD with int64 precision", func() {
					metric := GaugeMetric{Metric: "router__0.numCPUS", Value: 4}
					metric.Send(fakeStatsdClient)

					Expect(fakeStatsdClient.gaugeCalled).To(BeTrue())
					Expect(fakeStatsdClient.stat).To(Equal("router__0.numCPUS"))
					Expect(fakeStatsdClient.value).To(Equal(int64(4)))
				})
			})

			Context("with prefix", func() {
				It("sends the Metric to StatsD with int64 precision", func() {
					metric := GaugeMetric{Metric: "router__0.numCPUS", Value: 4}
					metric.Send(fakeStatsdClient)

					Expect(fakeStatsdClient.gaugeCalled).To(BeTrue())
					Expect(fakeStatsdClient.stat).To(Equal("router__0.numCPUS"))
					Expect(fakeStatsdClient.value).To(Equal(int64(4)))
				})
			})
		})

		Context("when the StatsdClient doesn't return an error", func() {
			It("doesn't return an error", func() {
				metric := GaugeMetric{Metric: "router__0.numCPUS", Value: 4}
				err := metric.Send(fakeStatsdClient)

				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
