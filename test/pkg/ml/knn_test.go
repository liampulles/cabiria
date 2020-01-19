package ml_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/liampulles/cabiria/pkg/ml"
)

func TestKNNClassifier_Fit_WhenInputNil_ExpectFail(t *testing.T) {
	// Setup fixture

	// Exercise SUT
	knn := ml.NewKNNClassifier()
	err := knn.Fit(nil)

	// Verify result
	if err == nil {
		t.Errorf("Expected SUT to throw error, but none thrown")
	}
}

func TestKNNClassifier_Fit_WhenInputEmpty_ExpectFail(t *testing.T) {
	// Setup fixture
	input := make([]ml.ClassificationSample, 0)

	// Exercise SUT
	knn := ml.NewKNNClassifier()
	err := knn.Fit(input)

	// Verify result
	if err == nil {
		t.Errorf("Expected SUT to throw error, but none thrown")
	}
}

func TestKNNClassifier_Fit_WhenInputNotEmpty_ExpectPass(t *testing.T) {
	// Setup fixture
	input := make([]ml.ClassificationSample, 1)
	input[0] = sample(topLeftInput(0.0, 0.0), topLeftClass())

	// Setup expectations
	expectedElem := sample(topLeftInput(0.0, 0.0), topLeftClass())

	// Exercise SUT
	knn := ml.NewKNNClassifier()
	err := knn.Fit(input)

	// Verify result
	if err != nil {
		t.Errorf("SUT threw error: %v", err)
	}
	if len(knn.Points) != len(input) {
		t.Errorf("Points has wrong length.\nActual: %d\nExpected: %d", len(knn.Points), len(input))
	}
	if !reflect.DeepEqual(knn.Points[0], expectedElem) {
		t.Errorf("Points has wrong element.\nActual: %v\nExpected: %v", knn.Points[0], expectedElem)
	}
}

func TestKNNClassifier_Save_ExpectPass(t *testing.T) {
	// Setup fixture
	path := "testdata/testSave.model"

	// Exercise SUT
	knn := knnClassifier(sample(topLeftInput(0.0, 0.0), topLeftClass()))
	err := knn.Save(path)

	// Verify result
	if err != nil {
		t.Errorf("SUT threw error: %v", err)
	}
}

func Test_LoadKNNClassifier_ExpectPass(t *testing.T) {
	// Setup fixture
	path := "testdata/testLoad.model"

	// Exercise SUT
	_, err := ml.LoadKNNClassfier(path)

	// Verify result
	if err != nil {
		t.Errorf("SUT threw error: %v", err)
	}
}

func Test_KNNSaveAndLoad_ExpectLoadedToMatchSaved(t *testing.T) {
	// Setup fixture
	path := "testdata/testRoundtrip.model"
	knnFixture := knnClassifier(
		sample(bottomRightInput(0.0, 0.0), bottomRightClass()),
		sample(topRightInput(0.0, 0.0), topRightClass()),
		sample(topLeftInput(0.0, 0.0), topLeftClass()),
		sample(bottomLeftInput(0.0, 0.0), bottomLeftClass()),
	)

	// Setup expectations
	expected := knnClassifier(
		sample(bottomRightInput(0.0, 0.0), bottomRightClass()),
		sample(topRightInput(0.0, 0.0), topRightClass()),
		sample(topLeftInput(0.0, 0.0), topLeftClass()),
		sample(bottomLeftInput(0.0, 0.0), bottomLeftClass()),
	)

	// Exercise SUT
	err := knnFixture.Save(path)
	if err != nil {
		t.Errorf("SUT threw error: %v", err)
	}
	actual, err := ml.LoadKNNClassfier(path)
	if err != nil {
		t.Errorf("SUT threw error: %v", err)
	}

	// Verify result
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Unexpected result.\nExpected: %v\nActual: %v", expected, actual)
	}

}

