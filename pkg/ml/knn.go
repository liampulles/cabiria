package ml

import (
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"math"
	"os"

	calibriaMath "github.com/liampulles/cabiria/pkg/math"
)

var gobRegistered bool

type KNNClassifier struct {
	Points []ClassificationSample
}

func NewKNNClassifier() *KNNClassifier {
	return &KNNClassifier{}
}

func LoadKNNClassfier(path string) (*KNNClassifier, error) {
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

func (kc *KNNClassifier) Fit(samples []ClassificationSample) error {
	if samples == nil || len(samples) == 0 {
		return fmt.Errorf("Input samples must have at least one element. Input: %v", samples)
	}
	kc.Points = make([]ClassificationSample, len(samples))
	copy(kc.Points, samples)
	return nil
}

func (kc *KNNClassifier) Predict(input []Input) ([]ClassificationOutput, error) {
	output := make([]ClassificationOutput, len(input))
	for i, elem := range input {
		target, err := kc.PredictSingle(elem)
		if err != nil {
			return nil, err
		}
		output[i] = target
	}
	return output, nil
}

func (kc *KNNClassifier) PredictSingle(input Input) (ClassificationOutput, error) {
	closeSample, err := findClosest(kc.Points, input)
	if err != nil {
		return nil, err
	}
	return closeSample.Output, nil
}

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

func findClosest(samples []ClassificationSample, closestTo Input) (ClassificationSample, error) {
	var closest ClassificationSample
	closestDist := math.MaxFloat64
	for _, sample := range samples {
		dist, err := calibriaMath.EuclideanDistance(closestTo, sample.Input)
		if err != nil {
			return ClassificationSample{}, err
		}
		if dist < closestDist {
			closestDist = dist
			closest = sample
		}
	}
	if closestDist == math.MaxFloat64 {
		return ClassificationSample{}, fmt.Errorf("No closest identified. Size of samples: %d", len(samples))
	}
	return closest, nil
}
