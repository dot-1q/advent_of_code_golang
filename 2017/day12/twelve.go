package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	graph := map[string][]string{}
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		graph[line[0]] = []string{}
		for _, n := range line[2:] {
			neigh, _ := strings.CutSuffix(n, ",")
			graph[line[0]] = append(graph[line[0]], neigh)
		}
	}
	// fmt.Println(graph)
	fmt.Println("Part1: ", findConnections(graph))
	fmt.Println("Part2: ", findGroups(graph))
}

func findConnections(graph map[string][]string) int {
	sum := 0
	for node := range graph {
		if dfs(graph, node, "0", []string{}) {
			sum++
		}
	}
	return sum
}

func findGroups(graph map[string][]string) int {
	groups := 0
	graphGroups := map[string]bool{}
	for node := range graph {
		// Connection between this node and all the others.
		if !graphGroups[node] {
			for endNode := range graph {
				// If there is a path, they are in the same group
				if dfs(graph, node, endNode, []string{}) && !graphGroups[endNode] {
					graphGroups[endNode] = true
				}
			}
			groups++
		}
	}
	return groups
}

func dfs(graph map[string][]string, current, end string, seen []string) bool {
	if slices.Contains(seen, current) {
		return false
	}
	if current == end {
		return true
	}

	seen = append(seen, current)
	// Neighbours
	for _, node := range graph[current] {
		if dfs(graph, node, end, seen) {
			return true
		}
	}
	return false
}