func TestKNNClassifier_PredictSingle_WhenClassifierAndInputIsValid_ExpectPass(t *testing.T) {
	// Setup fixture, expectations
	var tests = []struct {
		knnFixture   *ml.KNNClassifier
		inputFixture ml.Input
		expected     ml.ClassificationOutput
	}{
		// Single sample with matching input -> Expect sample class
		{
			knnClassifier(
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
			),
			topLeftInput(0.0, 0.0),
			topLeftClass(),
		},
		// Single sample with non-matching input -> Expect sample class
		{
			knnClassifier(
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
			),
			bottomRightInput(0.0, 0.0),
			topLeftClass(),
		},
		// Multiple samples with a matching input -> Expect matching class
		{
			knnClassifier(
				sample(bottomRightInput(0.0, 0.0), bottomRightClass()),
				sample(topRightInput(0.0, 0.0), topRightClass()),
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
				sample(bottomLeftInput(0.0, 0.0), bottomLeftClass()),
			),
			topLeftInput(0.0, 0.0),
			topLeftClass(),
		},
		// Multiple samples with a non-matching input -> Expect closest class
		{
			knnClassifier(
				sample(bottomRightInput(0.0, 0.0), bottomRightClass()),
				sample(topRightInput(0.0, 0.0), topRightClass()),
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
				sample(bottomLeftInput(0.0, 0.0), bottomLeftClass()),
			),
			topLeftInput(0.5, -0.5),
			topLeftClass(),
		},
		// -> Variation of above
		{
			knnClassifier(
				sample(bottomRightInput(0.0, 0.0), bottomRightClass()),
				sample(topRightInput(0.0, 0.0), topRightClass()),
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
				sample(bottomLeftInput(0.0, 0.0), bottomLeftClass()),
			),
			bottomLeftInput(0.5, 0.5),
			bottomLeftClass(),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {

			// Exercise SUT
			actual, err := test.knnFixture.PredictSingle(test.inputFixture)

			// Verify result
			if err != nil {
				t.Errorf("SUT threw error: %v", err)
			}
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Unexpected result.\nExpected: %v\nActual: %v", test.expected, actual)
			}
		})
	}
}

