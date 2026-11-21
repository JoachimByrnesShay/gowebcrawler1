package main_test

import (
	"fmt"
	"gowebcrawler/utilities"
	"strings"
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name          string
		inputURL      string
		expected      string
		errorContains string
	}{
		{name: "remove scheme https",
			inputURL: "https://www.boot.dev/blog/path",
			expected: "www.boot.dev/blog/path",
		},
		{name: "remove separator suffix",
			inputURL: "https://www.boot.dev/blog/path/",
			expected: "www.boot.dev/blog/path",
		},
		{name: "remove scheme http",
			inputURL: "http://www.boot.dev/blog/path",
			expected: "www.boot.dev/blog/path",
		},
		{name: "remove scheme http and separator suffix",
			inputURL: "http://www.boot.dev/blog/path/",
			expected: "www.boot.dev/blog/path",
		},
		{name: "handle capital letters",
			inputURL: "HTTP://www.BOOT.DEV/blog/path/",
			expected: "www.boot.dev/blog/path",
		},
		{
			name:          "handle missing scheme",
			inputURL:      "://www.boot.dev/blog/path/",
			expected:      "",
			errorContains: "could not parse url",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := utilities.NormalizeURL(tc.inputURL)
			if actual != tc.expected {
				fmt.Println("actual != tc.expected")
			}
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
