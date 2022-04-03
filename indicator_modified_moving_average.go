package techan

import (
	"github.com/sdcoffey/big"
	"sync"
)

type modifiedMovingAverageIndicator struct {
	indicator   Indicator
	window      int
	resultCache resultCache
	mu          *sync.RWMutex
}

// NewMMAIndicator returns a derivative indicator which returns the modified moving average of the underlying
// indicator. An in-depth explanation can be found here:
// https://en.wikipedia.org/wiki/Moving_average#Modified_moving_average
func NewMMAIndicator(indicator Indicator, window int) Indicator {
	return &modifiedMovingAverageIndicator{
		indicator:   indicator,
		window:      window,
		resultCache: make([]*big.Decimal, 10000),
		mu:          &sync.RWMutex{},
	}
}

func (mma *modifiedMovingAverageIndicator) Calculate(index int) big.Decimal {
	if cachedValue := returnIfCached(mma, index, func(i int) big.Decimal {
		return NewSimpleMovingAverage(mma.indicator, mma.window).Calculate(i)
	}); cachedValue != nil {
		return *cachedValue
	}

	todayVal := mma.indicator.Calculate(index)
	lastVal := mma.Calculate(index - 1)

	result := lastVal.Add(big.NewDecimal(1.0 / float64(mma.window)).Mul(todayVal.Sub(lastVal)))
	cacheResult(mma, index, result)

	return result
}

func (mma modifiedMovingAverageIndicator) cache() resultCache {
	return mma.resultCache
}

func (mma *modifiedMovingAverageIndicator) setCacheEntry(index int, val *big.Decimal) {
	mma.mu.Lock()
	defer mma.mu.Unlock()

	mma.resultCache[index] = val
}

func (mma *modifiedMovingAverageIndicator) getCacheEntry(index int) big.Decimal {
	mma.mu.RLock()
	defer mma.mu.RUnlock()

	if mma.resultCache[index] != nil {
		return *mma.resultCache[index]
	} else {
		return big.NaN
	}
}

func (mma *modifiedMovingAverageIndicator) setCache(cache resultCache) {
	mma.mu.Lock()
	defer mma.mu.Unlock()

	mma.resultCache = cache
}

func (mma modifiedMovingAverageIndicator) windowSize() int {
	return mma.window
}

func (mma *modifiedMovingAverageIndicator) Lock() {
	mma.mu.Lock()
}

func (mma *modifiedMovingAverageIndicator) Unlock() {
	mma.mu.Unlock()
}

func (mma *modifiedMovingAverageIndicator) RLock() {
	mma.mu.RLock()
}

func (mma *modifiedMovingAverageIndicator) RUnlock() {
	mma.mu.RUnlock()
}
