package intertitle

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/liampulles/cabiria/pkg/ml"
)

// Train loads a list of intensity stats pointed to by csvPath,
// and trains a predictive intertitle model which is saved to modelPath
func Train(csvPath string, modelPath string) error {
	// Load
	rawData, err := loadCsv(csvPath)
	if err != nil {
		return err
	}

	//Initializes a new KNN classifier
	cls := ml.NewKNNClassifier(1)

	//Do a training-test split
	trainData, testData := ml.Split(rawData, 0.66)
	err = cls.Fit(trainData)
	if err != nil {
		return err
	}

	// Determine pass-rate
	passRate, err := ml.Test(cls, testData)
	if err != nil {
		return err
	}
	fmt.Printf("Pass rate on test data: %f%%\n", passRate*100)

	// Save model
	err = cls.Save(modelPath)
	return err
}

func loadCsv(path string) ([]ml.Sample, error) {
	csvFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var samples []ml.Sample
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		sample := ml.Sample{}
		for i := 0; i < len(line)-1; i++ {
			f, err := strconv.ParseFloat(line[i], 64)
			if err != nil {
				return nil, err
			}
			sample.Input = append(sample.Input, f)
		}
		f, err := strconv.ParseFloat(line[len(line)-1], 64)
		if err != nil {
			return nil, err
		}
		sample.Output = []float64{f}
		samples = append(samples, sample)
	}
	return samples, nil
}
