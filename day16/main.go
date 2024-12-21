package main

import (
	_ "embed"
	"fmt"
	"math"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/core"
	"github.com/michael-janssen-dev/advent-of-code-2024-go/data-structures/set"
)

type Direction int

type QueueItem struct {
	point     core.Point
	direction core.Point
	score     int
	previous  core.Point
}

func (q QueueItem) Priority() int {
	return -q.score
}

func Part1(input string) int {
	grid := core.NewGridFromLines(input)
	queue := core.NewPriorityQueue[QueueItem]()
	queue.PushSafe(QueueItem{core.NewPoint(1, len(grid)-2), core.EAST, 0, core.NewPoint(-1, -1)})
	visited := set.NewSet[core.Point]()
	for {
		item := queue.PopSafe()
		if visited.Contains(item.point) {
			continue
		}
		visited.Add(item.point)
		char := grid.GetPoint(item.point)
		if char == 'E' {
			return item.score
		}
		next := item.point.Add(item.direction)
		if grid.GetPoint(next) != '#' {
			queue.PushSafe(QueueItem{next, item.direction, item.score + 1, item.point})
		}
		clockwise := item.point.Add(item.direction.Clockwise())
		if grid.GetPoint(clockwise) != '#' {
			queue.PushSafe(QueueItem{clockwise, item.direction.Clockwise(), item.score + 1001, item.point})
		}
		counterClockwise := item.point.Add(item.direction.CounterClockwise())
		if grid.GetPoint(counterClockwise) != '#' {
			queue.PushSafe(QueueItem{counterClockwise, item.direction.CounterClockwise(), item.score + 1001, item.point})
		}
	}
}

func Part2(input string) int {
	grid := core.NewGridFromLines(input)
	queue := core.NewPriorityQueue[QueueItem]()
	queue.PushSafe(QueueItem{core.NewPoint(1, len(grid)-2), core.EAST, 0, core.NewPoint(-1, -1)})
	best := math.MaxInt
	path := make(map[core.Point]set.Set[core.Point])
	visited := make(map[core.Point]int)
	for {
		item := queue.PopSafe()
		if item.score > best {
			break
		}
		if fastest, ok := visited[item.point]; ok {
			if fastest == item.score || fastest == item.score-1000 {
				for pathItem := range path[item.previous] {
					path[item.point].Add(pathItem)
				}
			}
			continue
		}
		visited[item.point] = item.score
		path[item.point] = set.NewSet[core.Point]()
		for pathItem := range path[item.previous] {
			path[item.point].Add(pathItem)
		}
		path[item.point].Add(item.point)
		char := grid.GetPoint(item.point)
		if char == 'E' {
			best = item.score
		}
		next := item.point.Add(item.direction)
		if grid.GetPoint(next) != '#' {
			queue.PushSafe(QueueItem{next, item.direction, item.score + 1, item.point})
		}
		clockwise := item.point.Add(item.direction.Clockwise())
		if grid.GetPoint(clockwise) != '#' {
			queue.PushSafe(QueueItem{clockwise, item.direction.Clockwise(), item.score + 1001, item.point})
		}
		counterClockwise := item.point.Add(item.direction.CounterClockwise())
		if grid.GetPoint(counterClockwise) != '#' {
			queue.PushSafe(QueueItem{counterClockwise, item.direction.CounterClockwise(), item.score + 1001, item.point})
		}
	}
	return path[core.NewPoint(len(grid[0])-2, 1)].Len()
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 16")
	input = input[:len(input)-1]
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
