package core

import (
	"fmt"
	"os"

	"github.com/liampulles/cabiria/pkg/subtitle/style"

	"github.com/liampulles/cabiria/pkg/subtitle"
)

// PrettyConfiguration provides configuration options which are needed
//  to stylize subtitles
type PrettyConfiguration interface {
	FrameOutputDirectory() string
	FontName() string
	FontSize() uint
}

// PrettyIntertitles can be exported to ASS.
type PrettyIntertitles struct {
	GlobalStyle style.Style
	Subtitles   []subtitle.Subtitle
}

// GeneratePrettyIntertitles uses extracted video and subtitle information
//  to generate PrettyIntertitles.
func GeneratePrettyIntertitles(
	videoInfo VideoInformation,
	subInfo SubtitlesInformation,
	config PrettyConfiguration) (PrettyIntertitles, error) {
	fmt.Print("Generating pretty intertitles")

	// Correct sub timing slice to intertitles, and copy style
	correctedSubs := subtitle.AlignSubtitles(subInfo.Subtitles, videoInfo.IntertitleRanges)
	printProgressDot()

	// Delete extracted frames
	err := os.RemoveAll(config.FrameOutputDirectory())
	if err != nil {
		return PrettyIntertitles{}, err
	}
	printProgressDot()

	printDone()
	return PrettyIntertitles{
		GlobalStyle: globalStyle(config),
		Subtitles:   correctedSubs,
	}, nil
}

func globalStyle(config PrettyConfiguration) style.Style {
	return style.Style{
		FontName: config.FontName(),
		FontSize: config.FontSize(),
	}
}
