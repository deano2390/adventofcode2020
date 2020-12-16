package main

import (
	"fmt"
	"strconv"
	"strings"
)

type BoardingPass struct {
	row    int64
	column int64
	seatID int64
}

func main() {
	partOne()
}

func partOne() {
	solution := getHighestSeatID()
	fmt.Printf("part one answer: %d\n", solution)
}

func getHighestSeatID() int {
	return -1 // TODO
}

func parseBoardingPass(boardingPass string) BoardingPass {
	rowStr := boardingPass[0:7]
	colStr := boardingPass[7:10]

	rowStr = strings.ReplaceAll(rowStr, "B", "1")
	rowStr = strings.ReplaceAll(rowStr, "F", "0")

	colStr = strings.ReplaceAll(colStr, "R", "1")
	colStr = strings.ReplaceAll(colStr, "L", "0")

	row, err := strconv.ParseInt(rowStr, 2, 64)

	if err != nil {
		fmt.Println(err)
	}

	col, err := strconv.ParseInt(colStr, 2, 64)

	if err != nil {
		fmt.Println(err)
	}

	seatID := (row * 8) + col

	fmt.Printf("boardingPass: %s, row:%d, col:%d, seatID:%d\n", boardingPass, row, col, seatID)
	return BoardingPass{row, col, seatID}
}
