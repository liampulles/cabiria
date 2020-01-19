package ml

type Datum []float64

type Sample struct {
	Input  Datum
	Output Datum
}

type Predictor interface {
	Fit(samples []Sample) error
	Predict(input []Datum) ([]Datum, error)
	PredictSingle(input Datum) (Datum, error)
	Save(path string) error
}

type Transformer interface {
	Fit(samples []Datum) error
	Transform(input []Datum) ([]Datum, error)
	TransformSingle(input Datum) (Datum, error)
	Save(path string) error
}
