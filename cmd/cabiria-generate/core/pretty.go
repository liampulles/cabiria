package core

import "os"

// PrettyConfiguration provides configuration options which are needed
//  to stylize subtitles
type PrettyConfiguration interface {
	FrameOutputDirectory() string
}

// PrettyIntertitles can be exported to ASS.
type PrettyIntertitles struct{}

// GeneratePrettyIntertitles uses extracted video and subtitle information
//  to generate PrettyIntertitles.
func GeneratePrettyIntertitles(
	videoInfo VideoInformation,
	subInfo SubtitlesInformation,
	config PrettyConfiguration) (PrettyIntertitles, error) {

	// TODO: Correct sub timing slice to intertitles

	// TODO: Extract intertitle color info(s)
	// TODO: Generate pretty text slice for text slice, given config and color info(s)

	// Delete extracted frames
	err := os.RemoveAll(config.FrameOutputDirectory())
	if err != nil {
		return PrettyIntertitles{}, err
	}

	return PrettyIntertitles{}, nil
}