func TestKNNClassifier_Predict_WhenClassifierAndInputIsValid_ExpectPass(t *testing.T) {
	// Setup fixture, expectations
	var tests = []struct {
		knnFixture   *ml.KNNClassifier
		inputFixture []ml.Input
		expected     []ml.ClassificationOutput
	}{
		// Single sample with matching input -> Expect sample class
		{
			knnClassifier(
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
			),
			[]ml.Input{
				topLeftInput(0.0, 0.0),
			},
			[]ml.ClassificationOutput{
				topLeftClass(),
			},
		},
		// Single sample with non-matching input -> Expect sample class
		{
			knnClassifier(
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
			),
			[]ml.Input{
				bottomRightInput(0.0, 0.0),
			},
			[]ml.ClassificationOutput{
				topLeftClass(),
			},
		},
		// Multiple samples with a matching input -> Expect matching class
		{
			knnClassifier(
				sample(bottomRightInput(0.0, 0.0), bottomRightClass()),
				sample(topRightInput(0.0, 0.0), topRightClass()),
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
				sample(bottomLeftInput(0.0, 0.0), bottomLeftClass()),
			),
			[]ml.Input{
				topLeftInput(0.0, 0.0),
			},
			[]ml.ClassificationOutput{
				topLeftClass(),
			},
		},
		// Multiple samples with a non-matching input -> Expect closest class
		{
			knnClassifier(
				sample(bottomRightInput(0.0, 0.0), bottomRightClass()),
				sample(topRightInput(0.0, 0.0), topRightClass()),
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
				sample(bottomLeftInput(0.0, 0.0), bottomLeftClass()),
			),
			[]ml.Input{
				topLeftInput(0.5, -0.5),
			},
			[]ml.ClassificationOutput{
				topLeftClass(),
			},
		},
		// -> Variation of above
		{
			knnClassifier(
				sample(bottomRightInput(0.0, 0.0), bottomRightClass()),
				sample(topRightInput(0.0, 0.0), topRightClass()),
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
				sample(bottomLeftInput(0.0, 0.0), bottomLeftClass()),
			),
			[]ml.Input{
				bottomLeftInput(0.5, 0.5),
			},
			[]ml.ClassificationOutput{
				bottomLeftClass(),
			},
		},
		// Single sample with multiple matching input -> Expect sample class
		{
			knnClassifier(
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
			),
			[]ml.Input{
				topLeftInput(0.0, 0.0),
				topLeftInput(0.0, 0.0),
				topLeftInput(0.0, 0.0),
			},
			[]ml.ClassificationOutput{
				topLeftClass(),
				topLeftClass(),
				topLeftClass(),
			},
		},
		// Single sample with multiple non-matching input -> Expect sample class
		{
			knnClassifier(
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
			),
			[]ml.Input{
				bottomLeftInput(0.0, 0.0),
				bottomRightInput(0.0, 0.0),
				topRightInput(0.0, 0.0),
			},
			[]ml.ClassificationOutput{
				topLeftClass(),
				topLeftClass(),
				topLeftClass(),
			},
		},
		// Multiple samples with mutiple matching input -> Expect matching class
		{
			knnClassifier(
				sample(bottomRightInput(0.0, 0.0), bottomRightClass()),
				sample(topRightInput(0.0, 0.0), topRightClass()),
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
				sample(bottomLeftInput(0.0, 0.0), bottomLeftClass()),
			),
			[]ml.Input{
				topLeftInput(0.0, 0.0),
				topLeftInput(0.0, 0.0),
				topLeftInput(0.0, 0.0),
			},
			[]ml.ClassificationOutput{
				topLeftClass(),
				topLeftClass(),
				topLeftClass(),
			},
		},
		// Multiple samples with multiple non-matching input -> Expect closest class
		{
			knnClassifier(
				sample(bottomRightInput(0.0, 0.0), bottomRightClass()),
				sample(topRightInput(0.0, 0.0), topRightClass()),
				sample(topLeftInput(0.0, 0.0), topLeftClass()),
				sample(bottomLeftInput(0.0, 0.0), bottomLeftClass()),
			),
			[]ml.Input{
				topLeftInput(0.5, -0.5),
				topRightInput(0.5, -0.5),
				bottomLeftInput(0.5, -0.5),
				bottomRightInput(0.5, -0.5),
			},
			[]ml.ClassificationOutput{
				topLeftClass(),
				topRightClass(),
				bottomLeftClass(),
				bottomRightClass(),
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {

			// Exercise SUT
			actual, err := test.knnFixture.Predict(test.inputFixture)

			// Verify result
			if err != nil {
				t.Errorf("SUT threw error: %v", err)
			}
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Unexpected result.\nExpected: %v\nActual: %v", test.expected, actual)
			}
		})
	}
}

func sample(input ml.Input, output ml.ClassificationOutput) ml.ClassificationSample {
	return ml.ClassificationSample{
		Input:  input,
		Output: output,
	}
}

func topLeftInput(varyX float64, varyY float64) ml.Input {
	return []float64{-1.0 + varyX, 1.0 + varyY}
}

func topRightInput(varyX float64, varyY float64) ml.Input {
	return []float64{1.0 + varyX, 1.0 + varyY}
}

func bottomLeftInput(varyX float64, varyY float64) ml.Input {
	return []float64{-1.0 + varyX, -1.0 + varyY}
}

func bottomRightInput(varyX float64, varyY float64) ml.Input {
	return []float64{1.0 + varyX, -1.0 + varyY}
}

func topLeftClass() ml.ClassificationOutput {
	return []uint{0, 0}
}

func topRightClass() ml.ClassificationOutput {
	return []uint{0, 1}
}

func bottomLeftClass() ml.ClassificationOutput {
	return []uint{1, 0}
}

func bottomRightClass() ml.ClassificationOutput {
	return []uint{1, 1}
}

func knnClassifier(samples ...ml.ClassificationSample) *ml.KNNClassifier {
	knn := ml.NewKNNClassifier()
	err := knn.Fit(samples)
	if err != nil {
		panic(err)
	}
	return knn
}
