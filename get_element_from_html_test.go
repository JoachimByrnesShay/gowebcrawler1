package main_test

import (
	"fmt"
	"gowebcrawler/utilities"
	"testing"
)

func TestGetHeadingFromHTML(t *testing.T) {

	tests := []struct {
		name          string
		inputHTML     string
		expected      string
		errorContains string
	}{
		{
			name:      "test basic",
			inputHTML: "<html><body><h1>Test Title</h1></body></html",
			expected:  "Test Title",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := utilities.GetHeadingFromHTML(tc.inputHTML)
			if actual != tc.expected {
				fmt.Println(err)
				fmt.Printf("Test %v -'%s' actual != tc.expected\n", i, tc.name)
			}

			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}
		})
	}

}

// func TestGetFirstParagraphFromHTML(t *testing.T) {
// 	fmt.Println(t)
// }
