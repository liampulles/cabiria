package core

type SubtitleConfiguration interface{}

type SubtitleInformation struct{}

func ExtractSubtitleInformation(config SubtitleConfiguration) (SubtitleInformation, error) {
	// TODO: Load sub
	// TODO: Extract text slice from sub
	// TODO: Extract sub timings slice
	return SubtitleInformation{}, nil
}

func SaveASS(prettyIntertitles PrettyIntertitles,
	config SubtitleConfiguration) error {

	// TODO: Combine pretty text and corrected timings into ASS
	// TODO: Save ASS
	return nil
}
