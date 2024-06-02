package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	two, three := countOccurrences(lines)
	fmt.Printf("Part 1 | Checksum: %d\n", two*three)
	fmt.Printf("Part 2 | Matching strings: %s\n", findOneCharDiff(lines))
}

func countOccurrences(ids []string) (int, int) {
	two := 0
	three := 0
	for _, id := range ids {
		seen := map[rune]bool{}
		countedTwo := false // If 2 or more letters appear twice, it still will only count as 1.
		countedThree := false
		for _, letter := range id {
			// If we haven't counted this letter
			if _, ok := seen[letter]; !ok {
				seen[letter] = true
				occurrences := strings.Count(id, string(letter))
				if occurrences == 2 {
					if !countedTwo {
						two++
						countedTwo = true
					}
				}
				if occurrences == 3 {
					if !countedThree {
						countedThree = true
						three++
					}
				}
			}
		}
	}
	return two, three
}

func findOneCharDiff(ids []string) string {
	for _, id := range ids {
		for _, second := range ids {
			// Don't compare equal strings.
			diff := 0
			diffChar := ' '
			for i, char := range id {
				// Check if the character in string 1 is present in string 2
				// If it isnt, it is different
				if char != rune(second[i]) {
					diffChar = char
					diff++
				}
				// IF there is more than two characters different. This is not a match. Escape early.
				if diff > 2 {
					break
				}
			}
			// If the difference was only one character, we found it. Escape early.
			if diff == 1 {
				before, after, _ := strings.Cut(id, string(diffChar))
				return before + after
			}
		}
	}
	return ""
}
