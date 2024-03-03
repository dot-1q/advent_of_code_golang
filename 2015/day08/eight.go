package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var bytes []byte
	var totalStringLiterals int
	var totalChars int
	for scanner.Scan() {
		// Get the full chars of the string as an array of bytes
		bytes = scanner.Bytes()
		// Add its length as the total chars of the string itself
		totalStringLiterals += len(bytes)
		// We should iterate only through the bytes inside the quotation marks ""
		totalChars += getNumberOfChars(bytes[1 : len(bytes)-1])
	}
	diff := totalStringLiterals - totalChars
	fmt.Printf("Part1: Total Literals: %d, Total Chars: %d, Diff: %d\n", totalStringLiterals, totalChars, diff)
}

func partTwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var bytes []byte
	var totalStringLiterals int
	var totalChars int
	for scanner.Scan() {
		// Get the full chars of the string as an array of bytes
		bytes = scanner.Bytes()
		// Add its length as the total chars of the string itself
		totalStringLiterals += len(bytes)
		// Now we iterate through all chars
		// +2 because this iteration adds two " to each string
		totalChars += getNumberOfCharsV2(bytes) + 2

	}
	diff := totalChars - totalStringLiterals
	fmt.Printf("Part2: Total Literals: %d, Total Chars: %d, Diff: %d\n", totalStringLiterals, totalChars, diff)

}

func getNumberOfChars(bytes []byte) int {
	// Iterate through the bytes
	// TODO: DON'T FORGET TO UPDATE THE INDEX AFTER WE CHECK THE TOKENS, AS ESCAPED CHARS DONT NEED TO BE CHECKED
	var numberOfChars int
	for i := 0; i < len(bytes); i++ {
		switch bytes[i] {
		// Escape character
		case '\\':
			// Since we will be checking the next index (\" or \\), and the next two indexes
			// in the case of (\xXX), we need to make sure we don't go out of bounds
			n1 := i + 1
			n2 := i + 3
			// Checking \" and \\, if we're out of bounds, the first check fails
			if n1 < len(bytes) && ((bytes[n1] == '\\') || bytes[n1] == '"') {
				numberOfChars++
				// Update index, since we checked this char already
				i++
			}
			// Checking \xBB, if we're out of bounds, the first check fails
			if n2 < len(bytes) && (bytes[n1] == 'x') {
				// Check if the next two chars can be a hex string
				if isValidHex(bytes[i+2]) && isValidHex(bytes[i+3]) {
					numberOfChars++
					// Update index, since we checked this char and the next already
					i = i + 3
				}

			}
		//Else its just a regular char
		default:
			numberOfChars++
		}
	}
	return numberOfChars
}

// In this version we just need to doubly count the number of escape characters,
// since they themselves will be escaped
func getNumberOfCharsV2(bytes []byte) int {
	// Iterate through the bytes
	// TODO: DON'T FORGET TO UPDATE THE INDEX AFTER WE CHECK THE TOKENS, AS ESCAPED CHARS DONT NEED TO BE CHECKED
	var numberOfChars int
	for i := 0; i < len(bytes); i++ {
		switch bytes[i] {
		// Escape character
		case '\\':
			// Since we will be checking the next index (\" or \\), and the next two indexes
			// in the case of (\xXX), we need to make sure we don't go out of bounds
			n1 := i + 1
			n2 := i + 3
			// Checking \" and \\, if we're out of bounds, the first check fails
			if n1 < len(bytes) && ((bytes[n1] == '\\') || bytes[n1] == '"') {
				numberOfChars += 4
				// Update index, since we checked this char already
				i++
			}
			// Checking \xBB, if we're out of bounds, the first check fails
			if n2 < len(bytes) && (bytes[n1] == 'x') {
				numberOfChars += 2
				// Check if the next two chars can be a hex string
				if isValidHex(bytes[i+2]) && isValidHex(bytes[i+3]) {
					numberOfChars += 3
					// Update index, since we checked this char and the next already
					i = i + 3
				}

			}
		case '"':
			// +2 because it would need to be escaped with a '\'
			numberOfChars += 2
		//Else its just a regular char
		default:
			numberOfChars++
		}
	}
	return numberOfChars
}

// Check if a given char can be an HEX value
func isValidHex(char byte) bool {

	if ('0' <= char && char <= '9') || ('a' <= char && char <= 'f') {
		return true
	}

	return false
}
