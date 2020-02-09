package intertitle

import (
	"image"

	"github.com/liampulles/cabiria/pkg/image/intertitle"
	"github.com/liampulles/cabiria/pkg/ml"
)

// Predictor can be used to predict whether given frames are intertitles
//  or not.
type Predictor struct {
	Predictor ml.Predictor
}

// Load loads and constructs the intertitle predictor from disk.
func Load(predictorPath string) (Predictor, error) {
	knn, err := ml.LoadKNNClassifier(predictorPath)
	if err != nil {
		return Predictor{}, err
	}
	return Predictor{Predictor: knn}, nil
}

// Predict guesses what frames are intertitles, and outputs associated intertitles
func (p Predictor) Predict(frames []image.Image) ([]bool, error) {
	datum := mapFramesToInput(frames)
	predictions, err := p.Predictor.Predict(datum)
	if err != nil {
		return nil, err
	}
	return mapPredictionsToIntertitles(predictions), nil
}

func mapFramesToInput(frames []image.Image) []ml.Datum {
	stats := make([]ml.Datum, len(frames))
	for i, elem := range frames {
		stats[i] = intertitle.GetIntensityStats(elem).AsInput()
	}
	return stats
}

func mapPredictionsToIntertitles(predictions []ml.Datum) []bool {
	intertitles := make([]bool, len(predictions))
	for i, elem := range predictions {
		// The output should just be a single element, which is 1.0 if
		//  an intertitle is predicted, 0.0 if not.
		intertitles[i] = elem[0] > 0.5
	}
	return intertitles
}
