package cluster_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/liampulles/cabiria/pkg/ml"
	"github.com/liampulles/cabiria/pkg/ml/cluster"
)

func TestKMeansClassifier_Fit_WhenClassifierAndInputIsValid_ExpectPass(t *testing.T) {
	// Setup fixture, expectations
	var tests = []struct {
		kMeansFixture     *cluster.KMeansClassifier
		inputFixture      []ml.Datum
		expectedCounts    []int
		expectedCentroids []ml.Datum
	}{
		// K = 1
		// -> One value
		{
			cluster.NewKMeansClassifier(1, 999999),
			datums(
				datum(),
			),
			counts(1),
			datums(
				datum(),
			),
		},
		{
			cluster.NewKMeansClassifier(1, 999999),
			datums(
				datum(0.0),
			),
			counts(1),
			datums(
				datum(0.0),
			),
		},
		{
			cluster.NewKMeansClassifier(1, 999999),
			datums(
				datum(1.0),
			),
			counts(1),
			datums(
				datum(1.0),
			),
		},
		{
			cluster.NewKMeansClassifier(1, 999999),
			datums(
				datum(0.0, 0.0),
			),
			counts(1),
			datums(
				datum(0.0, 0.0),
			),
		},
		{
			cluster.NewKMeansClassifier(1, 999999),
			datums(
				datum(1.0, 2.0),
			),
			counts(1),
			datums(
				datum(1.0, 2.0),
			),
		},
		// -> Many values
		{
			cluster.NewKMeansClassifier(1, 999999),
			datums(
				datum(),
				datum(),
				datum(),
			),
			counts(3),
			datums(
				datum(),
			),
		},
		{
			cluster.NewKMeansClassifier(1, 999999),
			datums(
				datum(0.0),
				datum(0.0),
				datum(0.0),
			),
			counts(3),
			datums(
				datum(0.0),
			),
		},
		{
			cluster.NewKMeansClassifier(1, 999999),
			datums(
				datum(1.0),
				datum(1.0),
				datum(1.0),
			),
			counts(3),
			datums(
				datum(1.0),
			),
		},
		{
			cluster.NewKMeansClassifier(1, 999999),
			datums(
				datum(1.0),
				datum(2.0),
				datum(3.0),
			),
			counts(3),
			datums(
				datum(2.0),
			),
		},
		{
			cluster.NewKMeansClassifier(1, 999999),
			datums(
				datum(2.0, 4.0),
				datum(2.0, 4.0),
				datum(2.0, 4.0),
			),
			counts(3),
			datums(
				datum(2.0, 4.0),
			),
		},
		{
			cluster.NewKMeansClassifier(1, 999999),
			datums(
				datum(1.0, 2.0),
				datum(2.0, 4.0),
				datum(3.0, 3.0),
			),
			counts(3),
			datums(
				datum(2.0, 3.0),
			),
		},
		// K = 3
		// -> 3 values
		{
			cluster.NewKMeansClassifier(3, 999999),
			datums(
				datum(),
				datum(),
				datum(),
			),
			counts(3, 0, 0),
			datums(
				datum(),
				datum(),
				datum(),
			),
		},
		{
			cluster.NewKMeansClassifier(3, 999999),
			datums(
				datum(0.0),
				datum(0.0),
				datum(0.0),
			),
			counts(3, 0, 0),
			datums(
				datum(0.0),
				datum(0.0),
				datum(0.0),
			),
		},
		{
			cluster.NewKMeansClassifier(3, 999999),
			datums(
				datum(1.0),
				datum(2.0),
				datum(3.0),
			),
			counts(1, 1, 1),
			datums(
				datum(3.0),
				datum(1.0),
				datum(2.0),
			),
		},
		{
			cluster.NewKMeansClassifier(3, 999999),
			datums(
				datum(1.0, 2.0),
				datum(2.0, 7.0),
				datum(3.0, 1.0),
			),
			counts(1, 1, 1),
			datums(
				datum(3.0, 1.0),
				datum(2.0, 7.0),
				datum(1.0, 2.0),
			),
		},
		// Many values
		{
			cluster.NewKMeansClassifier(3, 999999),
			datums(
				datum(), datum(), datum(),
				datum(), datum(), datum(),
				datum(), datum(), datum(),
			),
			counts(9, 0, 0),
			datums(
				datum(),
				datum(),
				datum(),
			),
		},
		{
			cluster.NewKMeansClassifier(3, 999999),
			datums(
				datum(0.0), datum(0.0), datum(0.0),
				datum(0.0), datum(0.0), datum(0.0),
				datum(0.0), datum(0.0), datum(0.0),
			),
			counts(9, 0, 0),
			datums(
				datum(0.0),
				datum(0.0),
				datum(0.0),
			),
		},
		{
			cluster.NewKMeansClassifier(3, 999999),
			datums(
				datum(1.0), datum(2.0), datum(3.0),
				datum(3.0), datum(1.0), datum(2.0),
				datum(2.0), datum(3.0), datum(1.0),
			),
			counts(3, 3, 3),
			datums(
				datum(2.0),
				datum(1.0),
				datum(3.0),
			),
		},
		{
			cluster.NewKMeansClassifier(3, 999999),
			datums(
				datum(1.0, 1.0), datum(2.0, 2.0), datum(3.0, 3.0),
				datum(3.0, 3.0), datum(1.0, 1.0), datum(2.0, 2.0),
				datum(2.0, 2.0), datum(3.0, 3.0), datum(1.0, 1.0),
			),
			counts(3, 3, 3),
			datums(
				datum(2.0, 2.0),
				datum(1.0, 1.0),
				datum(3.0, 3.0),
			),
		},
		{
			cluster.NewKMeansClassifier(3, 999999),
			datums(
				datum(0.5, 1.2), datum(2.4, 1.9), datum(3.2, 2.7),
				datum(3.4, 3.1), datum(1.2, 0.8), datum(2.1, 2.4),
				datum(1.8, 1.9), datum(2.9, 3.9), datum(0.9, 1.9),
			),
			counts(3, 3, 3),
			datums(
				datum(2.1, 2.0666666666666664),
				datum(0.8666666666666667, 1.3),
				datum(3.1666666666666665, 3.233333333333334),
			),
		},
		// Hits iteration cap
		{
			cluster.NewKMeansClassifier(3, 1),
			datums(
				datum(0.5, 1.2), datum(2.4, 1.9), datum(3.2, 2.7),
				datum(3.4, 3.1), datum(1.2, 0.8), datum(2.1, 2.4),
				datum(1.8, 1.9), datum(2.9, 3.9), datum(0.9, 1.9),
			),
			counts(4, 3, 2),
			datums(
				datum(2.375, 2.225),
				datum(0.8666666666666667, 1.3),
				datum(3.15, 3.5),
			),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			actualCounts, _, err := test.kMeansFixture.Fit(test.inputFixture)

			// Verify result
			if err != nil {
				t.Errorf("SUT threw error: %v", err)
			}
			if !reflect.DeepEqual(actualCounts, test.expectedCounts) {
				t.Errorf("Unexpected result.\nExpected counts: %v\nActual counts: %v", test.expectedCounts, actualCounts)
			}
			actualCentroids := test.kMeansFixture.ClusterCentroids()
			if !reflect.DeepEqual(actualCentroids, test.expectedCentroids) {
				t.Errorf("Unexpected result.\nExpected centroids: %v\nActual centroids: %v", test.expectedCentroids, actualCentroids)
			}
		})
	}
}

