package day8

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type programResult struct {
	acc    int
	failed bool
}

//DayEight run day eight
func DayEight() {
	content, err := ioutil.ReadFile("./day8/day8.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := string(content)
	instructions := strings.Split(text, "\n")
	partOne(instructions)
	partTwo(instructions)
}

func partOne(instructions []string) {
	fmt.Println("Part one", runSet(instructions).acc)
}

func partTwo(instructions []string) {
	for i := 0; i < len(instructions); i++ {
		instruction := strings.Split(instructions[i], " ")[0]
		arg := strings.Split(instructions[i], " ")[1]
		alteredInstructions := make([]string, len(instructions))
		copy(alteredInstructions, instructions)
		if instruction == "jmp" {
			alteredInstructions[i] = "nop " + arg
		} else if instruction == "nop" {
			alteredInstructions[i] = "jmp " + arg
		} else {
			continue
		}
		result := runSet(alteredInstructions)
		if !result.failed {
			fmt.Println("Part two", result.acc)
			return
		}
	}
}

func runSet(instructions []string) programResult {
	result := programResult{acc: 0, failed: false}
	acc := 0
	instructionsRan := make(map[int]bool)
	index := 0
	for {
		if index >= len(instructions) {
			break
		}
		if instructionsRan[index] {
			result.acc = acc
			result.failed = true
			return result
		}
		instructionsRan[index] = true
		instructionArray := strings.Split(instructions[index], " ")
		instruction := instructionArray[0]
		arg, err := strconv.Atoi(instructionArray[1])
		if err != nil {
			log.Fatal(err)
		}
		switch instruction {
		case "jmp":
			index += arg
			break
		case "acc":
			acc += arg
			index++
			break
		case "nop":
			index++
			break
		default:
			index++
			break
		}

	}
	result.acc = acc
	result.failed = false
	return result
}
