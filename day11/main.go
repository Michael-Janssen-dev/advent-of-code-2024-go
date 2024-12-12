package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/core"
)

type Pair struct {
	left, right int
}

func Stones(pair Pair, inner func(Pair) int) int {
	nr, iterations := pair.left, pair.right
	if iterations == 0 {
		return 1
	}

	if nr == 0 {
		return inner(Pair{1, iterations - 1})
	}
	nrOfCharacters := int(math.Log10(float64(nr))) + 1
	if nrOfCharacters%2 == 0 {
		leftHalf := nr / int(math.Pow10(nrOfCharacters/2))
		rightHalf := nr % int(math.Pow10(nrOfCharacters/2))
		left := inner(Pair{leftHalf, iterations - 1})
		right := inner(Pair{rightHalf, iterations - 1})
		return left + right
	}
	return inner(Pair{nr * 2024, iterations - 1})

}

func Part1(input string) int {
	stonesRaw := strings.Split(input[:len(input)-1], " ")
	stones := make([]int, len(stonesRaw))
	for i, stone := range stonesRaw {
		stones[i], _ = strconv.Atoi(stone)
	}

	result := 0
	dpStones := core.DP(Stones)
	for _, stone := range stones {
		result += dpStones(Pair{stone, 25})
	}
	return result

}

func Part2(input string) int {
	stonesRaw := strings.Split(input[:len(input)-1], " ")
	stones := make([]int, len(stonesRaw))
	for i, stone := range stonesRaw {
		stones[i], _ = strconv.Atoi(stone)
	}

	result := 0
	dpStones := core.DP(Stones)
	for _, stone := range stones {
		result += dpStones(Pair{stone, 75})
	}
	return result

}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 11")
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
