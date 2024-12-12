package main

import (
	_ "embed"
	"fmt"
)

func Part1(input string) int {
	input = input[:len(input)-1]
	disk := make([]int, len(input))
	for i, char := range input {
		disk[i] = int(char - '0')
	}
	frontPtr := 0
	backPtr := len(input) - 1
	sum := 0
	position := 0
	done := false
	for {
		if frontPtr == backPtr {
			done = true
		}
		if frontPtr%2 == 0 {
			for range disk[frontPtr] {
				sum += (frontPtr / 2) * position
				position += 1
			}
			frontPtr += 1
		} else {
			if disk[backPtr] > disk[frontPtr] {
				disk[backPtr] = disk[backPtr] - disk[frontPtr]
				for range disk[frontPtr] {
					sum += (backPtr / 2) * position
					position += 1
				}
				frontPtr += 1
			} else if disk[frontPtr] > disk[backPtr] {
				disk[frontPtr] = disk[frontPtr] - disk[backPtr]
				for range disk[backPtr] {
					sum += (backPtr / 2) * position
					position += 1
				}
				backPtr -= 2
			} else {
				for range disk[frontPtr] {
					sum += (backPtr / 2) * position
					position += 1
				}
				backPtr -= 2
				frontPtr += 1
			}
		}
		if done {
			break
		}
	}
	return sum
}

func Part2(input string) int {
	input = input[:len(input)-1]
	disk := make([]int, len(input))
	for i, char := range input {
		disk[i] = int(char - '0')
	}
	positions := make([]int, len(disk))
	positions[0] = 0
	for i, l := range disk[:len(disk)-1] {
		positions[i+1] = positions[i] + l
	}
	sum := 0
	for backPtr := len(disk) - 1; backPtr >= 0; backPtr -= 2 {
		fit := false
		for frontPtr := 1; frontPtr < backPtr; frontPtr += 2 {
			if disk[frontPtr] >= disk[backPtr] {
				disk[frontPtr] = disk[frontPtr] - disk[backPtr]
				for i := range disk[backPtr] {
					sum += (backPtr / 2) * (positions[frontPtr] + i)
				}
				positions[frontPtr] += disk[backPtr]
				fit = true
				break
			}
		}
		if !fit {
			for i := range disk[backPtr] {
				sum += (backPtr / 2) * (positions[backPtr] + i)
			}
		}
	}
	return sum
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 9")
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
