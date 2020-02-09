package test_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/liampulles/cabiria/pkg/ml"
	"github.com/liampulles/cabiria/pkg/ml/test"
)

func TestFit(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture  test.DummyPredictor
		samples  []ml.Sample
		expected test.DummyPredictor
	}{
		// Empty cases
		{
			dummy(nil),
			nil,
			dummy(map[string]ml.Datum{}),
		},
		{
			dummy(map[string]ml.Datum{}),
			nil,
			dummy(map[string]ml.Datum{}),
		},
		{
			dummy(map[string]ml.Datum{}),
			samples(),
			dummy(map[string]ml.Datum{}),
		},
		// Single sample
		{
			dummy(map[string]ml.Datum{}),
			samples(
				sample(datum(1.01), datum(2.2)),
			),
			dummy(map[string]ml.Datum{
				"1.01": datum(2.2),
			}),
		},
		{
			dummy(map[string]ml.Datum{}),
			samples(
				sample(datum(1.01, 3.023), datum(2.2, 4.002, 5.0101)),
			),
			dummy(map[string]ml.Datum{
				"1.01,3.023": datum(2.2, 4.002, 5.0101),
			}),
		},
		// Many samples
		{
			dummy(map[string]ml.Datum{}),
			samples(
				sample(datum(1.01, 3.023), datum(2.2, 4.002, 5.0101)),
				sample(datum(2.01, 4.251, 2.0005), datum(7.06)),
			),
			dummy(map[string]ml.Datum{
				"1.01,3.023":        datum(2.2, 4.002, 5.0101),
				"2.01,4.251,2.0005": datum(7.06),
			}),
		},
	}

	for i, testI := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			err := testI.fixture.Fit(testI.samples)

			// Verify result
			if err != nil {
				t.Errorf("SUT threw an error: %v", err)
			}
			if !reflect.DeepEqual(testI.fixture, testI.expected) {
				t.Errorf("Unexpected result.\nExpected: %v\nActual: %v", testI.expected, testI.fixture)
			}
		})
	}
}

func TestPredict_WhenFittedForInput(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture  test.DummyPredictor
		input    []ml.Datum
		expected []ml.Datum
	}{
		// Empty case
		{
			dummy(map[string]ml.Datum{}),
			datums(),
			datums(),
		},
		// Single case
		{
			dummy(map[string]ml.Datum{
				"1.01,3.023": datum(2.2, 4.002, 5.0101),
			}),
			datums(
				datum(1.01, 3.023),
			),
			datums(
				datum(2.2, 4.002, 5.0101),
			),
		},
		{
			dummy(map[string]ml.Datum{
				"4.2":        datum(8.7),
				"1.01,3.023": datum(2.2, 4.002, 5.0101),
			}),
			datums(
				datum(1.01, 3.023),
			),
			datums(
				datum(2.2, 4.002, 5.0101),
			),
		},
		// Many case
		{
			dummy(map[string]ml.Datum{
				"4.2":        datum(8.7),
				"1.01,3.023": datum(2.2, 4.002, 5.0101),
			}),
			datums(
				datum(4.2),
				datum(1.01, 3.023),
			),
			datums(
				datum(8.7),
				datum(2.2, 4.002, 5.0101),
			),
		},
		{
			dummy(map[string]ml.Datum{
				"5.7,8.2":    datum(1.01, 8.53),
				"4.2":        datum(8.7),
				"1.01,3.023": datum(2.2, 4.002, 5.0101),
			}),
			datums(
				datum(4.2),
				datum(1.01, 3.023),
			),
			datums(
				datum(8.7),
				datum(2.2, 4.002, 5.0101),
			),
		},
	}

	for i, testI := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual, err := testI.fixture.Predict(testI.input)

			// Verify result
			if err != nil {
				t.Errorf("SUT threw an error: %v", err)
			}
			if !reflect.DeepEqual(actual, testI.expected) {
				t.Errorf("Unexpected result.\nExpected: %v\nActual: %v", testI.expected, actual)
			}
		})
	}
}

