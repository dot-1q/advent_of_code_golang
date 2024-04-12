package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	r, _ := os.ReadFile("input.txt")
	ranges := strings.Split(string(r), "\n")
	// Discard the new line added from the previous split
	ranges = ranges[:len(ranges)-1]
	lowRange := []int{}
	highRange := []int{}
	part := 1

	for _, ranges := range ranges {
		line := strings.Split(ranges, "-")
		low, _ := strconv.Atoi(line[0])
		high, _ := strconv.Atoi(line[1])
		lowRange = append(lowRange, low)
		highRange = append(highRange, high)
	}
	slices.Sort(highRange)
	slices.Sort(lowRange)
	// Since the highest ip value is a 32 bit integer (4294967295) and our highRange array is sorted
	// The max value - highest range, gives us some numbers that are allowed, since they are not contemplated in the exclusion ranges.
	total := 4294967295 - highRange[len(highRange)-1]
	for i := range len(lowRange) {
		// If the (Highest number in the range +1) is less than the very next range of numbers
		// means that we found a gap between the ranges, and as such, the lowest allowed number
		if i+1 < len(lowRange) && highRange[i]+1 < lowRange[i+1] {
			// This calculation is for part 2. basically we just get the number of IPs from the
			// highest range to the next low range, since there is a gap here, as noticed by the
			// if statement above
			total += lowRange[i+1] - (highRange[i] + 1)
			if part == 1 {
				fmt.Printf("Part1: Lowest IP value not blocked is %d\n", highRange[i]+1)
				break
			}
		}
		// fmt.Printf("Lowest %d to %d\n", lowRange[i], highRange[i])
	}
	fmt.Printf("Part2: Total allowed: %d\n", total)
}
