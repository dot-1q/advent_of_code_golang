package main

import (
	"crypto/md5"
	"fmt"
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

type MazeState struct {
	path     string
	start    string
	position []int
}

func main() {
	puzzle := "vwbaicqe"
	solutions := bfs(puzzle)
	fmt.Printf("Part1 : Shortest Path is: %s\n", solutions[0].path)
	fmt.Printf("Part2 : Longest Path has length: %d\n", len(solutions[len(solutions)-1].path))
}

func bfs(start string) []MazeState {
	position := []int{0, 0}
	queue := pq.NewWith(priority)
	// Enqueue the start state, which is the puzzle input with no moves yet
	queue.Enqueue(MazeState{start: start, path: "", position: position})
	solutions := []MazeState{}

	for !queue.Empty() {
		s, _ := queue.Dequeue()
		maze := s.(MazeState)
		if maze.position[0] == 3 && maze.position[1] == 3 {
			// Found the exit. Append to all valid solutions.
			// This is needed for part 1
			solutions = append(solutions, maze)
			// If we'e reached the end, dont continue to iterate, since we would've been going in circles
			continue
		}
		status := getDoorStatus(maze.start + maze.path)
		moves := getValidMoves(status, maze.position)
		// For each valid move, append it to the queue to check its path later
		for _, move := range moves {
			newPos := []int{}
			switch move {
			case "U":
				newPos = append(newPos, maze.position[0])
				newPos = append(newPos, maze.position[1]-1)
			case "D":
				newPos = append(newPos, maze.position[0])
				newPos = append(newPos, maze.position[1]+1)
			case "L":
				newPos = append(newPos, maze.position[0]-1)
				newPos = append(newPos, maze.position[1])
			case "R":
				newPos = append(newPos, maze.position[0]+1)
				newPos = append(newPos, maze.position[1])
			}
			newPath := maze.path + move
			newMS := MazeState{start: start, position: newPos, path: newPath}
			queue.Enqueue(newMS)
		}
	}
	return solutions
}

// Generate the has and return the first four characters, which represent the state of each door
func getDoorStatus(input string) string {
	hash := md5.Sum([]byte(input))
	return fmt.Sprintf("%x", hash)[:4]
}

func getValidMoves(status string, position []int) []string {
	// The first for chars of the hash represent directions up/down/left/right
	// From those chars, return the list of valid moves (U,D,L,R)
	moves := []string{}
	dirs := []string{"U", "D", "L", "R"}
	for i := 0; i < 4; i++ {
		// Since status[i] returns a rune, which is an int32, we can check if its "bcdef" if its between the range [98-102]
		if status[i] >= 98 && status[i] <= 102 {
			// For the valid moves from the hash, check if they are truly valid when in position.
			// i.e, check if we dont go to out of bounds positions. The length of the grid is hard-coded (4x4)
			switch dirs[i] {
			case "U":
				if position[1]-1 >= 0 {
					moves = append(moves, dirs[i])
				}
			case "D":
				if position[1]+1 <= 3 {
					moves = append(moves, dirs[i])
				}
			case "L":
				if position[0]-1 >= 0 {
					moves = append(moves, dirs[i])
				}
			case "R":
				if position[0]+1 <= 3 {
					moves = append(moves, dirs[i])
				}
			}
		}
	}
	return moves
}

// Ascending priority, meaning the lower lens are prioritised
func priority(a, b interface{}) int {
	priorityA := len(a.(MazeState).path)
	priorityB := len(b.(MazeState).path)
	return utils.IntComparator(priorityA, priorityB)
}
