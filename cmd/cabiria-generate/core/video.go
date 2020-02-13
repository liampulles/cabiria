package core

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/liampulles/cabiria/pkg/array"
	"github.com/liampulles/cabiria/pkg/image"
	"github.com/liampulles/cabiria/pkg/intertitle"
	cabiriaMath "github.com/liampulles/cabiria/pkg/math"
	"github.com/liampulles/cabiria/pkg/video"

	"github.com/jinzhu/copier"
)

// VideoConfiguration provides configuration options necessary to extract video information
type VideoConfiguration interface {
	VideoPath() string
	FrameOutputDirectory() string
	PredictorPath() string
	SmoothingClosingThreshold() uint
	SmoothingOpeningThreshold() uint
}

// VideoInformation provides relevant information about the video (including
//  the intertitles)
type VideoInformation struct {
	VideoFPS         float64
	VideoWidth       int
	VideoHeight      int
	IntertitleRanges []intertitle.Range
}

// ExtractVideoInformation reads relevant information from the input video
func ExtractVideoInformation(config VideoConfiguration) (VideoInformation, error) {
	fmt.Print("Extracting video information")
	// Extract frames to configured dir
	framePaths, err := video.ExtractFrames(config.VideoPath(), config.FrameOutputDirectory())
	if err != nil {
		return VideoInformation{}, err
	}
	printProgressDot()

	// Predict intertitle frames
	predictions, err := predictIntertitles(framePaths, config.PredictorPath())
	if err != nil {
		return VideoInformation{}, err
	}
	printProgressDot()

	// Smooth intertitle frames
	smoothIntertitles(predictions, config.SmoothingClosingThreshold(), config.SmoothingOpeningThreshold())
	printProgressDot()

	// Get some basic video info
	basicInfo, err := video.GetBasicInformation(config.VideoPath())
	if err != nil {
		return VideoInformation{}, err
	}
	printProgressDot()

	// Extract intertitle timings
	interRanges, err := intertitle.MapRanges(predictions, basicInfo.FPS, framePaths)
	if err != nil {
		return VideoInformation{}, err
	}
	printDone()

	return VideoInformation{
		VideoFPS:         basicInfo.FPS,
		VideoHeight:      basicInfo.Height,
		VideoWidth:       basicInfo.Width,
		IntertitleRanges: interRanges,
	}, nil
}

func predictIntertitles(framePaths []string, predictorPath string) ([]bool, error) {
	predictor, err := intertitle.Load(predictorPath)
	if err != nil {
		return nil, err
	}

	// Split into workers
	_, workerCount := cabiriaMath.MinMaxInt(1, runtime.NumCPU()/2)
	var wg sync.WaitGroup
	framePathsDivided := divideStringArray(framePaths, workerCount)
	predictionsDivided := setupPredictionsArrays(len(framePaths), workerCount)
	errors := make([]error, workerCount)
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		// Setup vars
		var predictorCopy intertitle.Predictor
		copier.Copy(&predictorCopy, &predictor)

		go predictIntertitlesWorker(&predictorCopy, framePathsDivided[i], predictionsDivided[i], errors[i], &wg)
	}
	wg.Wait()

	// Check for errors
	for _, err := range errors {
		if err != nil {
			return nil, err
		}
	}

	// Construct predictions array
	var predictions []bool
	for _, elem := range predictionsDivided {
		predictions = append(predictions, elem...)
	}
	return predictions, nil
}

func predictIntertitlesWorker(predictor *intertitle.Predictor, framePaths []string, predictions []bool, err error, wg *sync.WaitGroup) {
	defer wg.Done()

	for i, path := range framePaths {
		// if i%3 == 0 {
		// 	fmt.Printf("Processing: %f%%\n", float64(i*100)/float64(len(framePaths)))
		// }
		img, potentialErr := image.GetPNG(path)
		if potentialErr != nil {
			err = potentialErr
			return
		}

		prediction, potentialErr := predictor.PredictSingle(img)
		if potentialErr != nil {
			err = potentialErr
			return
		}
		predictions[i] = prediction
	}
	printProgressDot()
}

func smoothIntertitles(intertitles []bool, closingThreshold, openingThreshold uint) {
	array.CloseBoolArray(intertitles, closingThreshold)
	array.OpenBoolArray(intertitles, openingThreshold)
}

func divideStringArray(many []string, parts int) [][]string {
	var divided [][]string
	chunkSize := (len(many) + parts - 1) / parts
	for i := 0; i < len(many); i += chunkSize {
		end := i + chunkSize
		if end > len(many) {
			end = len(many)
		}
		divided = append(divided, many[i:end])
	}
	return divided
}

func setupPredictionsArrays(total int, parts int) [][]bool {
	var divided [][]bool
	chunkSize := (total + parts - 1) / parts
	for i := 0; i < total; i += chunkSize {
		end := i + chunkSize
		if end > total {
			end = total
		}
		divided = append(divided, make([]bool, end-i))
	}
	return divided
}

func printProgressDot() {
	fmt.Print(".")
}

func printDone() {
	fmt.Print("DONE\n")
}
