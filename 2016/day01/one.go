package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// We only go right or left
	// The arrays represent (x,y)
	facingNorth := [][]int{{1, 0}, {-1, 0}}
	facingSouth := [][]int{{-1, 0}, {1, 0}}
	facingWest := [][]int{{0, -1}, {0, 1}}
	facingEast := [][]int{{0, 1}, {0, -1}}
	pos := []int{0, 0}
	visitedPos := map[[2]int]bool{
		{0, 0}: true,
	}

	f, _ := os.ReadFile("input.txt")
	line := string(f)
	directions := strings.Split(line, ",")

	partTwo := true
	facing := 'N'
	for _, direction := range directions {
		chars := []rune(strings.TrimSpace(direction))
		switch chars[0] {
		case 'R':
			steps, _ := strconv.Atoi(string(chars[1:]))
			switch facing {
			case 'N':
				facing = 'W'
				for range steps {
					pos[0] += facingNorth[0][0]
					pos[1] += facingNorth[0][1]
					if visitedPos[[2]int{pos[0], pos[1]}] && partTwo == true {
						manhattanDistance := math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
						fmt.Println("Part2: ", manhattanDistance)
						return
					}
					visitedPos[[2]int{pos[0], pos[1]}] = true
				}
			case 'S':
				facing = 'E'
				for range steps {
					pos[0] += facingSouth[0][0]
					pos[1] += facingSouth[0][1]
					if visitedPos[[2]int{pos[0], pos[1]}] && partTwo == true {
						manhattanDistance := math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
						fmt.Println("Part2: ", manhattanDistance)
						return
					}
					visitedPos[[2]int{pos[0], pos[1]}] = true
				}
			case 'W':
				facing = 'S'
				for range steps {
					pos[0] += facingWest[0][0]
					pos[1] += facingWest[0][1]
					if visitedPos[[2]int{pos[0], pos[1]}] && partTwo == true {
						manhattanDistance := math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
						fmt.Println("Part2: ", manhattanDistance)
						return
					}
					visitedPos[[2]int{pos[0], pos[1]}] = true
				}
			case 'E':
				facing = 'N'
				for range steps {
					pos[0] += facingEast[0][0]
					pos[1] += facingEast[0][1]
					if visitedPos[[2]int{pos[0], pos[1]}] && partTwo == true {
						manhattanDistance := math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
						fmt.Println("Part2: ", manhattanDistance)
						return
					}
					visitedPos[[2]int{pos[0], pos[1]}] = true
				}
			}
		case 'L':
			steps, _ := strconv.Atoi(string(chars[1:]))
			switch facing {
			case 'N':
				facing = 'E'
				for range steps {
					pos[0] += facingNorth[1][0]
					pos[1] += facingNorth[1][1]
					if visitedPos[[2]int{pos[0], pos[1]}] && partTwo == true {
						manhattanDistance := math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
						fmt.Println("Part2: ", manhattanDistance)
						return
					}
					visitedPos[[2]int{pos[0], pos[1]}] = true
				}
			case 'S':
				facing = 'W'
				for range steps {
					pos[0] += facingSouth[1][0]
					pos[1] += facingSouth[1][1]
					if visitedPos[[2]int{pos[0], pos[1]}] && partTwo == true {
						manhattanDistance := math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
						fmt.Println("Part2: ", manhattanDistance)
						return
					}
					visitedPos[[2]int{pos[0], pos[1]}] = true
				}
			case 'W':
				facing = 'N'
				for range steps {
					pos[0] += facingWest[1][0]
					pos[1] += facingWest[1][1]
					if visitedPos[[2]int{pos[0], pos[1]}] && partTwo == true {
						manhattanDistance := math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
						fmt.Println("Part2: ", manhattanDistance)
						return
					}
					visitedPos[[2]int{pos[0], pos[1]}] = true
				}
			case 'E':
				facing = 'S'
				for range steps {
					pos[0] += facingEast[1][0]
					pos[1] += facingEast[1][1]
					if visitedPos[[2]int{pos[0], pos[1]}] && partTwo == true {
						manhattanDistance := math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
						fmt.Println("Part2: ", manhattanDistance)
						return
					}
					visitedPos[[2]int{pos[0], pos[1]}] = true
				}
			}
		}
	}

	// fmt.Println("Pos visited", visitedPos)
	manhattanDistance := math.Abs(float64(pos[0])) + math.Abs(float64(pos[1]))
	fmt.Println("Part1: ", manhattanDistance)
}
