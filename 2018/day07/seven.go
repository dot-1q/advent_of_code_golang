package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Graph map[string][]string

type Elf struct {
	letter    string
	waitTime  int
	isWorking bool
}

func main() {
	// This is a problem about topological sort. Dependency sort.
	graph, nodes := createGraph()
	noDeps := getNoDep(graph)
	fmt.Println(graph)
	final := BreadthFirstSearch(graph, nodes, noDeps)
	fmt.Printf("Part 1 | Final word: %s\n", final)
	finalString, seconds := BreadthFirstSearchWithHelp(graph, nodes, noDeps, 5)
	fmt.Printf("Part 2 | Final word: %s and took %d seconds\n", finalString, seconds)
}

func BreadthFirstSearch(graph Graph, nodes []string, startingPoints []string) string {
	queue := []string{} // empty
	seen := []string{}
	str := ""

	// Add first elements with cost zero
	queue = append(queue, startingPoints...)
	for len(queue) > 0 {
		slices.Sort(queue)
		current := dequeueValid(queue, graph, seen)[0]
		queue = remove(queue, current)
		str += current
		seen = append(seen, current)
		neighbours := getNext(graph, current)
		for _, neighbour := range neighbours {
			if !slices.Contains(queue, neighbour) && !slices.Contains(seen, neighbour) {
				queue = append(queue, neighbour)
			}
		}
	}
	return str
}

// Part 2
func BreadthFirstSearchWithHelp(graph Graph, nodes []string, startingPoints []string, workers int) (string, int) {
	queue := []string{} // empty
	seen := []string{}
	elves := createElves(workers)
	str := ""
	seconds := 0
	done := false

	// Add first elements with cost zero
	queue = append(queue, startingPoints...)
	for !done || len(queue) > 0 {
		slices.Sort(queue)
		// Only assign letters to workers that are available.
		for _, elf := range elves {
			if !elf.isWorking {
				valid := dequeueValid(queue, graph, seen)
				if len(valid) > 0 {
					elf.letter = valid[0]
					elf.waitTime = (int([]rune(valid[0])[0]) - 64) + 60 // This arithmetic is like this for readability
					elf.isWorking = true
					queue = remove(queue, valid[0])
				}
			}
		}
		tickElves(elves)
		letters := getLettersDone(elves)
		for _, letter := range letters {
			str += letter
			seen = append(seen, letter)
			// This letters that are done can now be used in the queue
			neighbours := getNext(graph, letter)
			for _, neighbour := range neighbours {
				if !slices.Contains(queue, neighbour) && !slices.Contains(seen, neighbour) {
					queue = append(queue, neighbour)
				}
			}
		}

		done = true
		for _, elf := range elves {
			if elf.isWorking {
				done = false
			}
		}
		seconds++
	}
	return str, seconds
}

func getLettersDone(elves []*Elf) []string {
	letters := []string{}
	for _, elf := range elves {
		if elf.waitTime == 0 {
			letters = append(letters, elf.letter)
			elf.isWorking = false
			elf.letter = "."
		}
	}
	return letters
}

func tickElves(elves []*Elf) {
	for _, elf := range elves {
		elf.waitTime--
	}
}

// Get the nodes which have this one as a dependency, which mean they should be next.
// They should have their dependencies met, meaning, their dependencies should be in `seen`.
func getNext(g Graph, current string) []string {
	next := []string{}

	for key, value := range g {
		if slices.Contains(value, current) {
			next = append(next, key)
		}
	}
	return next
}

// Dequeue the next value, in alphabetical order, but it has to have its dependencies met.
func dequeueValid(queue []string, graph Graph, seen []string) []string {
	// Check if the node has the dependencies met
	valid := []string{}
	for _, candidate := range queue {
		depMet := true
		// Check dependencies
		for _, dep := range graph[candidate] {
			if !slices.Contains(seen, dep) {
				depMet = false
			}
		}
		if depMet {
			// When we find a valid candidate with its dependencies met, save it
			valid = append(valid, candidate)
		}
	}
	slices.Sort(valid)
	return valid
}

func remove(queue []string, node string) []string {
	idx := slices.Index(queue, node)
	newSlice := make([]string, 0, len(queue)-1)

	// Copy the elements we want to keep into the new slice
	for i, item := range queue {
		if i != idx {
			newSlice = append(newSlice, item)
		}
	}
	// Copy the new slice back into the original slice
	return newSlice
}

func createElves(n int) []*Elf {
	e := []*Elf{}

	for range n {
		e = append(e, &Elf{".", 0, false})
	}
	return e
}

func createGraph() (Graph, []string) {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	g := Graph{}
	nodes := []string{}
	for _, line := range lines {
		sep := strings.Fields(line)
		// If entry exists
		if dependencies, ok := g[sep[7]]; ok {
			dependencies = append(dependencies, sep[1])
			g[sep[7]] = dependencies
			// Also create the dependency map. It may be empty if its the first node.
			if _, ok := g[sep[1]]; !ok {
				g[sep[1]] = []string{}
			}
		} else {
			g[sep[7]] = []string{sep[1]}
			// Also create the dependency map. It may be empty if its the first node.
			if _, ok := g[sep[1]]; !ok {
				g[sep[1]] = []string{}
			}
		}
	}
	for key := range g {
		nodes = append(nodes, key)
	}
	return g, nodes
}

// Return the nodes with no dependencies.
func getNoDep(g Graph) []string {
	noDep := []string{}
	for key, value := range g {
		if len(value) == 0 {
			noDep = append(noDep, key)
		}
	}
	return noDep
}
