package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/data-structures/set"
)

func mix(number, secret int) int {
	return number ^ secret
}

func prune(number int) int {
	return number % 16777216
}

func step(number int) int {
	number = prune(mix(number*64, number))
	number = prune(mix(number/32, number))
	number = prune(mix(number*2048, number))
	return number
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	buyers := make([]int, len(lines))
	for i, line := range lines {
		buyers[i], _ = strconv.Atoi(line)
	}
	result := 0
	for i := range buyers {
		for range 2000 {
			buyers[i] = step(buyers[i])
		}
		result += buyers[i]
	}
	return result
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	buyers := make([]int, len(lines))
	for i, line := range lines {
		buyers[i], _ = strconv.Atoi(line)
	}
	sequenceHashToBananas := make(map[int]int)
	for i := range buyers {
		secret := buyers[i]
		lastFourDiffs := make([]int, 0)
		visited := set.NewSet[int]()
		for range 2000 {
			next := step(secret)
			diff := next%10 - secret%10
			lastFourDiffs = append(lastFourDiffs, diff)
			if len(lastFourDiffs) > 4 {
				lastFourDiffs = lastFourDiffs[1:]
			}
			if len(lastFourDiffs) == 4 {
				hash := 0
				for k := range 4 {
					hash = hash*20 + lastFourDiffs[k]
				}
				if visited.Contains(hash) {
					secret = next
					continue
				}
				visited.Add(hash)
				if _, ok := sequenceHashToBananas[hash]; !ok {
					sequenceHashToBananas[hash] = next % 10
				} else {
					sequenceHashToBananas[hash] += next % 10
				}
			}
			secret = next
		}
	}
	max := 0
	for _, v := range sequenceHashToBananas {
		if v > max {
			max = v
		}
	}
	return max
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 22")
	input = input[:len(input)-1]
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
