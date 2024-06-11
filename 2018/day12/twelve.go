package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	initialState := strings.TrimPrefix(strings.TrimSpace(lines[0]), "initial state: ")
	rules := createRules(lines[2:])

	tunnel := make([]string, 500)
	applyInitialState(&tunnel, initialState)
	for range 20 {
		applyRules(&tunnel, rules)
	}
	fmt.Println("Part 1 | Sum: ", calculateSum(&tunnel))

	partTwo(initialState, rules)
}

func partTwo(initialState string, rules map[string]bool) {
	// The tunnel has to be up sized for the ammount of generations we're computing.
	tunnel := make([]string, 1500)
	applyInitialState(&tunnel, initialState)
	lastValue := 0
	for i := range 500 {
		temp := make([]string, len(tunnel))
		copy(temp, tunnel)
		for range i {
			applyRules(&temp, rules)
		}
		sum := calculateSum(&temp)
		//@NOTE: Debug purposes, uncomment following print statements.
		// fmt.Printf("Part 2 | For %d Generations, Sum is: %d\n", i, sum)
		// You'll see this value converges to 72 in my case.
		// fmt.Printf("Gen %d | Difference to the last generation: %d\n", i, sum-lastValue)
		lastValue = sum
	}
	// The fifty billion iterations, minus the ones we did above (499) times the constant.
	lastValue += ((50000000000 - 499) * (72))
	fmt.Printf("Part 2 | Answer: %d\n", lastValue)
}

func calculateSum(tunnel *[]string) int {
	center := len(*tunnel) / 2
	sum := 0
	for i, value := range *tunnel {
		if value == "#" {
			sum += (i - center)
		}
	}
	return sum
}

func applyRules(tunnel *[]string, rules map[string]bool) {
	str := strings.Builder{}
	// We cant write straight to the array, because we override the state of the adjacent positions while we iterate.
	// Instead we keep a note of the changes to apply after. Its better than to copy the initial state i think.
	changes := map[int]string{}

	for i := 2; i < len(*tunnel)-2; i++ {
		// Write to the temporary string the 5 adjacent values of this position
		str.WriteString((*tunnel)[i-2])
		str.WriteString((*tunnel)[i-1])
		str.WriteString((*tunnel)[i])
		str.WriteString((*tunnel)[i+1])
		str.WriteString((*tunnel)[i+2])

		// If this string matches any rule:
		if value, ok := rules[str.String()]; ok {
			if value {
				changes[i] = "#"
			} else {
				changes[i] = "."
			}
		} else {
			changes[i] = "."
		}
		// Reset the temp string
		str.Reset()
	}
	// Apply changes
	for key, value := range changes {
		(*tunnel)[key] = value
	}
}

func applyInitialState(tunnel *[]string, initialState string) {
	center := len(*tunnel) / 2
	for i := range len(*tunnel) {
		(*tunnel)[i] = "."
	}
	for i, char := range initialState {
		if char == '#' {
			(*tunnel)[center+i] = "#"
		} else {
			(*tunnel)[center+i] = "."
		}
	}

}

func createRules(lines []string) map[string]bool {
	rules := map[string]bool{}
	for _, line := range lines {
		fields := strings.Fields(strings.TrimSpace(line))
		if fields[2] == "#" {
			rules[fields[0]] = true
		} else {
			rules[fields[0]] = false
		}
	}
	return rules
}
