package image_test

import (
	"fmt"
	"image/color"
	"path/filepath"
	"testing"

	"github.com/liampulles/cabiria/pkg/image"
)

func TestLuminance(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture  color.Color
		expected float64
	}{
		{color.RGBA{0, 0, 0, 0}, 0.0},
		{color.RGBA{0, 0, 0, 255}, 0.0},
		{color.RGBA{255, 255, 255, 255}, 1.0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v => %f", test.fixture, test.expected), func(t *testing.T) {
			// Exercise SUT
			actual := image.Luminance(test.fixture)

			// Verify result
			if actual != test.expected {
				t.Errorf("Unexpected result.\nExpected: %f\nActual: %f", test.expected, actual)
			}
		})
	}
}

func TestDetectMinLum(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture  string
		expected float64
	}{
		{"gray_with_extremes.png", 0.0},
		{"gray.png", image.Luminance(color.RGBA{128, 128, 128, 255})},
		{"color_gradient.png", image.Luminance(color.RGBA{127, 0, 0, 255})}, // Dark red value
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v => %f", test.fixture, test.expected), func(t *testing.T) {
			img, err := image.GetImage(filepath.Join("testdata", test.fixture))
			if err != nil {
				t.Errorf("Encountered error loading fixture image: %v", err)
				return
			}

			// Exercise SUT
			actual := image.DetectMinLum(img)

			// Verify result
			if actual != test.expected {
				t.Errorf("Unexpected result.\nExpected: %f\nActual: %f", test.expected, actual)
			}
		})
	}
}

func TestDetectMaxLum(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture  string
		expected float64
	}{
		{"gray_with_extremes.png", 1.0},
		{"gray.png", image.Luminance(color.RGBA{128, 128, 128, 255})},
		{"color_gradient.png", image.Luminance(color.RGBA{127, 146, 255, 255})}, // Light blue
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v => %f", test.fixture, test.expected), func(t *testing.T) {
			img, err := image.GetImage(filepath.Join("testdata", test.fixture))
			if err != nil {
				t.Errorf("Encountered error loading fixture image: %v", err)
				return
			}

			// Exercise SUT
			actual := image.DetectMaxLum(img)

			// Verify result
			if actual != test.expected {
				t.Errorf("Unexpected result.\nExpected: %f\nActual: %f", test.expected, actual)
			}
		})
	}
}

func TestLevelImage(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		path string
		min  float64
		max  float64
	}{
		{"gray_with_extremes.png", 0.0, 1.0},
		{"color_gradient.png", 0.25, 0.74},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v, %f, %f", test.path, test.min, test.max), func(t *testing.T) {
			img, err := image.GetImage(filepath.Join("testdata", test.path))
			if err != nil {
				t.Errorf("Encountered error loading fixture image: %v", err)
				return
			}

			// Exercise SUT
			actual := image.LevelImage(img, test.min, test.max)

			// Verify result
			actualMin := image.DetectMinLum(actual)
			if actualMin != 0.0 {
				t.Errorf("Unexpected result.\nExpected min: %f\nActual min: %f", 0.0, actualMin)
			}
			actualMax := image.DetectMaxLum(actual)
			if actualMax != 1.0 {
				t.Errorf("Unexpected result.\nExpected max: %f\nActual max: %f", 1.0, actualMax)
			}
		})
	}
}

func scale(in uint8) float64 {
	return float64(in) / 255.0
}
