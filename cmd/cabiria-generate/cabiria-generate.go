package main

import (
	"fmt"
	"os"

	"github.com/liampulles/cabiria/cmd/cabiria-generate/input"

	"github.com/liampulles/cabiria/cmd/cabiria-generate/core"
)

func main() {
	config, err := input.GetGenerateConfiguration()
	failIf(err)
	videoInfo, err := core.ExtractVideoInformation(&config)
	failIf(err)
	subsInfo, err := core.ExtractSubtitlesInformation(&config)
	failIf(err)
	prettyIntertitles, err := core.GeneratePrettyIntertitles(videoInfo, subsInfo, &config)
	failIf(err)
	err = core.SaveASS(prettyIntertitles, &config)
	failIf(err)
}

func failIf(err error) {
	if err != nil {
		fmt.Printf("Encountered fatal error: %v\n", err)
		os.Exit(1)
	}
}
