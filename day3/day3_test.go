package main

import "testing"

func TestCalculateTreeCount(t *testing.T) {

	var expected = 7
	var treeCount = calculateTreeCount("sample.txt")

	if treeCount != expected {
		t.Errorf("treeCount was incorrect, got: %d, want: %d.", treeCount, expected)
	}
}
