package main

import (
	"reflect"
	"testing"
)

func TestSortCrawlEntries(t *testing.T) {
	inputMap1 := make(map[string]int)
	inputMap1["one"] = 1
	inputMap1["two"] = 2
	inputMap1["three"] = 3

	inputMap2 := make(map[string]int)
	inputMap2["two"] = 2
	inputMap2["three"] = 3
	inputMap2["one"] = 1

	inputMap3 := make(map[string]int)
	inputMap3["three"] = 3
	inputMap3["two"] = 2
	inputMap3["one"] = 1

	var expected []crawlEntry
	expected = append(expected, crawlEntry{"three", 3})
	expected = append(expected, crawlEntry{"two", 2})
	expected = append(expected, crawlEntry{"one", 1})

	tests := []struct {
		name     string
		inputMap map[string]int
		expected []crawlEntry
	}{
		{
			name:     "sort reverse-sorted",
			inputMap: inputMap1,
			expected: expected,
		},
		{
			name:     "sort shuffled",
			inputMap: inputMap2,
			expected: expected,
		},
		{
			name:     "sort sorted",
			inputMap: inputMap3,
			expected: expected,
		},
		// more tests here
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := sortCrawlEntries(tc.inputMap)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
