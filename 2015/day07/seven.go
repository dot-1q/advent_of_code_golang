package main

import (
	"bufio"
	"fmt"
	"os"
	c "strconv"
	s "strings"
)

func main() {
	partOne()
	partTwo()
}

func readLines() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Be mindful of trimming the white space from the various splits
		lines = append(lines, scanner.Text())
	}
	return lines
}

func partOne() {

	lines := readLines()

	// Map of the wires ( instruction -> out )
	// var wires = make(map[string]string)
	// Map the known values of each wire ex: (a -> 400)
	wires := map[string]string{}
	results := map[string]uint16{}

	for i := range lines {
		// Be mindful of trimming the white space from the various splits
		splitLine := s.Split(s.TrimSpace(lines[i]), "->")
		// This is the destination of the assignment of each operation, after "->"
		dest := s.TrimSpace(splitLine[1])

		// Add the wires and their insts to map
		wires[dest] = s.TrimSpace(splitLine[0])
	}

	fmt.Println("Part1: Value: ", traverseGraph(wires, "a", results))
}

func partTwo() {
	lines := readLines()

	// Map of the wires ( instruction -> out )
	// var wires = make(map[string]string)
	// Map the known values of each wire ex: (a -> 400)
	wires := map[string]string{}
	results := map[string]uint16{}

	for i := range lines {
		// Be mindful of trimming the white space from the various splits
		splitLine := s.Split(s.TrimSpace(lines[i]), "->")
		// This is the destination of the assignment of each operation, after "->"
		dest := s.TrimSpace(splitLine[1])

		// Add the wires and their insts to map
		wires[dest] = s.TrimSpace(splitLine[0])
	}

	//Do the circuit, values get saved to results
	traverseGraph(wires, "a", results)
	// Override this wire b with the previous result
	wires["b"] = c.Itoa(int(results["a"]))
	// Reset the signals
	results = map[string]uint16{}
	fmt.Println("Part2: Value: ", traverseGraph(wires, "a", results))

}

// We have to recursively trasrverse the set of instructions, which are all interconected
// And assemble the logic circuit. Our limit/stop condition is when we hit a uint16 literal or an
// already known wire
func traverseGraph(instructions map[string]string, wire string, results map[string]uint16) uint16 {
	// Check if it has been calculated. Saves memory
	if savedValue, ok := results[wire]; ok {
		return savedValue
	}

	// This is the operation itself, the part before "->"
	// Trim the whitespaces around the string
	expression := s.Split(s.TrimSpace(instructions[wire]), " ")

	switch len(expression) {
	// if theres only one parameter, it's a simple assignment ex:(120 -> x or lx -> x)
	case 1:
		//Check if its a literal uint16
		if number, err := c.Atoi(wire); err == nil {
			results[wire] = uint16(number)
			return results[wire]
		}
		// else, simply return its value
		number := traverseGraph(instructions, expression[0], results)
		results[wire] = number
		return results[wire]
	// if theres two parameters, it's a bitwise operation on a single variable ex:(not y -> z)
	case 2:
		// Get the value of the variable we want to NOT
		operator1 := traverseGraph(instructions, expression[1], results)
		results[wire] = ^operator1
		return results[wire]
	// if theres three parameters, it's a bitwise operation ex:(x and y -> z)
	case 3:
		// another switch statement for the various operators
		switch expression[1] {
		case "AND":
			operator1 := traverseGraph(instructions, expression[0], results)
			operator2 := traverseGraph(instructions, expression[2], results)
			results[wire] = operator1 & operator2
			return results[wire]
		case "OR":
			operator1 := traverseGraph(instructions, expression[0], results)
			operator2 := traverseGraph(instructions, expression[2], results)
			results[wire] = operator1 | operator2
			return results[wire]
		case "LSHIFT":
			// the second operand in this case is an integer, and not derived from a variable
			operator1 := traverseGraph(instructions, expression[0], results)
			ui16, _ := c.Atoi(expression[2])
			operator2 := uint16(ui16)
			results[wire] = operator1 << operator2
			return results[wire]
		case "RSHIFT":
			// the second operand in this case is an integer, and not derived from a variable
			operator1 := traverseGraph(instructions, expression[0], results)
			ui16, _ := c.Atoi(expression[2])
			operator2 := uint16(ui16)
			results[wire] = operator1 >> operator2
			return results[wire]
		}
	}
	// All conditions are checked
	return 0
}
