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
			fmt.Printf("Seconds passed %d\n", second)
		}
	}
}

func dispersion(positions [][2]int) string {
	// get bounds
	left := math.MaxInt16
	right := -math.MaxInt16
	top := math.MaxInt16
	bottom := -math.MaxInt16

	coords := map[[2]int]bool{}
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
		if bottom-top > 65 {
			return ""
		}
	}

	str := ""
	for row := top; row <= bottom; row++ {
		for col := left; col <= right; col++ {
			if coords[[2]int{row, col}] {
				str += "0"
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
