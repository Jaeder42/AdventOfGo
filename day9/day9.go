package day9

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

// DayNine run day nine
func DayNine() {
	content, err := ioutil.ReadFile("./day9/day9.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	codes := strings.Split(text, "\n")
	partOne(codes)
	partTwo(codes, 1639024365)
}

func partOne(codes []string) {
	for i := 26; i < len(codes); i++ {
		code, err := strconv.Atoi(codes[i])
		if err != nil {
			log.Fatal(err)
		}
		if !validateCode(codes[i-25:i], code) {
			fmt.Println("Part one, code:", code, "index:", i)
			return
		}
	}
}

func partTwo(codes []string, code int) {
	result := findCont(codes, 0, code)
	sort.Ints(result)
	fmt.Println("Part two:", result[0]+result[len(result)-1])
}

func validateCode(preamble []string, code int) bool {
	for index, first := range preamble {
		intFirst, err := strconv.Atoi(first)
		if err != nil {
			log.Fatal(err)
		}
		for _, second := range preamble[index:] {
			intSecond, err := strconv.Atoi(second)
			if err != nil {
				log.Fatal(err)
			}
			if intFirst+intSecond == code {
				return true
			}
		}
	}
	return false
}

func findCont(codes []string, index int, code int) []int {
	total := 0
	intCodes := []int{}
	for i := index; i < len(codes); i++ {
		current, err := strconv.Atoi(codes[i])
		if err != nil {
			log.Fatal(err)
		}
		total += current
		intCodes = append(intCodes, current)
		if total == code && i != 653 {
			return intCodes
		} else if total > code {
			return findCont(codes, index+1, code)
		}
	}
	return nil
}
