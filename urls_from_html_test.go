package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestURLFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "no anchor tags",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<span>Boot.dev</span>
		<span>Boot.dev</span>
	</body>
</html>
			`,
			expected: []string{},
		},
		{
			name:     "all absolute paths",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="https://blog.boot.dev/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
			`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "all relative paths",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="/path/two">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
			`,
			expected: []string{"https://blog.boot.dev/path/one", "https://blog.boot.dev/path/two"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
			`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		// more tests here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			inputURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}

			actual, err := getURLsFromHTML(tc.inputBody, inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URLs: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
