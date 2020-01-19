package intertitle

import (
	"fmt"
	"image"
	"io/ioutil"
	"path/filepath"
	"strings"

	cabiriaImage "github.com/liampulles/cabiria/pkg/image"
	"github.com/liampulles/cabiria/pkg/image/intertitle"
)

func ProcessData(framePath string, csvPath string) error {
	// get filenames
	files, err := ioutil.ReadDir(framePath)
	if err != nil {
		return err
	}

	// Collate associated data
	output := "" //"filename,avgIntensity,lowerIntensity,middleIntensity,upperIntensity,proportionLower,proportionMiddle,proportionUpper,isIntertitle\n"
	for i, file := range files {
		fmt.Printf("Progress: %d/%d\n", i+1, len(files))
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".png") {
			continue
		}
		currImg, err := cabiriaImage.GetPNG(filepath.Join(framePath, file.Name()))
		if err != nil {
			return err
		}

		output += dataLineForFrame(file.Name(), currImg)
		if err != nil {
			return err
		}
	}

	// Save data
	err = writeToFile(output, csvPath)
	return err
}

func dataLineForFrame(path string, curr image.Image) string {
	// Filename
	file := filepath.Base(path)

	// Inputs
	stats := intertitle.GetIntensityStats(curr)
	// Outputs
	isIntertitle := 0.0
	if strings.Contains(file, "intertitle") {
		isIntertitle = 1.0
	}
	// Print
	return fmt.Sprintf("%f,%f,%f,%f,%f,%f,%f,%f\n",
		stats.AvgIntensity, stats.LowerAvgIntensity, stats.MiddleAvgIntensity, stats.UpperAvgIntensity, stats.ProportionLower, stats.ProportionMiddle, stats.ProportionUpper,
		isIntertitle)
}

func writeToFile(data string, path string) error {
	err := ioutil.WriteFile(path, []byte(data), 0644)
	return err
}
