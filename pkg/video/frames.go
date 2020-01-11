package video

import (
	"os/exec"
	"path"
)

func ExtractFrames(videoPath string, outputDirectory string) error {
	cmd := exec.Command("ffmpeg", "-i", videoPath, "-r", "1", path.Join(outputDirectory, "$filename%06d.png"))
	err := cmd.Start()
	if err != nil {
		return err
	}
	return cmd.Wait()
}
