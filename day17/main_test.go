package main

import (
	"os"
	"testing"
)

const PART_1 = "4,6,3,5,6,3,5,2,1,0"

func TestDay1(t *testing.T) {
	file, err := os.ReadFile("input/test.txt")
	if err != nil {
		t.Fatalf("Could not find test file")
	}
	input := string(file)
	input = input[:len(input)-1]

	result := Part1(input)
	if result != PART_1 {
		t.Fatalf("Day 17 Part 1: expected %s, got %s", PART_1, result)
	}
}
