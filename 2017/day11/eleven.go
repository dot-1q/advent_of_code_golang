package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	moves := strings.Split(strings.TrimSpace(string(f)), ",")

	// https://www.redblobgames.com/grids/hexagons/#coordinates
	// double-coordinates system
	// {x,y}
	dirs := map[string][]int{
		"n":  {0, -2},  // north
		"s":  {0, +2},  // south
		"ne": {+1, -1}, // northeast
		"nw": {-1, -1}, // northwest
		"se": {+1, +1}, // southeast
		"sw": {-1, +1}, // southwest
	}

	x := 0
	y := 0
	maxDistance := 0
	for _, move := range moves {
		x += dirs[move][0]
		y += dirs[move][1]
		// Part2
		dist := distance(x, y)
		if dist > maxDistance {
			maxDistance = dist
		}
	}
	fmt.Printf("Ending coords: X: %d, Y: %d\n", x, y)
	fmt.Printf("Part1 |  Distance: %d\n", distance(x, y))
	fmt.Printf("Part2 |  Max Distance ever : %d\n", maxDistance)
}

// https://www.redblobgames.com/grids/hexagons/#distances
// Doubleheight formula
func distance(x, y int) int {
	col := math.Abs(float64(x))
	row := math.Abs(float64(y))
	return int(col + math.Max(0.0, (row-col)/2))
}
