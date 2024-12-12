package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	leftList := make([]int, len(lines)-1)
	rightList := make([]int, len(lines)-1)
	for i, line := range lines[:len(lines)-1] {
		splitted := strings.Split(line, "   ")
		left, _ := strconv.Atoi(splitted[0])
		right, _ := strconv.Atoi(splitted[1])
		leftList[i] = left
		rightList[i] = right
	}
	sort.Ints(leftList)
	sort.Ints(rightList)
	distance := 0
	for i := range leftList {
		distance += int(math.Abs(float64(leftList[i] - rightList[i])))
	}
	return distance
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	leftList := make([]int, len(lines)-1)
	rightList := make([]int, len(lines)-1)
	for i, line := range lines[:len(lines)-1] {
		splitted := strings.Split(line, "   ")
		left, _ := strconv.Atoi(splitted[0])
		right, _ := strconv.Atoi(splitted[1])
		leftList[i] = left
		rightList[i] = right
	}
	score := 0
	for i := range leftList {
		similarity := 0
		for j := range rightList {
			if rightList[j] == leftList[i] {
				similarity += 1
			}
		}
		score += similarity * leftList[i]
	}
	return score
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 1")
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
