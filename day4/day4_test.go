package main

import "testing"

func TestSampleInput(t *testing.T) {

	expected := 2
	actual := getNumberOfValidPassports("sample.txt")
	checkValue(actual, expected, t)

}

func checkValue(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Errorf("actual was incorrect, got: %d, want: %d.", actual, expected)
	}
}