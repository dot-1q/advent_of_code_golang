package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	m, _ := os.ReadFile("input.txt")
	moves := strings.Split(strings.TrimSpace(string(m)), ",")
	letters := [16]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}
	fmt.Printf("Part 1 | Positions after dance: '%s'\n", performDance(moves, 1, letters))
	fmt.Printf("Part 2 | Positions after dance: '%s'\n", performDance(moves, 1000000000, letters))

}

func performDance(moves []string, times int, letters [16]string) string {
	seen := map[string]int{}
	for i := range times {
		for _, move := range moves {
			switch move[0] {
			// Spin
			case 's':
				amount, _ := strconv.Atoi(move[1:])
				letters = spin(letters, amount)
			// Exchange
			case 'x':
				ref := strings.Split(move[1:], "/")
				src, _ := strconv.Atoi(ref[0])
				dst, _ := strconv.Atoi(ref[1])
				exchange(&letters, src, dst)
			// Partner
			case 'p':
				ref := strings.Split(move[1:], "/")
				partner(&letters, ref[0], ref[1])
			}
		}
		s := strings.Join(letters[:], "")
		// I dont know lol.
		// Link: https://www.reddit.com/r/adventofcode/comments/7k572l/comment/drbota0/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button
		if _, ok := seen[s]; ok {
			if (1000000000-(i+1))%(i-seen[s]) == 0 {
				fmt.Println("Answer: ", s)
				break
			}
		}
		seen[s] = i
	}
	return strings.Join(letters[:], "")

}

func spin(letters [16]string, amount int) [16]string {
	position := 16 - amount
	newA := []string{}
	newA = append(newA, letters[position:]...)
	newA = append(newA, letters[:position]...)
	return [16]string(newA)
}

func exchange(letters *[16]string, src, dst int) {
	letters[src], letters[dst] = letters[dst], letters[src]
}

// Simply find index of the letters and swap them.
func partner(letters *[16]string, a, b string) {
	src := slices.Index(letters[:], a)
	dst := slices.Index(letters[:], b)
	exchange(letters, src, dst)
}
