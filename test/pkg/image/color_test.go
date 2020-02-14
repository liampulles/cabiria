package image_test

import (
	"fmt"
	"image"
	"image/color"
	"testing"

	"github.com/lucasb-eyer/go-colorful"

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
			color.White,
			colorFromHSV(0.0, 0.0, 0.1),
		},
		{
			singleColorImage(color.White),
			color.White,
			colorFromHSV(0.0, 0.0, 0.1),
		},
		{
			singleColorImage(colorFromHSV(60.0, 1.0, 0.5)),
			colorFromHSV(60.0, 1.0, 1.0),
			colorFromHSV(60.0, 1.0, 0.1),
		},
		// Non-tinted image
		{
			loadImage("testdata/viking.png"),
			color.RGBA64{
				R: 65535,
				G: 64424,
				B: 63684,
				A: 65535,
			},
			color.RGBA64{
				R: 6554,
				G: 6554,
				B: 6554,
				A: 65535,
			},
		},
		// Tinted image
		{
			loadImage("testdata/godard.png"),
			color.RGBA64{
				R: 65535,
				G: 1347,
				B: 898,
				A: 65535,
			},
			color.RGBA64{
				R: 6554,
				G: 0,
				B: 0,
				A: 65535,
			},
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

func TestChangeValue(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		col      color.Color
		value    float64
		expected color.Color
	}{
		// Min value
		{
			colorFromHSV(0.0, 0.0, 0.0),
			0.0,
			colorFromHSV(0.0, 0.0, 0.0),
		},
		{
			colorFromHSV(0.0, 0.0, 1.0),
			0.0,
			colorFromHSV(0.0, 0.0, 0.0),
		},
		{
			colorFromHSV(120.0, 0.5, 0.6),
			0.0,
			colorFromHSV(120.0, 0.5, 0.0),
		},
		// Min value clamped
		{
			colorFromHSV(0.0, 0.0, 0.0),
			-1.0,
			colorFromHSV(0.0, 0.0, 0.0),
		},
		{
			colorFromHSV(0.0, 0.0, 1.0),
			-1.0,
			colorFromHSV(0.0, 0.0, 0.0),
		},
		{
			colorFromHSV(120.0, 0.5, 0.6),
			-1.0,
			colorFromHSV(120.0, 0.5, 0.0),
		},
		// Max value
		{
			colorFromHSV(0.0, 0.0, 0.0),
			1.0,
			colorFromHSV(0.0, 0.0, 1.0),
		},
		{
			colorFromHSV(0.0, 0.0, 1.0),
			1.0,
			colorFromHSV(0.0, 0.0, 1.0),
		},
		{
			colorFromHSV(120.0, 0.5, 0.6),
			1.0,
			colorFromHSV(120.0, 0.5, 1.0),
		},
		// Max value clamped
		{
			colorFromHSV(0.0, 0.0, 0.0),
			2.0,
			colorFromHSV(0.0, 0.0, 1.0),
		},
		{
			colorFromHSV(0.0, 0.0, 1.0),
			2.0,
			colorFromHSV(0.0, 0.0, 1.0),
		},
		{
			colorFromHSV(120.0, 0.5, 0.6),
			2.0,
			colorFromHSV(120.0, 0.5, 1.0),
		},
		// Midrange value
		{
			colorFromHSV(0.0, 0.0, 0.0),
			0.3,
			colorFromHSV(0.0, 0.0, 0.3),
		},
		{
			colorFromHSV(0.0, 0.0, 1.0),
			0.3,
			colorFromHSV(0.0, 0.0, 0.3),
		},
		{
			colorFromHSV(120.0, 0.5, 0.6),
			0.4,
			colorFromHSV(120.0, 0.5, 0.4),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual := cabiriaImage.ChangeValue(test.col, test.value)

			// Verify result
			if err := cabiriaImageTest.CompareColor(actual, test.expected); err != nil {
				t.Errorf("Unexpected result: %v", err)
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

func colorFromRGB(r, g, b uint8) color.Color {
	return color.RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}

func colorFromHSV(h, s, v float64) color.Color {
	return colorful.Hsv(h, s, v)
}
