package main

import (
	_ "embed"
	"fmt"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/core"
	"github.com/michael-janssen-dev/advent-of-code-2024-go/data-structures/set"
)

func Part1(input string) int {
	grid := core.NewGridFromLines(input)
	start := make([]core.Point, 0)
	for y := range grid {
		for x := range grid[y] {
			if grid.Get(x, y) == '0' {
				start = append(start, core.NewPoint(x, y))
			}
		}
	}

	result := 0
	for _, s := range start {
		visited := set.NewSet[core.Point]()
		queue := core.NewQueue[core.Point]()
		queue.Push(s)
		for {
			if queue.IsEmpty() {
				break
			}
			item := queue.Pop()
			gridNr := grid.GetPoint(item)
			for _, dir := range item.Cardinal() {
				if !grid.InGrid(dir) || visited.Contains(dir) {
					continue
				}
				dirNr := grid.GetPoint(dir)
				if dirNr-gridNr == 1 {
					queue.Push(dir)
					visited.Add(dir)
				}
			}
			if gridNr == '9' {
				result += 1
			}
		}

	}

	return result
}

func Part2(input string) int {
	grid := core.NewGridFromLines(input)
	start := make([]core.Point, 0)
	for y := range grid {
		for x := range grid[y] {
			if grid.Get(x, y) == '0' {
				start = append(start, core.NewPoint(x, y))
			}
		}
	}

	result := 0
	for _, s := range start {
		queue := core.NewQueue[core.Point]()
		queue.Push(s)
		for {
			if queue.IsEmpty() {
				break
			}
			item := queue.Pop()
			gridNr := grid.GetPoint(item)
			for _, dir := range item.Cardinal() {
				if !grid.InGrid(dir) {
					continue
				}
				dirNr := grid.GetPoint(dir)
				if dirNr-gridNr == 1 {
					queue.Push(dir)
				}
			}
			if gridNr == '9' {
				result += 1
			}
		}

	}

	return result
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 10")
	input = input[:len(input)-1]
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
