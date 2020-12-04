package day4

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// DayFour runs day two
func DayFour() {
	content, err := ioutil.ReadFile("./day4/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	passports := strings.Split(text, "\n\n")
	partOne(passports)
	partTwo(passports)
}
func partOne(passports []string) {
	totalValids := 0
	for _, passport := range passports {
		if validate(passport) {
			totalValids++
		}
	}
	fmt.Println("Passports ", totalValids)
}

func partTwo(passports []string) {
	totalValids := 0

	for _, passport := range passports {
		if validate(passport) {
			if validate2(passport) {
				totalValids++
			}
		}
	}
	fmt.Println("Passports ", totalValids)

}
func validate2(passport string) bool {
	fields := strings.Split(strings.ReplaceAll(passport, "\n", " "), " ")
	for _, field := range fields {
		if !validateField(field) {
			return false
		}
	}

	return true
}

func validateField(field string) bool {
	if field == "" {
		return true
	}
	split := strings.Split(field, ":")
	switch split[0] {
	case "byr":
		return validateNum(split[1], 1920, 2002)
	case "iyr":
		return validateNum(split[1], 2010, 2020)
	case "eyr":
		return validateNum(split[1], 2020, 2030)
	case "hgt":
		if strings.Contains(split[1], "in") {
			return validateNum(strings.ReplaceAll(split[1], "in", ""), 59, 76)
		} else if strings.Contains(split[1], "cm") {
			return validateNum(strings.ReplaceAll(split[1], "cm", ""), 150, 193)
		}
		return false
	case "hcl":
		matched, err := regexp.MatchString(`^#(?:[0-9a-f]{3}){2}$`, split[1])
		if err != nil {
			return false
		}
		return matched
	case "ecl":
		if len(split[1]) < 3 {
			return false
		}
		oneOf := "amb blu brn gry grn hzl oth"
		count := strings.Count(oneOf, split[1])
		if count == 1 {

			return true
		}
		return false
	case "pid":
		matched, err := regexp.MatchString(`\d{9}`, split[1])
		if err != nil {
			return false
		}
		return matched && len(split[1]) == 9
	case "cid":
		return true
	default:
		return false
	}
}

func validateNum(numberString string, min int, max int) bool {
	num, err := strconv.Atoi(numberString)
	if err != nil {
		return false
	}
	return num >= min && num <= max
}

func validate(passport string) bool {
	requireds := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	validCount := 0
	for _, required := range requireds {
		validCount += strings.Count(passport, required)
	}
	return validCount >= 7
}
