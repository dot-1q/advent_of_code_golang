package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	puzzle := "ckczppom"

	i, k, h := partone(puzzle)
	fmt.Printf("Part1: For puzzle %s, final string is %s, the number is %d and hash is %x\n", puzzle, k, i, h)
	i, k, h = parttwo(puzzle)
	fmt.Printf("Part2: For puzzle %s, final string is %s, the number is %d and hash is %x\n", puzzle, k, i, h)

}

func partone(puzzle string) (int, string, [16]byte) {
	i := 0
	key := ""
	for {
		key = puzzle + strconv.Itoa(i)
		hash := md5.Sum([]byte(key))
		s := bytes.Runes(hash[:3])

		// First two bytes need to be zero
		if s[0] == 0 && s[1] == 0 {
			// The third byte has to start with a 0, but can be any of these values
			// This may seem ugly but its lightning fast
			// We could've converted the hash to string and then check its slice[:5], but i wanted it this way :)
			if s[2] == 0 || s[2] == 1 || s[2] == 2 || s[2] == 3 ||
				s[2] == 4 || s[2] == 5 || s[2] == 6 || s[2] == 7 ||
				s[2] == 8 || s[2] == 9 || s[2] == 10 || s[2] == 11 ||
				s[2] == 12 || s[2] == 13 || s[2] == 14 || s[2] == 15 {
				// Found the hash
				return i, key, hash
			}
		}
		i++
	}
}

func parttwo(puzzle string) (int, string, [16]byte) {
	i := 0
	key := ""
	for {
		key = puzzle + strconv.Itoa(i)
		hash := md5.Sum([]byte(key))
		s := bytes.Runes(hash[:3])

		// First three bytes need to be zero
		if s[0] == 0 && s[1] == 0 && s[2] == 0 {
			// Found the hash
			return i, key, hash
		}
		i++
	}
}
