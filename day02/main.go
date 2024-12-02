package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	safe := 0
	for _, line := range lines[:len(lines)-1] {
		splitted := strings.Split(line, " ")
		parsedLine := make([]int, len(splitted))
		for i, s := range splitted {
			parsedLine[i], _ = strconv.Atoi(s)
		}
		if isIncreasingList(parsedLine, -1) || isDecreasingList(parsedLine, -1) {
			safe += 1
		}
	}
	return safe
}

func isIncreasing(a, b int) bool {
	return b-a > 0 && b-a < 4
}

func isDecreasing(a, b int) bool {
	return b-a < 0 && b-a > -4
}

func isIncreasingList(list []int, skipIndex int) bool {
	increasing := true
	for i := range list[:len(list)-1] {
		if i == skipIndex {
			continue
		}
		if i+1 == skipIndex {
			if skipIndex < len(list)-1 && !isIncreasing(list[i], list[i+2]) {
				increasing = false
			}
			continue
		}
		if !isIncreasing(list[i], list[i+1]) {
			increasing = false
		}
	}
	return increasing
}

func isDecreasingList(list []int, skipIndex int) bool {
	decreasing := true
	for i := range len(list) - 1 {
		if i == skipIndex {
			continue
		}
		if i+1 == skipIndex {
			if skipIndex < len(list)-1 && !isDecreasing(list[i], list[i+2]) {
				decreasing = false
			}
			continue
		}
		if !isDecreasing(list[i], list[i+1]) {
			decreasing = false
		}
	}
	return decreasing
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	safe := 0
	for _, line := range lines[:len(lines)-1] {
		splitted := strings.Split(line, " ")
		parsedLine := make([]int, len(splitted))
		for i, s := range splitted {
			parsedLine[i], _ = strconv.Atoi(s)
		}
		increasing := false
		for j := 0; j < len(parsedLine); j++ {
			if isIncreasingList(parsedLine, j) {
				increasing = true
				break
			}
		}
		decreasing := false
		for j := 0; j < len(parsedLine); j++ {
			if isDecreasingList(parsedLine, j) {
				decreasing = true
				break
			}
		}
		if increasing || decreasing {
			safe += 1
		}
	}
	return safe
}

func main() {
	fmt.Println("Day 2")
	file, err := os.ReadFile("input/inp.txt")
	if err != nil {
		log.Fatalln("Could not find file, exiting...")
	}
	input := string(file)
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
