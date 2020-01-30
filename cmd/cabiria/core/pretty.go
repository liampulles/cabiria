package core

type PrettyConfiguration interface{}

type PrettyIntertitles struct{}

func GeneratePrettyIntertitles(
	videoInfo VideoInformation,
	subInfo SubtitleInformation,
	config PrettyConfiguration) (PrettyIntertitles, error) {

	// TODO: Correct sub timing slice to intertitles
	// TODO: Generate pretty text slice for text slice, given config and color info(s)
	return PrettyIntertitles{}, nil
}
