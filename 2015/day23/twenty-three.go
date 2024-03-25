package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read the whole file as a list of instructions
	f, _ := os.ReadFile("input.txt")
	line := string(f)
	instructions := strings.Split(line, "\n")

	// The last position of the array is an empty line
	runProgram(instructions[:len(instructions)-1])
}

func runProgram(instructions []string) {

	regA := 1
	regB := 0
	// THis for loop will essentially be the Program Counter of this program
	for i := 0; i < len(instructions); {
		// fmt.Printf("Address: %d ", i)
		i += decodeInstr(instructions[i], &regA, &regB)
	}

	fmt.Printf("Part1: RegA: %d, RegB: %d\n", regA, regB)
}

// For any given instruction, return the address offset for the next instruction
// The address will only be used in jump instructions
func decodeInstr(instruction string, regA, regB *int) int {
	inst := strings.Split(instruction, " ")[0]
	reg := strings.Split(instruction, " ")[1]
	fmt.Printf("inst: %s for reg %s\n", inst, reg)

	switch inst {
	case "hlf":
		switch reg {
		case "a":
			*regA = *regA / 2
		case "b":
			*regB = *regB / 2
		}
	case "tpl":
		switch reg {
		case "a":
			*regA = *regA * 3
		case "b":
			*regB = *regB * 3
		}
	case "inc":
		switch reg {
		case "a":
			*regA = *regA + 1
		case "b":
			*regB = *regB + 1
		}
	case "jmp":
		address, _ := strconv.Atoi(reg)
		return address
	case "jie":
		reg = strings.TrimSuffix(reg, ",")
		offset := strings.TrimSpace(strings.Split(instruction, " ")[2])
		switch reg {
		case "a":
			// Is Even
			if *regA%2 == 0 {
				address, _ := strconv.Atoi(offset)
				return address
			}
		case "b":
			// Is Even
			if *regB%2 == 0 {
				address, _ := strconv.Atoi(offset)
				return address
			}
		}
	case "jio":
		reg = strings.TrimSuffix(reg, ",")
		offset := strings.TrimSpace(strings.Split(instruction, " ")[2])
		switch reg {
		case "a":
			// Is One
			if *regA == 1 {
				address, _ := strconv.Atoi(offset)
				return address
			}
		case "b":
			// Is One
			if *regB == 1 {
				address, _ := strconv.Atoi(offset)
				return address
			}
		}
	}
	return 1
}
