package main

import (
	_ "embed"
	"fmt"
	"strings"
)

func GetChar(matrix [][]byte, x, y int) byte {
	if !(x >= 0 && x < len(matrix[0]) && y >= 0 && y < len(matrix)) {
		return 0
	}
	return matrix[y][x]
}

func CheckWord(matrix [][]byte, x, y int, word []byte) int {
	l, r, u, d, ur, dr, ul, dl := 1, 1, 1, 1, 1, 1, 1, 1
	for i, c := range word {
		if GetChar(matrix, x-i, y) != c {
			l = 0
		}
		if GetChar(matrix, x+i, y) != c {
			r = 0
		}
		if GetChar(matrix, x, y-i) != c {
			u = 0
		}
		if GetChar(matrix, x, y+i) != c {
			d = 0
		}
		if GetChar(matrix, x+i, y+i) != c {
			dr = 0
		}
		if GetChar(matrix, x+i, y-i) != c {
			ur = 0
		}
		if GetChar(matrix, x-i, y+i) != c {
			dl = 0
		}
		if GetChar(matrix, x-i, y-i) != c {
			ul = 0
		}
	}
	return l + r + d + u + dr + ur + dl + ul
}

func Part1(input string) int {
	splitted := strings.Split(input, "\n")
	splitted = splitted[:len(splitted)-1]
	grid := make([][]byte, len(splitted))
	for i := range splitted {
		grid[i] = []byte(splitted[i])
	}
	word := []byte("XMAS")
	count := 0
	for i := range grid {
		for j := range grid[i] {
			count += CheckWord(grid, i, j, word)
		}
	}
	return count
}

func CheckMas(matrix [][]byte, x, y int) bool {
	if matrix[y][x] != 'A' {
		return false
	}
	ul := GetChar(matrix, x-1, y-1)
	ur := GetChar(matrix, x+1, y-1)
	dl := GetChar(matrix, x-1, y+1)
	dr := GetChar(matrix, x+1, y+1)
	if !((ul == 'M' && dr == 'S') || (ul == 'S' && dr == 'M')) {
		return false
	}
	if !((ur == 'M' && dl == 'S') || (ur == 'S' && dl == 'M')) {
		return false
	}
	return true
}

func Part2(input string) int {
	splitted := strings.Split(input, "\n")
	splitted = splitted[:len(splitted)-1]
	grid := make([][]byte, len(splitted))
	for i := range splitted {
		grid[i] = []byte(splitted[i])
	}
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if CheckMas(grid, i, j) {
				count += 1
			}
		}
	}
	return count
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 4")
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
