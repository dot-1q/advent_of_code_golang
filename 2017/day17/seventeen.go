package main

import (
	"fmt"
)

func main() {
	puzzle := 376
	buffer := []int{0, 0}
	position := 0

	// The length of the buffer will be the same as the i value.
	i := 1
	for i < 2018 {
		position = (position + puzzle) % i
		position++
		buffer = insertAt(buffer, i, position)
		i++
	}

	for i := range buffer {
		if buffer[i] == 2017 {
			fmt.Printf("Part 1 | Value after 2017: %d\n", buffer[i+1])
			break
		}
	}

	// If you notice, the value 0 is always at the front. SO then we want the second element. Instead of always
	// creating and moving array elements, simply keep track of this second element.
	for i < 50000001 {
		position = (position + puzzle) % i
		position++
		if position == 1 {
			buffer[1] = i
		}
		i++
	}
	fmt.Printf("Part 2 | Value after 0: %d\n", buffer[1])
}

func insertAt(buffer []int, number, position int) []int {
	newArray := []int{}
	newArray = append(newArray, buffer[:position]...)
	newArray = append(newArray, number)
	newArray = append(newArray, buffer[position:]...)
	return newArray
}
