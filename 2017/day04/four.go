package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	validPassphrases := 0
	part2 := true
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if part2 {
			if !hasAnagram(words) {
				validPassphrases++
			}
		} else {
			if !hasDuplicate(words) {
				validPassphrases++
			}
		}
	}

	fmt.Println("Valid passphrases:", validPassphrases)
}

func hasDuplicate(words []string) bool {
	for idx, word := range words {
		// Make two ranges, before and after the word in question. We don't want to include it in the search.
		slice1 := words[:idx]
		slice2 := words[idx+1:]
		if slices.Contains(slice1, word) || slices.Contains(slice2, word) {
			return true
		}
	}
	return false
}

func hasAnagram(words []string) bool {
	for idx, word := range words {
		// Make two ranges, before and after the word in question. We don't want to include it in the search.
		slice1 := words[:idx]
		slice2 := words[idx+1:]
		ordered := sortString(word)
		// Order the strings, and check if they are equal
		for _, w := range slice1 {
			wOrder := sortString(w)
			if wOrder == ordered {
				return true
			}
		}
		for _, w := range slice2 {
			wOrder := sortString(w)
			if wOrder == ordered {
				return true
			}
		}
	}
	return false
}

func sortString(word string) string {
	chars := []rune(word)
	slices.Sort(chars)
	return string(chars)
}
