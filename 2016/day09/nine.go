package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	line := string(f)
	line = strings.TrimSuffix(line, "\n")

	str := strings.Builder{}
	decompress(&str, line, 0)
	fmt.Println("Part1:", len(str.String()))

	partTwo(line)
}

func partTwo(compressed string) {

	fmt.Println("Part2:", countLen(compressed))

}

// Recursive function to decompress
func decompress(str *strings.Builder, compressed string, index int) {
	// Base case
	if index > len(compressed)-1 {
		return
	}

	if compressed[index] == '(' {
		// Characters inside parenthesis. This looks like nonsense im aware
		inside_paren, _ := strings.CutSuffix(strings.SplitAfterN(compressed[index+1:], ")", 2)[0], ")")
		expr := strings.Split(inside_paren, "x")
		// From the expression inside the parenthesis, get the number of letters to repeat, and the number
		num_letters_repeated, _ := strconv.Atoi(expr[0])
		reps, _ := strconv.Atoi(expr[1])
		// Get the index of the closing ')', so we know when the next letter is
		idx := strings.Index(compressed[index+1:], ")")
		letters := compressed[(index+1)+(idx+1) : (index+1)+(idx+1)+num_letters_repeated]
		// Write the letters repeated
		str.WriteString(strings.Repeat(letters, reps))
		// recurse
		// Index will be after all the stuff enclosed in the parenthesis, and after the letters that were repeated
		index += 1 + (idx + 1) + num_letters_repeated
		decompress(str, compressed, index)

	} else {
		str.WriteString(string(compressed[index]))
		decompress(str, compressed, index+1)
	}
}

func countLen(compressed string) int {
	strLen := 0
	i := 0
	for i < len(compressed) {
		if compressed[i] == '(' {
			// Characters inside parenthesis. This looks like nonsense im aware
			inside_paren, _ := strings.CutSuffix(strings.SplitAfterN(compressed[i+1:], ")", 2)[0], ")")
			expr := strings.Split(inside_paren, "x")
			// From the expression inside the parenthesis, get the number of letters to repeat, and the number
			num_letters_repeated, _ := strconv.Atoi(expr[0])
			reps, _ := strconv.Atoi(expr[1])
			// Get the index of the closing ')', so we know when the next letter is
			idx := strings.Index(compressed[i+1:], ")")
			letters := compressed[(i+1)+(idx+1) : (i+1)+(idx+1)+num_letters_repeated]
			// the length will be the repetitions times the length of this subsection, which may contain more nested operations
			strLen += reps * countLen(letters)
			// Update the index for the position after the nesting
			i += 1 + (idx + 1) + num_letters_repeated
		} else {
			strLen++
			i++
		}
	}
	return strLen
}
