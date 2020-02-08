package ml

// Datum is an array of floating point values, used as the input to a ML model.
type Datum []float64

// Sample holds the input of a model to its expected output.
type Sample struct {
	Input  Datum
	Output Datum
}

// Predictor defines the methods necessary to Train, Predict, and Save a
//  ML model. The design is influence by scikit-learn (https://scikit-learn.org/stable/about.html).
type Predictor interface {
	Fit(samples []Sample) error
	Predict(input []Datum) ([]Datum, error)
	PredictSingle(input Datum) (Datum, error)
	Save(path string) error
}

// Transformer defines the methods to Train, Predict, and Save a pre-processing
//  model. The key difference from a Predictor is that it doesn't fit against
//  Samples, but just raw input Datum.
type Transformer interface {
	Fit(samples []Datum) error
	Transform(input []Datum) ([]Datum, error)
	TransformSingle(input Datum) (Datum, error)
	Save(path string) error
}
