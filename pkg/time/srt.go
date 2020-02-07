package time

import "time"

const srtTimecodeFormat = "15:04:05"

func FromSRTTimecode(timecode string) (time.Time, error) {
	t, err := time.Parse(srtTimecodeFormat, timecode[:len(timecode)-4])
	if err != nil {
		return time.Now(), err
	}
	milli, err := time.ParseDuration(timecode[len(timecode)-3:] + "ms")
	if err != nil {
		return time.Now(), err
	}
	return t.Add(milli), nil

}
