package main

import "fmt"

func main() {
	SerialNumber := 5791
	grid := [300][300]int{}
	cellPower(&grid, SerialNumber)
	// partOne(grid)
	partTwo(grid)
}

func partOne(grid [300][300]int) {
	maxPower := 0
	coords := ""
	for row := range grid {
		for cell := range grid[row] {
			power := calculatePower(grid, row, cell) // Calculate the power level on a 3x3 square starting here
			if power > maxPower {
				maxPower = power
				coords = fmt.Sprintf("%d,%d", cell+1, row+1)
			}
		}
	}
	fmt.Printf("Part 1 | Max at %s with power %d\n", coords, maxPower)
}

func partTwo(grid [300][300]int) {
	maxPower := 0
	coords := ""
	size := 0
	for row := range grid {
		for cell := range grid[row] {
			maxSize := min(299-cell, 299-row) // Maximum size the square can have
			// fmt.Printf("%d,%d, max size %d\n", cell+1, row+1, maxSize)
			for s := range maxSize {
				power := calculatePowerWithSize(grid, row, cell, s) // Calculate the power level on a 3x3 square starting here
				if power > maxPower {
					maxPower = power
					coords = fmt.Sprintf("%d,%d", cell+1, row+1)
					size = s
				}
			}
		}
	}
	fmt.Printf("Part 2 | Max at %s with power %d and size %d\n", coords, maxPower, size)
}

func cellPower(grid *[300][300]int, SerialNumber int) {
	// Remember the coordinates have to be +1, because arrays start at 0.
	for row := range grid {
		for cell := range grid[row] {
			rackID := (cell + 1) + 10
			powerLevel := rackID * (row + 1)
			add := powerLevel + SerialNumber
			value := add * rackID
			hundreds := (value / 100) % 10
			result := hundreds - 5
			grid[row][cell] = result
		}
	}
}

func calculatePower(grid [300][300]int, row, cell int) int {
	s := 0
	// Bounds
	if (row+2 < 300) && (cell+2 < 300) {
		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				s += grid[row+y][cell+x]
			}
		}
	}
	return s
}

func calculatePowerWithSize(grid [300][300]int, row, cell, size int) int {
	s := 0
	// Bounds
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			s += grid[row+y][cell+x]
		}
	}
	return s
}
