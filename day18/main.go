package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/core"
)

type Place struct {
	core.Point
	path int
}

func (p Place) Priority() int {
	return -p.path
}

func Part1(input string, n, size int) int {
	grid := core.NewGridWithSize(size, size, '.')
	lines := strings.Split(input, "\n")
	for _, line := range lines[:n] {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		grid[y][x] = '#'
	}
	queue := core.NewPriorityQueue[Place]()
	queue.PushSafe(Place{core.NewPoint(0, 0), 0})
	visited := make([]bool, size*size)
	for {
		place := queue.PopSafe()
		if place.X == size-1 && place.Y == size-1 {
			return place.path
		}
		if visited[place.Y*size+place.X] {
			continue
		}
		visited[place.Y*size+place.X] = true
		for _, point := range place.Cardinal() {
			if grid.GetPoint(point) == '.' && !visited[point.Y*size+point.X] {
				queue.PushSafe(Place{point, place.path + 1})
			}
		}
	}
}

func Part2(input string, n, size int) string {
	grid := core.NewGridWithSize(size, size, '.')
	lines := strings.Split(input, "\n")
	for _, line := range lines[:n] {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		grid[y][x] = '#'
	}
	for _, line := range lines[n:] {
		split := strings.Split(line, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		grid[y][x] = '#'
		queue := core.NewPriorityQueue[Place]()
		queue.PushSafe(Place{core.NewPoint(0, 0), 0})
		visited := make([]bool, size*size)
		for {
			if len(queue) == 0 {
				return line
			}
			place := queue.PopSafe()
			if place.X == size-1 && place.Y == size-1 {
				break
			}
			if visited[place.Y*size+place.X] {
				continue
			}
			visited[place.Y*size+place.X] = true
			for _, point := range place.Cardinal() {
				if grid.GetPoint(point) == '.' && !visited[point.Y*size+point.X] {
					queue.PushSafe(Place{point, place.path + 1})
				}
			}
		}
	}
	return "Not found"
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 18")
	input = input[:len(input)-1]
	fmt.Printf("\tPart 1: %d\n", Part1(input, 1024, 71))
	fmt.Printf("\tPart 2: %s\n", Part2(input, 1024, 71))
}
