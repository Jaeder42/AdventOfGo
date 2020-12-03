package day3

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// DayThree runs day two
func DayThree() {
	content, err := ioutil.ReadFile("./day3/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	array := strings.Split(text, "\n")
	slope := make([][]string, len(array))
	for index, xString := range array {
		slope[index] = strings.Split(xString, "")
	}
	// PartOne
	fmt.Println("PartOne:", calcTree(slope, 3, 1))

	// PartTwo
	first := calcTree(slope, 1, 1)
	second := calcTree(slope, 3, 1)
	third := calcTree(slope, 5, 1)
	fourth := calcTree(slope, 7, 1)
	fifth := calcTree(slope, 1, 2)

	fmt.Println("PartTwo: ", first*second*third*fourth*fifth)
}

func calcTree(slope [][]string, xStep int, yStep int) int {
	xIndex := xStep
	totalTrees := 0
	for i := yStep; i < len(slope); i += yStep {
		if slope[i][xIndex] == "#" {
			totalTrees++
		}
		xIndex = xStep + xIndex
		if xIndex >= len(slope[i]) {
			xIndex = xIndex % (len(slope[i]))
		}
	}
	return totalTrees
}
