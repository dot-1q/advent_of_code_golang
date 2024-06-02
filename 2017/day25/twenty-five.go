package main

import (
	"fmt"
)

func main() {
	// Im not going to parse the input.txt. I'm just going to hand write the values. No point in doing otherwise.
	state := "A"           // starting state.
	const steps = 12656374 // steps
	tape := [steps]int{}
	// Start in the middle
	position := steps / 2
	for range steps {
		value := tape[position]
		switch state {
		case "A":
			if value == 0 {
				tape[position] = 1
				position++
				state = "B"
			} else {
				tape[position] = 0
				position--
				state = "C"
			}
		case "B":
			if value == 0 {
				tape[position] = 1
				position--
				state = "A"
			} else {
				position--
				state = "D"
			}
		case "C":
			if value == 0 {
				tape[position] = 1
				position++
				state = "D"
			} else {
				tape[position] = 0
				position++
				state = "C"
			}
		case "D":
			if value == 0 {
				position--
				state = "B"
			} else {
				tape[position] = 0
				position++
				state = "E"
			}
		case "E":
			if value == 0 {
				tape[position] = 1
				position++
				state = "C"
			} else {
				position--
				state = "F"
			}
		case "F":
			if value == 0 {
				tape[position] = 1
				position--
				state = "E"
			} else {
				position++
				state = "A"
			}
		}
	}
	ones := countOnes(tape[:])
	fmt.Printf("Part 1 | There's %d ones\n", ones)
}

func countOnes(tape []int) int {
	s := 0
	for _, n := range tape {
		if n == 1 {
			s++
		}
	}
	return s
}
