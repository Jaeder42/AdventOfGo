package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type color struct {
	Amt   uint
	Color string
}

// DaySeven run day seven
func DaySeven() {
	rules := parseRules("./day7/day7.txt")
	fmt.Printf("Part 1: %d\n", part1(rules))
	fmt.Printf("Part 2: %d\n", part2(rules))
}

func parseRules(path string) map[string][]color {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	rules := make(map[string][]color)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		mainColor := strings.Join(parts[:2], " ")
		rest := make([]color, 0)
		i := 5
		for {
			if len(parts) < i+2 {
				break
			}
			amt, err := strconv.Atoi(parts[i-1])
			if err != nil {
				amt = 0
			}

			rest = append(rest, color{
				Color: strings.Join(parts[i:i+2], " "),
				Amt:   uint(amt),
			})
			i += 4
		}
		rules[mainColor] = rest
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	return rules
}

func part1(rules map[string][]color) int {
	contains := make(map[string]struct{})
	check("shiny gold", rules, contains)
	return int(len(contains))
}

func check(color string, rules map[string][]color, contains map[string]struct{}) {
	for c := range getContains(rules, color) {
		contains[c] = struct{}{}
		check(c, rules, contains)
	}
}

func getContains(input map[string][]color, bag string) map[string]struct{} {
	contains := make(map[string]struct{})
	for mainColor, restColors := range input {
		for _, clr := range restColors {
			if clr.Color == bag {
				contains[mainColor] = struct{}{}
			}
		}
	}
	return contains
}

func part2(rules map[string][]color) uint {
	return getAmt("shiny gold", rules)
}

func getAmt(color string, rules map[string][]color) uint {
	var amt uint
	for _, clr := range rules[color] {
		amt += clr.Amt
		amt += clr.Amt * getAmt(clr.Color, rules)
	}
	return amt
}
