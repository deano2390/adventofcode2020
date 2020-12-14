package main

import "testing"

func TestSampleInput(t *testing.T) {

	// TEST sample.txt
	expected := 4
	passports := parsePassports("sample.txt")
	checkValue(len(passports), expected, t)

	expected = 0
	actual := getNumberOfValidPassports("sample.txt")
	checkValue(actual, expected, t)

	// TEST invalid_sample.txt
	expected = 4
	passports = parsePassports("invalid_sample.txt")
	checkValue(len(passports), expected, t)

	expected = 0
	actual = getNumberOfValidPassports("invalid_sample.txt")
	checkValue(actual, expected, t)

	// TEST valid_sample.txt
	expected = 4
	passports = parsePassports("valid_sample.txt")
	checkValue(len(passports), expected, t)

	expected = 0
	actual = getNumberOfValidPassports("valid_sample.txt")
	checkValue(actual, expected, t)

}

func checkValue(actual int, expected int, t *testing.T) {
	if actual != expected {
		t.Errorf("actual was incorrect, got: %d, want: %d.", actual, expected)
	}
}
