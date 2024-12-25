package main

import (
	_ "embed"
	"fmt"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/core"
)

type Place struct {
	core.Point
	length int
	end    *core.Point
}

func (p Place) Priority() int {
	return -(p.end.Manhattan(p.Point) + p.length)
}

func FastestRoute(grid core.Grid, start, end core.Point) []core.Point {
	var track []core.Point
	current := start
	for {
		track = append(track, current)
		if current == end {
			return track
		}
		for _, adj := range current.Cardinal() {
			if (len(track) == 1 || adj != track[len(track)-2]) && grid.GetPoint(adj) != '#' {
				current = adj
			}
		}
	}
}

func Part1(input string, N int) int {
	grid := core.NewGridFromLines(input)
	return CountCheats(grid, N, 2)
}

func CountCheats(grid core.Grid, N, C int) int {
	start := *grid.FindOne('S')
	end := *grid.FindOne('E')
	track := FastestRoute(grid, start, end)
	cheats := 0
	for l1, t1 := range track {
		for l2 := l1 + 3; l2 < len(track); l2++ {
			t2 := track[l2]
			diff := t1.Manhattan(t2)
			if diff <= C && l2-l1-diff >= N && l2-l1 > diff {
				cheats += 1
			}
		}
	}
	return cheats
}

func Part2(input string, N int) int {
	grid := core.NewGridFromLines(input)
	return CountCheats(grid, N, 20)
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 20")
	input = input[:len(input)-1]
	fmt.Printf("\tPart 1: %d\n", Part1(input, 100))
	fmt.Printf("\tPart 2: %d\n", Part2(input, 100))
}
