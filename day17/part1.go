package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Register int

type Computer struct {
	A int
	B int
	C int

	ctr     int
	output  []int
	program []int
}

func registerInfo(inp string) int {
	split := strings.Split(inp, " ")
	result, _ := strconv.Atoi(split[2])
	return result
}

func NewComputerFromInput(input string) Computer {
	split := strings.Split(input, "\n")
	a := registerInfo(split[0])
	b := registerInfo(split[1])
	c := registerInfo(split[2])

	programRaw := strings.Split(split[4], " ")[1]
	programString := strings.Split(programRaw, ",")
	program := make([]int, len(programString))
	for i, s := range programString {
		program[i], _ = strconv.Atoi(s)
	}

	return Computer{
		A: a,
		B: b,
		C: c,

		output:  make([]int, 0),
		program: program,
	}
}

func (c *Computer) run() {
	for {
		if c.ctr < 0 || c.ctr >= len(c.program) {
			return
		}
		inst := c.program[c.ctr]
		op := c.program[c.ctr+1]
		switch inst {
		case 0:
			c.adv(op)
		case 1:
			c.bxl(op)
		case 2:
			c.bst(op)
		case 3:
			c.jnz(op)
		case 4:
			c.bxc(op)
		case 5:
			c.out(op)
		case 6:
			c.bdv(op)
		case 7:
			c.cdv(op)
		}
	}
}

func (c *Computer) increaseCtr() {
	c.ctr += 2
}

func (c *Computer) adv(combo int) {
	num := c.combo(combo)
	c.A = c.A / (int(math.Round(math.Pow(2, float64(num)))))
	c.increaseCtr()
}

func (c *Computer) bxl(literal int) {
	c.B = c.B ^ literal
	c.increaseCtr()
}

func (c *Computer) bst(combo int) {
	num := c.combo(combo)
	c.B = num % 8
	c.increaseCtr()
}

func (c *Computer) jnz(literal int) {
	if c.A == 0 {
		c.increaseCtr()
		return
	}
	c.ctr = literal
}

func (c *Computer) bxc(_ int) {
	c.B = c.B ^ c.C
	c.increaseCtr()
}

func (c *Computer) out(combo int) {
	num := c.combo(combo)
	c.output = append(c.output, num%8)
	c.increaseCtr()
}

func (c *Computer) bdv(combo int) {
	num := c.combo(combo)
	c.B = c.A / (int(math.Round(math.Pow(2, float64(num)))))
	c.increaseCtr()
}

func (c *Computer) cdv(combo int) {
	num := c.combo(combo)
	c.C = c.A / (int(math.Round(math.Pow(2, float64(num)))))
	c.increaseCtr()
}

func (c *Computer) combo(num int) int {
	if num < 4 {
		return num
	}
	if num == 4 {
		return c.A
	}
	if num == 5 {
		return c.B
	}
	if num == 6 {
		return c.C
	}
	panic(fmt.Sprintf("Unknown combo %d", num))
}

func Part1(input string) string {
	computer := NewComputerFromInput(input)
	computer.run()
	result := make([]string, len(computer.output))
	for i, res := range computer.output {
		result[i] = strconv.Itoa(res)
	}
	return strings.Join(result, ",")
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 17")
	input = input[:len(input)-1]
	fmt.Printf("\tPart 1: %s\n", Part1(input))
}
