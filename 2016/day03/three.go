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
	f.Seek(0, 0)
	partTwo(f)

}

func partOne(f *os.File) {
	scanner := bufio.NewScanner(f)
	validTriangles := 0
	for scanner.Scan() {
		line := scanner.Text()
		sides := strings.Fields(strings.TrimSpace(line))

		one, _ := strconv.Atoi(strings.TrimSpace(sides[0]))
		two, _ := strconv.Atoi(strings.TrimSpace(sides[1]))
		three, _ := strconv.Atoi(strings.TrimSpace(sides[2]))
		if (one+two > three) && (two+three > one) && (one+three > two) {
			validTriangles++
		}
	}
	fmt.Println("Part1: ", validTriangles)
}

func partTwo(f *os.File) {
	scanner := bufio.NewScanner(f)
	validTriangles := 0
	sideOne := []int{}
	sideTwo := []int{}
	sideThree := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		sides := strings.Fields(strings.TrimSpace(line))

		// Separate each columns into different arrays
		one, _ := strconv.Atoi(strings.TrimSpace(sides[0]))
		two, _ := strconv.Atoi(strings.TrimSpace(sides[1]))
		three, _ := strconv.Atoi(strings.TrimSpace(sides[2]))
		sideOne = append(sideOne, one)
		sideTwo = append(sideTwo, two)
		sideThree = append(sideThree, three)
	}

	i := 0
	for (i + 2) < len(sideOne) {
		one := sideOne[i]
		two := sideOne[i+1]
		three := sideOne[i+2]
		if (one+two > three) && (two+three > one) && (one+three > two) {
			validTriangles++
		}
		one = sideTwo[i]
		two = sideTwo[i+1]
		three = sideTwo[i+2]
		if (one+two > three) && (two+three > one) && (one+three > two) {
			validTriangles++
		}
		one = sideThree[i]
		two = sideThree[i+1]
		three = sideThree[i+2]
		if (one+two > three) && (two+three > one) && (one+three > two) {
			validTriangles++
		}
		i += 3
	}
	fmt.Println("Part2: ", validTriangles)
}
