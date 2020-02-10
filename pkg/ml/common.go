package ml

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"

	cabiriaMath "github.com/liampulles/cabiria/pkg/math"
)

// Split can be used to split training ang test data by a percentage.
//  The resulting arrays are pseudorandom in composition; not consistent
//  between runs.
func Split(samples []Sample, split float64) ([]Sample, []Sample) {
	rand.Shuffle(len(samples), func(i, j int) { samples[i], samples[j] = samples[j], samples[i] })
	cutoff := int(math.Ceil(float64(len(samples)) * cabiriaMath.ClampFloat64(split, 0.0, 1.0)))
	return samples[:cutoff], samples[cutoff:]
}

// Test runs test samples against a trained predictor to see how accurate it is.
//  The result is percentage accuracy.
func Test(cls Predictor, testData []Sample) (float64, error) {
	passed := 0
	for _, datum := range testData {
		prediction, err := cls.PredictSingle(datum.Input)
		if err != nil {
			return -1.0, err
		}
		match, err := Match(datum.Output, prediction)
		if err != nil {
			return -1.0, err
		}
		if match {
			passed++
		}
	}
	return float64(passed) / float64(len(testData)), nil
}

// Match returns true if the Datum' are the same, otherwise false.
//  An error is returned if the Datum' have different lengths.
func Match(actual Datum, expected Datum) (bool, error) {
	if len(actual) != len(expected) {
		return false, fmt.Errorf("Cannot Match outputs with different lengths. Actual length: %d, expected length: %d",
			len(actual), len(expected))
	}
	for i, actualI := range actual {
		if actualI != expected[i] {
			return false, nil
		}
	}
	return true, nil
}

// InitializeKMeans implements the init step for KMeans++ (at least,
//  it implements whatever this GeeksForGeeks article says:
//	https://www.geeksforgeeks.org/ml-k-means-algorithm/)
func InitializeKMeans(input []Datum, k int) ([]Datum, error) {
	if k == 0 {
		return []Datum{}, nil
	}
	if len(input) < k {
		return nil, fmt.Errorf("need at least k input to initialize KMeans")
	}
	centroids := make([]Datum, 1)
	// Effectively choose a pseudo-random element, but don;t change in-between
	// runs so our tests remian deterministic
	centroids[0] = input[1820244659%len(input)]
	for len(centroids) < k {
		newCentroid, err := computeNextBestCentroid(centroids, input)
		if err != nil {
			return nil, err
		}
		centroids = append(centroids, newCentroid)
	}
	return centroids, nil
}

// AsCSV maps a datum to a CSV line for ML purposes.
func (d Datum) AsCSV() string {
	result := ""
	for i, elem := range d {
		if i > 0 {
			result = result + ","
		}
		result = result + strconv.FormatFloat(elem, 'f', -1, 64)
	}
	return result
}

func computeNextBestCentroid(centroids []Datum, input []Datum) (Datum, error) {
	largestSquareDist := -1.0
	largestIdx := -1
	for i, elem := range input {
		squareDist, err := squareDistanceToClosestCentroid(centroids, elem)
		if err != nil {
			return nil, err
		}
		if squareDist > largestSquareDist {
			largestSquareDist = squareDist
			largestIdx = i
		}
	}
	return input[largestIdx], nil
}

func squareDistanceToClosestCentroid(centroids []Datum, elem Datum) (float64, error) {
	min := math.MaxFloat64
	for _, centroid := range centroids {
		squareDist, err := cabiriaMath.SquareDistance(elem, centroid)
		if err != nil {
			return -1.0, err
		}
		if squareDist < min {
			min = squareDist
		}
	}
	return min, nil
}
