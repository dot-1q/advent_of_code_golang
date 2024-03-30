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
	scanner := bufio.NewScanner(f)

	instructions := []string{}
	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	screen := partOne(instructions)

	fmt.Println("Part2 Below:")
	printChars(&screen)

}

func partOne(instructions []string) [6][50]int {
	// 50 pixels wide and 6 pixels tall
	screen := [6][50]int{}

	for _, inst := range instructions {
		s := strings.Split(inst, " ")

		// Separate instructions
		switch s[0] {
		case "rotate":
			n, _ := strconv.Atoi(s[4])
			rotateScreen(&screen, s[2], n)
		case "rect":
			dim := strings.Split(s[1], "x")
			// Its inverted
			col, _ := strconv.Atoi(dim[0])
			row, _ := strconv.Atoi(dim[1])
			rect(&screen, row, col)
		}
	}
	for _, r := range screen {
		for _, c := range r {
			fmt.Printf("%d", c)
		}
		fmt.Println()
	}
	fmt.Println("Part1: Lit pixels: ", countLitPixels(&screen))
	return screen
}

// Rotate the screen a N number of times. The line argument differentiates between row and column, via y or x
// place = "x=30" or "y=30" or some variation of that
func rotateScreen(screen *[6][50]int, place string, number int) {
	location := strings.Split(place, "=")

	switch location[0] {
	// Column
	case "x":
		column, _ := strconv.Atoi(location[1])
		// Rotating a pixel by a number N, essentially means moving that pixel N number of place to the left
		// We just have to keep in mind out of bounds, which we can do via modulo
		// Copy the original column, so that the values aren+t lost on rotation
		original_column := copyColumn(screen, column)
		// Iterate over all the pixels of said column and move them
		for r, _ := range screen {
			// new position for this pixel is the current one plus the rotation. % 50 so we dont go out of bounds
			newPos := (r + number) % len(screen)
			originalValue := original_column[r]
			(*screen)[newPos][column] = originalValue
		}
	// Row
	case "y":
		row, _ := strconv.Atoi(location[1])
		// Rotating a pixel by a number N, essentially means moving that pixel N number of place to the left
		// We just have to keep in mind out of bounds, which we can do via modulo
		// Create an empty final column, which will hold the final values of the rotation. We dont want to override the original
		original_row := copyRow(screen, row)
		// Iterate over all the pixels of said row and move them
		for pos, _ := range screen[row] {
			// new position for this pixel is the current one plus the rotation. % 50 so we dont go out of bounds
			newPos := (pos + number) % len(screen[row])
			originalValue := original_row[pos]
			screen[row][newPos] = originalValue
		}
	}
}

// Create a rectangle in the screen
func rect(screen *[6][50]int, row, col int) {
	for r := range row {
		for c := range col {
			screen[r][c] = 1
		}
	}
}

func copyColumn(screen *[6][50]int, col int) [6]int {
	column := [6]int{}
	for i, row := range screen {
		column[i] = row[col]
	}
	return column
}

func copyRow(screen *[6][50]int, row int) [50]int {
	r := [50]int{}
	for i, cell := range screen[row] {
		r[i] = cell
	}
	return r
}

func countLitPixels(screen *[6][50]int) int {
	count := 0
	for _, row := range screen {
		for _, pixel := range row {
			if pixel == 1 {
				count++
			}
		}
	}
	return count
}

func printChars(screen *[6][50]int) {
	for _, r := range screen {
		for _, c := range r {
			// each char is 6 letters wide
			if c == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf("-")
			}
		}
		fmt.Println()
	}
}
