package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func IsPossibleEquationP1(left int, sum int, rhs []int) bool {
	if len(rhs) == 0 {
		return left == sum
	}
	plus := IsPossibleEquationP1(left, sum+rhs[0], rhs[1:])
	mult := IsPossibleEquationP1(left, sum*rhs[0], rhs[1:])
	return plus || mult
}

func Part1(input string) int {
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines[:len(lines)-1] {
		equation := strings.Split(line, ": ")
		left, _ := strconv.Atoi(equation[0])
		right := strings.Split(equation[1], " ")
		rhs := make([]int, len(right))
		for i, r := range right {
			rhs[i], _ = strconv.Atoi(r)
		}
		if IsPossibleEquationP1(left, rhs[0], rhs[1:]) {
			sum += left
		}
	}

	return sum
}

func IsPossibleEquationP2(left int, sum int, rhs []int) bool {
	if len(rhs) == 0 {
		return left == sum
	}
	plus := IsPossibleEquationP2(left, sum+rhs[0], rhs[1:])
	mult := IsPossibleEquationP2(left, sum*rhs[0], rhs[1:])
	concat, _ := strconv.Atoi(strconv.Itoa(sum) + strconv.Itoa(rhs[0]))
	isConcat := IsPossibleEquationP2(left, concat, rhs[1:])
	return plus || mult || isConcat
}

func AsyncIsPossibleEquationP2(left, sum int, rhs []int, c chan int) {
	if IsPossibleEquationP2(left, sum, rhs) {
		c <- left
	} else {
		c <- 0
	}
}

func Part2(input string) int {
	lines := strings.Split(input, "\n")
	channel := make(chan int)
	for _, line := range lines[:len(lines)-1] {
		equation := strings.Split(line, ": ")
		left, _ := strconv.Atoi(equation[0])
		right := strings.Split(equation[1], " ")
		rhs := make([]int, len(right))
		for i, r := range right {
			rhs[i], _ = strconv.Atoi(r)
		}
		go AsyncIsPossibleEquationP2(left, rhs[0], rhs[1:], channel)
	}

	sum := 0
	for range len(lines) - 1 {
		sum += <-channel
	}

	return sum
}

func main() {
	fmt.Println("Day 7")
	file, err := os.ReadFile("input/inp.txt")
	if err != nil {
		log.Fatalln("Could not find file, exiting...")
	}
	input := string(file)
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
