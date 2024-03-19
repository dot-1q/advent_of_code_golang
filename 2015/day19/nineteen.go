package main

import (
	"bufio"
	"fmt"
	"os"
	s "strings"
)

func main() {
	f, _ := os.Open("Input.txt")
	scanner := bufio.NewScanner(f)

	replacementes := map[string][]string{}
	molecule := ""
	for scanner.Scan() {
		line := s.Split(scanner.Text(), "=>")

		if len(line) == 2 {
			// Check if it has the value already
			m := s.TrimSpace(line[0])
			if _, ok := replacementes[m]; ok {
				replacementes[m] = append(replacementes[m], s.TrimSpace(line[1]))
			} else {
				replacementes[m] = []string{s.TrimSpace(line[1])}
			}
		} else if len(line) == 1 {
			// Its the molecule
			molecule = line[0]
		}
	}
	fmt.Println(replacementes)
	m := moleculeArray(molecule)
	partOne(replacementes, m)
	partTwo(m)

	// Be mindful of copy and deep copy
	// a := []string{"a", "b", "c", "d"}
	// fmt.Println(a)
	// s := a
	// s[0] = "b"
	// fmt.Println(a)
	// fmt.Println(s)
}

func partOne(replacements map[string][]string, molecule []string) {
	// Map of all the combinations of molecules that were replaced
	// Easier to find if there's collisions (repetitions)
	allCombinations := map[string]bool{}

	for i, mol := range molecule {
		// Check if this molecule has a replacements
		// If so, replace and create a new molecule combination
		if replacement, ok := replacements[mol]; ok {
			// For each replacement, create a new string
			str := []string{}
			str = append(str, molecule...)
			for _, r := range replacement {
				// Temporarily hold the original molecule array
				// Replace
				str[i] = r
				// Create new string and combination
				n := s.Join(str, "")
				allCombinations[n] = true
			}
		}
	}
	fmt.Println("Part1: ", len(allCombinations))
}

func partTwo(molecule []string) {
	// I am unable to do this
	// I take the solution from:
	// https://www.reddit.com/r/adventofcode/comments/3xflz8/comment/cy4etju/?utm_source=share&utm_medium=web3x&utm_name=web3xcss&utm_term=1&utm_content=share_button

	elements := len(molecule)
	RnAr := count(molecule, "Rn") + count(molecule, "Ar")
	y := count(molecule, "Y")
	steps := elements - RnAr - (2 * y) - 1
	fmt.Println("Part2: ", steps)
}

func count(slice []string, s string) int {
	count := 0
	for _, value := range slice {
		if value == s {
			count++
		}
	}
	return count
}

// From a molecule string, return the array of molecules that comprise the molecule
func moleculeArray(molecule string) []string {
	moleculeArray := []string{}

	mol := []rune(molecule)
	// For each char in the string
	for i := 0; i < len(molecule); i++ {
		// Make sure we dont 'index out of range'
		// Check if the next char is lowercase. If it is, means its a
		// molecule with the name 'Xx'
		if (i+1) < len(mol) && (mol[i+1] >= 97 && mol[i+1] <= 122) {

			moleculeArray = append(moleculeArray, s.Join([]string{string(mol[i]), string(mol[i+1])}, ""))
			i++
		} else {
			moleculeArray = append(moleculeArray, string(mol[i]))
		}
	}
	return moleculeArray
}
