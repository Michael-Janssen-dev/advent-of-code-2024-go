package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func preprocess(input string) (map[int]map[int]struct{}, [][]int) {
	splitted := strings.Split(input, "\n\n")
	ordering, updates_raw := splitted[0], splitted[1]
	after := make(map[int]map[int]struct{})
	for _, line := range strings.Split(ordering, "\n") {
		split := strings.Split(line, "|")
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])
		if _, ok := after[left]; !ok {
			after[left] = make(map[int]struct{})
		}

		after[left][right] = struct{}{}
	}

	updates := make([][]int, 0)
	for _, line := range strings.Split(updates_raw, "\n") {
		if line == "" {
			continue
		}
		splitted := strings.Split(line, ",")
		update := make([]int, len(splitted))
		for i, num := range splitted {
			update[i], _ = strconv.Atoi(num)
		}
		updates = append(updates, update)
	}
	return after, updates
}

func Part1(input string) int {
	sum := 0
	after, updates := preprocess(input)
	for _, update := range updates {
		u := NewUpdateSorter(update, after)
		if !sort.IsSorted(u) {
			continue
		}

		sum += u.Middle()
	}
	return sum
}

type UpdateSorter struct {
	pages []int
	order map[int]map[int]struct{}
}

func NewUpdateSorter(pages []int, order map[int]map[int]struct{}) UpdateSorter {
	return UpdateSorter{
		pages,
		order,
	}
}

func (b UpdateSorter) Len() int {
	return len(b.pages)
}

func (b UpdateSorter) Less(i, j int) bool {
	_, ok := b.order[b.pages[i]][b.pages[j]]
	return ok
}

func (b UpdateSorter) Swap(i, j int) {
	b.pages[i], b.pages[j] = b.pages[j], b.pages[i]
}

func (b *UpdateSorter) Middle() int {
	return b.pages[len(b.pages)/2]
}

func Part2(input string) int {
	after, updates := preprocess(input)

	sum := 0

	for _, update := range updates {
		u := NewUpdateSorter(update, after)

		if sort.IsSorted(u) {
			continue
		}
		sort.Sort(u)
		sum += u.Middle()
	}

	return sum
}

func main() {
	fmt.Println("Day 5")
	file, err := os.ReadFile("input/inp.txt")
	if err != nil {
		log.Fatalln("Could not find file, exiting...")
	}
	input := string(file)
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %d\n", Part2(input))
}
