package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

type Graph struct {
	keyToIdx map[string]int
	vertices []string
	adj      [][]int
	size     int
}

func NewGraph() *Graph {
	return &Graph{
		keyToIdx: make(map[string]int),
	}
}

func (g *Graph) adjacent(vertex int) []int {
	return g.adj[vertex]
}

func (g *Graph) name(vertex int) string {
	return g.vertices[vertex]
}

func (g *Graph) addVertex(vertex string) {
	if _, ok := g.keyToIdx[vertex]; !ok {
		g.keyToIdx[vertex] = g.size
		g.vertices = append(g.vertices, vertex)
		g.adj = append(g.adj, []int{})
		g.size++
	}
}

func (g *Graph) addEdge(left, right string) {
	g.addVertex(left)
	g.addVertex(right)
	g.adj[g.keyToIdx[left]] = append(g.adj[g.keyToIdx[left]], g.keyToIdx[right])
	g.adj[g.keyToIdx[right]] = append(g.adj[g.keyToIdx[right]], g.keyToIdx[left])
}

func (g *Graph) hasEdge(left, right int) bool {
	return slices.Contains(g.adj[left], right)
}

func Part1(input string) int {
	graph := NewGraph()
	for _, line := range strings.Split(input, "\n") {
		split := strings.Split(line, "-")
		left, right := split[0], split[1]
		graph.addEdge(left, right)
	}
	result := 0
	for node1 := range graph.vertices {
		for _, node2 := range graph.adjacent(node1) {
			for _, node3 := range graph.adjacent(node2) {
				if node1 == node3 {
					continue
				}
				if graph.hasEdge(node3, node1) {
					name1, name2, name3 := graph.name(node1), graph.name(node2), graph.name(node3)
					if name1[0] == 't' || name2[0] == 't' || name3[0] == 't' {
						result += 1
					}
				}
			}
		}
	}
	return result / 6
}

func isClique(graph *Graph, list []int) bool {
	for _, node := range list {
		for _, node2 := range list {
			if node != node2 && !graph.hasEdge(node, node2) {
				return false
			}
		}
	}
	return true
}

func maxClique(graph *Graph, current []int, idx int) []int {
	if idx == graph.size {
		maxPath := make([]int, len(current))
		copy(maxPath, current)
		return maxPath
	}
	max := len(current)
	maxPath := make([]int, len(current))
	copy(maxPath, current)
	for node := idx; node < graph.size; node++ {
		new := make([]int, len(current))
		copy(new, current)
		new = append(new, node)
		if !isClique(graph, new) {
			continue
		}
		result := maxClique(graph, new, node+1)
		if len(result) > max {
			maxPath = result
			max = len(result)
		}
	}
	return maxPath
}

func Part2(input string) string {
	graph := NewGraph()
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		split := strings.Split(line, "-")
		left, right := split[0], split[1]
		graph.addEdge(left, right)
	}
	result := maxClique(graph, []int{}, 0)
	names := make([]string, len(result))
	for i, idx := range result {
		names[i] = graph.name(idx)
	}
	slices.Sort(names)
	return strings.Join(names, ",")
}

//go:embed input/inp.txt
var input string

func main() {
	fmt.Println("Day 23")
	input = input[:len(input)-1]
	fmt.Printf("\tPart 1: %d\n", Part1(input))
	fmt.Printf("\tPart 2: %s\n", Part2(input))
}
