package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Operator int

var wires map[string][]chan bool = make(map[string][]chan bool)

const (
	AND Operator = iota
	OR
	XOR
)

type Gate struct {
	gateType Operator
	left     <-chan bool
	right    <-chan bool
	output   string
}

func fireWire(wire string, value bool) {
	for _, w := range wires[wire] {
		w <- value
	}
}

func (g *Gate) run() {
	receivedLeft := false
	receivedRight := false
	var left, right bool
	for {
		if receivedLeft && receivedRight {
			receivedLeft = false
			receivedRight = false
			var output bool
			switch g.gateType {
			case AND:
				output = left && right
			case OR:
				output = left || right
			case XOR:
				output = left != right
			}
			fireWire(g.output, output)
		}
		select {
		case left = <-g.left:
			receivedLeft = true
		case right = <-g.right:
			receivedRight = true
		}
	}
}

func BuildGates(input string) []*Gate {
	split := strings.Split(input, "\n\n")
	gates := make([]*Gate, 0)
	wires = make(map[string][]chan bool)
	for _, line := range strings.Split(split[1], "\n") {
		split := strings.Split(line, " ")
		left, gateTypeRaw, right, output := split[0], split[1], split[2], split[4]
		if _, ok := wires[output]; !ok {
			wires[output] = append(wires[output], make(chan bool, 1))
			asInt, _ := strconv.Atoi(output[1:])
			if asInt > maxOutputWire {
				maxOutputWire = asInt
			}
		}
		var gateType Operator
		switch gateTypeRaw {
		case "AND":
			gateType = AND
		case "OR":
			gateType = OR
		case "XOR":
			gateType = XOR
		}
		leftChannel := make(chan bool, 1)
		rightChannel := make(chan bool, 1)
		wires[left] = append(wires[left], leftChannel)
		wires[right] = append(wires[right], rightChannel)
		gate := Gate{gateType, leftChannel, rightChannel, output}
		gates = append(gates, &gate)
		go gate.run()
	}
	return gates
}

var maxOutputWire int

func output() int {
	result := 0
	for i := range maxOutputWire + 1 {
		wire := fmt.Sprintf("z%02d", i)
		value := <-wires[wire][0]
		if value {
			result = 1<<i + result
		}
	}
	return result
}

func Part1(input string) int {
	split := strings.Split(input, "\n\n")
	BuildGates(input)

	for _, line := range strings.Split(split[0], "\n") {
		split := strings.Split(line, ": ")
		var value bool
		if split[1] == "1" {
			value = true
		}
		fireWire(split[0], value)

	}

	return output()
}

type GateBlueprint struct {
	left, right, operator, output string
}

func (g *GateBlueprint) hasInput(input string) bool {
	return g.left == input || g.right == input
}

func wireLinks(input string) []*GateBlueprint {
	split := strings.Split(input, "\n\n")
	gates := make([]*GateBlueprint, 0)
	for _, line := range strings.Split(split[1], "\n") {
		split := strings.Split(line, " ")
		left, gateType, right, output := split[0], split[1], split[2], split[4]
		gates = append(gates, &GateBlueprint{left, right, gateType, output})
	}
	return gates
}

func isInputOrOutput(wire string) bool {
	return wire[0] == 'x' || wire[0] == 'y' || wire[0] == 'z'
}

func Part2(input string) string {
	gates := wireLinks(input)
	swaps := make([]string, 0)

	biggestWire := fmt.Sprintf("z%02d", maxOutputWire)
	for _, gate := range gates {
		if gate.output[0] == 'z' && gate.operator != "XOR" && gate.output != biggestWire {
			swaps = append(swaps, gate.output)
		} else if gate.operator == "XOR" && !isInputOrOutput(gate.left) && !isInputOrOutput(gate.right) && !isInputOrOutput(gate.output) {
			swaps = append(swaps, gate.output)
		} else if gate.operator == "AND" && !gate.hasInput("x00") {
			for _, subGate := range gates {
				if subGate.hasInput(gate.output) && subGate.operator != "OR" {
					swaps = append(swaps, gate.output)
				}
			}
		} else if gate.operator == "XOR" {
			for _, subGate := range gates {
				if subGate.hasInput(gate.output) && subGate.operator == "OR" {
					swaps = append(swaps, gate.output)
				}
			}
		}
	}
	slices.Sort(swaps)
	return strings.Join(swaps, ",")
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 24")
	input = input[:len(input)-1]
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %s\n", Part2(input))
}
