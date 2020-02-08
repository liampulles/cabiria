package read

import (
	"fmt"
	"strings"
	"time"

	"github.com/liampulles/cabiria/pkg/subtitle/style"

	"github.com/liampulles/cabiria/pkg/file"
	"github.com/liampulles/cabiria/pkg/subtitle"
	cabiriaTime "github.com/liampulles/cabiria/pkg/time"
)

type lineType int

const (
	index lineType = iota
	timecodes
	text
	blank
)

// SRT loads the SRT pointed to by path into the associated Subtitle slice.
func SRT(path string) ([]subtitle.Subtitle, error) {
	lines, err := file.ReadLinesFromTextFile(path)
	if err != nil {
		return nil, err
	}

	var subs []subtitle.Subtitle
	var currStart time.Time
	var currEnd time.Time
	var currLines []string
	lastLineType := blank
	for _, line := range lines {
		currLineType := getLineType(line, lastLineType)
		switch currLineType {
		case blank, index:
			subs = closeAndAddCurrent(currStart, currEnd, currLines, subs)
			currLines = nil
		case timecodes:
			currStart, currEnd, err = getTimecodes(line)
			if err != nil {
				return nil, err
			}
		case text:
			currLines = append(currLines, line)
		}
		lastLineType = currLineType
	}
	subs = closeAndAddCurrent(currStart, currEnd, currLines, subs)
	return subs, nil
}

func getLineType(line string, last lineType) lineType {
	if strings.TrimSpace(line) == "" {
		return blank
	}
	if last == blank {
		return index
	}
	if last == index {
		return timecodes
	}
	if last == timecodes {
		return text
	}
	if last == text {
		return text
	}
	return -1
}

func closeAndAddCurrent(start, end time.Time, lines []string, subs []subtitle.Subtitle) []subtitle.Subtitle {
	if lines == nil {
		return subs
	}

	joined := strings.Join(lines, "\n")
	joined = style.RemoveStylesFromSRTText(joined)
	new := subtitle.Subtitle{
		StartTime: start,
		EndTime:   end,
		Text:      joined,
	}

	subs = append(subs, new)
	return subs
}

func getTimecodes(line string) (time.Time, time.Time, error) {
	fields := strings.Fields(line)
	if len(fields) < 3 {
		return time.Now(), time.Now(), fmt.Errorf("the timecode line needs at least 3 fields in a SRT file. Received: %s", line)
	}
	start, err := cabiriaTime.FromSRTTimecode(fields[0])
	if err != nil {
		return time.Now(), time.Now(), err
	}
	end, err := cabiriaTime.FromSRTTimecode(fields[2])
	if err != nil {
		return time.Now(), time.Now(), err
	}
	return start, end, err
}
