package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	part2 := true

	scanner := bufio.NewScanner(f)
	checksum := 0
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		values := stringToIntArr(row)

		if !part2 {
			maxV := slices.Max(values)
			minV := slices.Min(values)
			checksum += (maxV - minV)
		} else {
			rem := findTwoEvenDivs(values)
			checksum += rem
		}
	}

	fmt.Println("Checksum: ", checksum)
}

func stringToIntArr(numbers []string) []int {
	values := make([]int, len(numbers))
	for i, n := range numbers {
		values[i], _ = strconv.Atoi(n)
	}
	return values
}

// return the two values where one evenly divides the other
func findTwoEvenDivs(numbers []int) int {
	for _, n1 := range numbers {
		for _, n2 := range numbers {
			if n1 != n2 {
				// Could be n1/n2 or n2/n1, in case one is bigger than the other
				rem1 := math.Remainder(float64(n1), float64(n2))
				rem2 := math.Remainder(float64(n2), float64(n1))

				if rem1 == 0.0 {
					return n1 / n2
				}
				if rem2 == 0.0 {
					return n2 / n1
				}
			}
		}
	}
	return 0
}

