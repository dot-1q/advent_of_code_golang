package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func main() {
	const puzzle = 503761
	recipes := [50000000]int{}
	recipes[0] = 3
	recipes[1] = 7
	elf1, elf2 := 0, 1
	index := 1

	// We want the recipes that are 10 after the puzzle
	for index+1 < len(recipes) {
		sum := recipes[elf1] + recipes[elf2]
		repr := strings.Split(strconv.Itoa(sum), "") // Convert the int to string so we can easily seperate the numbers
		digit1, _ := strconv.Atoi(repr[0])
		index++
		recipes[index] = digit1
		if len(repr) == 2 {
			digit2, _ := strconv.Atoi(repr[1])
			index++
			recipes[index] = digit2
		}
		elf1 = (1 + recipes[(elf1)] + elf1) % (index + 1) // Calculate the new elf positions
		elf2 = (1 + recipes[(elf2)] + elf2) % (index + 1) // Its 1 plus their recipe value, plus its current position. % for the bounds. Index +1 holds how many items we've input so far.
	}
	answer := getRecipes(recipes, puzzle)
	fmt.Printf("Part 1 | The Scores of the 10 recipes after are %s\n", answer)
	numbers := toInt(strconv.Itoa(puzzle))
	fmt.Printf("Part 2 | Numbers left of the puzzle sequence %d\n", findNumbersLeft(recipes, numbers))
}

func getRecipes(recipe [50000000]int, puzzle int) string {
	str := strings.Builder{}
	for i := range 10 {
		str.WriteString(strconv.Itoa(recipe[puzzle+i]))
	}
	return str.String()
}

func findNumbersLeft(recipe [50000000]int, number []int) int {
	for i := 0; i < len(recipe)-5; i++ {
		// Find a subslice that equals our puzzle input
		if slices.Equal(recipe[i:i+6], number) {
			return i
		}
	}
	return 0
}

func toInt(number string) []int {
	n := []int{}
	numbers := strings.Split(number, "")

	for _, v := range numbers {
		val, _ := strconv.Atoi(v)
		n = append(n, val)
	}
	return n
}
