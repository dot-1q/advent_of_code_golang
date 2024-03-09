package main

import (
	"fmt"
	"slices"
)

func main() {
	// PART 1
	puzzle := []rune("vzbxkghb")
	ans := calculatePassword(puzzle)
	fmt.Println("Part1: ", string(ans))
	// PART 2
	puzzle = []rune("vzbxxzaa")
	ans = calculatePassword(puzzle)
	fmt.Println("Part2: ", string(ans))

}

func calculatePassword(puzzle []rune) []rune {
	found := false
	for !found {
		// Do the loop while the last character is not z, only then exit and wrap around
		for puzzle[7] < 123 {
			// Check for each iteration
			if isValid(&puzzle) {
				found = true
				break
			}
			puzzle[7]++

		}

		// Wrap around and increment the right char
		// integer 123 == {, which is right after z, we increment after the validation
		if puzzle[7] == 123 {
			puzzle[7] = 'a'
			puzzle[6]++
		}
		if puzzle[6] == 123 {
			puzzle[6] = 'a'
			puzzle[5]++
		}
		if puzzle[5] == 123 {
			puzzle[5] = 'a'
			puzzle[4]++
		}
		if puzzle[4] == 123 {
			puzzle[4] = 'a'
			puzzle[3]++
		}
		if puzzle[3] == 123 {
			puzzle[3] = 'a'
			puzzle[2]++
		}
		if puzzle[2] == 123 {
			puzzle[2] = 'a'
			puzzle[1]++
		}
		if puzzle[1] == 123 {
			puzzle[1] = 'a'
			puzzle[0]++
		}
	}
	return puzzle
}

func isValid(runes *[]rune) bool {
	//String must not contain 'i','o','l'
	contains := slices.ContainsFunc(*runes, func(char rune) bool {
		return char == 'i' || char == 'o' || char == 'l'
	})

	//Must have three letters in a row in an ascending manner
	ascending := isAscending(runes)
	//Must have two pairs
	pairs := hasPairs(runes)
	// Return the bool algebra of the previous conditions
	return !contains && ascending && pairs
}

// Check if a given string sequence is ascending ex: abc, bcd, cde, ...
func isAscending(runes *[]rune) bool {
	for i, rune := range *runes {
		next_char := i + 1
		next_2char := i + 2
		// Make sure we dont overflow
		// If the sequence is ascending, the difference between the int32 (runes), will be 1 for the first, and 2 for the second character
		// Remember runes are just int32, which later get coded into utf-8 strings
		if next_2char < len(*runes) && ((*runes)[next_char]-rune == 1) && ((*runes)[next_2char]-rune == 2) {
			return true
		}
	}
	return false
}

// Check if a given string has two pairs of letters ex: 'aa', 'bb', ...
func hasPairs(runes *[]rune) bool {
	count := 0
	// Since we increase the index within the loop, we have to declare it outside
	i := 0
	for i < len(*runes) {
		next_char := i + 1
		// Make sure we dont overflow
		// This char and the next must be equal, and not overlap, meaning if we
		// find an equal char, we increase the index, so we dont count the repeated
		// character again
		if next_char < len(*runes) && ((*runes)[i] == (*runes)[next_char]) {
			count++
			i++
		}
		i++
	}
	// If the number of Pairs is 2, returns true, which is what we want
	return count >= 2
}
