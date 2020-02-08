package core

// SubtitlesConfiguration provides configuration options neccesary
//  to extract subtitles.
type SubtitlesConfiguration interface{}

// SubtitlesInformation is a representaion of the input subtitle,
//  with some post processing.
type SubtitlesInformation struct{}

// ExtractSubtitlesInformation will read in a subtitle given by the configuration,
//  and provide relevant information about the subtitle as output.
func ExtractSubtitlesInformation(config SubtitlesConfiguration) (SubtitlesInformation, error) {
	// TODO: Load sub
	// TODO: Extract text slice from sub
	// TODO: Extract sub timings slice
	return SubtitlesInformation{}, nil
}

// SaveASS takes a representation of "pretty" subtitles and writes them to disk,
//  in ASS format.
func SaveASS(prettyIntertitles PrettyIntertitles,
	config SubtitlesConfiguration) error {

	// TODO: Combine pretty text and corrected timings into ASS
	// TODO: Save ASS
	return nil
}
