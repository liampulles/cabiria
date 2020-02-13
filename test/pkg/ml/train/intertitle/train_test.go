package train_test

import (
	"testing"

	"github.com/liampulles/cabiria/pkg/ml/train/intertitle"
)

func TestTrain(t *testing.T) {
	// Setup fixture
	csvPath := "testdata/data.csv"
	modelPath := "testdata/intertitlePredictor.model"

	// Exercise SUT
	err := intertitle.Train(csvPath, modelPath)

	// Verify result
	if err != nil {
		t.Errorf("Encountered error while executing SUT: %v", err)
	}
}
