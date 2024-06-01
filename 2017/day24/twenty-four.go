package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Each bind has exactly two ends. Their ordering doesn't matter.
type Bind struct {
	id   int
	End1 int
	End2 int
}

func main() {
	binds := createBinds()
	// Start the search from port type 0
	bridges := [][]Bind{}
	used := make([]bool, len(binds))
	findBridges(0, binds, used, []Bind{}, &bridges)
	maximum := 0
	for _, bridge := range bridges {
		// for _, b := range bridge {
		// 	fmt.Printf("%d/%d -> ", b.End1, b.End2)
		// }
		str := calculateStrength(bridge)
		if str > maximum {
			maximum = str
		}
	}
	fmt.Println("Part 1 | Max Strength is: ", maximum)
	longest := longestSlice(bridges)
	longestStrongest := calculateStrongest(bridges, longest)
	fmt.Printf("Part 2 | Longest which is the strongest has %d strength\n", longestStrongest)

}

// Find all possible bridges
func findBridges(currentPort int, binds []Bind, used []bool, currentBridge []Bind, allBridges *[][]Bind) {
	// Store the current bridge
	bridgeCopy := make([]Bind, len(currentBridge))
	copy(bridgeCopy, currentBridge)
	*allBridges = append(*allBridges, bridgeCopy)

	// Explore all possible connections
	for i, bind := range binds {
		if !used[i] {
			if bind.End1 == currentPort {
				used[i] = true
				findBridges(bind.End2, binds, used, append(currentBridge, bind), allBridges)
				used[i] = false
			} else if bind.End2 == currentPort {
				used[i] = true
				findBridges(bind.End1, binds, used, append(currentBridge, bind), allBridges)
				used[i] = false
			}
		}
	}
}

func calculateStrength(bridge []Bind) int {
	strength := 0
	for _, bind := range bridge {
		strength += bind.End1 + bind.End2
	}
	return strength
}

func longestSlice(slices [][]Bind) int {
	longest := 0
	for _, slice := range slices {
		currentLength := len(slice)
		if currentLength > longest {
			longest = currentLength
		}
	}
	return longest
}

// Calculate the strongest bridge from the set of the longest ones.
func calculateStrongest(bridges [][]Bind, length int) int {
	max := 0
	for _, bridge := range bridges {
		// Its the longest/one of the longest.
		if len(bridge) == length {
			str := calculateStrength(bridge)
			if str > max {
				max = str
			}
		}
	}
	return max
}

func createBinds() []Bind {
	b := []Bind{}

	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	id := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "/")
		end0, _ := strconv.Atoi(line[0])
		end1, _ := strconv.Atoi(line[1])
		bind := Bind{id, end0, end1}
		id++
		b = append(b, bind)
	}
	return b
}
