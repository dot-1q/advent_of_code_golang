package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x    int
	y    int
	velx int
	vely int
}

func main() {
	points := createPoints()
	seconds := 100000

	for second := range seconds {
		positions := [][2]int{}
		for p := range points {
			x := points[p].x + (points[p].velx * second)
			y := points[p].y + (points[p].vely * second)
			positions = append(positions, [2]int{x, y})
		}
		if str := dispersion(positions); str != "" {
			fmt.Println(str)
			fmt.Printf("Part 2 | Seconds passed %d\n", second)
		}
	}
}

func dispersion(positions [][2]int) string {
	left := math.MaxInt16
	right := -math.MaxInt16
	top := math.MaxInt16
	bottom := -math.MaxInt16

	coords := map[[2]int]bool{}
	// Get max coords from each direction.
	for _, p := range positions {
		coords[p] = true
		if p[0] < top {
			top = p[0]
		}
		if p[0] > bottom {
			bottom = p[0]
		}
		if p[1] < left {
			left = p[1]
		}
		if p[1] > right {
			right = p[1]
		}
		// If the distance between the top and bttom is too big, it most likely is not a
		// grid with letters. This value was trial and error.
		if bottom-top > 65 {
			return ""
		}
	}

	// Assemble the coords into the letters
	str := ""
	for col := left; col <= right; col++ {
		for row := top; row <= bottom; row++ {
			if coords[[2]int{row, col}] {
				str += "#"
			} else {
				str += " "
			}
		}
		str += "\n"
	}
	return str
}

func createPoints() []Point {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	points := []Point{}

	for scanner.Scan() {
		// String parsing shenanigans. Dont mind it.
		before, after, _ := strings.Cut(scanner.Text(), "> ")
		position, _ := strings.CutPrefix(before, "position=<")
		velocity, _ := strings.CutPrefix(after, "velocity=<")
		pcoords := strings.Fields(strings.TrimSpace(position))
		vcoords := strings.Fields(strings.TrimSpace(strings.TrimSuffix(velocity, ">")))
		xcoord, _ := strconv.Atoi(strings.TrimSuffix(pcoords[0], ","))
		ycoord, _ := strconv.Atoi(pcoords[1])
		vx, _ := strconv.Atoi(strings.TrimSuffix(vcoords[0], ","))
		vy, _ := strconv.Atoi(vcoords[1])
		points = append(points, Point{xcoord, ycoord, vx, vy})
	}
	return points
}
