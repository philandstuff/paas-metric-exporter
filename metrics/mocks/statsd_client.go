// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"
	"time"

	"github.com/alphagov/paas-metric-exporter/metrics"
)

type FakeStatsdClient struct {
	GaugeStub        func(stat string, value int64) error
	gaugeMutex       sync.RWMutex
	gaugeArgsForCall []struct {
		stat  string
		value int64
	}
	gaugeReturns struct {
		result1 error
	}
	gaugeReturnsOnCall map[int]struct {
		result1 error
	}
	FGaugeStub        func(stat string, value float64) error
	fGaugeMutex       sync.RWMutex
	fGaugeArgsForCall []struct {
		stat  string
		value float64
	}
	fGaugeReturns struct {
		result1 error
	}
	fGaugeReturnsOnCall map[int]struct {
		result1 error
	}
	IncrStub        func(stat string, count int64) error
	incrMutex       sync.RWMutex
	incrArgsForCall []struct {
		stat  string
		count int64
	}
	incrReturns struct {
		result1 error
	}
	incrReturnsOnCall map[int]struct {
		result1 error
	}
	TimingStub        func(string, int64) error
	timingMutex       sync.RWMutex
	timingArgsForCall []struct {
		arg1 string
		arg2 int64
	}
	timingReturns struct {
		result1 error
	}
	timingReturnsOnCall map[int]struct {
		result1 error
	}
	PrecisionTimingStub        func(stat string, delta time.Duration) error
	precisionTimingMutex       sync.RWMutex
	precisionTimingArgsForCall []struct {
		stat  string
		delta time.Duration
	}
	precisionTimingReturns struct {
		result1 error
	}
	precisionTimingReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStatsdClient) Gauge(stat string, value int64) error {
	fake.gaugeMutex.Lock()
	ret, specificReturn := fake.gaugeReturnsOnCall[len(fake.gaugeArgsForCall)]
	fake.gaugeArgsForCall = append(fake.gaugeArgsForCall, struct {
		stat  string
		value int64
	}{stat, value})
	fake.recordInvocation("Gauge", []interface{}{stat, value})
	fake.gaugeMutex.Unlock()
	if fake.GaugeStub != nil {
		return fake.GaugeStub(stat, value)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.gaugeReturns.result1
}

func (fake *FakeStatsdClient) GaugeCallCount() int {
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	return len(fake.gaugeArgsForCall)
}

func (fake *FakeStatsdClient) GaugeArgsForCall(i int) (string, int64) {
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	return fake.gaugeArgsForCall[i].stat, fake.gaugeArgsForCall[i].value
}

func (fake *FakeStatsdClient) GaugeReturns(result1 error) {
	fake.GaugeStub = nil
	fake.gaugeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatsdClient) GaugeReturnsOnCall(i int, result1 error) {
	fake.GaugeStub = nil
	if fake.gaugeReturnsOnCall == nil {
		fake.gaugeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.gaugeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatsdClient) FGauge(stat string, value float64) error {
	fake.fGaugeMutex.Lock()
	ret, specificReturn := fake.fGaugeReturnsOnCall[len(fake.fGaugeArgsForCall)]
	fake.fGaugeArgsForCall = append(fake.fGaugeArgsForCall, struct {
		stat  string
		value float64
	}{stat, value})
	fake.recordInvocation("FGauge", []interface{}{stat, value})
	fake.fGaugeMutex.Unlock()
	if fake.FGaugeStub != nil {
		return fake.FGaugeStub(stat, value)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.fGaugeReturns.result1
}

func (fake *FakeStatsdClient) FGaugeCallCount() int {
	fake.fGaugeMutex.RLock()
	defer fake.fGaugeMutex.RUnlock()
	return len(fake.fGaugeArgsForCall)
}

func (fake *FakeStatsdClient) FGaugeArgsForCall(i int) (string, float64) {
	fake.fGaugeMutex.RLock()
	defer fake.fGaugeMutex.RUnlock()
	return fake.fGaugeArgsForCall[i].stat, fake.fGaugeArgsForCall[i].value
}

func (fake *FakeStatsdClient) FGaugeReturns(result1 error) {
	fake.FGaugeStub = nil
	fake.fGaugeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatsdClient) FGaugeReturnsOnCall(i int, result1 error) {
	fake.FGaugeStub = nil
	if fake.fGaugeReturnsOnCall == nil {
		fake.fGaugeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.fGaugeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatsdClient) Incr(stat string, count int64) error {
	fake.incrMutex.Lock()
	ret, specificReturn := fake.incrReturnsOnCall[len(fake.incrArgsForCall)]
	fake.incrArgsForCall = append(fake.incrArgsForCall, struct {
		stat  string
		count int64
	}{stat, count})
	fake.recordInvocation("Incr", []interface{}{stat, count})
	fake.incrMutex.Unlock()
	if fake.IncrStub != nil {
		return fake.IncrStub(stat, count)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.incrReturns.result1
}

func (fake *FakeStatsdClient) IncrCallCount() int {
	fake.incrMutex.RLock()
	defer fake.incrMutex.RUnlock()
	return len(fake.incrArgsForCall)
}

func (fake *FakeStatsdClient) IncrArgsForCall(i int) (string, int64) {
	fake.incrMutex.RLock()
	defer fake.incrMutex.RUnlock()
	return fake.incrArgsForCall[i].stat, fake.incrArgsForCall[i].count
}

func (fake *FakeStatsdClient) IncrReturns(result1 error) {
	fake.IncrStub = nil
	fake.incrReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatsdClient) IncrReturnsOnCall(i int, result1 error) {
	fake.IncrStub = nil
	if fake.incrReturnsOnCall == nil {
		fake.incrReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.incrReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatsdClient) Timing(arg1 string, arg2 int64) error {
	fake.timingMutex.Lock()
	ret, specificReturn := fake.timingReturnsOnCall[len(fake.timingArgsForCall)]
	fake.timingArgsForCall = append(fake.timingArgsForCall, struct {
		arg1 string
		arg2 int64
	}{arg1, arg2})
	fake.recordInvocation("Timing", []interface{}{arg1, arg2})
	fake.timingMutex.Unlock()
	if fake.TimingStub != nil {
		return fake.TimingStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.timingReturns.result1
}

func (fake *FakeStatsdClient) TimingCallCount() int {
	fake.timingMutex.RLock()
	defer fake.timingMutex.RUnlock()
	return len(fake.timingArgsForCall)
}

func (fake *FakeStatsdClient) TimingArgsForCall(i int) (string, int64) {
	fake.timingMutex.RLock()
	defer fake.timingMutex.RUnlock()
	return fake.timingArgsForCall[i].arg1, fake.timingArgsForCall[i].arg2
}

func (fake *FakeStatsdClient) TimingReturns(result1 error) {
	fake.TimingStub = nil
	fake.timingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatsdClient) TimingReturnsOnCall(i int, result1 error) {
	fake.TimingStub = nil
	if fake.timingReturnsOnCall == nil {
		fake.timingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.timingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatsdClient) PrecisionTiming(stat string, delta time.Duration) error {
	fake.precisionTimingMutex.Lock()
	ret, specificReturn := fake.precisionTimingReturnsOnCall[len(fake.precisionTimingArgsForCall)]
	fake.precisionTimingArgsForCall = append(fake.precisionTimingArgsForCall, struct {
		stat  string
		delta time.Duration
	}{stat, delta})
	fake.recordInvocation("PrecisionTiming", []interface{}{stat, delta})
	fake.precisionTimingMutex.Unlock()
	if fake.PrecisionTimingStub != nil {
		return fake.PrecisionTimingStub(stat, delta)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.precisionTimingReturns.result1
}

func (fake *FakeStatsdClient) PrecisionTimingCallCount() int {
	fake.precisionTimingMutex.RLock()
	defer fake.precisionTimingMutex.RUnlock()
	return len(fake.precisionTimingArgsForCall)
}

func (fake *FakeStatsdClient) PrecisionTimingArgsForCall(i int) (string, time.Duration) {
	fake.precisionTimingMutex.RLock()
	defer fake.precisionTimingMutex.RUnlock()
	return fake.precisionTimingArgsForCall[i].stat, fake.precisionTimingArgsForCall[i].delta
}

func (fake *FakeStatsdClient) PrecisionTimingReturns(result1 error) {
	fake.PrecisionTimingStub = nil
	fake.precisionTimingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatsdClient) PrecisionTimingReturnsOnCall(i int, result1 error) {
	fake.PrecisionTimingStub = nil
	if fake.precisionTimingReturnsOnCall == nil {
		fake.precisionTimingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.precisionTimingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStatsdClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.gaugeMutex.RLock()
	defer fake.gaugeMutex.RUnlock()
	fake.fGaugeMutex.RLock()
	defer fake.fGaugeMutex.RUnlock()
	fake.incrMutex.RLock()
	defer fake.incrMutex.RUnlock()
	fake.timingMutex.RLock()
	defer fake.timingMutex.RUnlock()
	fake.precisionTimingMutex.RLock()
	defer fake.precisionTimingMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStatsdClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ metrics.StatsdClient = new(FakeStatsdClient)
