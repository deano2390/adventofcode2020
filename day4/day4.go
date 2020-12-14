package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
		if isPassportValid(passport) {
			validCount++
		}
	}

	return validCount
}

func isPassportValid(passport Passport) bool {
	return isBirthYearValid(passport.BirthYear) &&
		isIssueYearValid(passport.IssueYear) &&
		isExpirationYearValid(passport.ExpirationYear) &&
		isHeightValid(passport.Height) &&
		isHairColorValid(passport.HairColor) &&
		isEyeColorValid(passport.EyeColor) &&
		isPassportIDValid(passport.PassportID)
}

func isBirthYearValid(birthYear string) bool {
	return validateNumber(birthYear, 1920, 2002)
}

func isIssueYearValid(birthYear string) bool {
	return validateNumber(birthYear, 2010, 2020)
}

func isExpirationYearValid(birthYear string) bool {
	return validateNumber(birthYear, 2020, 2030)
}

func isHeightValid(height string) bool {

	if strings.HasSuffix(height, "cm") {
		height = strings.Replace(height, "cm", "", 1)
		return validateNumber(height, 150, 193)
	} else if strings.HasSuffix(height, "in") {
		height = strings.Replace(height, "in", "", 1)
		return validateNumber(height, 59, 76)
	} else {
		return false
	}
}

func validateNumber(numberStr string, min int, max int) bool {
	if i, err := strconv.Atoi(numberStr); err == nil {
		return (i >= min && i <= max)
	}

	return false
}

func isHairColorValid(hairColor string) bool {
	if strings.HasPrefix(hairColor, "#") {
		hairColor = strings.Replace(hairColor, "#", "", 1)
		if len(hairColor) == 6 {
			if i, err := strconv.ParseUint(hairColor, 16, 64); err == nil {
				return (i >= 0 && i <= 16777215)
			}

		}
	}

	return false
}

func isEyeColorValid(eyeColor string) bool {

	isValid := false

	switch eyeColor {
	case "amb":
		fallthrough
	case "blu":
		fallthrough
	case "brn":
		fallthrough
	case "gry":
		fallthrough
	case "grn":
		fallthrough
	case "hzl":
		fallthrough
	case "oth":
		isValid = true
	}

	return isValid
}

func isPassportIDValid(passportID string) bool {
	if len(passportID) == 9 {
		return true
	}

	return false
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
