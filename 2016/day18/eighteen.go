package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	line, _ := os.ReadFile("input.txt")
	row, _ := strings.CutSuffix(string(line), "\n")
	part := 2
	nRows := 40
	if part == 2 {
		nRows = 400000
	}

	floor := []string{}
	floor = append(floor, row)
	for range nRows - 1 {
		row = createNextRow(row)
		floor = append(floor, row)
	}

	// for _, f := range floor {
	// 	fmt.Printf("%+v\n", f)
	// }
	fmt.Printf("Part %d | There's %d safe tiles\n", part, countSafe(floor))
}

func createNextRow(row string) string {
	newRow := strings.Builder{}
	// Check the very first tile of each row, given that its an edge case
	// From the rules, we can deduce that if the right tile is a trap, this next tile is always a trap, and if not, its safe
	if row[1] == '^' {
		newRow.WriteRune('^')
	} else {
		newRow.WriteRune('.')
	}
	// For the middle tiles, check them normally
	// We can also deduce that if the left and right previous tiles are the same, the next tile is *ALWAYS* a safe tile, because none of the trap rules apply.
	// If they differ, we can be sure that it is *ALWAYS* a trap, given that one of them is a trap, and the center tile will not matter, because
	// one of the 2 rules will apply, with center tile or not.
	for i := 1; i < len(row)-1; i++ {
		if row[i-1] != row[i+1] {
			newRow.WriteRune('^')
		} else {
			newRow.WriteRune('.')
		}
	}
	// Again, for the very last tile, given that its an edge case, determine it the same as the first tile.
	// In this case, for the right most tile, if its previous left tile is a trap, it is necessarily a trap.
	if row[len(row)-2] == '^' {
		newRow.WriteRune('^')
	} else {
		newRow.WriteRune('.')
	}

	return newRow.String()
}

func countSafe(floor []string) int {
	count := 0
	for _, f := range floor {
		count += strings.Count(f, ".")
	}
	return count
}
