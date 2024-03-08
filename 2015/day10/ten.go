package main

import (
	"fmt"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	puzzle := "1113222113"

	// In part one, we're supposed to run the puzzle 40 times
	// and get the length of the final number
	p := puzzle
	s := ""
	for range 40 {
		s = lookAndSay(p)
		p = s
	}
	fmt.Printf("Length of Part1 %d\n", len(p))

}

func partTwo() {
	puzzle := "1113222113"

	// In part two, we just run the same algorithm 50 times
	p := puzzle
	s := ""
	for range 50 {
		s = lookAndSay(p)
		p = s
	}
	fmt.Printf("Length of Part2 %d\n", len(p))
}

// Create the look and say number via a given string
func lookAndSay(number string) string {
	var lookAndSay strings.Builder

	runes := []rune(number)
	// Iterate through the string
	i := 0
	for i < len(runes) {
		// Count how many consecutive chars present
		count := 1
		// Start with the next char
		n := i + 1
		for {
			if (n < len(runes)) && (runes[n] == runes[i]) {
				// Increment the number of occurrences
				count++
			} else {
				// Stop this loop if the character is different
				break
			}
			// Increment the next char, and the index of the array, since we are in a nested loop
			n++
			i++
		}
		// Create the string with SprintF
		// Using a Builder for strings is way way way way faster than trying to
		// Concatenate each time , ex: lookAndSay += "2" + "1"
		lookAndSay.WriteString(fmt.Sprintf("%d%s", count, string(runes[i])))
		i++
	}
	return lookAndSay.String()
}
