package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	clean    = iota // 0
	infected        // 1
	weakened        // 2
	flagged         // 3
)

func main() {
	// The grid will be hold in a map, so we get 0(1) checking.
	// Since the grid is infinite, we only keep track of our current XY position, and the visited positions.
	grid1 := createGrid()
	grid2 := createGrid()
	// North, West, South and  East, (x,y) in Cartesian map.
	// Circular array for the movements. Every time we turn, either left or right
	// the facing direction gets shifted in the circular buffer, so we always know where we're facing
	var directions = [4][2]int{
		{-1, 0}, // Down
		{0, 1},  // Right
		{1, 0},  // Up
		{0, -1}, // Left
	}
	partOne(grid1, directions)
	partTwo(grid2, directions)
}

func partOne(grid map[[2]int]int, directions [4][2]int) {
	// Starting positions.
	x := 12
	y := 12
	bursts := 10000
	facing := 0
	infections := 0
	for range bursts {
		value := grid[[2]int{y, x}]
		switch value {
		case infected:
			facing = (facing + 1) % 4
			// Infect the node
			grid[[2]int{y, x}] = clean
			// Update position
			y += directions[facing][0]
			x += directions[facing][1]
		case clean:
			facing = (facing + 3) % 4
			grid[[2]int{y, x}] = infected
			// Update position
			y += directions[facing][0]
			x += directions[facing][1]
			infections++
		}
	}
	fmt.Printf("Part 1 : After %d bursts, there have been %d Infections\n", bursts, infections)
}

func partTwo(grid map[[2]int]int, directions [4][2]int) {
	// Starting positions.
	x := 12
	y := 12
	bursts := 10000000
	facing := 0
	infections := 0
	for range bursts {
		value := grid[[2]int{y, x}]
		switch value {
		case infected:
			facing = (facing + 1) % 4
			// Infect the node
			grid[[2]int{y, x}] = flagged
			// Update position
			y += directions[facing][0]
			x += directions[facing][1]
		case clean:
			facing = (facing + 3) % 4
			grid[[2]int{y, x}] = weakened
			// Update position
			y += directions[facing][0]
			x += directions[facing][1]
		case weakened:
			// Infect the node
			grid[[2]int{y, x}] = infected
			// Update position
			y += directions[facing][0]
			x += directions[facing][1]
			infections++
		case flagged:
			facing = (facing + 2) % 4
			// Infect the node
			grid[[2]int{y, x}] = clean
			// Update position
			y += directions[facing][0]
			x += directions[facing][1]
		}
	}
	fmt.Printf("Part 2 : After %d bursts, there have been %d Infections", bursts, infections)
}

func createGrid() map[[2]int]int {
	g := map[[2]int]int{}
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for cell, char := range line {
			pos := [2]int{row, cell}
			if char == '#' {
				g[pos] = infected
			} else {
				g[pos] = clean
			}
		}
		row++
	}
	return g
}
