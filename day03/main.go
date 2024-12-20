package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

const MUL_REGEX = `(?:mul\((\d+),(\d+)\))|(do\(\))|(don't\(\))`

func Part1(input string) int {
	regex := regexp.MustCompile(MUL_REGEX)
	matches := regex.FindAllStringSubmatch(input, -1)
	result := 0
	for i := range matches {
		left, _ := strconv.Atoi(matches[i][1])
		right, _ := strconv.Atoi(matches[i][2])
		result += left * right
	}
	return result
}

func Part2(input string) int {
	regex := regexp.MustCompile(MUL_REGEX)
	matches := regex.FindAllStringSubmatch(input, -1)
	result := 0
	enabled := true
	for i := range matches {
		if matches[i][1] != "" {
			// Must
			if !enabled {
				continue
			}
			left, _ := strconv.Atoi(matches[i][1])
			right, _ := strconv.Atoi(matches[i][2])
			result += left * right
		} else if matches[i][3] != "" {
			// Do
			enabled = true
		} else {
			// Don't
			enabled = false
		}
	}
	return result
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 3")
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
