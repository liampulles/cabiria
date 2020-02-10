package video_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/liampulles/cabiria/pkg/video"
)

// TODO: Travis cannot run this test, maybe check if ffmpeg available, or put
//  into separate test package?
func TestExtractFrames_ForExistingVideo(t *testing.T) {
	// Setup fixture
	expected := []string{
		"testFrames/extracted_frame000001.png",
		"testFrames/extracted_frame000002.png",
		"testFrames/extracted_frame000003.png",
		"testFrames/extracted_frame000004.png",
		"testFrames/extracted_frame000005.png",
		"testFrames/extracted_frame000006.png",
		"testFrames/extracted_frame000007.png",
		"testFrames/extracted_frame000008.png",
		"testFrames/extracted_frame000009.png",
		"testFrames/extracted_frame000010.png",
		"testFrames/extracted_frame000011.png",
		"testFrames/extracted_frame000012.png",
		"testFrames/extracted_frame000013.png",
		"testFrames/extracted_frame000014.png",
		"testFrames/extracted_frame000015.png",
	}

	err := os.Mkdir("testFrames", os.ModePerm)
	if err != nil {
		t.Errorf("Mkdir threw an error: %v", err)
	}

	// Exercise SUT
	actual, err := video.ExtractFrames("testdata/Po-zakonu.mkv", "testFrames")

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