func TestPredict_WhenNotFittedForInput(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture test.DummyPredictor
		input   []ml.Datum
	}{
		// Single case
		{
			dummy(map[string]ml.Datum{
				"1.01,3.023": datum(2.2, 4.002, 5.0101),
			}),
			datums(
				datum(2.2, 3.023),
			),
		},
		{
			dummy(map[string]ml.Datum{
				"4.2":        datum(8.7),
				"1.01,3.023": datum(2.2, 4.002, 5.0101),
			}),
			datums(
				datum(2.2, 3.023),
			),
		},
		// Many case
		{
			dummy(map[string]ml.Datum{
				"4.2":        datum(8.7),
				"1.01,3.023": datum(2.2, 4.002, 5.0101),
			}),
			datums(
				datum(4.2),
				datum(2.2, 3.023),
			),
		},
		{
			dummy(map[string]ml.Datum{
				"5.7,8.2":    datum(1.01, 8.53),
				"4.2":        datum(8.7),
				"1.01,3.023": datum(2.2, 4.002, 5.0101),
			}),
			datums(
				datum(4.2),
				datum(2.2, 3.023),
			),
		},
	}

	for i, testI := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			_, err := testI.fixture.Predict(testI.input)

			// Verify result
			if err == nil {
				t.Errorf("Expected SUT to return an error.")
			}
		})
	}
}

func TestPredictSingle_WhenFittedForInput(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture  test.DummyPredictor
		input    ml.Datum
		expected ml.Datum
	}{
		{
			dummy(map[string]ml.Datum{
				"1.01,3.023": datum(2.2, 4.002, 5.0101),
			}),
			datum(1.01, 3.023),
			datum(2.2, 4.002, 5.0101),
		},
		{
			dummy(map[string]ml.Datum{
				"4.2":        datum(8.7),
				"1.01,3.023": datum(2.2, 4.002, 5.0101),
			}),
			datum(1.01, 3.023),
			datum(2.2, 4.002, 5.0101),
		},
	}

	for i, testI := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actual, err := testI.fixture.PredictSingle(testI.input)

			// Verify result
			if err != nil {
				t.Errorf("SUT threw an error: %v", err)
			}
			if !reflect.DeepEqual(actual, testI.expected) {
				t.Errorf("Unexpected result.\nExpected: %v\nActual: %v", testI.expected, actual)
			}
		})
	}
}

func TestPredictSingle_WhenNotFittedForInput(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture test.DummyPredictor
		input   ml.Datum
	}{
		{
			dummy(map[string]ml.Datum{
				"1.01,3.023": datum(2.2, 4.002, 5.0101),
			}),
			datum(2.2, 3.023),
		},
		{
			dummy(map[string]ml.Datum{
				"4.2":        datum(8.7),
				"1.01,3.023": datum(2.2, 4.002, 5.0101),
			}),
			datum(2.2, 3.023),
		},
	}

	for i, testI := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			_, err := testI.fixture.PredictSingle(testI.input)

			// Verify result
			if err == nil {
				t.Errorf("Expected SUT to return an error.")
			}
		})
	}
}

func TestSave(t *testing.T) {
	// Setup fixture
	predictor := dummy(map[string]ml.Datum{
		"4.2":        datum(8.7),
		"1.01,3.023": datum(2.2, 4.002, 5.0101),
	})

	// Exercise SUT
	err := predictor.Save("a.model")

	// Verify result
	if err == nil {
		t.Errorf("Expected SUT to return an error.")
	}
}

func dummy(mock map[string]ml.Datum) test.DummyPredictor {
	return test.DummyPredictor{
		MockData: mock,
	}
}

func samples(samples ...ml.Sample) []ml.Sample {
	result := make([]ml.Sample, 0)
	return append(result, samples...)
}

func sample(input, output ml.Datum) ml.Sample {
	return ml.Sample{
		Input:  input,
		Output: output,
	}
}

func datums(items ...ml.Datum) []ml.Datum {
	result := make([]ml.Datum, 0)
	return append(result, items...)
}

func datum(items ...float64) []float64 {
	return items
}
