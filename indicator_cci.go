package techan

import "github.com/sdcoffey/big"

type commidityChannelIndexIndicator struct {
	ccIndicator Indicator
	window      int
}

// NewCCIIndicator Returns a new Commodity Channel Index Indicator
// http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:commodity_channel_index_cci
func NewCCIIndicator(indicator Indicator, window int) Indicator {
	return commidityChannelIndexIndicator{
		ccIndicator: indicator,
		window:      window,
	}
}

func (ccii commidityChannelIndexIndicator) Calculate(index int) big.Decimal {
	typicalPriceSma := NewSimpleMovingAverage(ccii.ccIndicator, ccii.window)
	meanDeviation := NewMeanDeviationIndicator(ccii.ccIndicator, ccii.window)

	return ccii.ccIndicator.Calculate(index).Sub(typicalPriceSma.Calculate(index)).Div(meanDeviation.Calculate(index).Mul(big.NewFromString("0.015")))
}
