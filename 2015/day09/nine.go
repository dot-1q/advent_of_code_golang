package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	Vertices map[string]*Vertex
}

// Represents a Vertex on a graph that will have a list of edges
// pointing to other vertices
type Vertex struct {
	Name string
	// Represents an edge, meaning, a map between a vertex name and its weight
	Edges map[string]int
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	partOne(f)
}

func partOne(f *os.File) {
	scanner := bufio.NewScanner(f)
	graph := makeGraph(scanner)
	min_cost, max_cost := exhaustiveSearch(graph)
	fmt.Printf("Minimum cost: %d and Maximum cost: %d to visit all cities\n", min_cost, max_cost)

}

// Exhaustive search of each map connection, from start to finish
// Visiting all cities
func exhaustiveSearch(graph map[string]map[string]int) (int, int) {

	min_cost := math.MaxInt32
	max_cost := 0
	for city := range graph {
		// Calculate cost, starting form all the different cities
		mincost, maxcost := depthFirstSearch(graph, map[string]bool{city: true}, city)
		min_cost = min(min_cost, mincost)
		max_cost = max(max_cost, maxcost)
	}
	return min_cost, max_cost
}

// Depth first search of the cities
func depthFirstSearch(graph map[string]map[string]int, visited map[string]bool, start string) (int, int) {
	// Base case of the recursion call
	// Means we visited all cities
	if len(visited) == len(graph) {
		return 0, 0
	}

	min_cost := math.MaxInt32
	max_cost := 0
	for city := range graph {
		if !visited[city] {
			visited[city] = true
			cost := graph[start][city]
			mincost, maxcost := depthFirstSearch(graph, visited, city)
			min_cost = min(min_cost, mincost+cost)
			max_cost = max(max_cost, maxcost+cost)

			// After the recursion calls, remove the cities from the visited ones,
			// so we can restart the loop
			delete(visited, city)
		}
	}

	return min_cost, max_cost
}

// Create graph
func makeGraph(scanner *bufio.Scanner) map[string]map[string]int {
	graph := map[string]map[string]int{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		weight, _ := strconv.Atoi(strings.TrimSpace(line[1]))
		paths := strings.Split(strings.TrimSpace(line[0]), " ")

		// If there's a value already we dont create a new map altogether
		// Also, create a mapping thats valid for both directions
		if _, ok := graph[paths[0]]; ok {
			graph[paths[0]][paths[2]] = weight
		} else {
			graph[paths[0]] = map[string]int{paths[2]: weight}
		}
		if _, ok := graph[paths[2]]; ok {
			graph[paths[2]][paths[0]] = weight
		} else {
			graph[paths[2]] = map[string]int{paths[0]: weight}
		}

	}
	return graph
}
