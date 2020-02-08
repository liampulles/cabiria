package image_test

import (
	"fmt"
	"image/color"
	"math"
	"path/filepath"
	"testing"

	"github.com/liampulles/cabiria/pkg/image"
)

func TestColorDiff(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixtureCol1 color.Color
		fixtureCol2 color.Color
		expected    float64
	}{
		{color.RGBA{0, 0, 0, 0}, color.RGBA{0, 0, 0, 0}, 0.0},
		{color.RGBA{0, 0, 0, 255}, color.RGBA{0, 0, 0, 255}, 0.0},
		{color.RGBA{255, 255, 255, 255}, color.RGBA{255, 255, 255, 255}, 0.0},
		{color.RGBA{0, 0, 0, 255}, color.RGBA{255, 255, 255, 255}, 1.0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("(%v, %v) => %f", test.fixtureCol1, test.fixtureCol2, test.expected), func(t *testing.T) {
			// Exercise SUT
			actual := image.ColorDiff(test.fixtureCol1, test.fixtureCol2)

			// Verify result (should be very close)
			if math.Abs(actual-test.expected) > 0.0001 {
				t.Errorf("Unexpected result.\nExpected: %f\nActual: %f", test.expected, actual)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		file1    string
		file2    string
		expected float64
	}{
		{"white.png", "white.png", 0.0},
		{"black.png", "black.png", 0.0},
		{"white.png", "black.png", 1.0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("(%v, %v) => %f", test.file1, test.file2, test.expected), func(t *testing.T) {
			// Setup Fixture
			img1, err := image.GetPNG(filepath.Join("testdata", test.file1))
			if err != nil {
				t.Errorf("Encountered error loading fixture img1: %v", err)
				return
			}
			img2, err := image.GetPNG(filepath.Join("testdata", test.file2))
			if err != nil {
				t.Errorf("Encountered error loading fixture img2: %v", err)
				return
			}

			// Exercise SUT
			actual, err := image.Diff(img1, img2)

			// Verify result (should be very close)
			if err != nil {
				t.Errorf("Encountered error executing SUT: %v", err)
				return
			}
			if math.Abs(actual-test.expected) > 0.0001 {
				t.Errorf("Unexpected result.\nExpected: %f\nActual: %f", test.expected, actual)
			}
		})
	}
}

func TestDiff_WhenImageBoundsDoNotMatch(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		file1 string
		file2 string
	}{
		{"white.png", "tiny.png"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("(%v, %v)", test.file1, test.file2), func(t *testing.T) {
			// Setup Fixture
			img1, err := image.GetPNG(filepath.Join("testdata", test.file1))
			if err != nil {
				t.Errorf("Encountered error loading fixture img1: %v", err)
				return
			}
			img2, err := image.GetPNG(filepath.Join("testdata", test.file2))
			if err != nil {
				t.Errorf("Encountered error loading fixture img2: %v", err)
				return
			}

			// Exercise SUT
			_, err = image.Diff(img1, img2)

			// Verify result (should be very close)
			if err == nil {
				t.Errorf("Expected an err, but none was returned.")
				return
			}
		})
	}
}
