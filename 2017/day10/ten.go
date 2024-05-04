package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}
func partOne() {
	l, _ := os.ReadFile("input.txt")
	n := strings.Split(string(strings.TrimSpace(string(l))), ",")
	array := array()

	index := 0
	skip := 0
	for _, n := range n {
		number, _ := strconv.Atoi(n)
		reverse(&array, index, number)
		index += (number + skip)
		skip++
	}
	fmt.Printf("Part1: %d\n", array[0]*array[1])
}

func partTwo() {
	l, _ := os.ReadFile("input.txt")
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
	fmt.Printf("Part2 : %s\n", hexdHash)
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
