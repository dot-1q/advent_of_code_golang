package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

func main() {
	fmt.Println("Hi")
	f, _ := os.Open("input.txt")

	weights := []int{}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		capacity, _ := strconv.Atoi(line)

		weights = append(weights, capacity)
	}
	partOne(weights)
	partTwo(weights)

}

func partOne(weights []int) {
	// The possible groups are the groups where theres a combination of 2 or more numbers ex(11,9),
	// Whose sum is equal to the sum*2 of the rest of all possible numbers
	// Ex: Group 1;             Group 2; Group 3
	// 11,9                     10 8 2;  7 5 4 3 1
	// Only store group 1 though, no need for the rest, it doesnt matter how they are split, only matters their sum.
	possibleGroupsOne := [][]int{}

	// Only calculate for max length of 9 elements for the first group. more than that is overkill
	// This is an assumption for this particular puzzle, and probably cannto be generalized for other similar problems
	for i := 3; i < 9; i++ {
		combs := generateCombinations(weights, i)
		// Sum the rest of the weights that are not present in the combination, since
		// those weights are for group 1
		for _, combination := range combs {
			sumGroupOne := 0
			for _, weight := range combination {
				sumGroupOne += weight
			}
			sumGroupTwoThree := sumWithout(combination, weights)

			if 2*sumGroupOne == (sumGroupTwoThree) {
				possibleGroupsOne = append(possibleGroupsOne, combination)
			}
		}
	}

	// The possible groups are already sorted by length of group 1
	lowestLen := len(possibleGroupsOne[0])
	lowestQE := findLowestQE(possibleGroupsOne, lowestLen)
	fmt.Println("Part1: ", lowestQE)

}

func partTwo(weights []int) {
	// Same as PartOne, but its groups of 4 now, so the sum has to be 3*GroupOne
	possibleGroupsOne := [][]int{}
	for i := 3; i < 10; i++ {
		combs := generateCombinations(weights, i)
		// Sum the rest of the weights that are not present in the combination, since
		// those weights are for group 1
		for _, combination := range combs {
			sumGroupOne := 0
			for _, weight := range combination {
				sumGroupOne += weight
			}
			sumGroupTwoThree := sumWithout(combination, weights)

			if 3*sumGroupOne == (sumGroupTwoThree) {
				possibleGroupsOne = append(possibleGroupsOne, combination)
			}
		}
	}

	// The possible groups are already sorted by length of group 1
	lowestLen := len(possibleGroupsOne[0])
	lowestQE := findLowestQE(possibleGroupsOne, lowestLen)
	fmt.Println("Part2: ", lowestQE)

}

func findLowestQE(possibleGroups [][]int, length int) int {
	// Quantum Entanglament is given by the product of the weights
	minQE := math.MaxInt64
	for i := 0; i < len(possibleGroups); i++ {
		// Remember we prefer smaller groups, so the ones exceeding the minimum valid group length dont matter
		if len(possibleGroups[i]) == length {
			// Given its a product, start with value 1
			qe := 1
			for _, weight := range possibleGroups[i] {
				qe *= weight
			}
			minQE = min(minQE, qe)
		} else {
			break
		}
	}
	return minQE
}

func sumWithout(combination []int, weights []int) int {
	sum := 0

	for _, weight := range weights {
		// Check if this weight is already present in group 1
		if !slices.Contains(combination, weight) {
			sum += weight
		}
	}
	return sum
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
