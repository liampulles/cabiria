package main

import (
	"os"
	"path"

	"github.com/liampulles/cabiria/pkg/intertitle"

	mlIntertitle "github.com/liampulles/cabiria/pkg/ml/train/intertitle"
)

func main() {
	csvPath := "data/intertitle/data.csv"
	if len(os.Args) >= 2 {
		csvPath = os.Args[1]
	}

	modelPath := path.Join("data/intertitle", intertitle.PredictorFilename)
	if len(os.Args) >= 3 {
		modelPath = os.Args[2]
	}

	err := mlIntertitle.Train(csvPath, modelPath)
	if err != nil {
		panic(err)
	}
}
