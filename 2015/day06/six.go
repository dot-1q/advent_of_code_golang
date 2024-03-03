package main

import (
	"bufio"
	"fmt"
	"os"
	c "strconv"
	s "strings"
)

func main() {
	partOne()
	partTwo()
}

// Check how many lights are turned on. Simply check of true values
func howManyLights(array [1000][1000]bool) int {
	num := 0
	for _, row := range array {
		for _, cell := range row {
			if cell == true {
				num++
			}
		}
	}

	return num
}

// Check the total brightness of the grid
func totalBrightness(array [1000][1000]int) int {
	num := 0
	for _, row := range array {
		for _, cell := range row {
			num += cell
		}
	}

	return num
}

func partOne() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	// a 1000x1000 grid of booleans, its easier to toggle them (!bool)
	// int could've been used, and simply assign 1 to lit and 0 to off
	var grid [1000][1000]bool
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		// Check for every command available. The last one could be just 'else', but i left it this way for readability
		if s.HasPrefix(line, "turn on") {
			// Grab the command and split it
			command := s.Split(line, " ")

			//Traverssing the 2D array based on the coords
			// Remember that 2D arrays are Array[y][x] == Array[row][col]
			// Command[2] = initial pos, Command[4] = final pos
			initial_x, _ := c.Atoi(s.Split(command[2], ",")[0])
			initial_y, _ := c.Atoi(s.Split(command[2], ",")[1])
			final_x, _ := c.Atoi(s.Split(command[4], ",")[0])
			final_y, _ := c.Atoi(s.Split(command[4], ",")[1])
			for y := initial_y; y <= final_y; y++ {
				for x := initial_x; x <= final_x; x++ {
					grid[y][x] = true
				}
			}

		} else if s.HasPrefix(line, "turn off") {
			command := s.Split(line, " ")
			initial_x, _ := c.Atoi(s.Split(command[2], ",")[0])
			initial_y, _ := c.Atoi(s.Split(command[2], ",")[1])
			final_x, _ := c.Atoi(s.Split(command[4], ",")[0])
			final_y, _ := c.Atoi(s.Split(command[4], ",")[1])
			for y := initial_y; y <= final_y; y++ {
				for x := initial_x; x <= final_x; x++ {
					grid[y][x] = false
				}
			}
		} else if s.HasPrefix(line, "toggle") {
			command := s.Split(line, " ")
			initial_x, _ := c.Atoi(s.Split(command[1], ",")[0])
			initial_y, _ := c.Atoi(s.Split(command[1], ",")[1])
			final_x, _ := c.Atoi(s.Split(command[3], ",")[0])
			final_y, _ := c.Atoi(s.Split(command[3], ",")[1])
			for y := initial_y; y <= final_y; y++ {
				for x := initial_x; x <= final_x; x++ {
					grid[y][x] = !grid[y][x]
				}
			}
		}
	}
	fmt.Printf("Part1: Turned %d lights\n", howManyLights(grid))
}

func partTwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	// a 1000x1000 grid of booleans, its easier to toggle them (!bool)
	// int could've been used, and simply assign 1 to lit and 0 to off
	var grid [1000][1000]int
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		// Check for every command available. The last one could be just 'else', but i left it this way for readability
		if s.HasPrefix(line, "turn on") {
			// Grab the command and split it
			command := s.Split(line, " ")

			//Traverssing the 2D array based on the coords
			// Remember that 2D arrays are Array[y][x] == Array[row][col]
			// Command[2] = initial pos, Command[4] = final pos
			initial_x, _ := c.Atoi(s.Split(command[2], ",")[0])
			initial_y, _ := c.Atoi(s.Split(command[2], ",")[1])
			final_x, _ := c.Atoi(s.Split(command[4], ",")[0])
			final_y, _ := c.Atoi(s.Split(command[4], ",")[1])
			for y := initial_y; y <= final_y; y++ {
				for x := initial_x; x <= final_x; x++ {
					grid[y][x] += 1
				}
			}

		} else if s.HasPrefix(line, "turn off") {
			command := s.Split(line, " ")
			initial_x, _ := c.Atoi(s.Split(command[2], ",")[0])
			initial_y, _ := c.Atoi(s.Split(command[2], ",")[1])
			final_x, _ := c.Atoi(s.Split(command[4], ",")[0])
			final_y, _ := c.Atoi(s.Split(command[4], ",")[1])
			for y := initial_y; y <= final_y; y++ {
				for x := initial_x; x <= final_x; x++ {
					// If its zero, do not decrease below 0
					// continue keyword exits this loop iteration
					if grid[y][x] == 0 {
						continue
					} else {
						grid[y][x] -= 1
					}
				}
			}
		} else if s.HasPrefix(line, "toggle") {
			command := s.Split(line, " ")
			initial_x, _ := c.Atoi(s.Split(command[1], ",")[0])
			initial_y, _ := c.Atoi(s.Split(command[1], ",")[1])
			final_x, _ := c.Atoi(s.Split(command[3], ",")[0])
			final_y, _ := c.Atoi(s.Split(command[3], ",")[1])
			for y := initial_y; y <= final_y; y++ {
				for x := initial_x; x <= final_x; x++ {
					grid[y][x] += 2
				}
			}
		}
	}
	fmt.Printf("Part2: Has total brightness %d \n", totalBrightness(grid))
}
