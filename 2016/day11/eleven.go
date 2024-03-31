package main

import (
	"fmt"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	utils "github.com/emirpasic/gods/utils"
)

type Floor struct {
	Level int
	// There's at most 5 elements. Just to optimize space complexity when working on the solution
	Generators []string
	Chips      []string
}

type GameState struct {
	Floor    []Floor
	Turns    int
	Elevator int
}

func main() {
	fmt.Println("Hi")

	// Make an array of floors
	floors := []Floor{}
	// Input the data manually. Don't feel like parsing 4 lines tbh
	floors = append(floors,
		Floor{Level: 0,
			Generators: []string{"thulium", "plutonium", "strontium"},
			Chips:      []string{"thulium"}},

		Floor{Level: 1,
			Generators: []string{},
			Chips:      []string{"plutonium,strontium"}},
		Floor{Level: 2,
			Generators: []string{"promethium", "ruthenium"},
			Chips:      []string{"promethium", "ruthenium"}},
		Floor{Level: 3,
			Generators: []string{},
			Chips:      []string{}})

	partOne(floors)
}

func partOne(floors []Floor) {
	pqueue := pq.NewWith(byPriority)

	f := []Floor{}
	copy(f, floors)
	gs := GameState{Floor: f, Turns: 0, Elevator: 0}
	pqueue.Enqueue(gs)
	for {
		i, _ := pqueue.Dequeue()
		gs := i.(GameState)

		nextState := generateNextState(gs)
		for _, gs := range nextState {
			fmt.Println(gs)
		}
	}
}

func isFinal(gamestate GameState) {

}

func generateNextState(gamestate GameState) []GameState {
	return []GameState{}
}

func byPriority(a, b interface{}) int {
	// The priority is the number of elements in the last floor. more is better
	priorityA := len(a.(GameState).Floor[3].Chips) + len(a.(GameState).Floor[3].Generators)
	priorityB := len(b.(GameState).Floor[3].Chips) + len(b.(GameState).Floor[3].Generators)
	return -utils.IntComparator(priorityA, priorityB) // "-" descending order
}
