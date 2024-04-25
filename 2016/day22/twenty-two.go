package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	x     int
	y     int
	size  int
	used  int
	avail int
}

func main() {
	l, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(l), "\n")
	// Discard first two lines and the last \n split
	lines = lines[2 : len(lines)-1]

	nodes := []Node{}
	grid := [28][38]Node{}
	for _, l := range lines {
		line := strings.Fields(l)
		coords := strings.Split(line[0], "-")
		x, _ := strconv.Atoi(coords[1][1:])
		y, _ := strconv.Atoi(coords[2][1:])
		s, _ := strings.CutSuffix(line[1], "T")
		size, _ := strconv.Atoi(s)
		u, _ := strings.CutSuffix(line[2], "T")
		used, _ := strconv.Atoi(u)
		a, _ := strings.CutSuffix(line[3], "T")
		avail, _ := strconv.Atoi(a)
		n := Node{x: x, y: y, size: size, used: used, avail: avail}
		nodes = append(nodes, n)
		grid[y][x] = n
	}
	v := validNodesPairs(nodes)
	fmt.Printf("Valid Node Pairs: %d\n", v)
	printGrid(grid)
	fmt.Println()
}

func validNodesPairs(nodes []Node) int {
	count := 0
	for _, n1 := range nodes {
		for _, n2 := range nodes[1:] {
			// If it fits and is not 0. n1 and n2 are never the same, but the loop is O(n2).
			if (n1.used <= n2.avail) && (n1.used != 0) {
				count++
			}
		}
	}
	return count
}

// Basically you have to look at the grid and solve manually
// XX Nodes are nodes to big to be moved around, so we have to go around them.
// Start at position E, which is an empty node, which means any node can be moved to it.
// Move it around until he passes the row of XX, and stops just left to the G node.
// Each move until that position counts at just one move.
// When the E is right next to the G, "[ ] [ ] [E] [G]", swap them, and begin to move the G node to the right.
// Every time the G node needs to move to the right, it costs 5 moves, because the E has to go around it, liberating space.
// So the calc. is: (Total moves to move E to the right of G) + (5* how many positions G needs to move right).
// Ans: 261
func printGrid(grid [28][38]Node) {
	for y, row := range grid {
		for x, cell := range row {
			if cell.used != 0 && cell.used <= 100 {
				if y == 0 && x == 37 {
					fmt.Printf("[G] ")
				} else {
					fmt.Printf("[ ] ")
				}
			} else if cell.used > 100 {
				fmt.Printf(" XX ")
			} else {
				fmt.Printf("[E] ")
			}
		}
		fmt.Println()
	}
}
