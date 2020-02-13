package core

import (
	"fmt"
	"os"

	"github.com/liampulles/cabiria/cmd/cabiria-generate/input"
)

// Run runs the main app for cabiria-generate
func Run(args []string) {
	config, err := input.GetGenerateConfiguration(args)
	failIf(err)
	videoInfo, err := ExtractVideoInformation(&config)
	failIf(err)
	subsInfo, err := ExtractSubtitlesInformation(&config)
	failIf(err)
	prettyIntertitles, err := GeneratePrettyIntertitles(videoInfo, subsInfo, &config)
	failIf(err)
	err = SaveASS(prettyIntertitles, &config, &config, videoInfo)
	failIf(err)
}

func failIf(err error) {
	if err != nil {
		fmt.Printf("Encountered fatal error: %v\n", err)
		os.Exit(1)
	}
}
