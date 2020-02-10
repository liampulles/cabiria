package ml

import (
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"os"
	"sort"

	cabiriaMath "github.com/liampulles/cabiria/pkg/math"
)

var gobRegistered bool

// KNNClassifier defines a KNN model for classifying input against known
//  samples. All training data is saved in the model, so be wary of size and
//  performance for large training sets.
type KNNClassifier struct {
	K      uint
	Points []Sample
}

// NewKNNClassifier constructs an untrained KNNClassifier.
func NewKNNClassifier(k uint) *KNNClassifier {
	return &KNNClassifier{
		K: k,
	}
}

// LoadKNNClassifier is a convenience method for loading a save KNNClassifier model.
func LoadKNNClassifier(path string) (*KNNClassifier, error) {
	// Register type
	if !gobRegistered {
		gobRegistered = true
		gob.Register(KNNClassifier{})
	}
	// Open file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Decompress with gzip
	fz, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer fz.Close()
	// Read the file
	dec := gob.NewDecoder(fz)
	kc := &KNNClassifier{}
	err = dec.Decode(kc)
	if err != nil {
		return nil, err
	}
	return kc, nil
}

// Fit trains a KNNClassifier using given Samples.
func (kc *KNNClassifier) Fit(samples []Sample) error {
	if samples == nil || len(samples) == 0 {
		return fmt.Errorf("Input samples must have at least one element. Input: %v", samples)
	}
	kc.Points = make([]Sample, len(samples))
	copy(kc.Points, samples)
	return nil
}

// Predict finds the "closest" known Sample for each given Datum, and returns
//  the associated Sample output.
func (kc *KNNClassifier) Predict(input []Datum) ([]Datum, error) {
	output := make([]Datum, len(input))
	for i, elem := range input {
		target, err := kc.PredictSingle(elem)
		if err != nil {
			return nil, err
		}
		output[i] = target
	}
	return output, nil
}

// PredictSingle finds the "closest" known Sample the given Datum, and returns
//  the associated output.
func (kc *KNNClassifier) PredictSingle(input Datum) (Datum, error) {
	closeSample, err := findClosest(kc.Points, input, kc.K)
	if err != nil {
		return nil, err
	}
	return closeSample.Output, nil
}

// Save saves a KNNClassifier to disk.
func (kc *KNNClassifier) Save(path string) error {
	// Register type
	if !gobRegistered {
		gobRegistered = true
		gob.Register(KNNClassifier{})
	}
	// Open file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	// Compress with gzip
	fz := gzip.NewWriter(file)
	defer fz.Close()
	// Write the file
	enc := gob.NewEncoder(fz)
	err = enc.Encode(kc)
	return err
}

type argDistPair struct {
	Arg  int
	Dist float64
}

type argDistPairs []argDistPair

func (adp argDistPairs) Len() int           { return len(adp) }
func (adp argDistPairs) Swap(i, j int)      { adp[i], adp[j] = adp[j], adp[i] }
func (adp argDistPairs) Less(i, j int) bool { return adp[i].Dist < adp[j].Dist }

func findClosest(samples []Sample, closestTo Datum, k uint) (Sample, error) {
	pairs := make([]argDistPair, len(samples))
	for i, sample := range samples {
		dist, err := cabiriaMath.SquareDistance(closestTo, sample.Input)
		if err != nil {
			return Sample{}, err
		}
		pairs[i] = argDistPair{i, dist}
	}
	closestArgs := minKDistArg(pairs, k)
	closestSamples := selectByArgs(samples, closestArgs)
	return mode(closestSamples)
}

func minKDistArg(pairs argDistPairs, k uint) []int {
	sort.Sort(pairs)
	var args []int
	for i := uint(0); i < k && i < uint(len(pairs)); i++ {
		args = append(args, pairs[i].Arg)
	}
	return args
}

func selectByArgs(samples []Sample, args []int) []Sample {
	var result []Sample
	for _, elem := range args {
		if elem < len(samples) {
			result = append(result, samples[elem])
		}
	}
	return result
}

func mode(samples []Sample) (Sample, error) {
	pairMatches := make([]uint, len(samples))
	for i := 0; i < len(samples)-1; i++ {
		for j := i; j < len(samples); j++ {
			match, err := Match(samples[i].Output, samples[j].Output)
			if err != nil {
				return Sample{}, err
			}
			if match {
				pairMatches[i]++
				pairMatches[j]++
			}
		}
	}
	return samples[maxArg(pairMatches)], nil
}

func maxArg(input []uint) int {
	largest := uint(0)
	largestI := -1
	for i, elem := range input {
		if elem >= largest {
			largest = elem
			largestI = i
		}
	}
	return largestI
}
