package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("hi")
	f, _ := os.Open("input.txt")
	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	tls := 0
	for _, line := range lines {
		in, out := separateAddress(line)
		if supportsTLS(in, out) {
			fmt.Printf("Line %s supports tls\n", line)
			tls++
		} else {
			fmt.Printf("Line %s does not support tls\n", line)
		}
	}
	fmt.Println("Part1: ", tls)
}

func partTwo(lines []string) {
	ssl := 0
	for _, line := range lines {
		in, out := separateAddress(line)
		if supportsSSL(in, out) {
			fmt.Printf("Line %s supports ssl\n", line)
			ssl++
		} else {
			fmt.Printf("Line %s does not support ssl\n", line)
		}
	}
	fmt.Println("Part2: ", ssl)
}

func separateAddress(address string) ([]string, []string) {
	str_inside := strings.Builder{}
	str_outside := strings.Builder{}

	strings_inside := []string{}
	strings_outside := []string{}
	i := 0
	for i < len(address) {

		// Is between []
		if address[i] == '[' {
			// An inside string started, the other outside one ended
			strings_outside = append(strings_outside, str_outside.String())
			str_outside.Reset()
			i++
			for address[i] != ']' {
				// Go to next char
				str_inside.WriteString(string(address[i]))
				i++
			}
			// Go to the char after the ']'
			i++
			strings_inside = append(strings_inside, str_inside.String())
			str_inside.Reset()
		}

		str_outside.WriteString(string(address[i]))
		i++
	}
	strings_outside = append(strings_outside, str_outside.String())
	return strings_inside, strings_outside
}

func supportsTLS(inside, outside []string) bool {
	// None of these strings can follow the rule
	for _, ins := range inside {
		if hasReversePair(ins) {
			return false
		}
	}
	// If the above test passes, and any of these strings has reverse pairs, pass
	for _, out := range outside {
		if hasReversePair(out) {
			return true
		}
	}
	return false
}

func supportsSSL(inside, outside []string) bool {
	// Has the ABA in any of the outside sequences
	found := false
	for _, out := range outside {
		// For all the sequences found, check if the inside strings have the BAB sequence
		f, seqs := hasABA(out)
		if f {
			fmt.Println(seqs)
			for _, ins := range inside {
				if hasBAB(ins, seqs) {
					found = true
				}
			}
		}
	}
	return found
}

func hasReversePair(str string) bool {
	i := 0
	for (i + 3) < len(str) {
		// Chars 1 and 4,and 2 and 3 need to be the same, ex: abba
		// they cannot be repeated, so char 1 and 2 cannot be the same
		if str[i] == str[i+3] && str[i+1] == str[i+2] && str[i] != str[i+1] {
			return true
		}
		i++
	}
	return false
}

// Return if a string has the pattern and the middle char thats between such pattern
// we could find more than one patters
func hasABA(str string) (bool, []string) {
	seqs := []string{}
	found := false
	i := 0
	for (i + 2) < len(str) {
		// Chars 1 and 3 must be the same, and the middle one a different char
		// ex: aba, xyx
		if str[i] == str[i+2] && str[i] != str[i+1] {
			// save the sequence in BAB format, to compare latter
			s := string(str[i+1]) + string(str[i]) + string(str[i+1])
			seqs = append(seqs, s)
			found = true
		}
		i++
	}
	return found, seqs
}

func hasBAB(str string, sequences []string) bool {
	// Check if any of the sequences are present in the string
	for _, seq := range sequences {
		if strings.Contains(str, seq) {
			return true
		}
	}
	return false
}
