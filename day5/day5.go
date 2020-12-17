package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BoardingPass struct {
	row    int64
	column int64
	seatID int64
}

func main() {
	part1 := getHighestSeatID()
	part2 := findMySeat()
	fmt.Printf("part one answer: %d\n", part1)
	fmt.Printf("part two answer: %d\n", part2)
}

func findMySeat() int {

	scanner := openFile()

	tickets := make(map[int64]bool)

	for scanner.Scan() {
		boardingPassStr := scanner.Text()
		boardingPass := parseBoardingPass(boardingPassStr)
		tickets[boardingPass.seatID] = true
	}

	fmt.Printf("map size: %d\n", len(tickets))

	for ticketIndex := 0; ticketIndex < len(tickets); ticketIndex++ {
		if tickets[int64(ticketIndex)] == false {

			value, isset := tickets[int64(ticketIndex)-1]
			if isset {
				fmt.Printf("value: %b\n", value)
				value, isset := tickets[int64(ticketIndex)+1]
				if isset {
					fmt.Printf("value: %b\n", value)
					return ticketIndex
				}
			}
		}
	}

	return -1
}

func getHighestSeatID() int64 {

	scanner := openFile()

	var highestSeatID int64 = 0

	for scanner.Scan() {
		boardingPassStr := scanner.Text()
		boardingPass := parseBoardingPass(boardingPassStr)
		printBoardingPass(boardingPass)
		if boardingPass.seatID > highestSeatID {
			highestSeatID = boardingPass.seatID
		}
	}
	return highestSeatID
}

func openFile() *bufio.Scanner {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	return scanner
}

func parseBoardingPass(boardingPassStr string) BoardingPass {
	rowStr := boardingPassStr[0:7]
	colStr := boardingPassStr[7:10]

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

	boardingPass := BoardingPass{row, col, seatID}
	printBoardingPass(boardingPass)
	return boardingPass
}

func printBoardingPass(boardingPass BoardingPass) {
	fmt.Printf("boardingPass row:%d, col:%d, seatID:%d\n", boardingPass.row, boardingPass.column, boardingPass.seatID)
}
