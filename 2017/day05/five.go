package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	bytes, _ := os.ReadFile("input.txt")
	array := strings.Split(strings.TrimSpace(string(bytes)), "\n")
	numbers := strToInt(array)
	part2 := true

	steps := 0
	i := 0
	for i < len(numbers) {
		jump := numbers[i]
		if jump >= 3 && part2 {
			numbers[i]--
		} else {
			numbers[i]++
		}
		steps++
		i += jump
		if i > len(numbers)-1 {
			break
		}
	}
	fmt.Println("Steps: ", steps)

}

func strToInt(array []string) []int {
	numbers := make([]int, len(array))
	for i, a := range array {
		n, _ := strconv.Atoi(a)
		numbers[i] = n
	}
	return numbers
}
