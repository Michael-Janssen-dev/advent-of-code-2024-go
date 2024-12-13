package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Machine struct {
	aX, aY, bX, bY, pX, pY int
}

func CloseEnough(num float64) bool {
	diff := math.Abs(num - float64(int(math.Round(num))))
	fmt.Println(diff)
	return diff < 0.0001 || diff > 0.9999
}

func (m *Machine) Optimal2() int {
	div := float64(m.aX) / float64(m.aY)
	top := float64(m.pX) - float64(m.pY)*div
	bottom := float64(m.bX) - float64(m.bY)*div
	b := top / bottom
	a := (float64(m.pX) - float64(m.bX)*b) / float64(m.aX)
	if CloseEnough(a) && CloseEnough(b) {
		return int(math.Round(a*3 + b))
	}
	return 0
}

func Part1(input string) int {
	objects := strings.Split(input[:len(input)-1], "\n\n")
	machines := make([]Machine, len(objects))
	for i, object := range objects {
		lines := strings.Split(object, "\n")
		aX, _ := strconv.Atoi(lines[0][strings.Index(lines[0], "+")+1 : strings.Index(lines[0], ",")])
		aY, _ := strconv.Atoi(lines[0][strings.LastIndex(lines[0], "+")+1:])
		bX, _ := strconv.Atoi(lines[1][strings.Index(lines[1], "+")+1 : strings.Index(lines[1], ",")])
		bY, _ := strconv.Atoi(lines[1][strings.LastIndex(lines[1], "+")+1:])
		pX, _ := strconv.Atoi(lines[2][strings.Index(lines[2], "=")+1 : strings.Index(lines[2], ",")])
		pY, _ := strconv.Atoi(lines[2][strings.LastIndex(lines[2], "=")+1:])
		machines[i] = Machine{aX, aY, bX, bY, pX, pY}
	}

	sum := 0
	for _, machine := range machines {
		sum += machine.Optimal2()
	}

	return sum
}

func Part2(input string) int {
	objects := strings.Split(input[:len(input)-1], "\n\n")
	machines := make([]Machine, len(objects))
	for i, object := range objects {
		lines := strings.Split(object, "\n")
		aX, _ := strconv.Atoi(lines[0][strings.Index(lines[0], "+")+1 : strings.Index(lines[0], ",")])
		aY, _ := strconv.Atoi(lines[0][strings.LastIndex(lines[0], "+")+1:])
		bX, _ := strconv.Atoi(lines[1][strings.Index(lines[1], "+")+1 : strings.Index(lines[1], ",")])
		bY, _ := strconv.Atoi(lines[1][strings.LastIndex(lines[1], "+")+1:])
		pX, _ := strconv.Atoi(lines[2][strings.Index(lines[2], "=")+1 : strings.Index(lines[2], ",")])
		pY, _ := strconv.Atoi(lines[2][strings.LastIndex(lines[2], "=")+1:])
		machines[i] = Machine{aX, aY, bX, bY, pX + 10000000000000, pY + 10000000000000}
	}

	sum := 0
	for _, machine := range machines {
		fmt.Println(machine)
		sum += machine.Optimal2()
	}

	return sum
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 13")
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
