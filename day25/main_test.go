package main

import (
	"os"
	"testing"
)

const PART_1 = 3

func TestDay1(t *testing.T) {
	file, err := os.ReadFile("input/test.txt")
	if err != nil {
		t.Fatalf("Could not find test file")
	}
	input := string(file)
	input = input[:len(input)-1]

	result := Part1(input)
	if result != PART_1 {
		t.Fatalf("Day 25 Part 1: expected %d, got %d", PART_1, result)
	}
}
