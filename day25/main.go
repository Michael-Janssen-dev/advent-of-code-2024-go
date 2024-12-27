package main

import (
	_ "embed"
	"fmt"
	"strings"
)

func Part1(input string) int {
	locks := make([][5]int, 0)
	keys := make([][5]int, 0)
	for _, item := range strings.Split(input, "\n\n") {
		grid := strings.Split(item, "\n")
		var size [5]int
		for i := range 5 {
			height := 0
			for j := range 5 {
				if grid[j+1][i] == '#' {
					height++
				}
			}
			size[i] = height
		}
		if grid[0][0] == '.' {
			keys = append(keys, size)
		} else {
			locks = append(locks, size)
		}
	}
	result := 0
	for _, key := range keys {
	lockLoop:
		for _, lock := range locks {
			for i := range 5 {
				if key[i]+lock[i] > 5 {
					continue lockLoop
				}
			}
			result += 1
		}
	}
	return result
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 25")
	input = input[:len(input)-1]
	fmt.Printf("\tPart 1: %d\n", Part1(input))
}
