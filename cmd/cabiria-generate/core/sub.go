package core

import (
	"fmt"

	"github.com/liampulles/cabiria/pkg/subtitle"
	"github.com/liampulles/cabiria/pkg/subtitle/read"
	"github.com/liampulles/cabiria/pkg/subtitle/write"
)

// SubtitlesConfiguration provides configuration options necessary
//  to extract subtitles.
type SubtitlesConfiguration interface {
	SRTPath() string
	ASSPath() string
}

// SubtitlesInformation is a representation of the input subtitle,
//  with some post processing.
type SubtitlesInformation struct {
	Subtitles []subtitle.Subtitle
}

// ExtractSubtitlesInformation will read in a subtitle given by the configuration,
//  and provide relevant information about the subtitle as output.
func ExtractSubtitlesInformation(config SubtitlesConfiguration) (SubtitlesInformation, error) {
	fmt.Print("Extracting subtitle information")
	// Load subs
	subs, err := read.SRT(config.SRTPath())
	if err != nil {
		return SubtitlesInformation{}, err
	}
	printProgressDot()
	printDone()

	return SubtitlesInformation{
		Subtitles: subs,
	}, nil
}

// SaveASS takes a representation of "pretty" subtitles and writes them to disk,
//  in ASS format.
func SaveASS(prettyIntertitles PrettyIntertitles,
	subConfig SubtitlesConfiguration,
	videoConfig VideoConfiguration,
	videoInfo VideoInformation) error {
	fmt.Print("Saving ASS")

	// Save ASS
	config := ASSConfiguration{
		videoPath:   videoConfig.VideoPath(),
		videoWidth:  videoInfo.VideoWidth,
		videoHeight: videoInfo.VideoHeight,
	}
	printProgressDot()
	write.ASS(prettyIntertitles.Subtitles, prettyIntertitles.GlobalStyle, &config, subConfig.ASSPath())
	printProgressDot()
	printDone()
	return nil
}

// ASSConfiguration is the config necessary to generate and save an ASS file
type ASSConfiguration struct {
	videoPath   string
	videoWidth  int
	videoHeight int
}

// VideoPath is the path to the input video
func (ac *ASSConfiguration) VideoPath() string {
	return ac.videoPath
}

// VideoWidth is the width of the video
func (ac *ASSConfiguration) VideoWidth() int {
	return ac.videoWidth
}

// VideoHeight is the height of the video
func (ac *ASSConfiguration) VideoHeight() int {
	return ac.videoHeight
}
