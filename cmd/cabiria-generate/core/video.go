package core

import (
	"github.com/liampulles/cabiria/pkg/array"
	"github.com/liampulles/cabiria/pkg/image"
	"github.com/liampulles/cabiria/pkg/intertitle"
	"github.com/liampulles/cabiria/pkg/video"
)

// ExtractVideoConfiguration provides configuration options necessary to extract video information
type ExtractVideoConfiguration interface {
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
func ExtractVideoInformation(config ExtractVideoConfiguration) (VideoInformation, error) {
	// Extract frames to configured dir
	framePaths, err := video.ExtractFrames(config.VideoPath(), config.FrameOutputDirectory())
	if err != nil {
		return VideoInformation{}, err
	}

	// Predict intertitle frames
	predictions, err := predictIntertitles(framePaths, config.PredictorPath())
	if err != nil {
		return VideoInformation{}, err
	}

	// Smooth intertitle frames
	smoothIntertitles(predictions, config.SmoothingClosingThreshold(), config.SmoothingOpeningThreshold())

	// Get some basic video info
	basicInfo, err := video.GetBasicInformation(config.VideoPath())
	if err != nil {
		return VideoInformation{}, err
	}

	// Extract intertitle timings
	interRanges, err := intertitle.MapRanges(predictions, basicInfo.FPS, framePaths)
	if err != nil {
		return VideoInformation{}, err
	}

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
	predictions := make([]bool, len(framePaths))
	// We predict the images one-by-one to avoid memory issues.
	for i, path := range framePaths {
		img, err := image.GetPNG(path)
		if err != nil {
			return nil, err
		}

		prediction, err := predictor.PredictSingle(img)
		if err != nil {
			return nil, err
		}
		predictions[i] = prediction
	}
	return predictions, nil
}

func smoothIntertitles(intertitles []bool, closingThreshold, openingThreshold uint) {
	array.CloseBoolArray(intertitles, closingThreshold)
	array.OpenBoolArray(intertitles, openingThreshold)
}
