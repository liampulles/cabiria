package core

type SubtitlesConfiguration interface{}

type SubtitlesInformation struct{}

func ExtractSubtitlesInformation(config SubtitlesConfiguration) (SubtitlesInformation, error) {
	// TODO: Load sub
	// TODO: Extract text slice from sub
	// TODO: Extract sub timings slice
	return SubtitlesInformation{}, nil
}

func SaveASS(prettyIntertitles PrettyIntertitles,
	config SubtitlesConfiguration) error {

	// TODO: Combine pretty text and corrected timings into ASS
	// TODO: Save ASS
	return nil
}
