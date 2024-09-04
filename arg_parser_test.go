package main

import (
	"testing"
)

func TestParseMaxConcurrency(t *testing.T) {
	defaultMaxConcurrency := 1

	tests := []struct {
		name      string
		inputArgs []string
		expected  int
	}{
		{
			name:      "no args",
			inputArgs: []string{},
			expected:  defaultMaxConcurrency,
		},
		{
			name:      "non-int arg",
			inputArgs: []string{"url", "wee", "woo"},
			expected:  defaultMaxConcurrency,
		},
		{
			name:      "missing page limit",
			inputArgs: []string{"url", "10"},
			expected:  10,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := parseMaxConcurrency(tc.inputArgs, defaultMaxConcurrency)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected limit: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}

func TestParseMaxPages(t *testing.T) {
	defaultMaxPages := 10

	tests := []struct {
		name      string
		inputArgs []string
		expected  int
	}{
		{
			name:      "no args",
			inputArgs: []string{},
			expected:  defaultMaxPages,
		},
		{
			name:      "non-int arg",
			inputArgs: []string{"url", "wee", "woo"},
			expected:  defaultMaxPages,
		},
		{
			name:      "missing concurrency limit",
			inputArgs: []string{"url", "wee", "100"},
			expected:  100,
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := parseMaxPages(tc.inputArgs, defaultMaxPages)
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected limit: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
