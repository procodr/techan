package techan

import (
	"math"

	"github.com/sdcoffey/big"
)

type relativeStrengthIndexIndicator struct {
	rsIndicator Indicator
	oneHundred  big.Decimal
}

// NewRelativeStrengthIndexIndicator returns a derivative Indicator which returns the relative strength index of the base indicator
// in a given time frame. A more in-depth explanation of relative strength index can be found here:
// https://www.investopedia.com/terms/r/rsi.asp
func NewRelativeStrengthIndexIndicator(indicator Indicator, timeframe int) Indicator {
	return relativeStrengthIndexIndicator{
		rsIndicator: NewRelativeStrengthIndicator(indicator, timeframe),
		oneHundred:  big.NewFromString("100"),
	}
}

func (rsi relativeStrengthIndexIndicator) Calculate(index int) big.Decimal {
	relativeStrength := rsi.rsIndicator.Calculate(index)

	return rsi.oneHundred.Sub(rsi.oneHundred.Div(big.ONE.Add(relativeStrength)))
}

type relativeStrengthIndicator struct {
	avgGain Indicator
	avgLoss Indicator
	window  int
}

// NewRelativeStrengthIndicator returns a derivative Indicator which returns the relative strength of the base indicator
// in a given time frame. Relative strength is the average again of up periods during the time frame divided by the
// average loss of down period during the same time frame
func NewRelativeStrengthIndicator(indicator Indicator, timeframe int) Indicator {
	return relativeStrengthIndicator{
		avgGain: NewMMAIndicator(NewGainIndicator(indicator), timeframe),
		avgLoss: NewMMAIndicator(NewLossIndicator(indicator), timeframe),
		window:  timeframe,
	}
}

func (rs relativeStrengthIndicator) Calculate(index int) big.Decimal {
	if index < rs.window-1 {
		return big.ZERO
	}

	avgGain := rs.avgGain.Calculate(index)
	avgLoss := rs.avgLoss.Calculate(index)

	if avgLoss.EQ(big.ZERO) {
		return big.NewDecimal(math.Inf(1))
	}

	return avgGain.Div(avgLoss)
}

type relativeStrengthIndexEmaSignalIndicator struct {
	emaRsi Indicator
	window int
}

func NewRelativeStrengthIndexEmaSignalIndicator(rsi Indicator, window int) Indicator {
	return &relativeStrengthIndexEmaSignalIndicator{
		emaRsi: NewEMAIndicator(rsi, window),
		window: window,
	}
}

func (rsiEmaS *relativeStrengthIndexEmaSignalIndicator) Calculate(index int) big.Decimal {
	if index < rsiEmaS.window-1 {
		return big.ZERO
	}

	return rsiEmaS.emaRsi.Calculate(index)
}

type relativeStrengthIndexSmaSignalIndicator struct {
	smaRsi Indicator
	window int
}

func NewRelativeStrengthIndexSmaSignalIndicator(rsi Indicator, window int) Indicator {
	return &relativeStrengthIndexSmaSignalIndicator{
		smaRsi: NewSimpleMovingAverage(rsi, window),
		window: window,
	}
}

func (rsiSmaS *relativeStrengthIndexSmaSignalIndicator) Calculate(index int) big.Decimal {
	if index < rsiSmaS.window-1 {
		return big.ZERO
	}

	return rsiSmaS.smaRsi.Calculate(index)
}
