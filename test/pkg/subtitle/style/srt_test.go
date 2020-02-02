package style_test

import (
	"fmt"
	"testing"

	"github.com/liampulles/cabiria/pkg/subtitle/style"
)

func TestRemoveStylesFromSRTText(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		fixture  string
		expected string
	}{
		// No tags
		{
			"",
			"",
		},
		{
			"a",
			"a",
		},
		{
			"text",
			"text",
		},
		// A tag
		{
			"<i>text",
			"text",
		},
		{
			"te<i>xt",
			"text",
		},
		{
			"text</i>",
			"text",
		},
		{
			"te<font color=\"FFFFFF\">xt",
			"text",
		},
		// Many tags
		{
			"<i>text</i>",
			"text",
		},
		{
			"<i>text<t>",
			"text",
		},
		{
			"<i><t>text",
			"text",
		},
		{
			"<i>t<b>ext</i>",
			"text",
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("[%s]", test.fixture), func(t *testing.T) {
			input := test.fixture

			// Exercise SUT
			actual := style.RemoveStylesFromSRTText(test.fixture)

			// Verify input unchanged
			if test.fixture != input {
				t.Errorf("Input changed. Modified input: %s", test.fixture)
			}
			// Verify result
			if actual != test.expected {
				t.Errorf("Text differs. Actual: %s, Expected: %s", actual, test.expected)
			}
		})
	}
}
