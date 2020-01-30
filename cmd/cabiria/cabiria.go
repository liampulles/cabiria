package main

import (
	"fmt"
	"os"

	"github.com/liampulles/cabiria/cmd/cabiria/core"
	"github.com/liampulles/cabiria/cmd/cabiria/input"
)

func main() {
	config, err := input.GetConfiguration(os.Args)
	failIf(err)
	videoInfo, err := core.ExtractVideoInformation(config)
	failIf(err)
	subInfo, err := core.ExtractSubtitleInformation(config)
	failIf(err)
	prettyIntertitles, err := core.GeneratePrettyIntertitles(videoInfo, subInfo, config)
	failIf(err)
	err = core.SaveASS(prettyIntertitles, config)
	failIf(err)
}

func failIf(err error) {
	if err != nil {
		fmt.Printf("Encountered fatal error: %v\n", err)
		os.Exit(1)
	}
}
