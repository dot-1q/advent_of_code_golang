package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// partOne()
	partTwo()
}

func partOne() {
	area := area()

	for row := range area {
		fmt.Printf("%c", area[row])
		fmt.Println()
	}

	for range 10 {
		area = changeState(area)
	}

	fmt.Println("After")
	for row := range area {
		fmt.Printf("%c", area[row])
		fmt.Println()
	}

	wood := countChar2D(area, '|')
	lumber := countChar2D(area, '#')

	fmt.Printf("Part 1 : Total resource value: %d\n", wood*lumber)

}

func partTwo() {
	area := area()
	lastGen := 0
	visited := map[[50][50]rune]int{}
	// Get the wood and lumber counts for each generation that we're testing to see
	// if we can discover a pattern
	for minutes := 0; minutes < 1000000000; minutes++ {
		area = changeState(area)

		if visited[area] != 0 {
			firstOccurrence := visited[area]
			period := minutes - firstOccurrence // Repeats after X steps
			fmt.Println("Period:", period)
			// Skip the minutes ahead
			minutes = ((1000000000-firstOccurrence)/period)*period + firstOccurrence
		}

		wood := countChar2D(area, '|')
		lumber := countChar2D(area, '#')
		resources := wood * lumber
		fmt.Printf("For %d Minutes, Total resource value: %d | Diff to previous: %d\n", minutes+1, resources, resources-lastGen)
		lastGen = resources

		visited[area] = minutes
	}
	// The cycle repeats every 28 iterations. The cycle is as follows:
	// For 49972 Minutes, Total resource value: 210824 | Diff to previous: 320
	// For 49973 Minutes, Total resource value: 207282 | Diff to previous: -3542
	// For 49974 Minutes, Total resource value: 205320 | Diff to previous: -1962
	// For 49975 Minutes, Total resource value: 204125 | Diff to previous: -1195
	// For 49976 Minutes, Total resource value: 197316 | Diff to previous: -6809
	// For 49977 Minutes, Total resource value: 192984 | Diff to previous: -4332
	// For 49978 Minutes, Total resource value: 188914 | Diff to previous: -4070
	// For 49979 Minutes, Total resource value: 181485 | Diff to previous: -7429
	// For 49980 Minutes, Total resource value: 177416 | Diff to previous: -4069
	// For 49981 Minutes, Total resource value: 173910 | Diff to previous: -3506
	// For 49982 Minutes, Total resource value: 167475 | Diff to previous: -6435
	// For 49983 Minutes, Total resource value: 164424 | Diff to previous: -3051
	// For 49984 Minutes, Total resource value: 164079 | Diff to previous: -345
	// For 49985 Minutes, Total resource value: 163631 | Diff to previous: -448
	// For 49986 Minutes, Total resource value: 163248 | Diff to previous: -383
	// For 49987 Minutes, Total resource value: 167090 | Diff to previous: 3842
	// For 49988 Minutes, Total resource value: 168562 | Diff to previous: 1472
	// For 49989 Minutes, Total resource value: 171588 | Diff to previous: 3026
	// For 49990 Minutes, Total resource value: 172852 | Diff to previous: 1264
	// For 49991 Minutes, Total resource value: 174900 | Diff to previous: 2048
	// For 49992 Minutes, Total resource value: 176012 | Diff to previous: 1112
	// For 49993 Minutes, Total resource value: 182574 | Diff to previous: 6562
	// For 49994 Minutes, Total resource value: 187272 | Diff to previous: 4698
	// For 49995 Minutes, Total resource value: 193888 | Diff to previous: 6616
	// For 49996 Minutes, Total resource value: 199167 | Diff to previous: 5279
	// For 49997 Minutes, Total resource value: 203648 | Diff to previous: 4481
	// For 49998 Minutes, Total resource value: 204832 | Diff to previous: 1184
	// For 49999 Minutes, Total resource value: 210504 | Diff to previous: 5672
	// ...
	// For 100000 Minutes, Total resource value: 176012 | Diff to previous: 1112
	// Since 1000000000-100000 = 999900000
	// 999900000 % 28 = 8
	// We know that at 1 bilion, the iteration will be the 8th, from where the 100000 is currently at.
	// So its the one with the Diff = 320
	// The above was my first solution, which was by hand. I later added a programmatic one, which skips steps ahead.

	wood := countChar2D(area, '|')
	lumber := countChar2D(area, '#')
	fmt.Printf("Part 2 : Total resource value: %d\n", wood*lumber)
}

func changeState(area [50][50]rune) [50][50]rune {
	deepCopy := deepCopy(area) // We need an exact copy, so we dont alter the original adjacent values while iterating

	for row := range area {
		for cell, space := range area[row] {
			adjacents := adjacentSquares(area, row, cell) // Get the adjancets from the original
			switch space {
			case '.':
				if n := countChar(adjacents, '|'); n >= 3 {
					deepCopy[row][cell] = '|'
				}
			case '|':
				if n := countChar(adjacents, '#'); n >= 3 {
					deepCopy[row][cell] = '#'
				}
			case '#':
				l := countChar(adjacents, '#')
				t := countChar(adjacents, '|')
				if l == 0 || t == 0 {
					deepCopy[row][cell] = '.'
				}
			}
		}
	}

	return deepCopy
}

func adjacentSquares(area [50][50]rune, row, cell int) []rune {
	a := []rune{}

	// Go along all the adjacent squares. be mindfull of edges.
	for i := row - 1; i <= row+1; i++ {
		if i < 0 || i >= len(area) { // Row edges, top and bottom
			continue
		}
		for j := cell - 1; j <= cell+1; j++ {
			if j < 0 || j >= len(area) { // Column edges, left and right
				continue
			}
			if i == row && j == cell { // Dont append the current character were checking
				continue
			}
			a = append(a, area[i][j]) // Append the rest
		}
	}
	return a
}

// Count occurrences of any given char.
func countChar(array []rune, char rune) int {
	c := 0
	for i := range array {
		if array[i] == char {
			c++
		}
	}
	return c
}

// Count occurrences of any given char.
func countChar2D(array [50][50]rune, char rune) int {
	c := 0
	for i := range array {
		for j := range array[i] {
			if array[i][j] == char {
				c++
			}
		}
	}
	return c
}

func area() [50][50]rune {
	area := [50][50]rune{}
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	for i, line := range lines {
		characters := []rune(line)
		copy(area[i][:], characters)
	}

	return area
}

func deepCopy(original [50][50]rune) [50][50]rune {
	c := [50][50]rune{}

	for i := range original {
		copy(c[i][:], original[i][:])
	}

	return c
}
