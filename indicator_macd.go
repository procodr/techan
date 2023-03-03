package techan

import "github.com/sdcoffey/big"

// NewMacdIndicator returns a derivative Indicator which returns the difference between two EMAIndicators with long and
// short windows. It's useful for gauging the strength of price movements. A more in-depth explanation can be found here:
// http://www.investopedia.com/terms/m/macd.asp
func NewMacdIndicator(baseIndicator Indicator, short, long int) Indicator {
	return NewDifferenceIndicator(NewEMAIndicator(baseIndicator, short), NewEMAIndicator(baseIndicator, long))
}

type macdEmaSignalIndicator struct {
	avg    Indicator
	window int
}

func NewMacdEmaSignalIndicator(macd Indicator, window int) Indicator {
	return &macdEmaSignalIndicator{
		avg:    NewEMAIndicator(macd, window),
		window: window,
	}
}

func (m *macdEmaSignalIndicator) Calculate(index int) big.Decimal {
	if index < m.window-1 {
		return big.ZERO
	}

	return m.avg.Calculate(index)
}

type macdSmaSignalIndicator struct {
	avg    Indicator
	window int
}

func NewMacdSmaSignalIndicator(macd Indicator, window int) Indicator {
	return &macdSmaSignalIndicator{
		avg:    NewSimpleMovingAverage(macd, window),
		window: window,
	}
}

func (m *macdSmaSignalIndicator) Calculate(index int) big.Decimal {
	if index < m.window-1 {
		return big.ZERO
	}

	return m.avg.Calculate(index)
}

// NewMACDHistogramIndicator returns a derivative Indicator based on the MACDIndicator, the result of which is
// the macd indicator minus it's signalLinewindow EMA. A more in-depth explanation can be found here:
// http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:macd-histogram
func NewMACDHistogramIndicator(macdIdicator Indicator, signalLinewindow int) Indicator {
	return NewDifferenceIndicator(macdIdicator, NewEMAIndicator(macdIdicator, signalLinewindow))
}
