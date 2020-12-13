package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	partOne()
}

func partOne() {
	solution := getNumberOfValidPassports("input.txt")
	fmt.Printf("part one answer: %d\n", solution)
}

func getNumberOfValidPassports(inputSource string) int {

	file, err := os.Open(inputSource)

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	return 0
}
