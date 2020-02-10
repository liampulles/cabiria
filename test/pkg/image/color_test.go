package image_test

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"testing"

	cabiriaImage "github.com/liampulles/cabiria/pkg/image"
	cabiriaImageTest "github.com/liampulles/cabiria/pkg/image/test"
)

func TestGetForegroundAndBackground(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		img                image.Image
		expectedForeground color.Color
		expectedBackground color.Color
	}{
		// Single case
		{
			singleColorImage(color.Black),
			color.Black,
			color.Black,
		},
		{
			singleColorImage(color.White),
			color.White,
			color.White,
		},
		{
			singleColorImage(darkPink()),
			darkPink(),
			darkPink(),
		},
		// Non-tinted image
		{
			loadImage("testdata/viking.png"),
			colorFromRGB(177, 174, 172),
			colorFromRGB(1, 1, 1),
		},
		// Tinted image
		{
			loadImage("testdata/godard.png"),
			colorFromRGB(146, 3, 2),
			colorFromRGB(2, 0, 0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actualForeground, actualBackground, err := cabiriaImage.GetForegroundAndBackground(test.img)

			// Verify result
			if err != nil {
				t.Errorf("SUT returned an error: %v", err)
			}
			if err := cabiriaImageTest.CompareColor(actualForeground, test.expectedForeground); err != nil {
				t.Errorf("Unexpected result for foreground: %v", err)
			}
			if err := cabiriaImageTest.CompareColor(actualBackground, test.expectedBackground); err != nil {
				t.Errorf("Unexpected result for background: %v", err)
			}
		})
	}
}

func loadImage(path string) image.Image {
	img, err := cabiriaImage.GetPNG(path)
	if err != nil {
		panic(err)
	}
	return img
}

func emptyImage() image.Image {
	return image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{0, 0},
	})
}

func singleColorImage(col color.Color) image.Image {
	img := image.NewRGBA(image.Rectangle{
		Min: image.Point{0, 0},
		Max: image.Point{1, 2},
	})
	img.Set(0, 0, col)
	img.Set(0, 1, col)
	return img
}

func darkPink() color.Color {
	return color.RGBA{
		R: uint8(50),
		G: uint8(10),
		B: uint8(10),
		A: math.MaxUint8,
	}
}
func colorFromRGB(r, g, b uint8) color.Color {
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}
