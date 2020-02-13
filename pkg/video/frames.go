package video

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

const extractedFramePrefix = "extracted_frame"

// ExtractFrames uses FFmpeg to extract a video into frames, and returns
// An array of ordered filepaths to the resulting PNG files.
func ExtractFrames(videoPath string, outputDirectory string) ([]string, error) {
	// Delete direcory, if it exists
	os.RemoveAll(outputDirectory)

	// Create directory path
	err := os.MkdirAll(outputDirectory, 0700)
	if err != nil {
		return nil, err
	}

	// Extract the frames
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-vf", "scale=64:48", path.Join(outputDirectory, extractedFramePrefix+"%06d.png"))
	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("failed to start ffmpeg: %v", err)
	}
	err = cmd.Wait()
	if err != nil {
		return nil, fmt.Errorf("failed to wait on ffmpeg: %v", err)
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
