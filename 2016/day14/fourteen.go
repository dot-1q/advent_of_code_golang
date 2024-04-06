package main

import (
	"crypto/md5"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	puzzle := "ahsbgdzn"
	fmt.Println(puzzle)
	partOne(puzzle)
	partTwo(puzzle)
}

func partOne(puzzle string) {
	found := 0
	hashes := make([]string, 0)

	// Create a list of 1000 ashes, because if we do find a triplet in any given hash index
	// we need go over the next 1000 indexes and check if there's a 5 letter sequence of the same triplet (XXXXX)
	// This just help keep the 1000 hashes in memory for better performance.
	for i := 1; i < 1000; i++ {
		h := createHash(puzzle, i)
		hashes = append(hashes, h)
	}

	idx := 1
	for found < 64 {
		// We've already created the 1000 first hashes, now create the N+1000
		currentHash := hashes[0]
		h := createHash(puzzle, idx+1000)
		hashes = append(hashes, h)
		// Discard first element
		hashes = hashes[1:]

		// If this hash has a triplet, check if it has a 5 repeating of that char in the next 1000 hashes
		if char := hasTriplet(currentHash); char != "" {
			// Check if the hashes array contain any has with this char repeating 5 times
			if slices.ContainsFunc(hashes, func(str string) bool {
				return strings.Contains(str, strings.Repeat(char, 5))
			}) {
				// If both checks pass, we've found a valid hash
				found++
			}
		}
		idx++
	}
	fmt.Println("Part1: The index of the 64th valid hash was:", idx)

}

func partTwo(puzzle string) {
	found := 0
	hashes := make([]string, 0)

	// Create a list of 1000 ashes, because if we do find a triplet in any given hash index
	// we need go over the next 1000 indexes and check if there's a 5 letter sequence of the same triplet (XXXXX)
	// This just help keep the 1000 hashes in memory for better performance.
	for i := 1; i < 1000; i++ {
		h := createNthHash(puzzle, i, 2016)
		hashes = append(hashes, h)
	}

	idx := 1
	for found < 64 {
		// We've already created the 1000 first hashes, now create the N+1000
		currentHash := hashes[0]
		h := createNthHash(puzzle, idx+1000, 2016)
		hashes = append(hashes, h)
		// Discard first element
		hashes = hashes[1:]

		// If this hash has a triplet, check if it has a 5 repeating of that char in the next 1000 hashes
		if char := hasTriplet(currentHash); char != "" {
			// Check if the hashes array contain any has with this char repeating 5 times
			if slices.ContainsFunc(hashes, func(str string) bool {
				return strings.Contains(str, strings.Repeat(char, 5))
			}) {
				// If both checks pass, we've found a valid hash
				found++
			}
		}
		idx++
	}
	fmt.Println("Part2: The index of the 64th valid hash was:", idx)

}

// Check if a given string has a triplet, and return the char of said triplet
func hasTriplet(hash string) string {
	for i := 2; i < len(hash); i++ {
		if hash[i-2] == hash[i-1] && hash[i-1] == hash[i] {
			return string(hash[i])
		}
	}
	return ""
}

func createHash(salt string, index int) string {
	s := md5.Sum([]byte(salt + strconv.Itoa(index)))
	// Convert to string, has to be encoded
	return fmt.Sprintf("%x", s[:])
}

// Create the hash 2016 more times
func createNthHash(salt string, index int, times int) string {
	str := salt + strconv.Itoa(index)
	for range times + 1 {
		str = fmt.Sprintf("%x", md5.Sum([]byte(str)))

	}
	return str
}
