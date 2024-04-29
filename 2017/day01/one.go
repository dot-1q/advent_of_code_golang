package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	numbers := []rune(string(f))

	fmt.Println("Part One | Sum is: ", partOne(numbers))
	fmt.Println("Part Two | Sum is: ", partTwo(numbers))
}

func partOne(numbers []rune) int {

	// Get the number of elements
	n_length := len(numbers) - 1
	sum := 0
	for i := 0; i <= n_length; i++ {
		// Apply some modulo shenanigans to not go out of bounds and still be circular.
		if numbers[i%n_length] == numbers[(i+1)%n_length] {
			n, _ := strconv.Atoi(string(numbers[i]))
			sum += n
		}
	}
	return sum
}

func partTwo(numbers []rune) int {

	// Get the halfway point of the list of numbers
	steps_ahead := len(numbers) / 2
	sum := 0
	for i := 0; i < len(numbers); i++ {
		// Apply some modulo shenanigans to not go out of bounds and still be circular.
		// fmt.Printf("Index: %d, steps: %d\n", i, (i+steps_ahead)%(len(numbers)-1)) // DEBUG
		if numbers[i] == numbers[(i+steps_ahead)%(len(numbers)-1)] {
			n, _ := strconv.Atoi(string(numbers[i]))
			sum += n
		}
	}
	return sum
}
