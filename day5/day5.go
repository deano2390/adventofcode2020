package main

import (
	"fmt"
)

type BoardingPass struct {
	row    int
	column int
	seatID int
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
	return BoardingPass{} // TODO
}
