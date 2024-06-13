package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Point struct {
	x int
	y int
}

type Cart struct {
	position  Point
	direction int
	turn      int
	crashed   bool
}

// Global directions
var directions = [][2]int{
	{1, 0},  // down
	{0, 1},  // right
	{-1, 0}, // up
	{0, -1}, // left
}

func main() {
	maze := createMaze()
	carts := getCarts(&maze)
	part := 2

	if part == 1 {
		for !tick(maze, carts) {
		}
		// Get the index of the cart that crashed
		idx := slices.IndexFunc(carts, func(cart *Cart) bool {
			return cart.crashed
		})
		fmt.Printf("Part 1 | Location of the first crash: (%d,%d)\n", carts[idx].position.x, carts[idx].position.y)
	} else {
		// We have to remove the crashing carts at each iteration.
		for cartsLeft(carts) > 1 { // Stop when there's only one cart left
			tick(maze, carts)
		}
		// Get the index of the cart not crashed
		idx := slices.IndexFunc(carts, func(cart *Cart) bool {
			return !cart.crashed
		})
		fmt.Printf("Part 2 | Location of the final cart: (%d,%d)\n", carts[idx].position.x, carts[idx].position.y)
	}
}

func tick(maze [][]rune, carts []*Cart) bool {
	// Sort the carts array from top to bottom and left to right. For that, calculate its distance to
	// the origin (0,0) i.e: top left.
	slices.SortFunc(carts, func(a, b *Cart) int {
		return ManhattanDistance(a.position, Point{0, 0}) - ManhattanDistance(b.position, Point{0, 0})
	})
	crashed := false
	for _, c := range carts {
		if !c.crashed {
			c.position.x += directions[c.direction][1]
			c.position.y += directions[c.direction][0]
			switch maze[c.position.y][c.position.x] {
			case '/': // Changing direction
				if c.direction%2 == 0 { // The direction was up/down
					c.direction = ((c.direction - 1) + 4) % 4 // the direction will be right or left now
				} else { // The direction was left/right
					c.direction = (c.direction + 1) % 4 // the direction will be up or down now
				}
			case '\\': // Changing direction
				if c.direction%2 == 0 { // The direction was up/down
					c.direction = (c.direction + 1) % 4 // the direction will be right or left now
				} else { // The direction was left/right
					c.direction = ((c.direction - 1) + 4) % 4 // the direction will be up or down now
				}
			case '+': // Make a turn
				switch c.turn {
				case 0: // Turn left
					if c.direction%2 == 0 {
						c.direction++
					} else {
						c.direction = (c.direction + 1) % 4
					}
					c.turn++
				case 1: // Keep straight
					c.turn++
				case 2: // Turn right
					c.turn = 0
					if c.direction%2 == 0 {
						// Golang gives a negative number when applying %, if the number is negative
						c.direction = ((c.direction - 1) + 4) % 4
					} else {
						c.direction--
					}
				}
			}
			// Check if the new position crashes and mark it as such
			for _, c2 := range carts {
				// Mindful to not check against already crashed carts
				if c != c2 && (!c2.crashed) {
					if (c.position.x == c2.position.x) && c.position.y == c2.position.y {
						c.crashed = true
						c2.crashed = true
						crashed = true
					}
				}
			}
		}
	}
	return crashed
}

func cartsLeft(carts []*Cart) int {
	s := 0
	for _, c := range carts {
		if !c.crashed {
			s++
		}
	}
	return s
}

// Return the coordinates of the carts.
func getCarts(maze *[][]rune) []*Cart {
	// NOTE: Instead of having the initial carts position on the maze (<>^v), replace
	// it with either '-|' or '/\', so that we don't have to account for that rune when calculating the moves.
	carts := []*Cart{}
	for row := range *maze {
		for cell := range (*maze)[row] {
			switch (*maze)[row][cell] {
			case '<':
				carts = append(carts, &Cart{Point{cell, row}, 3, 0, false})
				(*maze)[row][cell] = '-'
			case '>':
				carts = append(carts, &Cart{Point{cell, row}, 1, 0, false})
				(*maze)[row][cell] = '-'
			case 'v':
				carts = append(carts, &Cart{Point{cell, row}, 0, 0, false})
				(*maze)[row][cell] = '|'
			case '^':
				carts = append(carts, &Cart{Point{cell, row}, 2, 0, false})
				(*maze)[row][cell] = '|'
			}
		}
	}
	return carts
}

func createMaze() [][]rune {
	maze := [][]rune{}
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune{}
		for _, char := range line {
			row = append(row, char)
		}
		maze = append(maze, row)
	}
	return maze
}

func ManhattanDistance(p1, p2 Point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
