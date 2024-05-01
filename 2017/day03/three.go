package main

import (
	"fmt"
	"math"
)

func main() {
	puzzle := 289326
	partOne(puzzle)
	fmt.Println("______________________________")
	partTwo(puzzle)
}

func partOne(puzzle int) {
	// For the first circle around number 1, we need 8 more numbers:
	// 5   4   3
	// 6   1   2
	// 7   8   9
	// The length of the square's side goes: 3-5-7-9-11.....
	// Which is side_length = 2 + prev_side_length
	side_length := 3
	directions := [][2]int{
		{1, 0},  // right
		{0, 1},  // north
		{-1, 0}, // west
		{0, -1}, // south
	}
	// I'll start where the hint left off.
	x := 0
	y := 0
	number := 1
	side_length = 3
	for number <= puzzle {
		// First step is going right
		x += directions[0][0]
		number++
		if number == puzzle {
			fmt.Println("Distance:", math.Abs(float64(x))+math.Abs(float64(y)))
		}

		// Then we go up the length of the square side -1
		for range side_length - 2 {
			y += directions[1][1]
			number++
			if number == puzzle {
				fmt.Println("Distance:", math.Abs(float64(x))+math.Abs(float64(y)))
			}
		}
		// Then we go left
		for range side_length - 1 {
			x += directions[2][0]
			number++
			if number == puzzle {
				fmt.Println("Distance:", math.Abs(float64(x))+math.Abs(float64(y)))
			}
		}
		// Then we go down
		for range side_length - 1 {
			y += directions[3][1]
			number++
			if number == puzzle {
				fmt.Println("Distance:", math.Abs(float64(x))+math.Abs(float64(y)))
			}
		}
		// Then we go right
		for range side_length - 1 {
			x += directions[0][0]
			number++
			if number == puzzle {
				fmt.Println("Distance:", math.Abs(float64(x))+math.Abs(float64(y)))
			}
		}
		side_length += 2
	}
	fmt.Printf("X: %d, Y:%d\n", x, y)
}

func partTwo(puzzle int) {
	side_length := 3
	// Changed the up and down directions for partTwo() for better visualization
	directions := [][2]int{
		{1, 0},  // right
		{0, -1}, // north
		{-1, 0}, // west
		{0, 1},  // south
	}
	// Create the grid. It'll be big enough to find the solution.
	grid := [1000][1000]int{}
	// I'll start in the middle of the grid
	x := 500
	y := 500
	// First position
	grid[y][x] = 1
	side_length = 3
	found := false
	for !found {
		// First step is going right
		x += directions[0][0]
		number := getNeighboursSum(grid, x, y)
		grid[y][x] = number
		if number > puzzle {
			found = true
		}

		// Then we go up the length of the square side -1
		for range side_length - 2 {
			if found {
				break
			}
			y += directions[1][1]
			number = getNeighboursSum(grid, x, y)
			grid[y][x] = number
			if number > puzzle {
				found = true
			}
		}
		// Then we go left
		for range side_length - 1 {
			if found {
				break
			}
			x += directions[2][0]
			number = getNeighboursSum(grid, x, y)
			grid[y][x] = number
			if number > puzzle {
				found = true
			}
		}
		// Then we go down
		for range side_length - 1 {
			if found {
				break
			}
			y += directions[3][1]
			number = getNeighboursSum(grid, x, y)
			grid[y][x] = number
			if number > puzzle {
				found = true
			}
		}
		// Then we go right
		for range side_length - 1 {
			if found {
				break
			}
			x += directions[0][0]
			number = getNeighboursSum(grid, x, y)
			grid[y][x] = number
			if number > puzzle {
				found = true
			}
		}
		side_length += 2
	}
	// Uncomment to debug small grids
	// for _, line := range grid {
	// 	fmt.Printf("%5d\n", line)
	// }
	fmt.Println("First bigger number is:", grid[y][x])
	fmt.Printf("X: %d, Y:%d\n", x, y)
}

func getNeighboursSum(grid [1000][1000]int, x, y int) int {
	sum := 0
	neighbours := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	for _, n := range neighbours {
		sum += grid[y+n[0]][x+n[1]]
	}
	return sum
}
