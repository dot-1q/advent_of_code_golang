package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")

	containers := []int{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		capacity, _ := strconv.Atoi(line)

		containers = append(containers, capacity)
	}

	partOne(containers)
	partTwo(containers)
}

func partOne(containers []int) {

	// For every length, calculate how many combinations hold exactly 150L
	number := 0
	for length := range len(containers) {
		combinations := generateCombinations(containers, length)
		n, _ := calculateCapactity(combinations)
		number += n
	}
	fmt.Println("Part1: ", number)

}
func partTwo(containers []int) {
	// For every length, get the combinations that hold 150L, and filter for minimum values
	minimum := math.MaxInt32
	for length := range len(containers) {
		combinations := generateCombinations(containers, length)
		_, combs := calculateCapactity(combinations)

		for _, c := range combs {
			minimum = min(minimum, len(c))
		}
	}

	// Now, after finding the minimum number of containers that hold 150L, how many different ways are there
	// of that length, that hold 150L
	combinations := generateCombinations(containers, minimum)
	_, combs := calculateCapactity(combinations)

	// fmt.Println(combs)
	fmt.Println("Part2: ", len(combs))
}

// Calculate Capacity of the combinations, and return those combinations which hold 150L
func calculateCapactity(combs [][]int) (int, [][]int) {
	holds := 0
	combinations := [][]int{}
	for _, comb := range combs {
		sum := 0
		for _, capacity := range comb {
			sum += capacity
		}
		// The sum of the combinations of the various sizes has to be exactly 150L
		if sum == 150 {
			holds++
			combinations = append(combinations, comb)
		}

	}
	return holds, combinations
}

// Generate all combinations of a given length
func generateCombinations(numbers []int, length int) [][]int {
	if length == 0 {
		return [][]int{{}}
	}

	combs := [][]int{}
	for i, n := range numbers {
		// Create an empty list with this first element
		l := append([]int{}, n)

		// Generate the lists taking out the first element, and with -1 length
		lists := generateCombinations(numbers[i+1:], length-1)

		// For all the different lists of length -1, append the elements, to the original
		// list 'l', which only had the first element.
		// After that, append those lists to the list of all combinations
		for _, list := range lists {

			combs = append(combs, append(l, list...))
		}
	}
	return combs
}
