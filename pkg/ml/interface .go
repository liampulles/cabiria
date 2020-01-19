package ml

type Input []float64

type Output []float64

type Sample struct {
	Input  Input
	Output Output
}

type Predictor interface {
	Fit(samples []Sample) error
	Predict(input []Input) ([]Output, error)
	PredictSingle(input Input) (Output, error)
	Save(path string) error
}
