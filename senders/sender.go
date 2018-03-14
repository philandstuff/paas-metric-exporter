package senders

import (
	"time"
)

//go:generate counterfeiter -o mocks/sender.go . Sender
type Sender interface {
	Gauge(stat string, value int64) error
	FGauge(stat string, value float64) error
	Incr(stat string, count int64) error
	Timing(string, int64) error
	PrecisionTiming(stat string, delta time.Duration) error
}

