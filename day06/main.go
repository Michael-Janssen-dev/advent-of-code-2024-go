package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/data-structures/set"
)

type Coor struct {
	x, y int
}

type State struct {
	coor Coor
	dir  int
}

func GetCoorFromStates(states set.Set[State]) set.Set[Coor] {
	result := set.NewSet[Coor]()
	for s := range states {
		result.Add(s.coor)
	}
	return result
}

func Run(grid set.Set[Coor], start Coor, max_x, max_y int, extra_obstacle Coor) (bool, set.Set[Coor]) {
	location := State{start, 0}
	visited := set.NewSet[State]()

	for {
		if visited.Contains(location) {
			return true, GetCoorFromStates(visited)
		}
		visited.Add(location)
		var next State
		if location.dir == 0 {
			next = State{Coor{location.coor.x, location.coor.y - 1}, location.dir}
		} else if location.dir == 1 {
			next = State{Coor{location.coor.x + 1, location.coor.y}, location.dir}
		} else if location.dir == 2 {
			next = State{Coor{location.coor.x, location.coor.y + 1}, location.dir}
		} else {
			next = State{Coor{location.coor.x - 1, location.coor.y}, location.dir}
		}

		if grid.Contains(next.coor) || next.coor == extra_obstacle {
			location.dir = (location.dir + 1) % 4
			continue
		}

		location = next
		if next.coor.x < 0 || next.coor.y < 0 || next.coor.x > max_x || next.coor.y > max_y {
			return false, GetCoorFromStates(visited)
		}
	}
}

func preprocess(input string) (set.Set[Coor], Coor, int, int) {
	splitted := strings.Split(input, "\n")
	splitted = splitted[:len(splitted)-1]
	grid := set.NewSet[Coor]()
	var start Coor
	for y, line := range splitted {
		for x, char := range line {
			if char == '#' {
				grid.Add(Coor{x, y})
			} else if char == '^' {
				start = Coor{x, y}
			}
		}
	}

	max_x, max_y := len(splitted[0])-1, len(splitted)-1
	return grid, start, max_x, max_y
}

func Part1(input string) int {
	grid, start, max_x, max_y := preprocess(input)

	_, result := Run(grid, start, max_x, max_y, Coor{-1, -1})

	return result.Len()
}

func RunAsync(grid set.Set[Coor], start Coor, max_x, max_y int, obstacle Coor, c chan bool) {
	loop, _ := Run(grid, start, max_x, max_y, obstacle)
	c <- loop
}

func Part2(input string) int {
	grid, start, max_x, max_y := preprocess(input)

	result := 0
	channel := make(chan bool)
	send := 0
	_, visited := Run(grid, start, max_x, max_y, Coor{-1, -1})
	for obstacle := range visited {
		go RunAsync(grid, start, max_x, max_y, obstacle, channel)
		send += 1
	}

	for range send {
		if <-channel {
			result += 1
		}
	}

	return result

}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 6")
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
