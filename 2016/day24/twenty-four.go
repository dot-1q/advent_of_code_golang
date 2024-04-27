package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type NodeCost struct {
	moves int
	x     int
	y     int
}

func main() {
	grid := createGrid()
	graph := map[int]map[int]int{}

	// for _, v := range grid {
	// 	fmt.Printf("%c\n", v)
	// }

	for y, row := range grid {
		for x, cell := range row {
			// Its a number
			if cell != '#' && cell != '.' {
				costs := breadthFirstSearch(grid, x, y)
				fmt.Printf("For number %c, the cost map is: %v\n", cell, costs)
				graph[int(cell-'0')] = costs
			}
		}
	}
	// From the cost map of each node pair, perform a dfs to know the minimal cost to visit every node
	part := 2
	cost := depthFirstSearch(graph, map[int]bool{}, 0, part)
	fmt.Println(cost)
}

// Perform a bfs from each number node to the other, and record the distance
func breadthFirstSearch(grid [][]rune, x, y int) map[int]int {
	costs := map[int]int{}
	// Add the first node
	n := int(grid[y][x] - '0')
	costs[n] = 0
	// Add the first node as visited
	visited := map[[2]int]bool{}
	queue := []NodeCost{
		{0, x, y},
	}
	dirs := [4][2]int{
		{0, -1},
		{0, 1},
		{1, 0},
		{-1, 0}}
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		// Check that we haven't visited this position
		if visited[[2]int{state.x, state.y}] {
			continue
		}
		visited[[2]int{state.x, state.y}] = true

		for _, dir := range dirs {
			// Its a valid position
			newX := state.x + dir[0]
			newY := state.y + dir[1]
			char := grid[newY][newX]
			// Check that we haven't visited this position
			if visited[[2]int{newX, newY}] {
				continue
			}
			if char != '#' {
				// It is also a number
				if char != '.' {
					n := int(char - '0')
					costs[n] = state.moves + 1
				}
				queue = append(queue, NodeCost{state.moves + 1, newX, newY})
			}
		}
	}
	return costs
}

func createGrid() [][]rune {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	// 0 means a not walkable position
	// 1 means a walkable position
	grid := [][]rune{}
	// Positions that will need to be visited
	y := 0
	for scanner.Scan() {
		grid = append(grid, []rune{})
		line := scanner.Text()
		for _, char := range line {
			switch char {
			case '#':
				grid[y] = append(grid[y], '#')
			case '.':
				grid[y] = append(grid[y], '.')
			default:
				// Its a number
				grid[y] = append(grid[y], char)
			}
		}
		y++
	}
	return grid
}

// Depth first search of the cities
func depthFirstSearch(graph map[int]map[int]int, visited map[int]bool, start int, part int) int {
	// Base case of the recursion call
	// Means we visited all cities
	if len(visited) == len(graph) {
		// For part 2, add the cost of this terminal node to the starting position (0)
		if part == 2 {
			return graph[start][0]
		} else {
			return 0
		}
	}

	min_cost := math.MaxInt32
	for node := range graph {
		if !visited[node] {
			visited[node] = true
			cost := graph[start][node]
			mincost := depthFirstSearch(graph, visited, node, part)
			min_cost = min(min_cost, mincost+cost)
			// After the recursion calls, remove the cities from the visited ones,
			// so we can restart the loop
			delete(visited, node)
		}
	}
	return min_cost
}
