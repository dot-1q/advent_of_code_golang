package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	partone()
	parttwo()
}

func partone() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	i := 0
	for scanner.Scan() {
		if isGoodString_part1(scanner.Text()) {
			i++
		}
	}

	fmt.Printf("Part1: Input has %d nice strings\n", i)

}

func parttwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	i := 0
	for scanner.Scan() {
		if isGoodString_part2(scanner.Text()) {
			i++
		}
	}

	fmt.Printf("Part2: Input has %d nice strings\n", i)

}

func isGoodString_part1(s string) bool {

	// Can't contain these substrings
	has_specific, _ := containsSubstrings(s, "ab", "cd", "pq", "xy")
	// Needs to have two contiguous letters
	has_cont := hasContiguous(s)
	// Needs to contain three vowels
	has_sub, num := containsSubstrings(s, "a", "e", "i", "o", "u")

	if !has_specific && has_cont && (has_sub && num >= 3) {
		// Uncomment for verbose string check
		// fmt.Printf("String \"%s\" is a good string with %d vowels\n", s, num)
		return true
	}
	return false
}

func isGoodString_part2(s string) bool {

	// Needs to have two pairs of letters
	has_Pairs := hasPairs(s)
	repeats := repeatsInBetween(s)

	if has_Pairs && repeats {
		// Uncomment for verbose string check
		// fmt.Printf("String \"%s\" is a good string\n", s)
		return true
	}
	return false
}

// Check if a given string has any of the substrings, and if so, how many of them
func containsSubstrings(s string, subs ...string) (bool, int) {

	num := 0
	has_sub := false
	for i := range subs {
		// If this string contain the substring at least one time.
		occur := strings.Count(s, subs[i])
		if occur > 0 {
			num += occur
			has_sub = true
		}
	}

	return has_sub, num
}

// Check if a given string has a contiguous set of letters ex: (aa, bb, cc, ...)
func hasContiguous(s string) bool {

	runes := []rune(s)

	for i := 0; i < (len(runes) - 1); i++ {
		if runes[i] == runes[i+1] {
			return true
		}
	}
	return false
}

// Check if a given string has pairs that repeat but are not overlapped
func hasPairs(s string) bool {

	runes := []rune(s)

	for i := 0; i < (len(runes) - 1); i++ {
		// Little bit of magic
		// We check if the string contains a substring
		// that is made of its character and the next one
		// We do this by slicing
		// We have to be careful and check the string with the pair removed till its end[pair:],
		// so we don't count overlaps, ex: (string: aaa with pair [aa])

		slc := string(runes[i : i+2]) // Grab the pair
		str := string(runes[i+2:])    // Grab the string AFTER the pair
		contains, _ := containsSubstrings(str, slc)
		if contains {
			return true
		}

	}
	return false
}

// Check if string has a repeating char with any char in between ex: (xyx, aaa, bcb)
func repeatsInBetween(s string) bool {

	runes := []rune(s)

	for i := 0; i < (len(runes) - 2); i++ {
		// If its the same as the char as the one two places next
		if runes[i] == runes[i+2] {
			return true
		}
	}
	return false
}
