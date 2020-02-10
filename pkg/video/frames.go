package video

import (
	"io/ioutil"
	"os/exec"
	"path"
	"strings"
)

const extractedFramePrefix = "extracted_frame"

// ExtractFrames uses FFmpeg to extract a video into frames, and returns
// An array of ordered filepaths to the resulting PNG files.
func ExtractFrames(videoPath string, outputDirectory string) ([]string, error) {
	// Extract the frames
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-r", "1", path.Join(outputDirectory, extractedFramePrefix+"%06d.png"))
	err := cmd.Start()
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	// Discover files
	files, err := ioutil.ReadDir(outputDirectory)
	if err != nil {
		return nil, err
	}
	// Filter out irrelevant files, and map to filenames
	filenames := make([]string, 0)
	for _, file := range files {
		filename := file.Name()
		if !file.IsDir() &&
			strings.HasPrefix(filename, extractedFramePrefix) &&
			strings.HasSuffix(filename, ".png") {
			filenames = append(filenames, path.Join(outputDirectory, filename))
		}
	}
	return filenames, nil
}
