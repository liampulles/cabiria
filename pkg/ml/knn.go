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

type KNNPredictor struct {
	Points []Sample
}

func NewKNNPredictor() *KNNPredictor {
	return &KNNPredictor{}
}

func LoadKNNPredictor(path string) (*KNNPredictor, error) {
	// Register type
	if !gobRegistered {
		gobRegistered = true
		gob.Register(KNNPredictor{})
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
	kc := &KNNPredictor{}
	err = dec.Decode(kc)
	if err != nil {
		return nil, err
	}
	return kc, nil
}

func (kc *KNNPredictor) Fit(samples []Sample) error {
	if samples == nil || len(samples) == 0 {
		return fmt.Errorf("Input samples must have at least one element. Input: %v", samples)
	}
	kc.Points = make([]Sample, len(samples))
	copy(kc.Points, samples)
	return nil
}

func (kc *KNNPredictor) Predict(input []Input) ([]Output, error) {
	output := make([]Output, len(input))
	for i, elem := range input {
		target, err := kc.PredictSingle(elem)
		if err != nil {
			return nil, err
		}
		output[i] = target
	}
	return output, nil
}

func (kc *KNNPredictor) PredictSingle(input Input) (Output, error) {
	closeSample, err := findClosest(kc.Points, input)
	if err != nil {
		return nil, err
	}
	return closeSample.Output, nil
}

func (kc *KNNPredictor) Save(path string) error {
	// Register type
	if !gobRegistered {
		gobRegistered = true
		gob.Register(KNNPredictor{})
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

func findClosest(samples []Sample, closestTo Input) (Sample, error) {
	var closest Sample
	closestDist := math.MaxFloat64
	for _, sample := range samples {
		dist, err := calibriaMath.EuclideanDistance(closestTo, sample.Input)
		if err != nil {
			return Sample{}, err
		}
		if dist < closestDist {
			closestDist = dist
			closest = sample
		}
	}
	if closestDist == math.MaxFloat64 {
		return Sample{}, fmt.Errorf("No closest identified. Size of samples: %d", len(samples))
	}
	return closest, nil
}
