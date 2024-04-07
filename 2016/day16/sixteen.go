package main

import (
	"fmt"
	"slices"
)

func main() {
	part := 2
	puzzle := "10111100110001111"
	lengthDisk := 272
	if part == 2 {
		lengthDisk = 35651584
	}

	// Generate the data to fill the disk length
	data := generateData(puzzle)
	for len(data) < lengthDisk {
		data = generateData(string(data))
	}

	// Calculate the checksum while its length is even
	d := calculateChecksum(data[:lengthDisk+1])
	for len(d)%2 == 0 {
		d = calculateChecksum(d)
	}

	fmt.Printf("Part%d: Checksum: %s\n", part, string(d))
}

func generateData(data string) []rune {
	a := []rune(data)
	b := []rune(data)
	slices.Reverse(b)
	flipRunes(&b)
	a = append(a, '0')
	a = append(a, b...)

	return a
}

func flipRunes(runes *[]rune) {
	for i, r := range *runes {
		if r == '0' {
			(*runes)[i] = '1'
		} else {
			(*runes)[i] = '0'
		}
	}
}

func calculateChecksum(data []rune) []rune {
	checksum := []rune{}
	// Iterate over the data, and create the checksum according to the pairs
	//Non overlapping pairs, so the very last character wont be checked.
	// Also, iterate +=2, so that we go pair by pair, instead of the char by char default behaviour.
	for i := 0; i < len(data)-1; i += 2 {
		if data[i] == data[i+1] {
			checksum = append(checksum, '1')
		} else {
			checksum = append(checksum, '0')
		}
	}
	return checksum
}
