package cluster

import (
	"fmt"
	"math"

	cabiriaMath "github.com/liampulles/cabiria/pkg/math"

	"github.com/liampulles/cabiria/pkg/ml"
)

// KMeansClassifier finds the k cluster centroids which minimize the associative
//  distances with the training input.
type KMeansClassifier struct {
	k             int
	maxIterations int
	centroids     []ml.Datum
}

// NewKMeansClassifier constructs a KMeansClassifier
func NewKMeansClassifier(k int, maxIterations int) *KMeansClassifier {
	return &KMeansClassifier{
		k:             k,
		maxIterations: maxIterations,
	}
}

// Fit finds the KMeans centroids given the input and saves the result.
//  the number of values that associate with each centroid and the number of iterations
//  is returned
func (kmc *KMeansClassifier) Fit(input []ml.Datum) ([]int, int, error) {
	if len(input) < kmc.k {
		return nil, -1, fmt.Errorf("Cannot assign %d clusters to input of length %d", kmc.k, len(input))
	}
	noIterations := 0
	// Select the initial centroids randomly from input elements.
	currentCentroids, err := ml.InitializeKMeans(input, kmc.k)
	if err != nil {
		return nil, -1, err
	}
	lastCentroids := initLastCentroids(currentCentroids)

	// Keep iterating until there is no change to the centroids, or max iterations reached
	centroidCounts := make([]int, len(currentCentroids))
	// -> This is in-case the input has empty datums, we won't enter the loop so we'll return this
	centroidCounts[0] = len(input)
	for noIterations < kmc.maxIterations && someChangeInCentroids(currentCentroids, lastCentroids) {

		// We'll keep a running total of each element assigned to a centroid, so we can get
		// the mean later.
		centroidSums := initialCentroidSums(currentCentroids)
		centroidCounts = make([]int, len(currentCentroids))

		// Find the closest centroid for each element, then add to the totals for
		// that centroid.
		for _, elem := range input {
			centroidI, err := closestCentroidIndex(elem, currentCentroids)
			if err != nil {
				return nil, -1, err
			}
			newCentroidISum, err := cabiriaMath.Add(centroidSums[centroidI], elem)
			if err != nil {
				return nil, -1, fmt.Errorf("input must have consistently sized elements: %v", err)
			}
			centroidSums[centroidI] = newCentroidISum
			centroidCounts[centroidI]++
		}

		// Adjust centroids given assignments
		lastCentroids = currentCentroids
		currentCentroids = calculateNewCentroids(centroidSums, centroidCounts)
		noIterations++
	}

	// Adjust for NaN values, which can occur if e.g. there are less than k
	//  unique values in input.
	currentCentroids, err = adjustForNaN(currentCentroids)
	if err != nil {
		return nil, -1, err
	}

	// Store result
	kmc.centroids = currentCentroids
	return centroidCounts, noIterations, nil
}

// ClusterCentroids returns the centroids
func (kmc *KMeansClassifier) ClusterCentroids() []ml.Datum {
	return kmc.centroids
}

func initLastCentroids(initial []ml.Datum) []ml.Datum {
	lastCentroids := make([]ml.Datum, len(initial))
	for i, elem := range initial {
		lastCentroidsI := make(ml.Datum, len(elem))
		for j, elemI := range elem {
			lastCentroidsI[j] = elemI + 0.1
		}
		lastCentroids[i] = lastCentroidsI
	}
	return lastCentroids
}

func initialCentroidSums(centroids []ml.Datum) []ml.Datum {
	centroidSums := make([]ml.Datum, len(centroids))
	for i, centroid := range centroids {
		centroidSums[i] = make(ml.Datum, len(centroid))
	}
	return centroidSums
}

func someChangeInCentroids(current, last []ml.Datum) bool {
	sumSquareDelta := 0.0
	for i := 0; i < len(current); i++ {
		val, _ := cabiriaMath.SquareDistance(current[i], last[i])
		sumSquareDelta += val
	}
	return sumSquareDelta > 0.0
}

func closestCentroidIndex(datum ml.Datum, centroids []ml.Datum) (int, error) {
	if len(centroids) == 0 {
		return 0, fmt.Errorf("cannot have k = 0")
	}
	closestI := -1
	closestSquareDist := math.MaxFloat64
	for i, elem := range centroids {
		squareDist, err := cabiriaMath.SquareDistance(datum, elem)
		if err != nil {
			return -1, fmt.Errorf("input must have consistently sized elements: %v", err)
		}
		if squareDist < closestSquareDist {
			closestSquareDist = squareDist
			closestI = i
		}
	}
	return closestI, nil
}

func calculateNewCentroids(centroidSums []ml.Datum, centroidCounts []int) []ml.Datum {
	newCentroids := make([]ml.Datum, len(centroidSums))
	for i, centroidSum := range centroidSums {
		centroidCount := float64(centroidCounts[i])
		newCentroidI := make(ml.Datum, len(centroidSum))
		for j, centroidSumJ := range centroidSum {
			newCentroidI[j] = centroidSumJ / centroidCount
		}
		newCentroids[i] = newCentroidI
	}
	return newCentroids
}

func adjustForNaN(centroids []ml.Datum) ([]ml.Datum, error) {
	// Do a first pass to find at least one "legit" centroid
	var legit ml.Datum
	for _, elem := range centroids {
		if isLegit(elem) {
			legit = elem
			break
		}
	}
	if legit == nil {
		return nil, fmt.Errorf("no valid centroids could be found")
	}
	// Replace any non-legit elements with the legit one.
	newCentroids := make([]ml.Datum, len(centroids))
	for i, elem := range centroids {
		if isLegit(elem) {
			newCentroids[i] = elem
		} else {
			newCentroids[i] = legit
		}
	}
	return newCentroids, nil
}

func isLegit(centroid ml.Datum) bool {
	for _, elem := range centroid {
		if math.IsNaN(elem) {
			return false
		}
	}
	return true
}