func TestKMeansClassifier_Fit_WhenInputHasElementsOfVaryingSize(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		input []ml.Datum
		k     int
	}{
		{
			datums(
				datum(),
				datum(1.0),
			),
			2,
		},
		{
			datums(
				datum(1.0),
				datum(1.0),
				datum(1.0),
				datum(1.0, 2.0),
				datum(1.0),
			),
			3,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			_, _, err := cluster.NewKMeansClassifier(test.k, 99999999999).Fit(test.input)

			// Verify result
			if err == nil {
				t.Errorf("Expected SUT to return an error")
			}
		})
	}
}

func TestKMeansClassifier_Fit_WhenInputLengthLessThanK(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		input []ml.Datum
		k     int
	}{
		{
			datums(),
			1,
		},
		{
			datums(
				datum(1.0),
			),
			2,
		},
		{
			datums(
				datum(1.0),
				datum(2.0),
			),
			5,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			_, _, err := cluster.NewKMeansClassifier(test.k, 99999999999).Fit(test.input)

			// Verify result
			if err == nil {
				t.Errorf("Expected SUT to return an error")
			}
		})
	}
}

func datums(datums ...ml.Datum) []ml.Datum {
	result := make([]ml.Datum, 0)
	return append(result, datums...)
}

func datum(val ...float64) ml.Datum {
	result := make(ml.Datum, 0)
	return append(result, val...)
}

func counts(val ...int) []int {
	return val
}
