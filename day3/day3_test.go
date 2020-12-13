package main

import "testing"

func TestCalculateTreeCount(t *testing.T) {

	multipleSum := 1

	expected := 2
	treeCount := calculateTreeCount("sample.txt", 1, 1)
	checkValue(treeCount, expected, t)

	multipleSum = multipleSum * treeCount

	expected = 7
	treeCount = calculateTreeCount("sample.txt", 3, 1)
	checkValue(treeCount, expected, t)

	multipleSum = multipleSum * treeCount

	expected = 3
	treeCount = calculateTreeCount("sample.txt", 5, 1)
	checkValue(treeCount, expected, t)

	multipleSum = multipleSum * treeCount

	expected = 4
	treeCount = calculateTreeCount("sample.txt", 7, 1)
	checkValue(treeCount, expected, t)

	multipleSum = multipleSum * treeCount

	expected = 2
	treeCount = calculateTreeCount("sample.txt", 1, 2)
	checkValue(treeCount, expected, t)

	expected = 336
	multipleSum = multipleSum * treeCount
	checkValue(multipleSum, expected, t)
}

func checkValue(treeCount int, expected int, t *testing.T) {
	if treeCount != expected {
		t.Errorf("treeCount was incorrect, got: %d, want: %d.", treeCount, expected)
	}
}
