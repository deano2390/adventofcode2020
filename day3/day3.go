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
	var treeCount = calculateTreeCount("map.txt", 3, 1)
	fmt.Printf("part one treeCount: %d\n", treeCount)
}

func calculateTreeCount(inputSource string, right int, down int) int {

	file, err := os.Open(inputSource)

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var rows []string

	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	file.Close()

	runLength := len(rows)
	patternWidth := len(rows[0])
	requiredWidth := runLength * right

	fmt.Printf("%d runLength\n", runLength)
	fmt.Printf("%d patternWidth\n", patternWidth)
	fmt.Printf("%d requiredWidth\n", requiredWidth)

	rowPos := 0
	colPos := 0
	treeCount := 0

	for true {
		rowPos += down

		row := rows[rowPos]
		row = buildRow(row, colPos+right)

		colPos += right
		itemOnMap := row[colPos]

		if itemOnMap == '#' {
			treeCount++
		}

		if rowPos >= runLength-down {
			break
		}
	}

	return treeCount
}

func buildRow(row string, requiredWidth int) string {

	for len(row) < requiredWidth {
		row += row
	}

	return row
}
