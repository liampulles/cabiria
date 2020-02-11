package core

import (
	"github.com/liampulles/cabiria/pkg/subtitle"
	"github.com/liampulles/cabiria/pkg/subtitle/read"
)

// SubtitlesConfiguration provides configuration options necessary
//  to extract subtitles.
type SubtitlesConfiguration interface {
	SRTPath() string
}

// SubtitlesInformation is a representation of the input subtitle,
//  with some post processing.
type SubtitlesInformation struct {
	Subtitles []subtitle.Subtitle
}

// ExtractSubtitlesInformation will read in a subtitle given by the configuration,
//  and provide relevant information about the subtitle as output.
func ExtractSubtitlesInformation(config SubtitlesConfiguration) (SubtitlesInformation, error) {
	// Load subs
	subs, err := read.SRT(config.SRTPath())
	if err != nil {
		return SubtitlesInformation{}, err
	}

	return SubtitlesInformation{
		Subtitles: subs,
	}, nil
}

// SaveASS takes a representation of "pretty" subtitles and writes them to disk,
//  in ASS format.
func SaveASS(prettyIntertitles PrettyIntertitles,
	config SubtitlesConfiguration) error {

	// TODO: Combine pretty text and corrected timings into ASS
	// TODO: Save ASS
	return nil
}
