package techan

import (
	"github.com/sdcoffey/big"
)

type resultCache []*big.Decimal

type cachedIndicator interface {
	Indicator
	cache() resultCache
	getCacheEntry(index int) big.Decimal
	setCache(cache resultCache)
	setCacheEntry(index int, val *big.Decimal)
	windowSize() int
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

func cacheResult(indicator cachedIndicator, index int, val big.Decimal) {
	if index < len(indicator.cache()) {
		indicator.setCacheEntry(index, &val)
	} else if index == len(indicator.cache()) {
		indicator.setCache(append(indicator.cache(), &val))
	} else {
		expandResultCache(indicator, index+1)
		cacheResult(indicator, index, val)
	}
}

func expandResultCache(indicator cachedIndicator, newSize int) {
	sizeDiff := newSize - len(indicator.cache())

	expansion := make([]*big.Decimal, sizeDiff)
	indicator.setCache(append(indicator.cache(), expansion...))
}

func returnIfCached(indicator cachedIndicator, index int, firstValueFallback func(int) big.Decimal) *big.Decimal {
	if index >= len(indicator.cache()) {
		expandResultCache(indicator, index+1)
	} else if index < indicator.windowSize()-1 {
		return &big.ZERO
	} else if val := indicator.getCacheEntry(index); val != big.NaN {
		return &val
	} else if index == indicator.windowSize()-1 {
		value := firstValueFallback(index)
		cacheResult(indicator, index, value)
		return &value
	}

	return nil
}
