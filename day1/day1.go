package day1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// DayOne runs day one
func DayOne() {
	content, err := ioutil.ReadFile("./day1/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	var array = strings.Split(text, "\n")
	var answerPartOne = findPartOne(array)
	var answerPartTwo = findPartTwo(array)
	fmt.Print("Part1: ")
	fmt.Println(answerPartOne)
	fmt.Print("Part2: ")
	fmt.Println(answerPartTwo)

}

func findPartOne(array []string) int {
	for index, element := range array {
		num, err := strconv.Atoi(element)
		if err != nil {
		} else {
			for i := index; i < len(array); i++ {
				j, err := strconv.Atoi(array[i])
				if err != nil {
				} else {
					if num+j == 2020 {
						return num * j
					}
				}
			}
		}
	}
	return 0
}
func findPartTwo(array []string) int {
	for index, element := range array {
		num, err := strconv.Atoi(element)
		if err != nil {
		} else {
			for i := index; i < len(array); i++ {
				num2, err := strconv.Atoi(array[i])
				if err != nil {
				} else {
					for j := i; j < len(array); j++ {
						num3, err := strconv.Atoi(array[j])
						if err != nil {
						} else {
							if num+num2+num3 == 2020 {
								return num * num2 * num3
							}
						}
					}
				}
			}
		}
	}
	return 0
}
