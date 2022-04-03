package techan

import (
	"math"
	"testing"

	"github.com/sdcoffey/big"
	"github.com/stretchr/testify/assert"
)

func TestRelativeStrengthIndexIndicator(t *testing.T) {
	indicator := NewRelativeStrengthIndexIndicator(NewClosePriceIndicator(mockedTimeSeries), 3)

	expectedValues := []float64{
		0,
		0,
		0,
		0,
		0,
		0,
		57.9952,
		54.0751,
		21.451,
		44.7739,
		14.1542,
		21.2794,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestRelativeStrengthIndicator(t *testing.T) {
	indicator := NewRelativeStrengthIndicator(NewClosePriceIndicator(mockedTimeSeries), 3)

	expectedValues := []float64{
		0,
		0,
		0,
		0,
		0,
		0,
		1.3807,
		1.1775,
		0.2731,
		0.8107,
		0.1649,
		0.2703,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestRelativeStrengthIndicatorNoPriceChange(t *testing.T) {
	close := NewClosePriceIndicator(mockTimeSeries("42.0", "42.0"))
	rsInd := NewRelativeStrengthIndicator(close, 2)
	assert.Equal(t, big.NewDecimal(math.Inf(1)).FormattedString(2), rsInd.Calculate(1).FormattedString(2))
}

func TestRelativeStrengthIndexEmaSignalIndicator(t *testing.T) {
	indicator := NewRelativeStrengthIndexEmaSignalIndicator(
		NewRelativeStrengthIndexIndicator(NewClosePriceIndicator(mockedTimeSeries), 3),
		3)

	expectedValues := []float64{
		0, 0, 0, 0, 0, 0, 28.9976, 41.5364, 31.4937, 38.1338, 26.144, 23.7117,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestRelativeStrengthIndexSmaSignalIndicator(t *testing.T) {
	indicator := NewRelativeStrengthIndexSmaSignalIndicator(
		NewRelativeStrengthIndexIndicator(NewClosePriceIndicator(mockedTimeSeries), 3),
		3)

	expectedValues := []float64{
		0, 0, 0, 0, 0, 0, 19.3317, 37.3568, 44.5071, 40.1, 26.7931, 26.7359,
	}

	indicatorEquals(t, expectedValues, indicator)
}
