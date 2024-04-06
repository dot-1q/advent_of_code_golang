package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Disc struct {
	Number    int
	Positions int
	CurrPos   int
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	discs := []Disc{}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		number, _ := strconv.Atoi(strings.TrimPrefix(line[1], "#"))
		positions, _ := strconv.Atoi(line[3])
		currPos, _ := strconv.Atoi(strings.TrimSuffix(line[11], "."))
		disc := Disc{Number: number, Positions: positions, CurrPos: currPos}
		discs = append(discs, disc)
	}

	part := 2
	if part == 2 {
		discs = append(discs, Disc{Number: 7, Positions: 11, CurrPos: 0})
	}
	time := 0
	for {
		// Calculate the disposition had we drop at time N
		disposition := discPositionAt(discs, time)
		// Check if for any given disposition, we would be able to pass
		if canPassAllSlots(disposition) {
			break
		}
		time++
	}

	fmt.Printf("Part %d: Drop at time: %d to pass all slots:\n", part, time)
}

// At every time instant, calculate the disposition that each disc would have,
// had we drop the capsule at time N.
// Return said disposition.
func discPositionAt(discs []Disc, time int) []int {
	disp := []int{}
	for d := range discs {
		// Calculate the new position everytime we tick. Be mindfull of modulo over
		// all the positions, so we dont go out of bounds
		// Also add the Disc position(Number), to the time, as we need to wait N seconds to reach any given disc, because of its position
		disp = append(disp, (discs[d].CurrPos+time+discs[d].Number)%discs[d].Positions)
	}

	return disp
}

// Check if all discs are at slot 0, which allows the capsule passes
func canPassAllSlots(disp []int) bool {
	for _, p := range disp {
		// If any of the discs is not in the slot 0, it cant pass
		if p != 0 {
			return false
		}
	}
	return true
}
