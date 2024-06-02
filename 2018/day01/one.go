package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	numbers := toInt(lines)
	fmt.Printf("Part 1: Resulting Frequency %d\n", count(numbers))
	fmt.Printf("Part 2: Repeated Frequency %d\n", findRepeated(numbers))
}

func count(values []int) int {
	s := 0
	for _, value := range values {
		s += value
	}
	return s
}

func toInt(values []string) []int {
	numbers := make([]int, len(values))
	for i, value := range values {
		n, _ := strconv.Atoi(value)
		numbers[i] = n
	}
	return numbers
}

func findRepeated(values []int) int {
	seen := map[int]bool{}
	i := 0
	s := 0
	for {
		s += values[i]
		// If seen already, return here
		if _, ok := seen[s]; ok {
			return s
		} else {
			seen[s] = true
		}
		i = (i + 1) % len(values)
	}
}
