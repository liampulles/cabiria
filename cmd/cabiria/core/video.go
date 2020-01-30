package core

type VideoConfiguration interface{}

type VideoInformation struct{}

func ExtractVideoInformation(config VideoConfiguration) (VideoInformation, error) {
	// TODO: Extract frames to configured dir
	// TODO: Predict intertitle frames
	// TODO: Delete extracted frames
	// TODO: Smooth intertitle frames
	// TODO: Extract intiteritle timings
	// TODO: Extract intertitle color info(s)
	return VideoInformation{}, nil
}
