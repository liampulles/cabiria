package video

import (
	"fmt"
	"os/exec"
	"strings"
)

const seperator = "-SEPERATOR-"

// QueryWithMediaInfo queries the desired aspects of a video using
//  mediainfo. The corresponding results for each parameter are returned.
func QueryWithMediaInfo(videoPath string, videoParameters []string) ([]string, error) {
	videoParameters = filterOutEmptyAndTrimWhitespace(videoParameters)
	if len(videoParameters) == 0 {
		return []string{}, nil
	}
	finalArgs := outputArg(videoParameters)
	cmd := exec.Command("mediainfo", finalArgs, videoPath)
	bytes, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	result := string(bytes)
	output := filterOutEmptyAndTrimWhitespace(mapOutput(result))
	if len(output) != len(videoParameters) {
		return nil, fmt.Errorf("mediainfo failed to return all parameters for input. Input parameters: %v", videoParameters)
	}
	return output, nil
}

func filterOutEmptyAndTrimWhitespace(array []string) []string {
	var result []string
	for _, elem := range array {
		trimmed := strings.TrimSpace(elem)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}

func outputArg(videoParameters []string) string {
	result := `--Output=Video;`
	transformedParameters := make([]string, len(videoParameters))
	for i, videoParameter := range videoParameters {
		transformedParameters[i] = fmt.Sprintf(`%%%s%%`, videoParameter)
	}
	parameters := strings.Join(transformedParameters, seperator)
	return result + parameters
}

func mapOutput(output string) []string {
	return strings.Split(strings.Split(output, "\n")[0], seperator)
}
