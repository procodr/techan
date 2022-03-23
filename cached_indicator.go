package techan

import (
	"github.com/sdcoffey/big"
)

type resultCache []*big.Decimal

type cachedIndicator interface {
	Indicator
	cache() resultCache
	setCache(cache resultCache)
	windowSize() int
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

func cacheResult(indicator cachedIndicator, index int, val big.Decimal) {
	if index < len(indicator.cache()) {
		indicator.Lock()
		indicator.cache()[index] = &val
		indicator.Unlock()
	} else if index == len(indicator.cache()) {
		indicator.setCache(append(indicator.cache(), &val))
	} else {
		expandResultCache(indicator, index+1)
		cacheResult(indicator, index, val)
	}
}

func expandResultCache(indicator cachedIndicator, newSize int) {
	indicator.Lock()
	defer indicator.Unlock()

	sizeDiff := newSize - len(indicator.cache())

	expansion := make([]*big.Decimal, sizeDiff)
	indicator.setCache(append(indicator.cache(), expansion...))
}

func returnIfCached(indicator cachedIndicator, index int, firstValueFallback func(int) big.Decimal) *big.Decimal {
	if index >= len(indicator.cache()) {
		expandResultCache(indicator, index+1)
	} else if index < indicator.windowSize()-1 {
		return &big.ZERO
	} else if val := indicator.cache()[index]; val != nil {
		return val
	} else if index == indicator.windowSize()-1 {
		value := firstValueFallback(index)
		cacheResult(indicator, index, value)
		return &value
	}

	return nil
}
