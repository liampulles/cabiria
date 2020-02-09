package intertitle_test

import (
	"fmt"
	"image"
	"image/color"
	"reflect"
	"testing"

	"github.com/liampulles/cabiria/pkg/ml/test"

	"github.com/liampulles/cabiria/pkg/ml"

	"github.com/liampulles/cabiria/pkg/intertitle"

	imageIntertitle "github.com/liampulles/cabiria/pkg/image/intertitle"
)

func TestLoad_WhenNotAvailable(t *testing.T) {
	// Exercise SUT
	_, err := intertitle.Load("/this/path/does/not/exist")

	// Verify results
	if err == nil {
		t.Errorf("Expected SUT to return an error, but none was returned.")
	}
}

func TestLoad_WhenAvailable(t *testing.T) {
	// Setup fixture
	expected := intertitle.Predictor{
		Predictor: &ml.KNNClassifier{
			K: 1,
			Points: []ml.Sample{{
				Input:  ml.Datum{-1, 1},
				Output: ml.Datum{0, 0},
			}},
		},
	}

	// Exercise SUT
	actual, err := intertitle.Load("testdata/toLoad.model")

	// Verify results
	if err != nil {
		t.Errorf("SUT returned an error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Results differ: Expected %v, Actual %v", expected, actual)
	}
}

func TestPredict(t *testing.T) {
	// Setup fixture
	predictor := testPredictor()
	var tests = []struct {
		frames   []image.Image
		expected []bool
	}{
		{
			nil,
			intertitles(),
		},
		{
			frames(),
			intertitles(),
		},
		{
			frames(
				whiteImage(),
			),
			intertitles(1),
		},
		{
			frames(
				blackImage(),
			),
			intertitles(0),
		},
		{
			frames(
				blackImage(),
				whiteImage(),
				whiteImage(),
				blackImage(),
				blackImage(),
			),
			intertitles(0, 1, 1, 0, 0),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual, err := predictor.Predict(test.frames)

			// Verify result
			if err != nil {
				t.Errorf("SUT returned an error: %v", err)
			}
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Result differs. Actual: %v, Expected %v", actual, test.expected)
			}
		})
	}
}

func testPredictor() intertitle.Predictor {
	dummy := test.DummyPredictor{}
	dummy.Fit([]ml.Sample{
		{
			Input:  whiteImageStats().AsInput(),
			Output: ml.Datum{0.6},
		},
		{
			Input:  blackImageStats().AsInput(),
			Output: ml.Datum{0.3},
		},
	})
	return intertitle.Predictor{
		Predictor: &dummy,
	}
}

func blackImageStats() imageIntertitle.IntensityStats {
	return imageIntertitle.IntensityStats{
		AvgIntensity:       0,
		LowerAvgIntensity:  0,
		MiddleAvgIntensity: 0,
		UpperAvgIntensity:  0,
		ProportionLower:    0.9998000599820054,
		ProportionMiddle:   0.00009997000899730082,
		ProportionUpper:    0.00009997000899730082,
	}
}

func whiteImageStats() imageIntertitle.IntensityStats {
	return imageIntertitle.IntensityStats{
		AvgIntensity:       0.9997000899730081,
		LowerAvgIntensity:  0,
		MiddleAvgIntensity: 0,
		UpperAvgIntensity:  0.9999000099990001,
		ProportionLower:    0.00009997000899730082,
		ProportionMiddle:   0.00009997000899730082,
		ProportionUpper:    0.9998000599820054,
	}
}

func blackImage() image.Image {
	img := image.NewRGBA(rect(0, 0, 1, 1))
	img.Set(0, 0, color.Black)
	return img
}

func whiteImage() image.Image {
	img := image.NewRGBA(rect(0, 0, 1, 1))
	img.Set(0, 0, color.White)
	return img
}

func rect(X1, Y1, X2, Y2 int) image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: X1, Y: Y1},
		Max: image.Point{X: X2, Y: Y2},
	}
}

func intertitles(onOff ...int) []bool {
	result := make([]bool, len(onOff))
	for i, elem := range onOff {
		result[i] = elem == 1
	}
	return result
}

func frames(images ...image.Image) []image.Image {
	result := make([]image.Image, 0)
	return append(result, images...)
}
