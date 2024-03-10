package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	c "strconv"
	"strings"
)

// Each person has a Name, and a map of their happiness relative to their neighbours, a position, to generate their permutation
type Person struct {
	Name      string
	Position  int
	Happiness map[string]int
}

func CreatePerson(name string, pos int) *Person {
	p := Person{Name: name, Position: pos, Happiness: make(map[string]int)}
	return &p
}

func (person *Person) addNeighbour(name string, weight int) {
	person.Happiness[name] = weight
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	people := createMap(f)
	partOne(people)
	partTwo(people)

}

func createMap(file *os.File) []Person {
	// Declare an array of people
	people := []Person{}
	scanner := bufio.NewScanner(file)
	position := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		happiness, _ := c.Atoi(line[3])

		// Check if person has already been created
		if !slices.ContainsFunc(people, func(p Person) bool {
			return p.Name == line[0]
		}) {
			person := CreatePerson(line[0], position)
			// Give a number for a person, to later permutate between them
			position++
			people = append(people, *person)
		}
		// Else just add information
		i := slices.IndexFunc(people, func(p Person) bool {
			return p.Name == line[0]
		})
		// Add the neihbour and its weight. Trim the period ending the sentence
		// Careful it its a 'gain' or 'lose' happiness
		if line[2] == "gain" {
			people[i].addNeighbour(strings.TrimSuffix(line[10], "."), happiness)
		} else {
			people[i].addNeighbour(strings.TrimSuffix(line[10], "."), happiness*-1)
		}
	}
	return people
}

func partOne(people []Person) {

	max_happiness := 0
	// Generate and get the next permutations of an array
	// Taken from https://stackoverflow.com/a/30230552
	for p := make([]Person, len(people)); p[0].Position < len(p); nextPerm(p) {
		disposition := getPerm(people, p)
		happiness := calculateHappiness(disposition)
		max_happiness = max(max_happiness, happiness)
	}

	fmt.Println("Part1: Maximum happiness is: ", max_happiness)
}

func partTwo(people []Person) {
	// In part two, we simply add ourselves to the mix, with happiness 0
	me := CreatePerson("Tiago", len(people))
	people = append(people, *me)

	for _, people := range people {
		// Add me as neighbour with weight 0 to everyone, and vice versa
		people.addNeighbour("Tiago", 0)
		me.addNeighbour(people.Name, 0)
	}

	max_happiness := 0
	// Generate and get the next permutations of an array
	// Taken from https://stackoverflow.com/a/30230552
	for p := make([]Person, len(people)); p[0].Position < len(p); nextPerm(p) {
		disposition := getPerm(people, p)
		happiness := calculateHappiness(disposition)
		max_happiness = max(max_happiness, happiness)
	}

	fmt.Println("Part2: Maximum happiness is: ", max_happiness)

}

// Given a certain disposition of people, calculate
func calculateHappiness(people []Person) int {
	happiness := 0

	for position, person := range people {
		next_pos := position + 1
		if next_pos < len(people) {
			// Get the happiness of this person to sit next to the other one and vice versa
			// They both have an Happiness value of sitting next to eachother
			neigh1 := person.Happiness[people[next_pos].Name]
			neigh2 := people[next_pos].Happiness[person.Name]

			happiness += (neigh1 + neigh2)
		} else {
			// If the next_pos, is at the end of the array, the last
			// position is adjacent to the first one, think of it like a circular array
			neigh1 := person.Happiness[people[0].Name]
			neigh2 := people[0].Happiness[person.Name]
			happiness += (neigh1 + neigh2)
		}
	}
	return happiness
}

// Generate and get the next permutations of an array
// Taken from https://stackoverflow.com/a/30230552
func nextPerm(p []Person) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i].Position < len(p)-i-1 {
			p[i].Position++
			return
		}
		p[i].Position = 0
	}
}

func getPerm(orig, p []Person) []Person {
	result := append([]Person{}, orig...)
	for i, v := range p {
		result[i], result[i+v.Position] = result[i+v.Position], result[i]
	}
	return result
}
