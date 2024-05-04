package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	chars := strings.TrimSpace(string(f))
	fmt.Println(chars)
	sum, garbage := findGroups(chars)
	fmt.Println("Part1 : Sum:", sum)
	fmt.Println("Part2 : Garbage ", garbage)
}

func findGroups(line string) (int, int) {
	level := 0
	sum := 0
	garbage := false
	i := 0
	count := 0
	for i < len(line) {
		if garbage {
			count++
		}
		if line[i] == '{' && !garbage {
			level++
			sum += level
		}
		if line[i] == '}' && !garbage {
			level--
		}
		if line[i] == '<' && !garbage {
			garbage = true
		}
		if line[i] == '>' {
			garbage = false
			// We dont want to count this character
			count--
		}
		if line[i] == '!' {
			// We dont want to count this character
			count--
			i++
		}
		i++
	}
	return sum, count
}
