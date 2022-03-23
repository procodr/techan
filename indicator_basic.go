package techan

import "github.com/sdcoffey/big"

type volumeIndicator struct {
	*TimeSeries
}

// NewVolumeIndicator returns an indicator which returns the volume of a candle for a given index
func NewVolumeIndicator(series *TimeSeries) Indicator {
	return volumeIndicator{series}
}

func (vi volumeIndicator) Calculate(index int) big.Decimal {
	vi.Mu.RLock()
	defer vi.Mu.RUnlock()

	return vi.Candles[index].Volume
}

type closePriceIndicator struct {
	*TimeSeries
}

// NewClosePriceIndicator returns an Indicator which returns the close price of a candle for a given index
func NewClosePriceIndicator(series *TimeSeries) Indicator {
	return closePriceIndicator{series}
}

func (cpi closePriceIndicator) Calculate(index int) big.Decimal {
	cpi.Mu.RLock()
	defer cpi.Mu.RUnlock()

	return cpi.Candles[index].ClosePrice
}

type highPriceIndicator struct {
	*TimeSeries
}

// NewHighPriceIndicator returns an Indicator which returns the high price of a candle for a given index
func NewHighPriceIndicator(series *TimeSeries) Indicator {
	return highPriceIndicator{
		series,
	}
}

func (hpi highPriceIndicator) Calculate(index int) big.Decimal {
	hpi.Mu.RLock()
	defer hpi.Mu.RUnlock()

	return hpi.Candles[index].MaxPrice
}

type lowPriceIndicator struct {
	*TimeSeries
}

// NewLowPriceIndicator returns an Indicator which returns the low price of a candle for a given index
func NewLowPriceIndicator(series *TimeSeries) Indicator {
	return lowPriceIndicator{
		series,
	}
}

func (lpi lowPriceIndicator) Calculate(index int) big.Decimal {
	lpi.Mu.RLock()
	defer lpi.Mu.RUnlock()

	return lpi.Candles[index].MinPrice
}

type openPriceIndicator struct {
	*TimeSeries
}

// NewOpenPriceIndicator returns an Indicator which returns the open price of a candle for a given index
func NewOpenPriceIndicator(series *TimeSeries) Indicator {
	return openPriceIndicator{
		series,
	}
}

func (opi openPriceIndicator) Calculate(index int) big.Decimal {
	opi.Mu.RLock()
	defer opi.Mu.RUnlock()

	return opi.Candles[index].OpenPrice
}

type typicalPriceIndicator struct {
	*TimeSeries
}

// NewTypicalPriceIndicator returns an Indicator which returns the typical price of a candle for a given index.
// The typical price is an average of the high, low, and close prices for a given candle.
func NewTypicalPriceIndicator(series *TimeSeries) Indicator {
	return typicalPriceIndicator{series}
}

func (tpi typicalPriceIndicator) Calculate(index int) big.Decimal {
	tpi.Mu.RLock()
	defer tpi.Mu.RUnlock()

	numerator := tpi.Candles[index].MaxPrice.Add(tpi.Candles[index].MinPrice).Add(tpi.Candles[index].ClosePrice)
	return numerator.Div(big.NewFromString("3"))
}
