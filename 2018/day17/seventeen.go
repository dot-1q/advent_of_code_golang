package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	st "github.com/emirpasic/gods/stacks/linkedliststack"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	floor := floor()
	maxY, minY := CreateReservoirs(lines, &floor)
	floor[0][500] = '+'

	simulate(&floor)
	// for y := 1760; y < 1800; y++ {
	// 	for x := 370; x < 420; x++ {
	// 		fmt.Printf("%c", floor[y][x])
	// 	}
	// 	fmt.Println()
	// }
	// print2DSliceToFile(floor, "out.txt")
	fmt.Println("Part 1 | Water: ", countWater(floor, minY, maxY))
	fmt.Println("Part 2 | Water: ", countStillWater(floor, minY, maxY))
}

func simulate(floor *[2000][2000]rune) {
	st := st.New()

	st.Push([2]int{1, 500}) // First drop of water
	seen := [][2]int{}
	for !st.Empty() {
		p, _ := st.Pop()
		position := p.([2]int)
		// Go downwards and append every position to be checked later.
		pour(floor, position[0], position[1], st, &seen)
		// Backtrack and fill sideways the enclosed positions in the stack.
		fillSideways(floor, st)
	}
}

// Flow downwards and create a water tile every time we can
func pour(floor *[2000][2000]rune, y, x int, st *st.Stack, seen *[][2]int) {
	if y >= len(floor) {
		return
	}
	if floor[y][x] == '#' {
		return
	}
	*seen = append(*seen, [2]int{y, x})

	if floor[y][x] != '~' { // This is an edge case on my input, disregard.
		floor[y][x] = '|'
	}

	if !slices.Contains(*seen, [2]int{y + 1, x}) {
		st.Push([2]int{y, x}) // Append this current position to check later if we need to go sideways or down.
		pour(floor, y+1, x, st, seen)
	}
}

// Flow to the right and left until the ground is open, or we hit clay.
func overflow(floor *[2000][2000]rune, st *st.Stack, y, x int) {
	// Out of bounds check
	if y+1 >= len(floor) {
		return
	}

	for i := x; i > 0; i-- {
		// Stop when we find the wall
		if floor[y][i] == '#' {
			break
		}
		// The first position to the left which has an open bttom, it will pour from there.
		if floor[y+1][i] == '.' {
			st.Push([2]int{y, i})
			break
		}
		// Just checking if there is already water flowing, so we dont have two streams.
		if floor[y+1][i] == '|' {
			break
		}
		// Fill to the Left
		floor[y][i] = '|'
	}
	// Fill to the Right
	for i := x; i < len(floor)-1; i++ {
		// Stop when we find the wall
		if floor[y][i] == '#' {
			break
		}
		// Just checking if there is already water flowing, so we dont have two streams.
		if floor[y+1][i] == '|' {
			break
		}
		// The first position to the right which has an open bttom, it will pour from there.
		if floor[y+1][i] == '.' {
			st.Push([2]int{y, i})
			break
		}
		floor[y][i] = '|'
	}
}

// When we reach a settlement, we fill the sides.
func fillSideways(floor *[2000][2000]rune, st *st.Stack) {
	for {
		p, _ := st.Pop()
		position := p.([2]int)
		if isEnclosed(floor, position) {
			floor[position[0]][position[1]] = '~'
			for x := position[1]; x >= 0; x-- {
				// Stop when we find the wall
				if floor[position[0]][x] == '#' {
					break
				}
				// Fill to the Left
				floor[position[0]][x] = '~'
			}
			// Fill to the Right
			for x := position[1]; x <= len(floor); x++ {
				// Stop when we find the wall
				if floor[position[0]][x] == '#' {
					break
				}
				floor[position[0]][x] = '~'
			}
		} else { // When the first position in the stack that is not enclosed, it overflows to the sides.
			overflow(floor, st, position[0], position[1])
			break
		}
	}
}

// Figure out if a position is enclosed by clay, so that it can hold water.
func isEnclosed(floor *[2000][2000]rune, position [2]int) bool {
	// Enclosed to the left
	enclosed := true
	// Out of bounds check
	if position[0]+1 >= len(floor) {
		return false
	}
	for x := position[1]; x > 0; x-- {
		// If the floor of the position to the left is not clay, it is not enclosed.
		if floor[position[0]+1][x] == '.' {
			enclosed = false
		}
		// Stop when we find the wall
		if floor[position[0]][x] == '#' {
			break
		}
	}
	// Enclosed to the right
	for x := position[1]; x < len(floor); x++ {
		if floor[position[0]+1][x] == '.' {
			enclosed = false
		}
		// Stop when we find the wall
		if floor[position[0]][x] == '#' {
			break
		}
	}
	return enclosed
}

func CreateReservoirs(lines []string, floor *[2000][2000]rune) (int, int) {
	maxY := 0
	minY := 1000
	for _, line := range lines {
		coordinates := strings.Split(line, ",")
		coord1 := strings.Split(coordinates[0], "=")
		coord2 := strings.Split(strings.TrimSpace(coordinates[1]), "=")
		ranges := strings.Split(coord2[1], "..")
		position, _ := strconv.Atoi(coord1[1])
		start, _ := strconv.Atoi(ranges[0])
		end, _ := strconv.Atoi(ranges[1])

		switch coord1[0] {
		case "x": // Fill in the Y position
			for i := start; i <= end; i++ {
				floor[i][position] = '#'
			}
			maxY = max(maxY, end)
			minY = min(minY, start)
		case "y": // Fill in the X position
			for i := start; i <= end; i++ {
				floor[position][i] = '#'
			}
			maxY = max(maxY, position)
			minY = min(minY, position)
		}
	}
	return maxY, minY
}

func floor() [2000][2000]rune {
	dots := [2000][2000]rune{}
	for i := range dots {
		dots[i] = [2000]rune{}
		for j := range dots[i] {
			dots[i][j] = '.'
		}
	}
	return dots
}

func print2DSliceToFile(slice2D [2000][2000]rune, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, row := range slice2D {
		for _, col := range row {
			file.WriteString(string(col))
		}
		file.WriteString("\n")
	}
}

func countWater(floor [2000][2000]rune, minY, maxY int) int {
	s := 0
	for row := minY; row <= maxY; row++ {
		for _, char := range floor[row] {
			if char == '~' || char == '|' {
				s++
			}
		}
	}
	return s
}

func countStillWater(floor [2000][2000]rune, minY, maxY int) int {
	s := 0
	for row := minY; row <= maxY; row++ {
		for _, char := range floor[row] {
			if char == '~' {
				s++
			}
		}
	}
	return s
}
