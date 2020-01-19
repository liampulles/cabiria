package train_test

import (
	"testing"

	"github.com/liampulles/cabiria/pkg/ml/train/intertitle"
)

func TestProcessData(t *testing.T) {
	// Setup fixture
	framePath := "testdata/frames"
	csvPath := "testdata/data.csv"

	// Exercise SUT
	err := intertitle.ProcessData(framePath, csvPath)

	// Verify result
	if err != nil {
		t.Errorf("Encountered error while executing SUT: %v", err)
	}
}
