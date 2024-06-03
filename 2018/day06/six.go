package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	id int
	x  int
	y  int
}

func main() {
	// Create a grid of 400x400
	// Input the points, and for each grid position find the closest points
	// After that, go around the edges of the grid (the sides), and the numbers that are along
	// the sides correspond to areas that are infinite. All the other areas are enclosed by others.
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	points := getPoints(lines)
	grid := [400][400]int{}
	inputPoints(&grid, points)
	fillDistances(&grid, points)
	inf := findInifite(&grid)
	area := findLargestNotInf(&grid, inf)
	fmt.Printf("Part 1 | Largest area not infinite: %d\n", area)
	region := findRegionLessThan(&grid, points)
	fmt.Printf("Part 2 | Largest region with distance <10000: %d\n", region)
}

// Find the region which is less than 10000 distance to the other points.
func findRegionLessThan(grid *[400][400]int, points []Point) int {
	for row, line := range grid {
		for cell := range line {
			distance := 0
			// Calculate the sum of the distance of this point to all the rest
			for _, point := range points {
				distance += ManhattanDistance(Point{0, cell, row}, point)
			}
			if distance < 10000 {
				// Set this position with the number 99, so as to be different.
				grid[row][cell] = 99
			}
		}
	}
	count := countOccur(grid, 99)
	return count
}

func findLargestNotInf(grid *[400][400]int, infinites map[int]bool) int {
	area := 0
	for _, row := range grid {
		for _, cell := range row {
			// If it is not in the infinite list, check its area
			if _, ok := infinites[cell]; !ok {
				a := countOccur(grid, cell)
				if a > area {
					area = a
				}
			}
		}
	}
	return area
}

// Find the infinite areas. Its the areas which touch the grid edges. Return all the numbers that touch.
func findInifite(grid *[400][400]int) map[int]bool {
	ids := map[int]bool{}

	// Top row
	for _, cell := range grid[0] {
		ids[cell] = true
	}
	// Bottom row
	for _, cell := range grid[len(grid)-1] {
		ids[cell] = true
	}
	// // Left side
	for row := range grid {
		ids[grid[row][0]] = true
	}
	// Right side
	for row := range grid {
		ids[grid[row][len(grid)-1]] = true
	}
	return ids
}

func fillDistances(grid *[400][400]int, points []Point) {
	for row, line := range grid {
		for cell := range line {
			// Find the closest point to this coordinate
			closest := findClosest(Point{0, cell, row}, points)
			grid[row][cell] = closest.id
		}
	}
}

func findClosest(point Point, points []Point) Point {
	minDist := math.MaxInt16
	closest := Point{}

	for _, p := range points {
		dist := ManhattanDistance(point, p)
		if dist < minDist {
			minDist = dist
			closest = p
		}
	}
	return closest
}

func inputPoints(grid *[400][400]int, points []Point) {
	for _, point := range points {
		grid[point.y][point.x] = point.id
	}
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

func countOccur(grid *[400][400]int, id int) int {
	s := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == id {
				s++
			}
		}
	}
	return s
}
func getPoints(lines []string) []Point {
	points := make([]Point, len(lines))

	for i, line := range lines {
		sep := strings.Split(line, ", ")
		x, _ := strconv.Atoi(sep[0])
		y, _ := strconv.Atoi(sep[1])
		points[i] = Point{i + 1, x, y}
	}
	return points
}
