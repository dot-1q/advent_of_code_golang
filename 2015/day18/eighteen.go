package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	lights := [100][100]bool{}
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)

		for i, rune := range runes {
			// Turn on or off the array of lights
			switch rune {
			case '#':
				lights[row][i] = true
			case '.':
				lights[row][i] = false
			}
		}
		row++
	}
	partOne(&lights)
	partTwo(&lights)
}

func partOne(lights *[100][100]bool) {

	state := lights
	// Consecutively generate the new light states
	for range 100 {
		newState := calculateLightState(state, 1)
		state = &newState
	}

	fmt.Println("Part1: ", howManyOn(state))
}

func partTwo(lights *[100][100]bool) {
	state := lights
	// Consecutively generate the new light states
	for range 100 {
		newState := calculateLightState(state, 2)
		state = &newState
	}

	fmt.Println("Part2: ", howManyOn(state))

}

// Calculate the *NEW* light state given a current light state
func calculateLightState(lights *[100][100]bool, part int) [100][100]bool {
	newState := [100][100]bool{}
	for row, lightRow := range lights {
		for cell, light := range lightRow {
			adjacents := getNeighbours(row, cell, lights)
			newState[row][cell] = newLightState(adjacents, light)

			// For part 2, the corners are always ON
			if part == 2 {
				if row == 0 && cell == 0 {
					newState[row][cell] = true
				}
				if row == (len(lights)-1) && cell == 0 {
					newState[row][cell] = true
				}
				if row == 0 && cell == (len(lightRow)-1) {
					newState[row][cell] = true
				}
				if row == (len(lights)-1) && cell == (len(lightRow)-1) {
					newState[row][cell] = true
				}

			}
		}
	}
	return newState
}

// Calculate new light state given its neighbours and current state
func newLightState(neighbours []bool, state bool) bool {

	// If its ON
	if state == true {
		on := 0
		for _, neighbour := range neighbours {
			if neighbour == true {
				on++
			}
		}
		// Light stays on if 2 or 3 neighbours are On
		if on == 2 || on == 3 {
			return true
		}
	} else {
		on := 0
		for _, neighbour := range neighbours {
			if neighbour == true {
				on++
			}
		}
		// Light stays on if 3 neighbours are On
		if on == 3 {
			return true
		}
	}
	return false
}

// Get neighbours if a light
func getNeighbours(row, cell int, lights *[100][100]bool) []bool {
	adjacents := []bool{}
	// If the light is in the top row
	if row == 0 {
		// If the light is in the rightmost position
		if cell == 0 {
			// Right, bottom and diagonal-right neighbour
			adjacents = append(adjacents, (*lights)[row][cell+1])
			adjacents = append(adjacents, (*lights)[row+1][cell])
			adjacents = append(adjacents, (*lights)[row+1][cell+1])

		} else if cell == (len(lights[row]) - 1) {
			// If the light is in the leftmost position
			// Left, bottom and diagonal-left neighbour
			adjacents = append(adjacents, (*lights)[row][cell-1])
			adjacents = append(adjacents, (*lights)[row+1][cell])
			adjacents = append(adjacents, (*lights)[row+1][cell-1])

		} else {
			// Else, this light is in the middle of the row, and has 5 neighbours
			adjacents = append(adjacents, (*lights)[row][cell-1])
			adjacents = append(adjacents, (*lights)[row][cell+1])
			adjacents = append(adjacents, (*lights)[row+1][cell-1])
			adjacents = append(adjacents, (*lights)[row+1][cell])
			adjacents = append(adjacents, (*lights)[row+1][cell+1])
		}
	} else if row == (len(lights) - 1) {
		// If the light is in the bottom row
		// If the light is in the leftmost position
		if cell == 0 {
			// Right, bottom and diagonal-right neighbour
			adjacents = append(adjacents, (*lights)[row-1][cell])
			adjacents = append(adjacents, (*lights)[row-1][cell+1])
			adjacents = append(adjacents, (*lights)[row][cell+1])

		} else if cell == (len(lights[row]) - 1) {
			// If the light is in the rightmost position
			// Left, bottom and diagonal-left neighbour
			adjacents = append(adjacents, (*lights)[row-1][cell])
			adjacents = append(adjacents, (*lights)[row-1][cell-1])
			adjacents = append(adjacents, (*lights)[row][cell-1])

		} else {
			// Else, this light is in the middle of the row, and has 5 neighbours
			adjacents = append(adjacents, (*lights)[row][cell-1])
			adjacents = append(adjacents, (*lights)[row][cell+1])
			adjacents = append(adjacents, (*lights)[row-1][cell-1])
			adjacents = append(adjacents, (*lights)[row-1][cell])
			adjacents = append(adjacents, (*lights)[row-1][cell+1])
		}
	} else if cell == 0 {
		// If the light is in the left most column, has 5 neighbours
		adjacents = append(adjacents, (*lights)[row-1][cell])
		adjacents = append(adjacents, (*lights)[row-1][cell+1])
		adjacents = append(adjacents, (*lights)[row][cell+1])
		adjacents = append(adjacents, (*lights)[row+1][cell])
		adjacents = append(adjacents, (*lights)[row+1][cell+1])

	} else if cell == (len(lights[row]) - 1) {
		// If the light is in the right most column, has 5 neighbours
		adjacents = append(adjacents, (*lights)[row-1][cell])
		adjacents = append(adjacents, (*lights)[row-1][cell-1])
		adjacents = append(adjacents, (*lights)[row][cell-1])
		adjacents = append(adjacents, (*lights)[row+1][cell])
		adjacents = append(adjacents, (*lights)[row+1][cell-1])

	} else {
		// Light is in the middle and has 8 neighbours
		adjacents = append(adjacents, (*lights)[row-1][cell])
		adjacents = append(adjacents, (*lights)[row-1][cell-1])
		adjacents = append(adjacents, (*lights)[row-1][cell+1])
		adjacents = append(adjacents, (*lights)[row][cell-1])
		adjacents = append(adjacents, (*lights)[row][cell+1])
		adjacents = append(adjacents, (*lights)[row+1][cell])
		adjacents = append(adjacents, (*lights)[row+1][cell-1])
		adjacents = append(adjacents, (*lights)[row+1][cell+1])
	}
	return adjacents
}

// Calculate how many lights are turned on
func howManyOn(lights *[100][100]bool) int {
	on := 0
	for _, lightRow := range lights {
		for _, light := range lightRow {
			if light == true {
				on++
			}
		}
	}
	return on
}
