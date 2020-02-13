package write_test

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"testing"
	"time"

	"github.com/liampulles/cabiria/pkg/intertitle"

	"github.com/liampulles/cabiria/pkg/subtitle"
	"github.com/liampulles/cabiria/pkg/subtitle/style"

	"github.com/liampulles/cabiria/pkg/subtitle/write"
)

func TestASS(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		subs     []subtitle.Subtitle
		sty      style.Style
		vidInfo  write.VideoInformation
		expected string
	}{
		// No subs
		{
			subs(),
			sty("Arial", 20),
			vidInfo("City Lights", 1280, 576),
			`[Script Info]
; Script generated by Cabiria v0.0.1
; https://github.com/liampulles/cabiria
Title: Cabiria Styled Subs - City Lights
ScriptType: v4.00+
WrapStyle: 0
PlayResX: 1280
PlayResY: 576
Video Aspect Ratio: 0
Video Zoom: 6
Video Position: 0
Collisions: Normal

[V4+ Styles]
Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
Style: cabiria,Arial,20,&HFFFFFF,&HFF000000,&H00000000,&H000000,0,0,0,0,100,100,0,0,3,1000,0,5,10,10,10,1

[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text

`,
		},
		// One sub
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "Hello\nWorld", interSty(color.White, color.Black)),
			),
			sty("Arial", 20),
			vidInfo("City Lights", 1280, 576),
			`[Script Info]
; Script generated by Cabiria v0.0.1
; https://github.com/liampulles/cabiria
Title: Cabiria Styled Subs - City Lights
ScriptType: v4.00+
WrapStyle: 0
PlayResX: 1280
PlayResY: 576
Video Aspect Ratio: 0
Video Zoom: 6
Video Position: 0
Collisions: Normal

[V4+ Styles]
Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
Style: cabiria,Arial,20,&HFFFFFF,&HFF000000,&H00000000,&H000000,0,0,0,0,100,100,0,0,3,1000,0,5,10,10,10,1

[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
Dialogue: 0,0:00:01.00,0:00:02.00,cabiria,,0000,0000,0000,,{\c&HFFFFFF&\3c&H000000&}Hello\NWorld

`,
		},
		// Many subs
		{
			subs(
				sub(timestamp(0, 0, 1, 0), timestamp(0, 0, 2, 0), "Hello\nWorld", interSty(color.White, color.Black)),
				sub(timestamp(0, 1, 12, 354), timestamp(0, 12, 32, 90), "How is it going?", interSty(greenishPink(), color.Black)), //&H00D0E0FF
			),
			sty("Arial", 20),
			vidInfo("City Lights", 1280, 576),
			`[Script Info]
; Script generated by Cabiria v0.0.1
; https://github.com/liampulles/cabiria
Title: Cabiria Styled Subs - City Lights
ScriptType: v4.00+
WrapStyle: 0
PlayResX: 1280
PlayResY: 576
Video Aspect Ratio: 0
Video Zoom: 6
Video Position: 0
Collisions: Normal

[V4+ Styles]
Format: Name, Fontname, Fontsize, PrimaryColour, SecondaryColour, OutlineColour, BackColour, Bold, Italic, Underline, StrikeOut, ScaleX, ScaleY, Spacing, Angle, BorderStyle, Outline, Shadow, Alignment, MarginL, MarginR, MarginV, Encoding
Style: cabiria,Arial,20,&HFFFFFF,&HFF000000,&H00000000,&H000000,0,0,0,0,100,100,0,0,3,1000,0,5,10,10,10,1

[Events]
Format: Layer, Start, End, Style, Name, MarginL, MarginR, MarginV, Effect, Text
Dialogue: 0,0:00:01.00,0:00:02.00,cabiria,,0000,0000,0000,,{\c&HFFFFFF&\3c&H000000&}Hello\NWorld
Dialogue: 0,0:01:12.35,0:12:32.09,cabiria,,0000,0000,0000,,{\c&HD0E0FF&\3c&H000000&}How is it going?

`,
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			// Exercise SUT
			err := write.ASS(test.subs, test.sty, test.vidInfo, "/tmp/cabiria/assTest.ass")

			// Verify result
			if err != nil {
				t.Errorf("SUT returned an error: %v", err)
			}
			actual := readActual()
			if actual != test.expected {
				t.Errorf("Result differs. Actual:\n%sExpected:\n%s", actual, test.expected)
			}
		})
	}
}

func readActual() string {
	content, err := ioutil.ReadFile("/tmp/cabiria/assTest.ass")
	if err != nil {
		panic(err)
	}
	return string(content)
}

func subs(subs ...subtitle.Subtitle) []subtitle.Subtitle {
	return subs
}

func sub(start, end time.Time, text string, style intertitle.Style) subtitle.Subtitle {
	return subtitle.Subtitle{
		StartTime: start,
		EndTime:   end,
		Text:      text,
		Style:     style,
	}
}

func interSty(foreground, background color.Color) intertitle.Style {
	return intertitle.Style{
		ForegroundColor: foreground,
		BackgroundColor: background,
	}
}

func sty(fontName string, fontSize uint) style.Style {
	return style.Style{
		FontName: fontName,
		FontSize: fontSize,
	}
}

func vidInfo(videoName string, videoWidth, videoHeight int) testVideoInformation {
	return testVideoInformation{
		videoName:   videoName,
		videoWidth:  videoWidth,
		videoHeight: videoHeight,
	}
}

func timestamp(hour, min, sec, milli int) time.Time {
	return time.Date(0, time.January, 1, hour, min, sec, milli*1e+6, time.UTC)
}

func greenishPink() color.Color {
	return color.RGBA{
		R: 0xFF,
		G: 0xE0,
		B: 0xD0,
		A: 0x00,
	}
}

type testVideoInformation struct {
	videoName   string
	videoWidth  int
	videoHeight int
}

func (t testVideoInformation) VideoPath() string {
	return t.videoName
}

func (t testVideoInformation) VideoWidth() int {
	return t.videoWidth
}

func (t testVideoInformation) VideoHeight() int {
	return t.videoHeight
}
