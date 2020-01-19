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
		clsFixture      ml.Predictor
		testDataFixture []ml.Sample
		expected        float64
	}{
		// No input - we cannot say.
		{
			MockClassifier{},
			[]ml.Sample{},
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
			[]ml.Sample{
				falseNegative(),
			},
			0.0,
		},
		// Single pass - all pass
		{
			MockClassifier{},
			[]ml.Sample{
				truePositive(),
			},
			1.0,
		},
		// Multiple fail - all fail
		{
			MockClassifier{},
			[]ml.Sample{
				falseNegative(),
				falsePositive(),
				falseNegative(),
			},
			0.0,
		},
		// Multiple pass - all fail
		{
			MockClassifier{},
			[]ml.Sample{
				trueNegative(),
				truePositive(),
				trueNegative(),
			},
			1.0,
		},
		// Mixed bag - mixed results
		{
			MockClassifier{},
			[]ml.Sample{
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
		samplesFixture    []ml.Sample
		splitFixture      float64
		expectedTrainSize int
		expectedTestSize  int
	}{
		// Empty input - Empty outputs.
		{
			[]ml.Sample{},
			0.5,
			0,
			0,
		},
		// Single input when split > 0 -> all in Train
		{
			[]ml.Sample{
				aSample(),
			},
			0.01,
			1,
			0,
		},
		// Single input when split == 0 -> all in Test
		{
			[]ml.Sample{
				aSample(),
			},
			0.0,
			0,
			1,
		},
		// Multiple input when split == 0 -> all in Test
		{
			[]ml.Sample{
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
			[]ml.Sample{
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
			[]ml.Sample{
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
			[]ml.Sample{
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
	a := []float64{0.0}
	b := []float64{0.0, 0.0}

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
		a        ml.Output
		b        ml.Output
		expected bool
	}{
		{
			[]float64{},
			[]float64{},
			true,
		},
		{
			[]float64{0},
			[]float64{0},
			true,
		},
		{
			[]float64{0},
			[]float64{1},
			false,
		},
		{
			[]float64{0, 1},
			[]float64{0, 1},
			true,
		},
		{
			[]float64{1, 0},
			[]float64{0, 1},
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

func (mc MockClassifier) Fit(samples []ml.Sample) error {
	panic(fmt.Errorf("SUT should not call Fit"))
}

func (mc MockClassifier) Predict(input []ml.Input) ([]ml.Output, error) {
	results := make([]ml.Output, len(input))
	for i, elem := range input {
		result, err := mc.PredictSingle(elem)
		if err != nil {
			panic(err)
		}
		results[i] = result
	}
	return results, nil
}

func (mc MockClassifier) PredictSingle(input ml.Input) (ml.Output, error) {
	if input[0] < 0.0 {
		return []float64{0}, nil
	}
	return []float64{1}, nil
}

func (mc MockClassifier) Save(path string) error {
	panic(fmt.Errorf("SUT should not call Save"))
}

func aSample() ml.Sample {
	return sample([]float64{0.0}, []float64{0})
}

func truePositive() ml.Sample {
	return sample([]float64{1.0}, []float64{1})
}

func trueNegative() ml.Sample {
	return sample([]float64{-1.0}, []float64{0})
}

func falsePositive() ml.Sample {
	return sample([]float64{-1.0}, []float64{1})
}

func falseNegative() ml.Sample {
	return sample([]float64{1.0}, []float64{0})
}
