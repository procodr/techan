package techan

import "testing"

func TestTrueStrengthIndexIndicator(t *testing.T) {
	indicator := NewTrueStrengthIndexIndicator(NewClosePriceIndicator(mockedTimeSeries), 3, 2)

	expectedValues := []float64{
		0, 0, 0, -100, -100, -100, 17.4829, 24.6987, -49.2538, -15.542, -69.33, -62.4557,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestTrueStrengthIndexEmaSignalIndicator(t *testing.T) {
	indicator := NewTrueStrengthIndexEmaSignalIndicator(
		NewTrueStrengthIndexIndicator(NewClosePriceIndicator(mockedTimeSeries), 3, 2), 2)

	expectedValues := []float64{
		0, 0, 0, -66.6667, -88.8889, -96.2963, -20.4435, 9.6513, -29.6188, -20.2343, -52.9648, -59.292,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestTrueStrengthIndexSmaSignalIndicator(t *testing.T) {
	indicator := NewTrueStrengthIndexSmaSignalIndicator(
		NewTrueStrengthIndexIndicator(NewClosePriceIndicator(mockedTimeSeries), 3, 2), 2)

	expectedValues := []float64{
		0, 0, 0, -50, -100, -100, -41.2585, 21.0908, -12.2776, -32.3979, -42.436, -65.8928,
	}

	indicatorEquals(t, expectedValues, indicator)
}
