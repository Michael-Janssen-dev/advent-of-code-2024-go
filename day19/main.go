package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/core"
)

func Part1(input string) int {
	split := strings.Split(input, "\n\n")
	towels := strings.Split(split[0], ", ")
	lines := strings.Split(split[1], "\n")
	result := 0
	isPossible := func(pattern string, inner func(string) bool) bool {
		if len(pattern) == 0 {
			return true
		}
	towelLoop:
		for _, towel := range towels {
			if len(towel) > len(pattern) {
				continue
			}
			for i := range towel {
				if pattern[i] != towel[i] {
					continue towelLoop
				}
			}
			if inner(pattern[len(towel):]) {
				return true
			}
		}
		return false
	}
	dpIsPossible := core.DP(isPossible)
	for _, line := range lines {
		if dpIsPossible(line) {
			result += 1
		}
	}
	return result
}

func Part2(input string) int {
	split := strings.Split(input, "\n\n")
	towels := strings.Split(split[0], ", ")
	lines := strings.Split(split[1], "\n")
	result := 0
	isPossible := func(pattern string, inner func(string) int) int {
		if len(pattern) == 0 {
			return 1
		}
		result := 0
	towelLoop:
		for _, towel := range towels {
			if len(towel) > len(pattern) {
				continue
			}
			for i := range towel {
				if pattern[i] != towel[i] {
					continue towelLoop
				}
			}
			result += inner(pattern[len(towel):])
		}
		return result
	}
	dpIsPossible := core.DP(isPossible)
	for _, line := range lines {
		result += dpIsPossible(line)
	}
	return result
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 19")
	input = input[:len(input)-1]
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
