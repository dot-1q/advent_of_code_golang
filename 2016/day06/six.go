package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hi")

	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)
	columns := [8][]rune{}
	for scanner.Scan() {
		// array of chars for each line
		chars := []rune(scanner.Text())

		for i, char := range chars {
			columns[i] = append(columns[i], rune(char))
		}
	}

	partOne(columns)
	//gddd
	partTwo(columns)
}

func partOne(columns [8][]rune) {
	letters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	str := ""

	for _, col := range columns {
		occurrence := map[rune]int{}
		for _, letter := range letters {
			n := strings.Count(string(col[:]), string(letter))
			occurrence[letter] = n
		}
		str += string(findMostCommon(occurrence))
	}

	fmt.Println("Part1: ", str)
}

func partTwo(columns [8][]rune) {
	letters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	str := ""

	for _, col := range columns {
		occurrence := map[rune]int{}
		// asd as
		for _, letter := range letters {
			n := strings.Count(string(col[:]), string(letter))
			occurrence[letter] = n
		}
		str += string(findLeastCommon(occurrence))
	}

	fmt.Println("Part2: ", str)

}

func findMostCommon(occurrences map[rune]int) rune {
	maxV := 0
	char := '0'

	for key, value := range occurrences {
		if value > maxV {
			maxV = value
			char = key
		}
	}
	return char
}

func findLeastCommon(occurrences map[rune]int) rune {
	minV := math.MaxInt32
	char := '0'

	for key, value := range occurrences {
		if value < minV {
			minV = value
			char = key
		}
	}
	return char
}
