package test

import (
	"fmt"

	"github.com/liampulles/cabiria/pkg/ml"
)

// DummyPredictor can be used in tests as a Predictor fixture.
type DummyPredictor struct {
	MockData map[string]ml.Datum
}

// Fit creates mock data for samples, such that the associated
//  output for a given input is always returned.
func (p *DummyPredictor) Fit(samples []ml.Sample) error {
	if p.MockData == nil {
		p.MockData = make(map[string]ml.Datum)
	}
	for i := range samples {
		p.MockData[samples[i].Input.AsCSV()] = samples[i].Output
	}
	return nil
}

// Predict returns mock data for input.
func (p *DummyPredictor) Predict(input []ml.Datum) ([]ml.Datum, error) {
	result := make([]ml.Datum, len(input))
	for i, elem := range input {
		predicted, err := p.PredictSingle(elem)
		if err != nil {
			return nil, err
		}
		result[i] = predicted
	}
	return result, nil
}

// PredictSingle returns mock data for input
func (p *DummyPredictor) PredictSingle(input ml.Datum) (ml.Datum, error) {
	if p.MockData == nil {
		return nil, fmt.Errorf("dummyPredictor was not fitted with any mock data")
	}
	elem, ok := p.MockData[input.AsCSV()]
	if !ok {
		return nil, fmt.Errorf("dummyPredictor was not fitted with mock data for %s", input.AsCSV())
	}
	return elem, nil
}

// Save doe snot work for DummyPredictors, do not use it.
func (p *DummyPredictor) Save(path string) error {
	return fmt.Errorf("dummyPredictor cannot be saved")
}
