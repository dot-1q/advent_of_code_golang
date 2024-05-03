package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	numbers := strings.Fields(string(f))
	blocks := createBlocks(numbers)
	seen := map[[16]int]bool{}

	// First configuration
	seen[blocks] = true
	redistributions := 0
	for {
		redistribute(&blocks)
		redistributions++
		// If we've seen, break
		if _, ok := seen[blocks]; ok {
			break
		}
		seen[blocks] = true
	}
	fmt.Printf("Part1 | Redistributions made: %d\n", redistributions)

	loops := 0
	firstSeen := blocks
	for {
		redistribute(&blocks)
		loops++
		// If we've seen, break
		if blocks == firstSeen {
			break
		}
	}

	fmt.Printf("Part2 | Cycles done: %d\n", loops)
}

func redistribute(blocks *[16]int) {
	maximum, idx := findMax(blocks[:])

	blocks[idx] = 0
	for maximum > 0 {
		idx = (idx + 1) % len(blocks)
		blocks[idx]++
		maximum--

	}
}

func createBlocks(array []string) [16]int {
	blocks := [16]int{}
	for i, n := range array {
		blocks[i], _ = strconv.Atoi(n)
	}
	return blocks
}

func findMax(array []int) (int, int) {
	maximum := array[0]
	idx := 0
	for i := range array {
		if array[i] > maximum {
			maximum = array[i]
			idx = i
		}
	}
	return maximum, idx
}
