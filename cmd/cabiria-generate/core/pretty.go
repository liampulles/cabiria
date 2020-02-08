package core

// PrettyConfiguration provides configuration options which are needed
//  to stylize subtitles
type PrettyConfiguration interface{}

// PrettyIntertitles can be exported to ASS.
type PrettyIntertitles struct{}

// GeneratePrettyIntertitles uses extracted video and subtitle information
//  to generate PrettyIntertitles.
func GeneratePrettyIntertitles(
	videoInfo VideoInformation,
	subInfo SubtitlesInformation,
	config PrettyConfiguration) (PrettyIntertitles, error) {

	// TODO: Correct sub timing slice to intertitles
	// TODO: Generate pretty text slice for text slice, given config and color info(s)
	return PrettyIntertitles{}, nil
}
