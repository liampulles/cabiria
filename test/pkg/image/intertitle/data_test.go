package intertitle_test

import (
	"fmt"
	"image/color"
	"math"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/liampulles/cabiria/pkg/ml"

	"github.com/liampulles/cabiria/pkg/image"

	"github.com/liampulles/cabiria/pkg/image/intertitle"
)

func TestIntensity(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture  color.Color
		expected float64
	}{
		{color.RGBA{0, 0, 0, 255}, 0.0},
		{color.RGBA{255, 255, 255, 255}, 1.0},
		{color.RGBA{128, 128, 128, 255}, 0.535850},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v => %f", test.fixture, test.expected), func(t *testing.T) {

			// Exercise SUT
			actual := intertitle.Intensity(test.fixture)

			// Verify result (must be very close)
			if !veryClose(actual, test.expected) {
				t.Errorf("Unexpected result.\nExpected: %f\nActual: %f", test.expected, actual)
			}
		})
	}
}

func TestGetIntensityStats(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture  string
		expected intertitle.IntensityStats
	}{
		{"frame000003.png", intertitle.IntensityStats{
			AvgIntensity:       0.109066,
			LowerAvgIntensity:  0.001877,
			MiddleAvgIntensity: 0.563076,
			UpperAvgIntensity:  0.827755,
			ProportionLower:    0.857025,
			ProportionMiddle:   0.041146,
			ProportionUpper:    0.101829,
		}},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v => %f", test.fixture, test.expected), func(t *testing.T) {
			// Srtup fixture
			img, err := image.GetPNG(filepath.Join("testdata", test.fixture))
			if err != nil {
				t.Errorf("encountered error while loading fixture: %v", err)
				return
			}

			// Exercise SUT
			actual := intertitle.GetIntensityStats(img)

			// Verify result (must be very close)
			checks := [][2]float64{
				{actual.AvgIntensity, test.expected.AvgIntensity},
				{actual.LowerAvgIntensity, test.expected.LowerAvgIntensity},
				{actual.MiddleAvgIntensity, test.expected.MiddleAvgIntensity},
				{actual.UpperAvgIntensity, test.expected.UpperAvgIntensity},
				{actual.ProportionLower, test.expected.ProportionLower},
				{actual.ProportionMiddle, test.expected.ProportionMiddle},
				{actual.ProportionUpper, test.expected.ProportionUpper},
			}
			for _, check := range checks {
				if !veryClose(check[0], check[1]) {
					t.Errorf("Unexpected result.\nExpected: %f\nActual: %f", check[1], check[0])
				}
			}
		})
	}
}

func TestAsInput(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture  intertitle.IntensityStats
		expected ml.Datum
	}{
		{
			intertitle.IntensityStats{
				AvgIntensity:       1.0,
				LowerAvgIntensity:  2.0,
				MiddleAvgIntensity: 3.0,
				UpperAvgIntensity:  4.0,
				ProportionLower:    5.0,
				ProportionMiddle:   6.0,
				ProportionUpper:    7.0,
			},
			[]float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v => %v", test.fixture, test.expected), func(t *testing.T) {

			// Exercise SUT
			actual := test.fixture.AsInput()

			// Verify result (must be very close)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Unexpected result.\nExpected: %f\nActual: %f", test.expected, actual)
			}
		})
	}
}

func veryClose(actual float64, expected float64) bool {
	return math.Abs(actual-expected) < 0.00001
}
