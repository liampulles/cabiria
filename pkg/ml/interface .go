package ml

type Input []float64

type ClassificationOutput []uint

type ClassificationSample struct {
	Input  Input
	Output ClassificationOutput
}

type ClassificationPredictor interface {
	Fit(samples []ClassificationSample) error
	Predict(input []Input) ([]ClassificationOutput, error)
	PredictSingle(input Input) (ClassificationOutput, error)
	Save(path string) error
}

type RegressionOutput []float64

type RegressionSample struct {
	Input  Input
	Output RegressionOutput
}

type RegressionPredictor interface {
	Fit(samples []RegressionSample) error
	Predict(input []Input) ([]RegressionOutput, error)
	PredictSingle(input Input) (RegressionOutput, error)
	Save(path string) error
}
