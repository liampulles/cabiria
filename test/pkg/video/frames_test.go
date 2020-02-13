package video_test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/liampulles/cabiria/pkg/video"
)

func TestExtractFrames_ForExistingVideo(t *testing.T) {
	// Setup fixture
	expected := expectedFrames()

	err := os.Mkdir("testFrames", os.ModePerm)
	if err != nil {
		t.Errorf("Mkdir threw an error: %v", err)
	}

	// Exercise SUT
	actual, err := video.ExtractFrames("testdata/By-The-Law.mkv", "testFrames")

	// Verify result
	if err != nil {
		t.Errorf("SUT threw an error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Result differs. Actual: %s, Expected %s", actual, expected)
	}

	// Cleanup
	err = os.RemoveAll("testFrames")
	if err != nil {
		t.Errorf("RemoveAll threw an error: %v", err)
	}
}

func TestExtractFrames_ForNonExistingVideo(t *testing.T) {
	// Exercise SUT
	_, err := video.ExtractFrames("this/path/does/not.exist", ".")

	// Verify result
	if err == nil {
		t.Errorf("Expected SUT to return an error")
	}
}

func expectedFrames() []string {
	result := make([]string, 330)
	for i := 0; i < 330; i++ {
		result[i] = fmt.Sprintf("testFrames/extracted_frame%06d.png", i+1)
	}
	return result
}
