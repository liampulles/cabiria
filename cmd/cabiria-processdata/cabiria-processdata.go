package main

import (
	"os"

	"github.com/liampulles/cabiria/pkg/ml/train/intertitle"
)

func main() {
	framePath := "data/intertitle/frames"
	if len(os.Args) >= 2 {
		framePath = os.Args[1]
	}

	csvPath := "data/intertitle/data.csv"
	if len(os.Args) >= 3 {
		csvPath = os.Args[2]
	}

	err := intertitle.ProcessData(framePath, csvPath)
	if err != nil {
		panic(err)
	}
}
