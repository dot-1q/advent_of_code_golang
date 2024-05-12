package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	puzzle := "hxtvlmkl"
	grid := [128][128]rune{}

	for i := range 128 {
		str := fmt.Sprintf("%s-%d", puzzle, i)
		hash := knotHash(str)
		repr := binaryRepr(hash)
		insertInGrid(&grid, i, repr)
	}
	// printGrid(&grid)
	fmt.Printf("Part 1 | Used squares: %d\n", usedSquares(&grid))
	fmt.Printf("Part 2 | Islands: %d\n", numberOfIslands(grid))
}

func binaryRepr(input string) string {
	hash := strings.Builder{}
	for _, char := range strings.Split(input, "") {
		bits, _ := strconv.ParseInt(char, 16, 32)
		hash.WriteString(fmt.Sprintf("%04b", bits))
	}
	return hash.String()
}

func insertInGrid(grid *[128][128]rune, row int, hash string) {
	for idx, rune := range hash {
		if rune == '0' {
			grid[row][idx] = '.'
		} else {
			grid[row][idx] = '#'
		}
	}
}

func knotHash(input string) string {
	l := input
	n := strings.TrimSpace(string(l))
	numbers := getASCII(n)
	array := array()
	index := 0
	skip := 0

	for range 64 {
		for _, n := range numbers {
			reverse(&array, index, n)
			index += (n + skip)
			skip++
		}
	}

	denseHash := []int{}
	for i := 0; i < 16; i++ {
		result := 0
		for j := i * 16; j < (i+1)*16; j++ {
			result ^= array[j]
		}
		denseHash = append(denseHash, result)
	}

	var hexdHash string
	for _, dense := range denseHash {
		// use %x to get hexadecimal version & 02 ensures leading 0 if needed
		hexdHash += fmt.Sprintf("%02x", dense)
	}
	return hexdHash
}

func array() [256]int {
	array := [256]int{}
	for i := range array {
		array[i] = i
	}
	return array
}

func getASCII(numbers string) []int {
	array := []int{}
	concat := []int{17, 31, 73, 47, 23}
	for _, rune := range numbers {
		array = append(array, int(rune))
	}
	return slices.Concat(array, concat)
}

func reverse(array *[256]int, index, number int) {
	sequence := make([]int, number)
	// Get the sequence to reverse
	i := 0
	for i < number {
		sequence[i] = array[(index+i)%len(array)]
		i++
	}
	// Reverse it
	slices.Reverse(sequence)
	// Apply it
	for _, n := range sequence {
		array[index%len(array)] = n
		index++
	}
}

func printGrid(grid *[128][128]rune) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

func usedSquares(grid *[128][128]rune) int {
	sum := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == '#' {
				sum++
			}
		}
	}
	return sum
}

func numberOfIslands(grid [128][128]rune) int {
	sum := 0
	visited := [128][128]int{}
	directions := [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	for y, row := range grid {
		for x := range row {
			if dfs(grid, &visited, x, y, directions) {
				sum++
			}
		}
	}
	return sum
}

func dfs(grid [128][128]rune, visited *[128][128]int, x, y int, directions [4][2]int) bool {
	if !valid(x, y) {
		return false
	}
	if visited[y][x] == 1 {
		return false
	}
	visited[y][x] = 1

	if grid[y][x] == '.' {
		return false
	}

	// Every direction
	for _, dir := range directions {
		newY := y + dir[0]
		newX := x + dir[1]
		dfs(grid, visited, newX, newY, directions)
	}
	return true
}

func valid(x, y int) bool {
	if y < 0 || x < 0 || y > 127 || x > 127 {
		return false
	}
	return true
}
