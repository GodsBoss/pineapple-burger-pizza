package text_test

import (
	"testing"

	"github.com/GodsBoss/pineapple-burger-pizza/pkg/text"
)

func TestLines(t *testing.T) {
	testcases := map[string]struct {
		maxLineWidth  int
		content       string
		expectedLines []string
	}{
		"empty": {
			maxLineWidth:  5,
			content:       "",
			expectedLines: []string{},
		},
		"no split, one word": {
			maxLineWidth: 10,
			content:      "12345",
			expectedLines: []string{
				"12345",
			},
		},
		"no split, multiple words": {
			maxLineWidth: 20,
			content:      "123 123 123",
			expectedLines: []string{
				"123 123 123",
			},
		},
		"split, many words": {
			maxLineWidth: 10,
			content:      "Hello, world!",
			expectedLines: []string{
				"Hello,",
				"world!",
			},
		},
		"split, edge case": {
			maxLineWidth: 10,
			content:      "foobar 123 shooshoo",
			expectedLines: []string{
				"foobar 123",
				"shooshoo",
			},
		},
		"very long word": {
			maxLineWidth: 5,
			content:      "ThisIsAVeryLongWord",
			expectedLines: []string{
				"ThisI",
				"sAVer",
				"yLong",
				"Word",
			},
		},
		"blanks": {
			maxLineWidth: 10,
			content:      "xx   yy",
			expectedLines: []string{
				"xx   yy",
			},
		},
	}

	for name := range testcases {
		testcase := testcases[name]

		t.Run(
			name,
			func(t *testing.T) {
				lines := text.Lines(testcase.maxLineWidth, testcase.content)

				if len(lines) != len(testcase.expectedLines) {
					t.Fatalf("expected %d lines, got %+v", len(testcase.expectedLines), lines)
				}

				for i := range lines {
					if lines[i] != testcase.expectedLines[i] {
						t.Errorf("expected line #%d to be '%s', got '%s'", i, testcase.expectedLines[i], lines[i])
					}
				}
			},
		)
	}
}
