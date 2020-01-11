package test

import (
	"path/filepath"
	"testing"

	"github.com/liampulles/cabiria/pkg/image/test"

	"github.com/liampulles/cabiria/pkg/image"
)

func TestCompare_SinglePixelImage_WhenImagesAreTheSame(t *testing.T) {
	// Setup fixture
	path := filepath.Join("testdata", "singlepixel_1.png")
	img1, err := image.GetImage(path)
	if err != nil {
		t.Errorf("Failed to load img1: %v", err)
	}
	img2, err := image.GetImage(path)
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
	img1, err := image.GetImage(path1)
	if err != nil {
		t.Errorf("Failed to load img1: %v", err)
	}
	path2 := filepath.Join("testdata", "singlepixel_2.png")
	img2, err := image.GetImage(path2)
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
	img1, err := image.GetImage(path)
	if err != nil {
		t.Errorf("Failed to load img1: %v", err)
	}
	img2, err := image.GetImage(path)
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
	img1, err := image.GetImage(path1)
	if err != nil {
		t.Errorf("Failed to load img1: %v", err)
	}
	path2 := filepath.Join("testdata", "normal_2.png")
	img2, err := image.GetImage(path2)
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
