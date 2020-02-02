package input

import (
	"flag"
	"fmt"
	"path"
)

type GenerateConfiguration struct {
	VideoPath string
	SRTPath   string
	ASSPath   string
}

func GetGenerateConfiguration() (GenerateConfiguration, error) {
	video := flag.String("video", "", "Silent film to analyze for intertitles")
	srt := flag.String("srt", "", "SRT subtitles to source for text")
	ass := flag.String("ass", "", "(Optional) ASS file to save to. Default is the SRT path with ASS extension")

	flag.Parse()

	if *video == "" {
		return GenerateConfiguration{}, fmt.Errorf("you must provide a -video parameter")
	}
	if *srt == "" {
		return GenerateConfiguration{}, fmt.Errorf("you must provide a -srt parameter")
	}
	if *ass == "" {
		ass = defaultASS(srt)
	}

	return GenerateConfiguration{
		VideoPath: *video,
		SRTPath:   *srt,
		ASSPath:   *ass,
	}, nil
}

func defaultASS(srt *string) *string {
	base := path.Base(*srt)
	ext := path.Ext(*srt)
	base = base[:len(base)-len(ext)]
	base += ".cabiria"
	dir := path.Dir(*srt)
	ass := path.Join(dir, base+".ass")
	return &ass
}
