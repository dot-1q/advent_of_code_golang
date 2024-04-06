package main

import (
	"fmt"
	q "github.com/emirpasic/gods/queues/linkedlistqueue"
	"strconv"
	"strings"
)

func main() {
	puzzle := 1358
	grid := createGrid(puzzle)
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%s", cell)
		}
		fmt.Println()
	}

	cost, visited := findShortestPath(grid, [2]int{1, 1}, [2]int{31, 39})
	fmt.Println("Part1 Cost:", cost)
	fmt.Println("Part2 Visited in less than 50:", len(visited))
}

func createGrid(number int) [50][50]string {
	grid := [50][50]string{}
	for y, row := range grid {
		for x := range row {
			if isOpenSpace(x, y, number) {
				grid[y][x] = "."
			} else {
				grid[y][x] = "#"
			}
		}
	}
	return grid
}

func isOpenSpace(x, y, number int) bool {
	calc := x*x + 3*x + 2*x*y + y + y*y
	calc += number
	binary := strconv.FormatInt(int64(calc), 2)
	onebits := strings.Count(binary, "1")
	// if the number of 1 bits is even, its an openspace
	if onebits%2 == 0 {
		return true
	} else {
		// Its a wall
		return false
	}
}

func findShortestPath(grid [50][50]string, start, end [2]int) (int, [][2]int) {
	queue := q.New()
	// The seen array will be: x,y, and will hold the cost to visit that coordinate
	seen := map[[2]int]int{}
	visited := [][2]int{}
	queue.Enqueue(start)
	// Mark start position and seen
	seen[[2]int{start[0], start[1]}] = 0
	// This is for part2, how many locations we reach with less that 50 moves
	visited = append(visited, start)

	for !queue.Empty() {
		p, _ := queue.Dequeue()
		point := p.([2]int)
		if point[0] == end[0] && point[1] == end[1] {
			// Return the cost calculated
			return seen[[2]int{point[0], point[1]}], visited
		}

		n := neighbours(grid, point)
		for _, neighbour := range n {
			// If it has not been visited
			if _, ok := seen[[2]int{neighbour[0], neighbour[1]}]; !ok {
				cost := seen[[2]int{point[0], point[1]}]
				seen[[2]int{neighbour[0], neighbour[1]}] = cost + 1
				if cost+1 <= 50 {
					visited = append(visited, neighbour)
				}
				queue.Enqueue(neighbour)
			}
		}
	}
	return 0, visited
}

func neighbours(grid [50][50]string, current [2]int) [][2]int {
	dirs := [][]int{
		{1, 0},  // Go right
		{-1, 0}, // Go left
		{0, 1},  // Go Up
		{0, -1}} // Go Down
	neigh := [][2]int{}
	for _, dir := range dirs {
		// Don't go out of bounds
		if current[0]+dir[0] >= 0 && current[1]+dir[1] >= 0 {
			if grid[current[1]+dir[1]][current[0]+dir[0]] != "#" {
				neigh = append(neigh, [2]int{current[0] + dir[0], current[1] + dir[1]})
			}

		}
	}
	return neigh
}
