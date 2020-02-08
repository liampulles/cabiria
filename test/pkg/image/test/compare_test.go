package test

import (
	"fmt"
	"image"
	"path/filepath"
	"testing"

	"github.com/liampulles/cabiria/pkg/image/test"

	cabiriaImage "github.com/liampulles/cabiria/pkg/image"
)

func TestCompare_SinglePixelImage_WhenImagesAreTheSame(t *testing.T) {
	// Setup fixture
	path := filepath.Join("testdata", "singlepixel_1.png")
	img1, err := cabiriaImage.GetPNG(path)
	if err != nil {
		t.Errorf("Failed to load img1: %v", err)
	}
	img2, err := cabiriaImage.GetPNG(path)
	if err != nil {
		t.Errorf("Failed to load img2: %v", err)
	}

	// Exercise SUT
	err = test.CompareImage(img1, img2)

	// Verify result
	if err != nil {
		t.Errorf("Comparison failed: %v", err)
	}
}

func TestCompare_SinglePixelImage_WhenImagesAreDifferent(t *testing.T) {
	// Setup fixture
	path1 := filepath.Join("testdata", "singlepixel_1.png")
	img1, err := cabiriaImage.GetPNG(path1)
	if err != nil {
		t.Errorf("Failed to load img1: %v", err)
	}
	path2 := filepath.Join("testdata", "singlepixel_2.png")
	img2, err := cabiriaImage.GetPNG(path2)
	if err != nil {
		t.Errorf("Failed to load img2: %v", err)
	}

	// Exercise SUT
	err = test.CompareImage(img1, img2)

	// Verify result
	if err == nil {
		t.Errorf("Expected comparison failure, but none occured.")
	}
}

func TestCompare_NormalImage_WhenImagesAreTheSame(t *testing.T) {
	// Setup fixture
	path := filepath.Join("testdata", "normal_1.png")
	img1, err := cabiriaImage.GetPNG(path)
	if err != nil {
		t.Errorf("Failed to load img1: %v", err)
	}
	img2, err := cabiriaImage.GetPNG(path)
	if err != nil {
		t.Errorf("Failed to load img2: %v", err)
	}

	// Exercise SUT
	err = test.CompareImage(img1, img2)

	// Verify result
	if err != nil {
		t.Errorf("Comparison failed: %v", err)
	}
}

func TestCompare_NormalImage_WhenImagesAreDifferent(t *testing.T) {
	// Setup fixture
	path1 := filepath.Join("testdata", "normal_1.png")
	img1, err := cabiriaImage.GetPNG(path1)
	if err != nil {
		t.Errorf("Failed to load img1: %v", err)
	}
	path2 := filepath.Join("testdata", "normal_2.png")
	img2, err := cabiriaImage.GetPNG(path2)
	if err != nil {
		t.Errorf("Failed to load img2: %v", err)
	}

	// Exercise SUT
	err = test.CompareImage(img1, img2)

	// Verify result
	if err == nil {
		t.Errorf("Expected comparison failure, but none occured.")
	}
}

func TestValidateBoundsMatch_WhenTheyDoMatch(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		actual   image.Rectangle
		expected image.Rectangle
	}{
		{
			rect(0, 0, 0, 0),
			rect(0, 0, 0, 0),
		},
		{
			rect(0, 1, 1, 2),
			rect(0, 1, 1, 2),
		},
		{
			rect(-1, 1, 0, 2),
			rect(0, 0, 1, 1),
		},
	}

	for i, testI := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			err := test.ValidateBoundsMatch(testI.actual, testI.expected)

			// Verify result
			if err != nil {
				t.Errorf("SUT returned an error: %v", err)
			}
		})
	}
}

func TestValidateBoundsMatch_WhenTheyDontMatch(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		actual   image.Rectangle
		expected image.Rectangle
	}{
		{
			rect(0, 0, 0, 0),
			rect(1, 0, 0, 0),
		},
		{
			rect(0, 0, 0, 0),
			rect(0, 1, 0, 0),
		},
		{
			rect(0, 0, 0, 0),
			rect(0, 0, 1, 0),
		},
		{
			rect(0, 0, 0, 0),
			rect(0, 0, 0, 1),
		},
	}

	for i, testI := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			err := test.ValidateBoundsMatch(testI.actual, testI.expected)

			// Verify result
			if err == nil {
				t.Errorf("Expected SUT to throw an error.")
			}
		})
	}
}

func rect(X1, Y1, X2, Y2 int) image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: X1, Y: Y1},
		Max: image.Point{X: X2, Y: Y2},
	}
}
