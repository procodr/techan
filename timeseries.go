package techan

import (
	"fmt"
	"sync"
)

// TimeSeries represents an array of candles
type TimeSeries struct {
	Candles []*Candle
	Mu      *sync.RWMutex
}

// NewTimeSeries returns a new, empty, TimeSeries
func NewTimeSeries() (t *TimeSeries) {
	t = new(TimeSeries)
	t.Candles = make([]*Candle, 0)
	t.Mu = new(sync.RWMutex)

	return t
}

// AddCandle adds the given candle to this TimeSeries if it is not nil and after the last candle in this timeseries.
// If the candle is added, AddCandle will return true, otherwise it will return false.
func (ts *TimeSeries) AddCandle(candle *Candle) bool {
	if candle == nil {
		panic(fmt.Errorf("error adding Candle: candle cannot be nil"))
	}

	if ts.LastCandle() == nil || candle.Period.Since(ts.LastCandle().Period) >= 0 {
		ts.Mu.Lock()
		ts.Candles = append(ts.Candles, candle)
		ts.Mu.Unlock()
		return true
	}

	return false
}

// LastCandle will return the lastCandle in this series, or nil if this series is empty
func (ts *TimeSeries) LastCandle() *Candle {
	ts.Mu.RLock()
	defer ts.Mu.RUnlock()

	if len(ts.Candles) > 0 {
		return ts.Candles[len(ts.Candles)-1]
	}

	return nil
}

// LastIndex will return the index of the last candle in this series
func (ts *TimeSeries) LastIndex() int {
	ts.Mu.RLock()
	defer ts.Mu.RUnlock()

	return len(ts.Candles) - 1
}
