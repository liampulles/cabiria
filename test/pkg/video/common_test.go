package video_test

import (
	"testing"

	"github.com/liampulles/cabiria/pkg/video"
)

func TestGetBasicInformation_ForExistingVideo(t *testing.T) {
	// Setup expectations
	expected := video.Information{
		Width:  656,
		Height: 526,
		FPS:    25.0,
	}

	// Exercise SUT
	actual, err := video.GetBasicInformation("testdata/Po-zakonu.mkv")

	// Verify result
	if err != nil {
		t.Errorf("SUT returned an error: %v", err)
	}
	if actual != expected {
		t.Errorf("Results differ - Expected: %v, Actual: %v", expected, actual)
	}
}

func TestGetBasicInformation_ForNonExistingVideo(t *testing.T) {
	// Exercise SUT
	_, err := video.GetBasicInformation("this/path/does/not.exist")

	// Verify result
	if err == nil {
		t.Errorf("Expected SUT to return an error")
	}
}
