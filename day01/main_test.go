package main

import (
	"os"
	"testing"
)

const PART_1 = 11
const PART_2 = 31

func TestDay1(t *testing.T) {
	file, err := os.ReadFile("input/test.txt")
	if err != nil {
		t.Fatalf("Could not find test file")
	}
	input := string(file)

	result := Part1(input)
	if result != PART_1 {
		t.Fatalf("Day 1 Part 1: expected %d, got %d", PART_1, result)
	}

	result = Part2(input)
	if result != PART_2 {
		t.Fatalf("Day 1 Part 2: expected %d, got %d", PART_2, result)
	}
}
