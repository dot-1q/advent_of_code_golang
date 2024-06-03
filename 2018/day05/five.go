package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	polymer := []rune(strings.TrimSpace(string(f)))

	// Provide a copy of the original polymer
	tempPoly := make([]rune, len(polymer))
	copy(tempPoly, polymer)
	result := reactions(tempPoly)
	fmt.Printf("Part 1 | Units remaining: %d\n", len(result))

	// Go over the alphabet and remove each unit and measure
	minimum := math.MaxInt16
	culprit := 'a'
	for i := 'a'; i <= 'z'; i++ {
		tempPoly := make([]rune, len(polymer))
		copy(tempPoly, polymer)
		capital := i - 32
		removeUnit(&tempPoly, i, capital)
		// Measure
		length := len(reactions(tempPoly))
		if length < minimum {
			minimum = length
			culprit = i
		}
	}
	fmt.Printf("Part 2 | Culprit Unit is %c, and the remaining length is: %d\n", culprit, minimum)
}

func reactions(polymer []rune) string {
	replaced := true
	for replaced {
		i := 0
		replaced = false
		for i < (len(polymer) - 1) {
			// Given that the current char can be small or capital, its corresponding reaction
			// Is 32 positions after/before, if its capital or small. ASCII table.
			currentIdx := i

			// We replace deleted reactions with '.' so as not to be continuously dimensioning the array.
			// So we have to find the next index which has not been deleted.
			if polymer[i+1] == '.' {
				for polymer[i+1] == '.' {
					i++
					// Out of bounds check
					if i == (len(polymer) - 1) {
						break
					}
				}
			}
			// Out of bounds check
			if i == (len(polymer) - 1) {
				break
			}
			next := polymer[i+1]
			small := next + 32
			capital := next - 32
			// Its a reaction
			if polymer[currentIdx] == small || polymer[currentIdx] == capital {
				replaced = true
				polymer[currentIdx] = +'.'
				polymer[i+1] = '.'
				// Escape early. Only do one reaction per cycle.
				break
			}
			i += 1
		}
	}
	// Remove the dots and return the string
	result := string(polymer)
	return strings.ReplaceAll(result, ".", "")
}

func removeUnit(polymer *[]rune, small, capital rune) {
	for i := 0; i < len(*polymer); i++ {
		if (*polymer)[i] == small || (*polymer)[i] == capital {
			(*polymer)[i] = '.'
		}
	}
}
