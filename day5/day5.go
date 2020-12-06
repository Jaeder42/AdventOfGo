package day5

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strings"
)

// DayFive runs day 5
func DayFive() {
	content, err := ioutil.ReadFile("./day5/day5.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)

	boardingPasses := strings.Split(text, "\n")
	partBoth(boardingPasses)
}

func partBoth(boardingPasses []string) {
	ids := []int{}
	for _, pass := range boardingPasses {
		if len(pass) > 1 {

			passArray := strings.Split(pass, "")
			row := traverse(passArray[0:7], 0, 127)
			col := traverse(passArray[7:10], 0, 7)

			ids = append(ids, (row*8)+col)
		}
	}
	sort.Ints(ids)
	fmt.Println("Part one: ", ids[len(ids)-1])
	for i := 1; i < len(ids); i++ {
		if ids[i]-ids[i-1] > 1 {
			fmt.Println("Part two:", ids[i]-1)
		}
	}
}

func traverse(rowString []string, start float64, end float64) int {
	d, rowString := rowString[0], rowString[1:]
	if len(rowString) > 0 {
		if d == "F" || d == "L" {
			end = end - math.Ceil(((end - start) / 2))
		} else {
			start = end - math.Floor((end-start)/2)
		}
		return traverse(rowString, start, end)
	}

	if d == "F" || d == "L" {
		return int(start)
	}
	return int(end)
}
