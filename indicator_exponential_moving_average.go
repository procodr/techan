package techan

import (
	"github.com/sdcoffey/big"
	"sync"
)

type emaIndicator struct {
	indicator   Indicator
	window      int
	alpha       big.Decimal
	resultCache resultCache
	mu          *sync.RWMutex
}

// NewEMAIndicator returns a derivative indicator which returns the average of the current and preceding values in
// the given windowSize, with values closer to current index given more weight. A more in-depth explanation can be found here:
// http://www.investopedia.com/terms/e/ema.asp
func NewEMAIndicator(indicator Indicator, window int) Indicator {
	return &emaIndicator{
		indicator:   indicator,
		window:      window,
		alpha:       big.ONE.Frac(2).Div(big.NewFromInt(window + 1)),
		resultCache: make([]*big.Decimal, 1000),
		mu:          &sync.RWMutex{},
	}
}

func (ema *emaIndicator) Calculate(index int) big.Decimal {
	if cachedValue := returnIfCached(ema, index, func(i int) big.Decimal {
		return NewSimpleMovingAverage(ema.indicator, ema.window).Calculate(i)
	}); cachedValue != nil {
		return *cachedValue
	}

	todayVal := ema.indicator.Calculate(index).Mul(ema.alpha)
	result := todayVal.Add(ema.Calculate(index - 1).Mul(big.ONE.Sub(ema.alpha)))

	cacheResult(ema, index, result)

	return result
}

func (ema emaIndicator) cache() resultCache {
	return ema.resultCache
}

func (ema *emaIndicator) setCacheEntry(index int, val *big.Decimal) {
	ema.mu.Lock()
	defer ema.mu.Unlock()

	ema.resultCache[index] = val
}

func (ema *emaIndicator) getCacheEntry(index int) big.Decimal {
	ema.mu.RLock()
	defer ema.mu.RUnlock()

	if ema.resultCache[index] != nil {
		return *ema.resultCache[index]
	} else {
		return big.NaN
	}
}

func (ema *emaIndicator) setCache(newCache resultCache) {
	ema.mu.Lock()
	defer ema.mu.Unlock()

	ema.resultCache = newCache
}

func (ema emaIndicator) windowSize() int { return ema.window }

func (ema *emaIndicator) Lock() {
	ema.mu.Lock()
}

func (ema *emaIndicator) Unlock() {
	ema.mu.Unlock()
}

func (ema *emaIndicator) RLock() {
	ema.mu.RLock()
}

func (ema *emaIndicator) RUnlock() {
	ema.mu.RUnlock()
}
