package core

// VideoConfiguration provides configuration options about the input video.
type VideoConfiguration interface{}

// VideoInformation provides relvant informaiton about the video (including
//  the intertitles)
type VideoInformation struct{}

// ExtractVideoInformation reads relevant information from the input video
func ExtractVideoInformation(config VideoConfiguration) (VideoInformation, error) {
	// TODO: Extract frames to configured dir
	// TODO: Predict intertitle frames
	// TODO: Delete extracted frames
	// TODO: Smooth intertitle frames
	// TODO: Extract intiteritle timings
	// TODO: Extract intertitle color info(s)
	return VideoInformation{}, nil
}
