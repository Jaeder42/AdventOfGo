package day6

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// DaySix run day six
func DaySix() {
	content, err := ioutil.ReadFile("./day6/day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	groupAnswers := strings.Split(text, "\n\n")
	partOne(groupAnswers)
	partTwo(groupAnswers)
}

func partOne(groupAnswers []string) {
	total := 0
	for _, answer := range groupAnswers {
		uniqueArray := uniqueNonEmptyElementsOf(strings.Split(strings.ReplaceAll(answer, "\n", ""), ""))
		total += len(uniqueArray)
	}
	fmt.Println("Part one: ", total)
}
func uniqueNonEmptyElementsOf(s []string) []string {
	unique := make(map[string]bool, len(s))
	us := make([]string, len(unique))
	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				us = append(us, elem)
				unique[elem] = true
			}
		}
	}
	return us
}

func partTwo(groupAnswers []string) {
	total := 0
	for _, answers := range groupAnswers {
		options := make(map[string]int)
		answersArr := strings.Split(answers, "\n")
		size := len(answersArr)
		for _, answer := range answersArr {
			if len(answer) < 1 {
				size--
			}
			split := strings.Split(answer, "")

			for _, ans := range split {
				options[ans] = options[ans] + 1
			}
		}
		for _, element := range options {
			if element == size {
				total++
			}
		}

	}
	fmt.Println("Part two: ", total)
}
