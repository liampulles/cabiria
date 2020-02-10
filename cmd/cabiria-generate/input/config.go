package input

import (
	"flag"
	"fmt"
	"path"
)

// GenerateConfiguration provides configuration options necessary
//  for generating pretty subtitles from an input video and subtitle
type GenerateConfiguration struct {
	videoPath string
	srtPath   string
	assPath   string
}

// GetGenerateConfiguration parses the command line to provide config
//  for the core application
func GetGenerateConfiguration() (GenerateConfiguration, error) {
	video := flag.String("video", "", "Silent film to analyze for intertitles")
	srt := flag.String("srt", "", "SRT subtitles to source for text")
	ass := flag.String("ass", "", "(Optional) ASS file to save to. Default is the SRT path with ASS extension")

	flag.Parse()

	if *video == "" {
		return GenerateConfiguration{}, fmt.Errorf("you must provide a -video parameter")
	}
	if *srt == "" {
		return GenerateConfiguration{}, fmt.Errorf("you must provide a -srt parameter")
	}
	if *ass == "" {
		ass = defaultASS(srt)
	}

	return GenerateConfiguration{
		videoPath: *video,
		srtPath:   *srt,
		assPath:   *ass,
	}, nil
}

// VideoPath is the path to the input video
func (gc *GenerateConfiguration) VideoPath() string {
	return gc.videoPath
}

// FrameOutputDirectory is the temporary directory where frames will be extracted to.
func (gc *GenerateConfiguration) FrameOutputDirectory() string {
	//TODO: Check this folder gets deleted
	return "/tmp/cabiria/extractedFrames"
}

// PredictorPath points to the ml.Predictor model used to predict intertitles
func (gc *GenerateConfiguration) PredictorPath() string {
	return "/etc/cabiria/intertitlePredictor.model"
}

// SmoothingClosingThreshold defines the upper bound for a gap in intertitles
//  to be closed
func (gc *GenerateConfiguration) SmoothingClosingThreshold() uint {
	return 15
}

// SmoothingOpeningThreshold defines the minimum length of an intertitle to be
//  kept
func (gc *GenerateConfiguration) SmoothingOpeningThreshold() uint {
	return 15
}

func defaultASS(srt *string) *string {
	base := path.Base(*srt)
	ext := path.Ext(*srt)
	base = base[:len(base)-len(ext)]
	base += ".cabiria"
	dir := path.Dir(*srt)
	ass := path.Join(dir, base+".ass")
	return &ass
}
