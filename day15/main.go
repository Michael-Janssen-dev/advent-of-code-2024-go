package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/core"
)

type BoxGrid struct {
	core.Grid
	pos core.Point
}

const PLAYER = '@'
const BOX = 'O'
const WALL = '#'

func (g *BoxGrid) Copy() BoxGrid {
	return BoxGrid{
		g.Grid.Copy(),
		g.pos,
	}
}

func (g *BoxGrid) move(char rune) {
	var dir core.Point
	if char == '<' {
		dir = core.WEST
	} else if char == '>' {
		dir = core.EAST
	} else if char == '^' {
		dir = core.NORTH
	} else if char == 'v' {
		dir = core.SOUTH
	} else {
		return
	}
	startPos := g.pos
	newPos := g.pos
	move := '@'
	for {
		newPos = newPos.Add(dir)
		point := g.GetPoint(&newPos)
		if point == WALL {
			g.pos = startPos
			return
		}
		if point == BOX {
			if move == '@' {
				g.pos = newPos
				move = 'O'
			}
		} else {
			if move == '@' {
				g.pos = newPos
			} else {
				g.SetPoint(&newPos, 'O')
			}
			g.SetPoint(&g.pos, '@')
			g.SetPoint(&startPos, '.')
			return
		}
	}
}

func (g *BoxGrid) moveWide(char rune) {
	var dir core.Point
	if char == '<' {
		dir = core.WEST
	} else if char == '>' {
		dir = core.EAST
	} else if char == '^' {
		dir = core.NORTH
	} else if char == 'v' {
		dir = core.SOUTH
	} else {
		return
	}
	queue := core.NewQueue[core.Point]()
	changes := make(map[core.Point]rune, 0)
	changes[g.pos] = '.'
	queue.Push(g.pos.Add(dir))
	changes[g.pos.Add(dir)] = '@'
	for !queue.IsEmpty() {
		pos := queue.Pop()
		point := g.GetPoint(&pos)
		if point == WALL {
			return
		}
		if point == '[' {
			changes[pos.Add(dir)] = '['
			if dir == core.NORTH || dir == core.SOUTH {
				if _, ok := changes[pos.East()]; !ok {
					changes[pos.East()] = '.'
				}
				queue.Push(pos.Add(dir))
				queue.Push(pos.Add(dir).East())
				changes[pos.Add(dir).East()] = ']'
			} else {
				queue.Push(pos.Add(dir).Add(dir))
				changes[pos.Add(dir)] = '['
				changes[pos.Add(dir).Add(dir)] = ']'
			}
		} else if point == ']' {
			changes[pos.Add(dir)] = ']'
			if dir == core.NORTH || dir == core.SOUTH {
				if _, ok := changes[pos.West()]; !ok {
					changes[pos.West()] = '.'
				}
				queue.Push(pos.Add(dir))
				queue.Push(pos.Add(dir).West())
				changes[pos.Add(dir).West()] = '['
			} else {
				queue.Push(pos.Add(dir).Add(dir))
				changes[pos.Add(dir)] = ']'
				changes[pos.Add(dir).Add(dir)] = '['
			}
		} else {
			continue
		}
	}
	for change, val := range changes {
		g.SetPoint(&change, core.Char(val))
	}
	g.pos = g.pos.Add(dir)
}

func (g *BoxGrid) GPSCoordinates() int {
	points := g.FindAll('O')
	result := 0
	for point := range points {
		result += point.Y*100 + point.X
	}

	points = g.FindAll('[')
	for point := range points {
		result += point.Y*100 + point.X
	}
	return result
}

func NewBoxGridFromFile(input string) BoxGrid {
	inner := core.NewGridFromLines(input)
	pos := inner.FindOne('@')
	return BoxGrid{
		inner,
		*pos,
	}
}

func (g *BoxGrid) WiderBoxGrid() BoxGrid {
	newGrid := make(core.Grid, len(g.Grid))
	for y := range g.Grid {
		newGrid[y] = make([]core.Char, len((g.Grid)[y])*2)
		for x, char := range (g.Grid)[y] {
			if char == WALL {
				newGrid[y][2*x] = WALL
				newGrid[y][2*x+1] = WALL
			} else if char == BOX {
				newGrid[y][2*x] = '['
				newGrid[y][2*x+1] = ']'
			} else if char == PLAYER {
				newGrid[y][2*x] = '@'
				newGrid[y][2*x+1] = '.'
			} else {
				newGrid[y][2*x] = '.'
				newGrid[y][2*x+1] = '.'
			}
		}
	}
	return BoxGrid{
		newGrid,
		core.Point{X: g.pos.X * 2, Y: g.pos.Y},
	}
}

func Part1(input string) int {
	splitted := strings.Split(input, "\n\n")
	rawGrid, rawInstructions := splitted[0]+"\n", splitted[1][:len(splitted[1])-1]
	grid := NewBoxGridFromFile(rawGrid)
	for _, move := range rawInstructions {
		grid.move(move)
	}
	return grid.GPSCoordinates()
}

func Part2(input string) int {
	splitted := strings.Split(input, "\n\n")
	rawGrid, rawInstructions := splitted[0]+"\n", splitted[1][:len(splitted[1])-1]
	grid := NewBoxGridFromFile(rawGrid)
	grid = grid.WiderBoxGrid()
	for _, move := range rawInstructions {
		grid.moveWide(move)
	}
	return grid.GPSCoordinates()
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 15")
	input = input[:len(input)-1]
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
