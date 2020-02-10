package write

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/liampulles/cabiria/pkg/file"

	"github.com/liampulles/cabiria/pkg/meta"

	"github.com/liampulles/cabiria/pkg/subtitle"
	"github.com/liampulles/cabiria/pkg/subtitle/style"

	cabiriaTime "github.com/liampulles/cabiria/pkg/time"
)

// VideoInformation provides necessary info about the video for generating an
//  ASS file.
type VideoInformation interface {
	VideoName() string
	VideoWidth() int
	VideoHeight() int
}

// ASS saves subtitles with a given style to ASS format at path
func ASS(subs []subtitle.Subtitle, sty style.Style, vidInfo VideoInformation, path string) error {
	text := ""
	text += assHeader(vidInfo.VideoName(), vidInfo.VideoWidth(), vidInfo.VideoHeight())
	text += assStyles(sty)
	text += assEvents(subs)

	return file.SaveTextToFile(path, text)
}

func assHeader(videoName string, videoWidth, videoHeight int) string {
	return fmt.Sprintf(`[Script Info]
; Script generated by %s %s
; %s
Title: %s Styled Subs - %s
ScriptType: v4.00+
WrapStyle: 0
PlayResX: %d
PlayResY: %d
Video Aspect Ratio: 0
Video Zoom: 6
Video Position: 0
Collisions: Normal

`,
		meta.ProgramName,
		meta.ProgramVersion,
		meta.ProgramURL,
		meta.ProgramName,
		videoName,
		videoWidth,
		videoHeight)
}

func assStyles(sty style.Style) string {
	return fmt.Sprintf(`[V4+ Styles]
Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
Style: cabiria,%s,%d,%s,%s,%s,%s,0,0,0,0,100,100,0,0,3,1000,0,5,10,10,10,1

`,
		sty.FontName,
		sty.FontSize,
		assColor(sty.FontColor),
		assColor(color.Transparent),
		assColor(color.Transparent),
		assColor(color.Black))
}

func assColor(col color.Color) string {
	r, g, b, a := col.RGBA()
	return fmt.Sprintf("&H%02X%02X%02X%02X",
		int(a/257),
		int(b/257),
		int(g/257),
		int(r/257))
}

func assEvents(subs []subtitle.Subtitle) string {
	result := `[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
`
	for _, sub := range subs {
		result += assDialogueLine(sub)
	}
	result += "\n"
	return result
}

func assDialogueLine(sub subtitle.Subtitle) string {
	return fmt.Sprintf("Dialogue: 0,%s,%s,cabiria,,0000,0000,0000,,%s\n",
		cabiriaTime.ToASSTimecode(sub.StartTime),
		cabiriaTime.ToASSTimecode(sub.EndTime),
		replaceNewlineWithSlashN(sub.Text))
}

func replaceNewlineWithSlashN(lined string) string {
	return strings.ReplaceAll(lined, "\n", "\\N")
}