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

	fmt.Println(containers)

	c := generateCombs(containers, []int{}, 0)
	n := calculateCapactity(c)
	fmt.Println("Part1: ", n)
	// Find minimum number of containers that can hold the 150 litters
	n = calculateMinCapactity(c)
	// Find the different ways that those 4 containers appear
	d := differentWays(c, n)
	fmt.Println("Part2: ", d)

}

func generateCombs(numbers []int, combination []int, start int) [][]int {
	if start == len(numbers) {
		// (reached end of list after selecting/not selecting)
		return [][]int{append([]int{}, combination...)}
	} else {
		// (element at ndx not included)
		c := [][]int{}
		c = append(c, generateCombs(numbers, combination, start+1)...)
		// (include element at ndx)
		c = append(c, generateCombs(numbers, append(combination, numbers[start]), start+1)...)
		return c
	}
}

func calculateCapactity(combs [][]int) int {

	diff_combs := 0
	for _, comb := range combs {
		sum := 0
		for _, capacity := range comb {
			sum += capacity
		}
		// The sum of the combinations of the various sizes has to be exactly
		// 150
		if sum == 150 {
			diff_combs++
		}
		sum = 0
	}

	return diff_combs
}

func calculateMinCapactity(combs [][]int) int {

	min_containers := math.MaxInt32
	for _, comb := range combs {
		sum := 0
		for _, capacity := range comb {
			sum += capacity
		}
		// The sum of the combinations of the various sizes has to be exactly
		// 150
		if sum == 150 {
			// get the minimum amount of containers
			if len(comb) < min_containers {
				min_containers = len(comb)
			}
		}
		sum = 0
	}

	return min_containers
}

func differentWays(combs [][]int, ways int) int {
	diff_way := 0
	for _, comb := range combs {
		if len(comb) == ways {
			sum := 0
			for _, capacity := range comb {
				sum += capacity
			}
			// The sum of the combinations of the various sizes has to be exactly
			// 150
			if sum == 150 {
				// get the minimum amount of containers
				diff_way++
			}
			sum = 0

		}
	}
	return diff_way
}
