package main

import (
	"fmt"
	"log"
	"os"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/core"
	"github.com/michael-janssen-dev/advent-of-code-2024-go/data-structures/set"
)

type Queue[T any] struct {
	List []T
}

func (q *Queue[T]) Push(item T) {
	q.List = append(q.List, item)
}

func (q *Queue[T]) Pop() T {
	item := q.List[0]
	q.List = q.List[1:]
	return item
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.List) == 0
}

func NewQueue[T any]() Queue[T] {
	return Queue[T]{make([]T, 0)}
}

func Part1(input string) int {
	grid := core.NewGridFromFile(input)
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
		queue := NewQueue[core.Point]()
		queue.Push(s)
		for {
			if queue.IsEmpty() {
				break
			}
			item := queue.Pop()
			gridNr := grid.GetPoint(&item)
			for _, dir := range item.Cardinal() {
				if !grid.InGrid(dir) || visited.Contains(dir) {
					continue
				}
				dirNr := grid.GetPoint(&dir)
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
	grid := core.NewGridFromFile(input)
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
		queue := NewQueue[core.Point]()
		queue.Push(s)
		for {
			if queue.IsEmpty() {
				break
			}
			item := queue.Pop()
			gridNr := grid.GetPoint(&item)
			for _, dir := range item.Cardinal() {
				if !grid.InGrid(dir) {
					continue
				}
				dirNr := grid.GetPoint(&dir)
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

func main() {
	fmt.Println("Day 10")
	file, err := os.ReadFile("input/inp.txt")
	if err != nil {
		log.Fatalln("Could not find file, exiting...")
	}
	input := string(file)
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
