package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("passwords.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	validCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		password := parts[1]

		parts = strings.Split(parts[0], " ")
		requiredChar := parts[1]

		parts = strings.Split(parts[0], "-")

		min, err := strconv.Atoi(parts[0])

		if err != nil {
			log.Fatalf("failed to parse min")
		}

		max, err := strconv.Atoi(parts[1])

		fmt.Printf("pass: %s requiredChar: %s min: %d max: %d\n", password, requiredChar, min, max)

		if isPasswordValid(password, requiredChar, min, max) {
			validCount++
		}
	}

	fmt.Printf("%d passwords are valid\n", validCount)

	file.Close()

}

func isPasswordValid(password string, requiredChar string, min int, max int) bool {
	count := strings.Count(password, requiredChar)
	return (count >= min && count <= max)
}
