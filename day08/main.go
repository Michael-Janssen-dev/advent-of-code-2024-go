package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/michael-janssen-dev/advent-of-code-2024-go/data-structures/set"
)

type Coor struct {
	x, y int
}

func (c Coor) diff(other Coor) Coor {
	return Coor{c.x - other.x, c.y - other.y}
}

func (c Coor) add(other Coor) Coor {
	return Coor{c.x + other.x, c.y + other.y}
}

func (c Coor) sub(other Coor) Coor {
	return Coor{c.x - other.x, c.y - other.y}
}

func (c Coor) inBounds(maxX, maxY int) bool {
	return c.x >= 0 && c.x < maxX && c.y >= 0 && c.y < maxY
}

func Part1(input string) int {
	antennas := make(map[rune][]Coor)
	lines := strings.Split(input[:len(input)-1], "\n")
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				_, ok := antennas[char]
				if !ok {
					antennas[char] = make([]Coor, 0)
				}
				antennas[char] = append(antennas[char], Coor{x, y})
			}
		}
	}

	maxX := len(lines[0])
	maxY := len(lines)

	places := set.NewSet[Coor]()

	for _, values := range antennas {
		for i, a1 := range values[:len(values)-1] {
			for _, a2 := range values[i+1:] {
				al, ar := a1, a2
				if al.y < ar.y || (al.y == ar.y && al.x < ar.x) {
					al, ar = ar, al
				}
				diff := ar.diff(al)

				p1 := al.sub(diff)
				if p1.inBounds(maxX, maxY) {
					places.Add(p1)
				}
				p2 := ar.add(diff)
				if p2.inBounds(maxX, maxY) {
					places.Add(p2)
				}
			}
		}
	}

	return places.Len()
}

func Part2(input string) int {
	antennas := make(map[rune][]Coor)
	lines := strings.Split(input[:len(input)-1], "\n")
	for y, line := range lines {
		for x, char := range line {
			if char != '.' {
				_, ok := antennas[char]
				if !ok {
					antennas[char] = make([]Coor, 0)
				}
				antennas[char] = append(antennas[char], Coor{x, y})
			}
		}
	}

	maxX := len(lines[0])
	maxY := len(lines)

	places := set.NewSet[Coor]()

	for _, values := range antennas {
		for i, a1 := range values[:len(values)-1] {
			for _, a2 := range values[i+1:] {
				al, ar := a1, a2
				if al.y < ar.y || (al.y == ar.y && al.x < ar.x) {
					al, ar = ar, al
				}
				diff := ar.diff(al)
				total := diff
				places.Add(al)
				places.Add(ar)
				for {
					changed := false
					p1 := al.sub(total)
					if p1.inBounds(maxX, maxY) {
						places.Add(p1)
						changed = true
					}
					p2 := ar.add(total)
					if p2.inBounds(maxX, maxY) {
						places.Add(p2)
						changed = true
					}
					if !changed {
						break
					}
					total = total.add(diff)
				}
			}
		}
	}

	return places.Len()
}

func main() {
	fmt.Println("Day 8")
	file, err := os.ReadFile("input/inp.txt")
	if err != nil {
		log.Fatalln("Could not find file, exiting...")
	}
	input := string(file)
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
