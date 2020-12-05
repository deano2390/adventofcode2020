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
		parts := strings.Split(line, ": ")
		password := parts[1]

		parts = strings.Split(parts[0], " ")
		requiredChar := parts[1]

		parts = strings.Split(parts[0], "-")

		pos1, err := strconv.Atoi(parts[0])

		if err != nil {
			log.Fatalf("failed to parse min")
		}

		pos2, err := strconv.Atoi(parts[1])

		fmt.Printf("pass: %s requiredChar: %s pos1: %d pos2: %d\n", password, requiredChar, pos1, pos2)

		if isPasswordValid(password, requiredChar, pos1, pos2) {
			validCount++
		}
	}

	fmt.Printf("%d passwords are valid\n", validCount)

	file.Close()

}

func isPasswordValid(password string, requiredChar string, pos1 int, pos2 int) bool {

	char1 := getCharAt(password, pos1-1)
	char2 := getCharAt(password, pos2-1)

	return (char1 == requiredChar) != (char2 == requiredChar)
}

func getCharAt(password string, position int) string {
	return string(password[position])
}
