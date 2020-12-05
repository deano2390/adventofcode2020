package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("map.txt")

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

	downStride := 1
	strafe := 3
	runLength := len(rows)
	patternWidth := len(rows[0])
	requiredWidth := runLength * strafe

	fmt.Printf("%d runLength\n", runLength)
	fmt.Printf("%d patternWidth\n", patternWidth)
	fmt.Printf("%d requiredWidth\n", requiredWidth)

	rowPos := 0
	colPos := 0
	treeCount := 0

	for true {
		rowPos += downStride

		row := rows[rowPos]
		row = buildRow(row, colPos+strafe)

		colPos += strafe
		itemOnMap := row[colPos]

		if itemOnMap == '#' {
			treeCount++
		}

		if rowPos >= runLength-downStride {
			break
		}
	}

	fmt.Printf("treeCount: %d\n", treeCount)
}

func buildRow(row string, requiredWidth int) string {

	for len(row) < requiredWidth {
		row += row
	}

	return row
}
