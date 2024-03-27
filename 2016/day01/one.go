package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// We only go right or left
	// The arrays represent (x,y)
	pos := []int{0, 0}
	visitedPos := map[[2]int]bool{
		{0, 0}: true,
	}

	// North, West, South and  East, (x,y) in Cartesian map.
	// Circular array for the movements. Every time we turn, either left or right
	// the facing direction gets shifted in the circular buffer, so we always know where we're facing
	facing := 0
	var directions = [4][2]int{
		{1, 0},  // North
		{0, 1},  // West
		{-1, 0}, // South
		{0, -1}, // East
	}

	f, _ := os.ReadFile("input.txt")
	line := string(f)
	dirs := strings.Split(line, ",")

	partTwo := true
	for _, direction := range dirs {
		chars := []rune(strings.TrimSpace(direction))
		switch chars[0] {
		case 'R':
			steps, _ := strconv.Atoi(string(chars[1:]))
			// Update wherewere facing
			facing = (facing + 3) % 4
			for range steps {
				pos[0] += directions[facing][0]
				pos[1] += directions[facing][1]
				if visitedPos[[2]int{pos[0], pos[1]}] && partTwo == true {
					manhattanDistance := math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
					fmt.Println("Part2: ", manhattanDistance)
					return
				}
				visitedPos[[2]int{pos[0], pos[1]}] = true
			}
		case 'L':
			steps, _ := strconv.Atoi(string(chars[1:]))
			// Update where were facing
			facing = (facing + 1) % 4
			for range steps {
				pos[0] += directions[facing][0]
				pos[1] += directions[facing][1]
				if visitedPos[[2]int{pos[0], pos[1]}] && partTwo == true {
					manhattanDistance := math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
					fmt.Println("Part2: ", manhattanDistance)
					return
				}
				visitedPos[[2]int{pos[0], pos[1]}] = true
			}
		}
	}

	manhattanDistance := math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
	fmt.Println("Part1: ", manhattanDistance)
}
