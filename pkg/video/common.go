package video

import (
	"strconv"
)

// Information defines attributes of a video.
type Information struct {
	Width  int
	Height int
	FPS    float64
}

// GetBasicInformation  extracts some basic attributes form the video pointed
//  to by videoPath
func GetBasicInformation(videoPath string) (Information, error) {
	stringResults, err := QueryWithMediaInfo(videoPath, []string{"Width", "Height", "FrameRate"})
	if err != nil {
		return Information{}, err
	}
	width, err := strconv.ParseInt(stringResults[0], 0, 0)
	if err != nil {
		return Information{}, err
	}
	height, err := strconv.ParseInt(stringResults[1], 0, 0)
	if err != nil {
		return Information{}, err
	}
	fps, err := strconv.ParseFloat(stringResults[2], 64)
	if err != nil {
		return Information{}, err
	}
	return Information{
		Width:  int(width),
		Height: int(height),
		FPS:    fps,
	}, nil
}
