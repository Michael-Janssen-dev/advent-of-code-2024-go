package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/core"
	"github.com/michael-janssen-dev/advent-of-code-2024-go/data-structures/set"
)

type Robot struct {
	x, y, vx, vy int
}

const WIDTH = 101
const HEIGHT = 103

func (r *Robot) PositionAfter(t int) (int, int) {
	x := (r.x + (t%WIDTH)*r.vx) % WIDTH
	y := (r.y + (t%HEIGHT)*r.vy) % HEIGHT
	if x < 0 {
		x += WIDTH
	}
	if y < 0 {
		y += HEIGHT
	}
	return x, y
}

func (r *Robot) Move() {
	r.x = (r.x + r.vx) % WIDTH
	r.y = (r.y + r.vy) % HEIGHT
	if r.x < 0 {
		r.x += WIDTH
	}
	if r.y < 0 {
		r.y += HEIGHT
	}
}

func (r *Robot) PositionHash() int {
	return r.y*WIDTH + r.x
}
func (r *Robot) Position() core.Point {
	return core.Point{X: r.x, Y: r.y}
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	robots := make([]Robot, len(lines))
	for i, line := range lines {
		robots[i] = ParseRobot(line)
	}
	quadrants := make([]int, 4)
	for _, robot := range robots {
		x, y := robot.PositionAfter(100)
		if x < WIDTH/2 && y < HEIGHT/2 {
			quadrants[0] += 1
		} else if x > WIDTH/2 && y < HEIGHT/2 {
			quadrants[1] += 1
		} else if x < WIDTH/2 && y > HEIGHT/2 {
			quadrants[2] += 1
		} else if x > WIDTH/2 && y > HEIGHT/2 {
			quadrants[3] += 1
		}
	}
	product := 1
	for _, q := range quadrants {
		product *= q
	}
	return product
}

func ParseRobot(line string) Robot {
	x, _ := strconv.Atoi(line[strings.Index(line, "=")+1 : strings.Index(line, ",")])
	y, _ := strconv.Atoi(line[strings.Index(line, ",")+1 : strings.Index(line, " ")])
	vx, _ := strconv.Atoi(line[strings.LastIndex(line, "=")+1 : strings.LastIndex(line, ",")])
	vy, _ := strconv.Atoi(line[strings.LastIndex(line, ",")+1:])
	return Robot{
		x, y, vx, vy,
	}
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	robots := make([]Robot, len(lines))
	for i, line := range lines {
		robots[i] = ParseRobot(line)
	}
	t := 0
	for {
		t += 1
		if t%100 == 0 {
			fmt.Println(t)
		}
		hash := set.NewSet[int]()
		// positions := set.NewSet[core.Point]()
		colUsed := make([]bool, WIDTH)
		rowUsed := make([]bool, HEIGHT)
		for i := range robots {
			robots[i].Move()
			colUsed[robots[i].x] = true
			rowUsed[robots[i].y] = true
			hash.Add(robots[i].PositionHash())
			// positions.Add(robots[i].Position())
		}

		nrColUnUsed := 0
		for _, col := range colUsed {
			if !col {
				nrColUnUsed += 1
			}
		}
		nrRowUnUsed := 0
		for _, row := range rowUsed {
			if !row {
				nrRowUnUsed += 1
			}
		}

		if nrColUnUsed > 10 && nrRowUnUsed > 10 {
			for y := range HEIGHT {
				for x := range WIDTH {
					point := core.Point{X: x, Y: y}
					if hash.Contains(point.Y*WIDTH + point.X) {
						fmt.Print("X")
					} else {
						fmt.Print(".")
					}
				}
				fmt.Println()
			}
			fmt.Println(t)
			fmt.Scanln()
		}
	}
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 14")
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
