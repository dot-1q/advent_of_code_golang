package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	start := []rune("abcdefgh")
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(f), "\n")
	// Remove the \n from the array
	lines = lines[:len(lines)-1]
	decoded := decodePassword(start, lines, false)
	fmt.Printf("Part1: %s\n", decoded)
	// Part 2 we have to unscramble
	// So the order of the instructions have to be reversed
	slices.Reverse(lines)
	scramble := []rune("fbgdceah")
	fmt.Printf("Part2: %s\n", decodePassword(scramble, lines, true))
}

// Decode a password based on a given instruction. For part2, we should reverse the transformation. Some reversals are trivial
// Moves,Swaps and ReverseStrings are the same if reverse or normal order, they undo eachother
// Rotations have to be their symmetric.
func decodePassword(start []rune, lines []string, reverse bool) string {
	for _, inst := range lines {
		line := strings.Split(inst, " ")

		switch line[0] {
		case "swap":
			if line[1] == "position" {
				src, _ := strconv.Atoi(line[2])
				dst, _ := strconv.Atoi(line[5])
				swapPosition(start, src, dst)
			} else if line[1] == "letter" {
				// crazy conversions
				swapLetters(start, []rune(line[2])[0], []rune(line[5])[0])
			}
		case "reverse":
			src, _ := strconv.Atoi(line[2])
			dst, _ := strconv.Atoi(line[4])
			reverseString(start, src, dst)
		case "rotate":
			if line[1] == "based" {
				if !reverse {
					start = rotateBased(start, []rune(line[6])[0])
				} else {
					start = rotateBasedReverse(start, []rune(line[6])[0])
				}
			} else if line[1] == "right" {
				n, _ := strconv.Atoi(line[2])
				if !reverse {
					start = rotateRight(start, n)
				} else {
					start = rotateLeft(start, n)
				}
			} else if line[1] == "left" {
				n, _ := strconv.Atoi(line[2])
				if !reverse {
					start = rotateLeft(start, n)
				} else {
					start = rotateRight(start, n)
				}
			}
		case "move":
			src, _ := strconv.Atoi(line[2])
			dst, _ := strconv.Atoi(line[5])
			if !reverse {
				start = move(start, src, dst)
			} else {
				start = move(start, dst, src)
			}
		}
	}
	return string(start)
}

// Rotate array of runes. It creates a new one
func rotateLeft(str []rune, positions int) []rune {
	r := (positions % len(str))
	return append(str[r:], str[:r]...)
}

// Rotate array of runes. It creates a new one
func rotateRight(str []rune, positions int) []rune {
	r := len(str) - (positions % len(str))
	return append(str[r:], str[:r]...)
}

func rotateBased(str []rune, letter rune) []rune {
	idxLetter := slices.Index(str, letter)
	numberOfRotations := idxLetter + 1
	if idxLetter >= 4 {
		numberOfRotations++
	}
	return rotateRight(str, numberOfRotations)
}

func rotateBasedReverse(str []rune, letter rune) []rune {
	idxLetter := slices.Index(str, letter)
	numberOfRotations := idxLetter
	if numberOfRotations == 0 || numberOfRotations%2 == 1 {
		numberOfRotations = ((numberOfRotations / 2) + 1)
	} else {
		numberOfRotations = ((numberOfRotations / 2) + 5)
	}
	return rotateLeft(str, numberOfRotations)
}

// Reverse a given string from a start position till the end
// It *DOES NOT* create a new array, simply arranges the underlying one.
func reverseString(str []rune, start, end int) {
	for i, j := start, end; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
}

func swapPosition(str []rune, src, dst int) {
	str[src], str[dst] = str[dst], str[src]
}

func swapLetters(str []rune, src, dst rune) {
	idxSource := slices.Index(str, src)
	idxDest := slices.Index(str, dst)
	str[idxSource], str[idxDest] = str[idxDest], str[idxSource]
}

// Move one position to another
func move(str []rune, src, dst int) []rune {
	// Make a new array that will hold the transformation
	newStr := make([]rune, len(str))
	newStr[dst] = str[src]

	// Hold the index that we want to copy from the original string
	copyIdx := 0
	for i := 0; copyIdx < len(newStr); i++ {
		// If its the destination, we've already copied, so skip this position
		if i == dst {
			i++
		}
		// When we arrive at the copy position from where we moved, also skip
		// because this letter is in a new position. This is hard to explain and better visualized.
		if copyIdx == src {
			copyIdx++
		}
		if copyIdx < len(newStr) {
			newStr[i] = str[copyIdx]
		}
		copyIdx++
	}
	return newStr
}
