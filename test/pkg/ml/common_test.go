package ml_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/liampulles/cabiria/pkg/ml"
)

func TestTest_WhenGivenValidInput_ExpectPass(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		clsFixture      ml.ClassificationPredictor
		testDataFixture []ml.ClassificationSample
		expected        float64
	}{
		// No input - we cannot say.
		{
			MockClassifier{},
			[]ml.ClassificationSample{},
			math.NaN(),
		},
		// -> Variation of above
		{
			MockClassifier{},
			nil,
			math.NaN(),
		},
		// Single fail - all fail
		{
			MockClassifier{},
			[]ml.ClassificationSample{
				falseNegative(),
			},
			0.0,
		},
		// Single pass - all pass
		{
			MockClassifier{},
			[]ml.ClassificationSample{
				truePositive(),
			},
			1.0,
		},
		// Multiple fail - all fail
		{
			MockClassifier{},
			[]ml.ClassificationSample{
				falseNegative(),
				falsePositive(),
				falseNegative(),
			},
			0.0,
		},
		// Multiple pass - all fail
		{
			MockClassifier{},
			[]ml.ClassificationSample{
				trueNegative(),
				truePositive(),
				trueNegative(),
			},
			1.0,
		},
		// Mixed bag - mixed results
		{
			MockClassifier{},
			[]ml.ClassificationSample{
				trueNegative(),
				falseNegative(),
				truePositive(),
				falsePositive(),
				trueNegative(),
			},
			0.6,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual, err := ml.Test(test.clsFixture, test.testDataFixture)

			// Verify result (must be very close)
			if err != nil {
				t.Errorf("SUT threw an error: %v", err)
			}
			if math.IsNaN(test.expected) {
				if !math.IsNaN(actual) {
					t.Errorf("Unexpected result.\nExpected: %f\nActual: %f", test.expected, actual)
				}
			} else if actual != test.expected {
				t.Errorf("Unexpected result.\nExpected: %f\nActual: %f", test.expected, actual)
			}
		})
	}
}

func TestSplit(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		samplesFixture    []ml.ClassificationSample
		splitFixture      float64
		expectedTrainSize int
		expectedTestSize  int
	}{
		// Empty input - Empty outputs.
		{
			[]ml.ClassificationSample{},
			0.5,
			0,
			0,
		},
		// Single input when split > 0 -> all in Train
		{
			[]ml.ClassificationSample{
				aSample(),
			},
			0.01,
			1,
			0,
		},
		// Single input when split == 0 -> all in Test
		{
			[]ml.ClassificationSample{
				aSample(),
			},
			0.0,
			0,
			1,
		},
		// Multiple input when split == 0 -> all in Test
		{
			[]ml.ClassificationSample{
				aSample(),
				aSample(),
				aSample(),
			},
			0.0,
			0,
			3,
		},
		// Multiple input when split == 1 -> all in Train
		{
			[]ml.ClassificationSample{
				aSample(),
				aSample(),
				aSample(),
			},
			1.0,
			3,
			0,
		},
		// Odd input when split == 0.5 -> extra in Train
		{
			[]ml.ClassificationSample{
				aSample(),
				aSample(),
				aSample(),
			},
			0.5,
			2,
			1,
		},
		// Even input when split == 0.5 -> even amount
		{
			[]ml.ClassificationSample{
				aSample(),
				aSample(),
				aSample(),
				aSample(),
				aSample(),
				aSample(),
				aSample(),
				aSample(),
			},
			0.5,
			4,
			4,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			trainSplitActual, testSplitActual := ml.Split(test.samplesFixture, test.splitFixture)

			// Verify result
			if len(trainSplitActual) != test.expectedTrainSize {
				t.Errorf("Unexpected split. Actual training set length: %d, Expected: %d", len(trainSplitActual), test.expectedTrainSize)
			}
			if len(testSplitActual) != test.expectedTestSize {
				t.Errorf("Unexpected split. Actual test set length: %d, Expected: %d", len(testSplitActual), test.expectedTestSize)
			}
		})
	}
}

func TestMatch_WhenDifferingInputLength_ExpectFail(t *testing.T) {
	// Setup fixture
	a := []uint{0}
	b := []uint{0, 0}

	// Exercise SUT
	_, err := ml.Match(a, b)

	// Verify result
	if err == nil {
		t.Errorf("Expected SUT to throw error")
	}
}

func TestMatch_WhenGivenValidInput_ExpectToPass(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		a        ml.ClassificationOutput
		b        ml.ClassificationOutput
		expected bool
	}{
		{
			[]uint{},
			[]uint{},
			true,
		},
		{
			[]uint{0},
			[]uint{0},
			true,
		},
		{
			[]uint{0},
			[]uint{1},
			false,
		},
		{
			[]uint{0, 1},
			[]uint{0, 1},
			true,
		},
		{
			[]uint{1, 0},
			[]uint{0, 1},
			false,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual, err := ml.Match(test.a, test.b)

			// Verify result
			if err != nil {
				t.Errorf("SUT threw an error: %v", err)
			}
			if actual != test.expected {
				t.Errorf("Unexpected result.\nExpected: %v\nActual: %v", test.expected, actual)
			}
		})
	}
}

type MockClassifier struct{}

func (mc MockClassifier) Fit(samples []ml.ClassificationSample) error {
	panic(fmt.Errorf("SUT should not call Fit"))
}

func (mc MockClassifier) Predict(input []ml.Input) ([]ml.ClassificationOutput, error) {
	results := make([]ml.ClassificationOutput, len(input))
	for i, elem := range input {
		result, err := mc.PredictSingle(elem)
		if err != nil {
			panic(err)
		}
		results[i] = result
	}
	return results, nil
}

func (mc MockClassifier) PredictSingle(input ml.Input) (ml.ClassificationOutput, error) {
	if input[0] < 0.0 {
		return []uint{0}, nil
	}
	return []uint{1}, nil
}

func (mc MockClassifier) Save(path string) error {
	panic(fmt.Errorf("SUT should not call Save"))
}

func aSample() ml.ClassificationSample {
	return sample([]float64{0.0}, []uint{0})
}

func truePositive() ml.ClassificationSample {
	return sample([]float64{1.0}, []uint{1})
}

func trueNegative() ml.ClassificationSample {
	return sample([]float64{-1.0}, []uint{0})
}

func falsePositive() ml.ClassificationSample {
	return sample([]float64{-1.0}, []uint{1})
}

func falseNegative() ml.ClassificationSample {
	return sample([]float64{1.0}, []uint{0})
}
