package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// DayTwo runs day two
func DayTwo() {
	content, err := ioutil.ReadFile("./day2/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	array := strings.Split(text, "\n")
	totalFirst := 0
	totalSecond := 0
	for _, entry := range array {
		if validateEntry(entry, 1) {
			totalFirst++
		}
		if validateEntry(entry, 2) {
			totalSecond++
		}
	}
	fmt.Println("First: ", totalFirst, "second: ", totalSecond)
}

func validateEntry(entry string, rule int) bool {
	// Extract the data
	array := strings.Split(entry, ":")
	rules := strings.Split(array[0], " ")
	pass := array[1]
	maxMin := strings.Split(rules[0], "-")
	target := strings.ReplaceAll(rules[1], " ", "")
	min, err := strconv.Atoi(maxMin[0])
	max, err := strconv.Atoi(maxMin[1])

	if err != nil {
		return false
	}
	if rule == 1 {
		return rulesPartOne(target, pass, max, min)
	} else if rule == 2 {
		return rulesPartTwo(target, pass, max, min)
	}
	return false
}

func rulesPartOne(target string, pass string, max int, min int) bool {
	count := strings.Count(pass, target)
	return count <= max && count >= min
}
func rulesPartTwo(target string, pass string, max int, min int) bool {
	passArr := strings.Split(pass, "")
	return (passArr[max] == target || passArr[min] == target) && passArr[max] != passArr[min]

}
