package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	c "strconv"
	s "strings"
)

type Reindeer struct {
	Name     string
	Speed    int
	Time     int
	RestTime int
	Distance int
	// The number of "ticks" a reindeer can make before resting
	Stamina     int
	CurrentRest int
	Resting     bool
	Points      int
}

// This method will be called every "second" (every loop of the array), and calculate the reindeers position
func (r *Reindeer) run() {
	// Check if reindeer is resting or not
	if r.Resting {
		r.CurrentRest--
		// Limit case
		if r.CurrentRest == 1 {
			r.CurrentRest = r.RestTime
			r.Resting = false
		}
	} else { // If the reindeer is not resting ,it can run
		// While the reindeer can run...
		if r.Stamina > 0 {
			r.Stamina--
			r.Distance += r.Speed
		} else {
			// When she runs out of stamina, restart the stamina counter, but now
			// she has to rest
			r.Stamina = r.Time
			r.Resting = true
		}
	}
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	reindeers := []Reindeer{}
	for scanner.Scan() {
		line := s.Split(scanner.Text(), " ")
		name := line[0]
		speed, _ := c.Atoi(s.TrimSpace(line[3]))
		time, _ := c.Atoi(s.TrimSpace(line[6]))
		rest, _ := c.Atoi(s.TrimSpace(line[13]))

		reindeers = append(reindeers, Reindeer{name, speed, time, rest, 0, time, rest, false, 0})
	}

	// partOne(reindeers)
	partTwo(reindeers)
}

func partOne(reindeers []Reindeer) {

	// Do the race
	for range 2503 {
		for r := range len(reindeers) {
			reindeers[r].run()
		}
	}

	max_travel := 0
	name := ""
	for _, r := range reindeers {
		if r.Distance > max_travel {
			max_travel = r.Distance
			name = r.Name
		}
	}
	fmt.Printf("Reindeer %s has won by travelling %d\n", name, max_travel)

}

func partTwo(reindeers []Reindeer) {
	// Do the race
	for range 2503 {
		for r := range len(reindeers) {
			reindeers[r].run()
		}
		// Find and give point to lead. Have to pass a pointer
		givePointToLead(&reindeers)
	}

	max_points := 0
	name := ""
	for _, r := range reindeers {
		if r.Points > max_points {
			max_points = r.Points
			name = r.Name
		}
	}

	fmt.Printf("Reindeer %s has won by having %d points\n", name, max_points)
}

func givePointToLead(reindeers *[]Reindeer) {

	// Sort the reindeer array by the max points
	slices.SortFunc(*reindeers, func(r1, r2 Reindeer) int {
		return cmp.Compare(r1.Distance, r2.Distance)
	})
	// Give a point to the first place
	// The slices is sorted in ascending order,
	// So we iterate the over the slice from back to front,
	// because there could be ties
	(*reindeers)[len(*reindeers)-1].Points++
	for i := len(*reindeers) - 2; i >= 0; i-- {
		if (*reindeers)[i].Distance == (*reindeers)[len(*reindeers)-1].Distance {
			(*reindeers)[i].Points++

		} else {
			break
		}
	}

}
