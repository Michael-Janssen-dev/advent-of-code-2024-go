package main

import (
	"os"
	"testing"
)

const PART_1 = 7
const PART_2 = "co,de,ka,ta"

func TestDay1(t *testing.T) {
	file, err := os.ReadFile("input/test.txt")
	if err != nil {
		t.Fatalf("Could not find test file")
	}
	input := string(file)
	input = input[:len(input)-1]

	result := Part1(input)
	if result != PART_1 {
		t.Fatalf("Day 23 Part 1: expected %d, got %d", PART_1, result)
	}

	if PART_2 == "" {
		return
	}

	result2 := Part2(input)
	if result2 != PART_2 {
		t.Fatalf("Day 23 Part 2: expected %s, got %s", PART_2, result2)
	}
}
