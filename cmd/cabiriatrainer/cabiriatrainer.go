package main

import (
	"os"

	"github.com/liampulles/cabiria/pkg/ml/train/intertitle"
)

func main() {
	csvPath := "data/intertitle/data.csv"
	if len(os.Args) >= 2 {
		csvPath = os.Args[1]
	}

	modelPath := "data/intertitle/knn.model"
	if len(os.Args) >= 3 {
		modelPath = os.Args[2]
	}

	err := intertitle.Train(csvPath, modelPath)
	if err != nil {
		panic(err)
	}
}
