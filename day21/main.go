package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/core"
)

var Seven core.Point = core.NewPoint(0, 0)
var Eight core.Point = core.NewPoint(1, 0)
var Nine core.Point = core.NewPoint(2, 0)
var Four core.Point = core.NewPoint(0, 1)
var Five core.Point = core.NewPoint(1, 1)
var Six core.Point = core.NewPoint(2, 1)
var One core.Point = core.NewPoint(0, 2)
var Two core.Point = core.NewPoint(1, 2)
var Three core.Point = core.NewPoint(2, 2)
var ANum core.Point = core.NewPoint(2, 3)

var ADir core.Point = core.NewPoint(2, 0)
var Left core.Point = core.NewPoint(0, 1)
var Down core.Point = core.NewPoint(1, 1)
var Right core.Point = core.NewPoint(2, 1)
var Up core.Point = core.NewPoint(1, 0)

var EmptyNumeric core.Point = core.NewPoint(0, 3)
var EmptyDirectional core.Point = core.NewPoint(0, 0)

func numToPoint(num rune) core.Point {
	switch num {
	case '7':
		return core.NewPoint(0, 0)
	case '8':
		return core.NewPoint(1, 0)
	case '9':
		return core.NewPoint(2, 0)
	case '4':
		return core.NewPoint(0, 1)
	case '5':
		return core.NewPoint(1, 1)
	case '6':
		return core.NewPoint(2, 1)
	case '1':
		return core.NewPoint(0, 2)
	case '2':
		return core.NewPoint(1, 2)
	case '3':
		return core.NewPoint(2, 2)
	case '0':
		return core.NewPoint(1, 3)
	case 'A':
		return core.NewPoint(2, 3)
	default:
		panic("Unimplemented")
	}
}

func dirToPoint(dir rune) core.Point {
	switch dir {
	case '^':
		return core.NewPoint(1, 0)
	case 'A':
		return core.NewPoint(2, 0)
	case '<':
		return core.NewPoint(0, 1)
	case 'v':
		return core.NewPoint(1, 1)
	case '>':
		return core.NewPoint(2, 1)
	default:
		panic("Unimplemented")
	}
}

func pointToDir(point core.Point) rune {
	switch point {
	case core.NORTH:
		return '^'
	case core.SOUTH:
		return 'v'
	case core.WEST:
		return '<'
	case core.EAST:
		return '>'
	default:
		panic("Unimplemented")
	}
}

func pathFind(start, end core.Point, isNumPad bool) []rune {
	var ver, hor core.Point
	diff := end.Subtract(start)
	if diff.Y < 0 {
		ver = core.NORTH
	} else if diff.Y > 0 {
		ver = core.SOUTH
	}
	if diff.X < 0 {
		hor = core.WEST
	} else if diff.X > 0 {
		hor = core.EAST
	}
	var isOnGap bool
	if isNumPad {
		isOnGap = (start.Y == 3 && end.X == 0) || (start.X == 0 && end.Y == 3)
	} else {
		isOnGap = (start.Y == 0 && end.X == 0) || (start.X == 0 && end.Y == 0)
	}
	isGoingLeft := end.X < start.X
	movesV := make([]rune, 0)
	movesH := make([]rune, 0)
	for range int(math.Abs(float64(diff.Y))) {
		movesV = append(movesV, pointToDir(ver))
	}
	for range int(math.Abs(float64(diff.X))) {
		movesH = append(movesH, pointToDir(hor))
	}
	var path []rune
	if isOnGap != isGoingLeft {
		movesV, movesH = movesH, movesV
	}
	path = append(append([]rune{}, movesV...), movesH...)
	path = append(path, 'A')
	return path
}

func shortestPathNumeric(code string) []rune {
	previous := ANum
	result := make([]rune, 0)
	for _, char := range code {
		position := numToPoint(char)
		result = append(result, pathFind(previous, position, true)...)
		previous = position
	}
	return result
}

type cacheKey struct {
	path  string
	level int
}

func shortestPathDirectional(args cacheKey, inner func(cacheKey) int) int {
	level, path := args.level, args.path
	if level == 0 {
		return len(path)
	}
	previous := ADir
	result := 0
	for _, current := range path {
		directional := dirToPoint(current)
		shortestPath := pathFind(previous, directional, false)
		result += inner(cacheKey{string(shortestPath), level - 1})
		previous = directional
	}
	return result
}

var cachedShortestPathDirectional = core.DP(shortestPathDirectional)

func shortestPath(code string, directionalPads int) int {
	topLevelDirPath := shortestPathNumeric(code)
	shortest := cachedShortestPathDirectional(cacheKey{string(topLevelDirPath), directionalPads})
	return shortest
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	result := 0
	for _, code := range lines {
		minimalLength := shortestPath(code, 2)
		number, _ := strconv.Atoi(code[:len(code)-1])
		result += minimalLength * number
	}
	return result
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	result := 0
	for _, code := range lines {
		minimalLength := shortestPath(code, 25)
		number, _ := strconv.Atoi(code[:len(code)-1])
		result += minimalLength * number
	}
	return result
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 21")
	input = input[:len(input)-1]
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
