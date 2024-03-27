package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, _ := os.Open("input.txt")
	partOne(f)
	// Reset file pointer
	f.Seek(0, 0)
	partTwo(f)

}
func partOne(file *os.File) {
	scanner := bufio.NewScanner(file)
	// Add a padding surrounding the keypad, that represent invalid movements
	keypad := [][]int{
		{0, 0, 0, 0, 0},
		{0, 1, 2, 3, 0},
		{0, 4, 5, 6, 0},
		{0, 7, 8, 9, 0},
		{0, 0, 0, 0, 0}}
	// We start in the middle of the keypad
	lastNumber := 5
	code := strings.Builder{}
	for scanner.Scan() {
		line := scanner.Text()
		lastNumber = findNumber(lastNumber, []rune(line), &keypad)
		code.WriteString(strconv.Itoa(lastNumber))
	}
	fmt.Println("Part1: ", code.String())
}

func partTwo(file *os.File) {
	scanner := bufio.NewScanner(file)
	// Golang, runes are treated as int32, so valid
	// Add a padding surrounding the keypad, that represent invalid movements
	keypad := [][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0},
		{0, 0, 2, 3, 4, 0, 0},
		{0, 5, 6, 7, 8, 9, 0},
		{0, 0, 'A', 'B', 'C', 0, 0},
		{0, 0, 0, 'D', 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0}}

	// We start in the middle of the keypad
	lastNumber := 5
	code := strings.Builder{}
	for scanner.Scan() {
		line := scanner.Text()
		lastNumber = findNumber(lastNumber, []rune(line), &keypad)
		// If it returns a number bigger than 65, its an ascii repr of the letter A,B,C or D
		if lastNumber >= 65 {
			code.WriteString((string(lastNumber)))
		} else {
			code.WriteString(strconv.Itoa(lastNumber))
		}
	}
	fmt.Println("Part2: ", code.String())
}

// Specific to this annormal keypad
func findNumber(start int, moves []rune, keypad *[][]int) int {
	row, col := 0, 0
	// Translate number to keypad position
	for i, r := range *keypad {
		for j, number := range r {
			if number == start {
				row, col = i, j
			}
		}
	}

	for _, move := range moves {
		switch move {
		case 'U':
			if (*keypad)[row-1][col] != 0 {
				row--
			}
		case 'D':
			if (*keypad)[row+1][col] != 0 {
				row++
			}
		case 'R':
			if (*keypad)[row][col+1] != 0 {
				col++
			}
		case 'L':
			if (*keypad)[row][col-1] != 0 {
				col--
			}
		}
	}
	return (*keypad)[row][col]
}
