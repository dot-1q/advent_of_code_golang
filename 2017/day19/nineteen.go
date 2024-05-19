package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := createGrid()
	str, steps := walk(grid)
	fmt.Println("Part 1 | Letters: ", string(str))
	fmt.Println("Part 2 | Steps: ", steps)
}

func walk(grid [][]rune) ([]rune, int) {
	dirs := [4][2]int{ // (row,col)
		{1, 0},  // Down
		{0, -1}, // Left
		{-1, 0}, // Up
		{0, 1},  // Right
	}
	lastDir := 0
	// This is the starting position
	position := []int{0, 109}
	// Append for every rune found
	runes := []rune{}

	steps := 0
	end := false
	for !end {
		// Next position
		newR := position[0] + dirs[lastDir][0]
		newC := position[1] + dirs[lastDir][1]
		steps++

		switch grid[newR][newC] {
		case '|': // Go up or down
			// If the last direction is up/down, and we encounter this char,
			// means its an intersecting one, and we should skip to the next char
			if lastDir%2 == 1 {
				position[0] = newR + dirs[lastDir][0]
				position[1] = newC + dirs[lastDir][1]
			}
		case '-': // Go left or right
			// If the last direction is up/down, and we encounter this char,
			// means its an intersecting one, and we should skip to the next char
			if lastDir%2 == 0 {
				position[0] = newR + dirs[lastDir][0]
				position[1] = newC + dirs[lastDir][1]
			}
		case '+': // Change dir
			// It was going UP/DOWN, not must go either left or right
			if lastDir%2 == 0 {
				// Get the right and left characters.
				left := grid[newR][newC-1]
				if left != ' ' { // we must go left
					lastDir = 1
				} else { // we must go right
					lastDir = 3
				}
			} else { // It was going LEFT/RIGHT, not must go either up or down
				up := grid[newR-1][newC]
				if up != ' ' { // we must go up
					lastDir = 2
				} else { // we must go down
					lastDir = 0
				}
			}
		default: // Its a letter we have to append
			runes = append(runes, grid[newR][newC])
			r, c := position[0]+dirs[lastDir][0], position[1]+dirs[lastDir][1]
			// If the position after the letter, following our path is empty, its the end.
			if grid[r][c] == ' ' {
				end = true
			}
		}
		position[0] = newR
		position[1] = newC
	}
	return runes, steps
}

func createGrid() [][]rune {
	grid := [][]rune{}
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		row := []rune{}
		line := scanner.Text()
		for _, r := range line {
			row = append(row, r)
		}
		grid = append(grid, row)
	}
	return grid
}
