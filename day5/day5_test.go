package main

import "testing"

func TestSampleInput(t *testing.T) {

	boardingPass := parseBoardingPass("FBFBBFFRLR")
	checkValue(boardingPass.row, 44, t)
	checkValue(boardingPass.column, 5, t)
	checkValue(boardingPass.seatID, 357, t)

	boardingPass = parseBoardingPass("BFFFBBFRRR")
	checkValue(boardingPass.row, 70, t)
	checkValue(boardingPass.column, 7, t)
	checkValue(boardingPass.seatID, 567, t)

	boardingPass = parseBoardingPass("FFFBBBFRRR")
	checkValue(boardingPass.row, 14, t)
	checkValue(boardingPass.column, 7, t)
	checkValue(boardingPass.seatID, 119, t)

	boardingPass = parseBoardingPass("BBFFBBFRLL")
	checkValue(boardingPass.row, 102, t)
	checkValue(boardingPass.column, 4, t)
	checkValue(boardingPass.seatID, 820, t)

}

func checkValue(actual int64, expected int64, t *testing.T) {
	if actual != expected {
		t.Errorf("actual was incorrect, got: %d, want: %d.", actual, expected)
	}
}
