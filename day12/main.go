package main

import (
	"fmt"
	"log"
	"os"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/core"
)

func CalculateSides(grid core.Grid, partOf [][]bool, plant core.Char) int {
	sides := 0
	for y := range partOf {
		side := false
		for x := range partOf[y] {
			if !partOf[y][x] {
				side = false
				continue
			}
			if grid.Get(x, y-1) == plant {
				side = false
				continue
			}
			if !side {
				side = true
				sides += 1
			}
		}
	}
	for y := len(partOf) - 1; y >= 0; y-- {
		side := false
		for x := range partOf[y] {
			if !partOf[y][x] {
				side = false
				continue
			}
			if grid.Get(x, y+1) == plant {
				side = false
				continue
			}
			if !side {
				side = true
				sides += 1
			}
		}
	}
	for x := range partOf[0] {
		side := false
		for y := range partOf {
			if !partOf[y][x] {
				side = false
				continue
			}
			if grid.Get(x+1, y) == plant {
				side = false
				continue
			}
			if !side {
				side = true
				sides += 1
			}
		}
	}
	for x := len(partOf[0]) - 1; x >= 0; x-- {
		side := false
		for y := range partOf {
			if !partOf[y][x] {
				side = false
				continue
			}
			if grid.Get(x-1, y) == plant {
				side = false
				continue
			}
			if !side {
				side = true
				sides += 1
			}
		}
	}
	return sides
}

func FloodFill(grid core.Grid, x, y int, visited [][]bool) (int, int, [][]bool) {
	plant := grid[y][x]
	perimeter := 0
	area := 1
	queue := core.NewQueue[core.Point]()
	queue.Push(core.NewPoint(x, y))
	visited[y][x] = true
	partOf := make([][]bool, len(visited))
	for i := range visited {
		partOf[i] = make([]bool, len(visited[i]))
	}
	partOf[y][x] = true
	for {
		if queue.IsEmpty() {
			break
		}
		p := queue.Pop()
		x, y := p.X, p.Y
		perimeter += 4
		for _, d := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
			dx, dy := x+d[0], y+d[1]
			if grid.Get(dx, dy) == plant {
				perimeter -= 1
				if !visited[dy][dx] {
					area += 1
					partOf[dy][dx] = true
					visited[dy][dx] = true
					queue.Push(core.NewPoint(dx, dy))
				}
			}
		}
	}
	return area, perimeter, partOf
}

func Part1(input string) int {
	grid := core.NewGridFromFile(input)
	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}

	sum := 0
	for j := range grid {
		for i := range grid[j] {
			if visited[j][i] {
				continue
			}
			area, perimiter, _ := FloodFill(grid, i, j, visited)
			sum += area * perimiter
		}
	}
	return sum
}

func Part2(input string) int {
	grid := core.NewGridFromFile(input)
	visited := make([][]bool, len(grid))
	for i := range grid {
		visited[i] = make([]bool, len(grid[i]))
	}

	sum := 0
	for j := range grid {
		for i := range grid[j] {
			if visited[j][i] {
				continue
			}
			area, _, partOf := FloodFill(grid, i, j, visited)
			sum += area * CalculateSides(grid, partOf, grid.Get(i, j))
		}
	}
	return sum
}

func main() {
	fmt.Println("Day 12")
	file, err := os.ReadFile("input/inp.txt")
	if err != nil {
		log.Fatalln("Could not find file, exiting...")
	}
	input := string(file)
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
