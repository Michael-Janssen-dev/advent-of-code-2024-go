package main

import (
	"os"
	"testing"
)

const PART_1 = 22
const PART_2 = "6,1"

func TestDay1(t *testing.T) {
	file, err := os.ReadFile("input/test.txt")
	if err != nil {
		t.Fatalf("Could not find test file")
	}
	input := string(file)
	input = input[:len(input)-1]

	result := Part1(input, 12, 7)
	if result != PART_1 {
		t.Fatalf("Day 18 Part 1: expected %d, got %d", PART_1, result)
	}

	if PART_2 == "" {
		return
	}

	res := Part2(input, 12, 7)
	if res != PART_2 {
		t.Fatalf("Day 18 Part 2: expected %s, got %s", PART_2, res)
	}
}
