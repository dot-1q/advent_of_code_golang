package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func main() {
	fmt.Println("Hi")
	area := createMap()
	fmt.Println(area)
	// Im not doing some convoluted game simulation for the sake of being convoluted.
	// It is not fun nor challenging and has too many edge cases that will just burn me out.
}

func createMap() [][]rune {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	area := [][]rune{}

	for scanner.Scan() {
		line := scanner.Text()
		row := []rune{}
		for _, char := range line {
			row = append(row, char)
		}
		area = append(area, row)
	}
	return area
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
