package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	// partOne(f)
	f.Seek(0, 0)
	partTwo(f)
}

func partOne(f *os.File) {
	scanner := bufio.NewScanner(f)
	totalIds := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		// id and hash are always the last elements of the string
		idAndHash := line[len(line)-1]
		id := strings.Split(idAndHash, "[")
		hash, _ := strings.CutSuffix(id[1], "]")

		str := strings.Join(line[0:len(line)-1], "")
		m := occurrenceMap(str)
		realHash := createHash(m)

		// If the calculated hash is the same as the inputted one
		if realHash == hash {
			num, _ := strconv.Atoi(id[0])
			totalIds += num
		}
	}

	fmt.Println("Part1: ", totalIds)
}

func partTwo(f *os.File) {
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		// id and hash are always the last elements of the string
		idAndHash := line[len(line)-1]
		id, _ := strconv.Atoi(strings.Split(idAndHash, "[")[0])
		roomName := decodeName(strings.Join(line[0:len(line)-1], " "), id)
		if strings.Contains(roomName, "north") {
			fmt.Printf("Part2: Room name: '%s', id: %d\n", roomName, id)
		}
	}
}

func occurrenceMap(s string) map[string]int {
	occurrence := map[string]int{}

	for _, char := range s {
		n := strings.Count(s, string(char))
		occurrence[string(char)] = n
	}
	return occurrence
}

func createHash(occurrence map[string]int) string {
	hash := []string{}

	// Append the keys to the hash string
	for key := range occurrence {
		hash = append(hash, key)
	}

	// First sort the letters alphabetically
	slices.Sort(hash)
	// Then sort by occurrence
	slices.SortStableFunc(hash, func(s1, s2 string) int {
		return occurrence[s2] - occurrence[s1]
	})

	// Only return the first 5 elements
	return strings.Join(hash, "")[0:5]
}

func decodeName(sequence string, id int) string {
	str := strings.Builder{}
	// There's 26 letters in the alphabet
	rotateBy := (id % 26)
	letters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for _, char := range sequence {
		// Only rotate the letters and not whitespace
		if char != ' ' {
			// Find the index of the char in the char array
			idx := slices.Index(letters, char)
			// Rotate it by the number we want
			idx += rotateBy
			// Modulo 26 because the char array is 26 chars long
			str.WriteRune(letters[idx%26])
		} else {
			str.WriteRune(' ')
		}
	}

	return str.String()
}
