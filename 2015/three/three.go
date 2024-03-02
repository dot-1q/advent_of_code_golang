package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

func main() {
	// Houses that got more than one present
	diff_houses := 1
	// Map each move to a cartesian coordinate, and store it in a map,
	// so we can efficiently check if we've visited that house before or not
	visited := make(map[point]int)
	// Starting position
	position := point{0, 0}
	// Mark it as visited
	visited[position] = 1

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewReader(f)

	for {
		move, _, e := scanner.ReadRune()
		if e != nil {
			// end of file
			break
		}

		// For each move, calculate its new position, mark it has visited if it has not been yet visited
		// Also, if it has not been visited, increase the number of different houses visited
		switch move {
		case '>':
			position.x += 1
			if visited[position] == 0 {
				visited[position] = 1
				diff_houses += 1
			}
		case '<':
			position.x -= 1
			// Meaning we havent visited
			if visited[position] == 0 {
				visited[position] = 1
				diff_houses += 1
			}
		case '^':
			position.y += 1
			if visited[position] == 0 {
				visited[position] = 1
				diff_houses += 1
			}
		case 'v':
			position.y -= 1
			if visited[position] == 0 {
				visited[position] = 1
				diff_houses += 1
			}
		}
	}

	fmt.Printf("Part1: Number of houses that received at least one present: %d\n", diff_houses)

	// Start part two
	parttwo()

}

func parttwo() {
	// Houses that got more than one present
	diff_houses := 1
	// Map each move to a cartesian coordinate, and store it in a map,
	// so we can efficiently check if we've visited that house before or not
	visited := make(map[point]int)
	// Starting position
	santa_position := point{0, 0}
	robo_position := point{0, 0}
	// Mark it as visited
	visited[santa_position] = 1
	visited[robo_position] = 1

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewReader(f)

	// Flag to differentiate between santa's and robots turn (Santa = True, Robot = False)
	santaOrRobot := true
	for {
		move, _, e := scanner.ReadRune()
		if e != nil {
			// end of file
			break
		}

		// For each move, calculate its new position, mark it has visited if it has not been yet visited
		// Also, if it has not been visited, increase the number of different houses visited
		switch move {
		case '>':
			if santaOrRobot {
				santa_position.x += 1
				if visited[santa_position] == 0 {
					visited[santa_position] = 1
					diff_houses += 1
				}
				santaOrRobot = !santaOrRobot
			} else {
				robo_position.x += 1
				if visited[robo_position] == 0 {
					visited[robo_position] = 1
					diff_houses += 1
				}
				santaOrRobot = !santaOrRobot
			}
		case '<':
			if santaOrRobot {
				santa_position.x -= 1
				// Meaning we havent visited
				if visited[santa_position] == 0 {
					visited[santa_position] = 1
					diff_houses += 1
				}
				santaOrRobot = !santaOrRobot
			} else {
				robo_position.x -= 1
				if visited[robo_position] == 0 {
					visited[robo_position] = 1
					diff_houses += 1
				}
				santaOrRobot = !santaOrRobot
			}
		case '^':
			if santaOrRobot {
				santa_position.y += 1
				if visited[santa_position] == 0 {
					visited[santa_position] = 1
					diff_houses += 1
				}
				santaOrRobot = !santaOrRobot
			} else {
				robo_position.y += 1
				if visited[robo_position] == 0 {
					visited[robo_position] = 1
					diff_houses += 1
				}
				santaOrRobot = !santaOrRobot
			}
		case 'v':
			if santaOrRobot {
				santa_position.y -= 1
				if visited[santa_position] == 0 {
					visited[santa_position] = 1
					diff_houses += 1
				}
				santaOrRobot = !santaOrRobot
			} else {
				robo_position.y -= 1
				if visited[robo_position] == 0 {
					visited[robo_position] = 1
					diff_houses += 1
				}
				santaOrRobot = !santaOrRobot
			}
		}
	}
	fmt.Printf("Part2: Number of houses that received at least one present: %d\n", diff_houses)
}
