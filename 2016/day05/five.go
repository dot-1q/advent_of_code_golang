package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	puzzle := "ffykfhsq"
	partOne(puzzle)
	partTwo(puzzle)
}

func partOne(puzzle string) {
	password := strings.Builder{}

	n := 0
	for password.Len() < 8 {
		// This breaks after we've done part two
		char, _, i := getHash(puzzle, n+1)
		password.WriteString(char)
		n = i

	}
	fmt.Println("Part1: ", password.String())

}

func partTwo(puzzle string) {
	password := make([]string, 8)

	n := 0
	charsComputed := 0
	for charsComputed < 8 {
		p, char, number := getHash(puzzle, n+1)
		n = number
		pos, err := strconv.Atoi(p)
		if err != nil || pos > 7 {
			// Is not a number, it a hex char
			// Or is a position bigger than the password
			continue
		}

		// Don't override
		if password[pos] == "" {
			password[pos] = char
			charsComputed++
		}
	}
	fmt.Println("Part2: ", strings.Join(password[:], ""))
}

// Get the first character of the first hash with five 0, and return the number
// so we can star the next search from that previous number
func getHash(str string, number int) (string, string, int) {
	i := number
	for {
		input := fmt.Sprintf("%s%d", str, i)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input)))
		if strings.HasPrefix(hash, "00000") {
			// fmt.Println(input)
			// fmt.Println(hash)
			return string(hash[5]), string(hash[6]), i
		}
		// If first 5 chars are zero
		i++
	}
}
