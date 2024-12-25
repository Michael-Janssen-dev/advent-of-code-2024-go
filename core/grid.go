package core

import (
	"fmt"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/data-structures/set"
)

type Char byte

type Grid [][]Char

func NewGridFromLines(input string) Grid {
	lines := strings.Split(input, "\n")
	grid := make([][]Char, len(lines))
	for i, line := range lines {
		grid[i] = []Char(line)
	}
	return grid
}

func NewGridWithSize(x, y int, defaultChar Char) Grid {
	grid := make([][]Char, y)
	for i := range y {
		grid[i] = make([]Char, x)
		for j := range x {
			grid[i][j] = defaultChar
		}
	}
	return grid
}

func (g Grid) FindAll(c Char) set.Set[Point] {
	points := set.NewSet[Point]()
	for y := range g {
		for x, val := range g[y] {
			if val == c {
				points.Add(NewPoint(x, y))
			}
		}
	}
	return points
}

func (g Grid) FindOne(c Char) *Point {
	for y := range g {
		for x, val := range g[y] {
			if val == c {
				point := NewPoint(x, y)
				return &point
			}
		}
	}
	return nil
}

func (g Grid) InGrid(p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.X < len(g[0]) && p.Y < len(g)
}

func (g Grid) Values() set.Set[Char] {
	result := set.NewSet[Char]()
	for _, line := range g {
		for _, val := range line {
			result.Add(val)
		}
	}
	return result
}

func (g Grid) Get(x, y int) Char {
	if x < 0 || x >= len(g[0]) || y < 0 || y >= len(g) {
		return 0
	}
	return g[y][x]
}

func (g Grid) GetPoint(p Point) Char {
	return g.Get(p.X, p.Y)
}

func (g Grid) SetPoint(p Point, char Char) {
	g[p.Y][p.X] = char
}

func (g Grid) Print() {
	for y := range g {
		for _, char := range g[y] {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g Grid) Copy() Grid {
	new := make(Grid, len(g))
	for y := range g {
		new[y] = make([]Char, len(g[y]))
		copy(new[y], g[y])
	}
	return new
}
