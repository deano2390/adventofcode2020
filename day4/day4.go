package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Passport struct {
	BirthYear      string // byr
	IssueYear      string // iyr
	ExpirationYear string // eyr
	Height         string // hgt
	HairColor      string // hcl
	EyeColor       string // ecl
	PassportID     string // pid
	CountryID      string // cid
}

func main() {
	partOne()
}

func partOne() {
	solution := getNumberOfValidPassports("input.txt")
	fmt.Printf("part one answer: %d\n", solution)
}

func getNumberOfValidPassports(inputSource string) int {
	passports := parsePassports(inputSource)

	validCount := 0

	for _, passport := range passports {
		if passport.BirthYear != "" &&
			passport.IssueYear != "" &&
			passport.ExpirationYear != "" &&
			passport.Height != "" &&
			passport.HairColor != "" &&
			passport.EyeColor != "" &&
			passport.PassportID != "" {
			validCount++
		}
	}

	return validCount
}

func parsePassports(inputSource string) []Passport {
	file, err := os.Open(inputSource)

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	passports := []Passport{}
	var passport Passport

	for scanner.Scan() {

		line := scanner.Text()

		if line == "" {
			passports = append(passports, passport)
			passport = Passport{}
			continue
		}

		parts := strings.Split(line, " ")
		for _, part := range parts {
			keyValuePair := strings.Split(part, ":")
			if len(keyValuePair) == 2 {
				key := keyValuePair[0]
				value := keyValuePair[1]

				switch key {
				case "byr":
					passport.BirthYear = value
				case "iyr":
					passport.IssueYear = value
				case "eyr":
					passport.ExpirationYear = value
				case "hgt":
					passport.Height = value
				case "hcl":
					passport.HairColor = value
				case "ecl":
					passport.EyeColor = value
				case "pid":
					passport.PassportID = value
				case "cid":
					passport.CountryID = value
				}
			}
		}

	}

	passports = append(passports, passport)

	return passports
}
