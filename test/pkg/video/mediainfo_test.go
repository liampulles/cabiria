package video_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/liampulles/cabiria/pkg/video"
)

func TestQueryWithMediaInfo_WhenExistsAndParametersValid(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		videoParameters []string
		expected        []string
	}{
		// Empty cases
		{
			nil,
			[]string{},
		},
		{
			[]string{},
			[]string{},
		},
		// Empty parameters
		{
			[]string{""},
			[]string{},
		},
		{
			[]string{" "},
			[]string{},
		},
		{
			[]string{"  ", "", " "},
			[]string{},
		},
		// Single valid parameter
		{
			[]string{"Width"},
			[]string{"656"},
		},
		{
			[]string{" Width"},
			[]string{"656"},
		},
		{
			[]string{"Width "},
			[]string{"656"},
		},
		{
			[]string{"  Width "},
			[]string{"656"},
		},
		// Many valid parameters
		{
			[]string{"Width", "Height"},
			[]string{"656", "526"},
		},
		{
			[]string{"Height ", " Width"},
			[]string{"526", "656"},
		},
		{
			[]string{"Width", "Height", "FrameRate"},
			[]string{"656", "526", "25.000"},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v -> %v", test.videoParameters, test.expected), func(t *testing.T) {
			// Exercise SUT
			actual, err := video.QueryWithMediaInfo("testdata/By-The-Law.mkv", test.videoParameters)

			// Verify result
			if err != nil {
				t.Errorf("SUT threw an error: %v", err)
			}
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Result differs. Actual: %s, Expected %s", actual, test.expected)
			}
		})
	}
}

func TestQueryWithMediaInfo_WhenParameterInvalid(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		videoParameters []string
	}{
		{
			[]string{"Invalid"},
		},
		{
			[]string{"Height", "Invalid"},
		},
		{
			[]string{"Invalid", "Height"},
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v", test.videoParameters), func(t *testing.T) {
			// Exercise SUT
			_, err := video.QueryWithMediaInfo("testdata/By-The-Law.mkv", test.videoParameters)

			// Verify result
			if err == nil {
				t.Errorf("Expected SUT to return an error")
			}
		})
	}
}

func TestQueryWithMediaInfo_WhenPathDoesNotExist(t *testing.T) {
	// Exercise SUT
	_, err := video.QueryWithMediaInfo("this/path/does/not.exist", []string{"Width"})

	// Verify result
	if err == nil {
		t.Errorf("Expected SUT to return an error")
	}
}
