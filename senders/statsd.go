package senders

import (
	"time"

	quipo_statsd "github.com/quipo/statsd"
)

type StatsdSender struct {
	Client *quipo_statsd.StatsdClient
}

func NewStatsdSender(statsdEndpoint string, statsdPrefix string) StatsdSender {
	return StatsdSender{Client: quipo_statsd.NewStatsdClient(statsdEndpoint, statsdPrefix)}
}

func (s StatsdSender) Start() {
	s.Client.CreateSocket()
}

func (s StatsdSender) Gauge(stat string, value int64) error {
	return s.Client.Gauge(stat, value)
}

func (s StatsdSender) FGauge(stat string, value float64) error {
	return s.Client.FGauge(stat, value)
}

func (s StatsdSender) Incr(stat string, count int64) error {
	return s.Client.Incr(stat, count)
}

func (s StatsdSender) Timing(stat string, delta int64) error {
	return s.Client.Timing(stat, delta)
}

func (s StatsdSender) PrecisionTiming(stat string, delta time.Duration) error {
	return s.Client.PrecisionTiming(stat, delta)
}
