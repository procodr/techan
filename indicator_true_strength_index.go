package techan

import (
	"github.com/sdcoffey/big"
	"math"
)

type trueStrengthIndexIndicator struct {
	avgPriceChange    Indicator
	avgAbsPriceChange Indicator
	oneHundred        big.Decimal
	long, short       int
}

func NewTrueStrengthIndexIndicator(indicator Indicator, long, short int) Indicator {
	return &trueStrengthIndexIndicator{
		avgPriceChange:    NewEMAIndicator(NewEMAIndicator(NewChangeIndicator(indicator), long), short),
		avgAbsPriceChange: NewEMAIndicator(NewEMAIndicator(NewAbsoluteChangeIndicator(indicator), long), short),
		long:              long,
		short:             short,
		oneHundred:        big.NewFromString("100"),
	}
}

func (tsi *trueStrengthIndexIndicator) Calculate(index int) big.Decimal {
	if index < tsi.long+tsi.short-2 {
		return big.ZERO
	}

	avgPriceChange := tsi.avgPriceChange.Calculate(index)
	avgAbsPriceChange := tsi.avgAbsPriceChange.Calculate(index)

	if avgAbsPriceChange.EQ(big.ZERO) {
		return big.NewDecimal(math.Inf(1))
	}

	return tsi.oneHundred.Mul(avgPriceChange.Div(avgAbsPriceChange))
}

type trueStrengthIndexSignalIndicator struct {
	avgTsi Indicator
	window int
}

func NewTrueStrengthIndexSignalIndicator(tsi Indicator, window int) Indicator {
	return &trueStrengthIndexSignalIndicator{
		avgTsi: NewEMAIndicator(tsi, window),
		window: window,
	}
}

func (ts *trueStrengthIndexSignalIndicator) Calculate(index int) big.Decimal {
	if index < ts.window-1 {
		return big.ZERO
	}

	return ts.avgTsi.Calculate(index)
}
