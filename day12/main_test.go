package main

import (
	"os"
	"testing"
)

const PART_1 = 1930
const PART_2 = 1206

func TestDay1(t *testing.T) {
	file, err := os.ReadFile("input/test.txt")
	if err != nil {
		t.Fatalf("Could not find test file")
	}
	input := string(file)
	input = input[:len(input)-1]

	result := Part1(input)
	if result != PART_1 {
		t.Fatalf("Day 12 Part 1: expected %d, got %d", PART_1, result)
	}

	if PART_2 == -1 {
		return
	}

	result = Part2(input)
	if result != PART_2 {
		t.Fatalf("Day 12 Part 2: expected %d, got %d", PART_2, result)
	}
}
