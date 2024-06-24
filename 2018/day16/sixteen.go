package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n\n\n\n")

	// Seperate each state, which is comprised of "Before, Instruction, After"
	states := strings.Split(lines[0], "\n\n")
	program := strings.Split(lines[1], "\n")

	// Create an array of functions, so that we can go over each one of the possible opeations and check if the result
	// could be from them.
	funcNames := map[int]string{0: "addr", 1: "addi", 2: "mulr", 3: "muli", 4: "banr", 5: "bani", 6: "borr", 7: "bori", 8: "setr", 9: "seti", 10: "gtir", 11: "gtri", 12: "gtrr", 13: "eqir", 14: "eqri", 15: "eqrr"}
	funcs := []func([4]int, [4]int) [4]int{addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}

	behaveThree := 0
	opcodesPossible := map[int][]string{} //Map of each opcode to its instruction
	for _, line := range states {
		before := [4]int{}
		inst := [4]int{}
		after := [4]int{}
		// Extract the values from each string that represents a state. Mostly Sscanf shenanigans.
		fmt.Sscanf(line, "Before: [%2d, %2d, %2d, %2d]\n%2d %2d %2d %2d\nAfter:  [%2d, %2d, %2d, %2d]",
			&before[0], &before[1], &before[2], &before[3],
			&inst[0], &inst[1], &inst[2], &inst[3],
			&after[0], &after[1], &after[2], &after[3])
		// Check this state against all functions.
		passed, possilbleFuncs := Validate(funcs, funcNames, before, inst, after)
		if passed >= 3 {
			behaveThree++
		}
		appendUnique(opcodesPossible, inst[0], possilbleFuncs)
	}

	fmt.Printf("Part 1 | %d samples behave like three or more opcodes\n", behaveThree)
	opcodes := deduceOpcodes(opcodesPossible)
	registers := [4]int{}
	registers = runProgram(program, registers, opcodes)
	fmt.Printf("Part 2 | Register 0 has value %d\n", registers[0])
}

// Execute each of the 16 instructions and check how many give equivalent results. Also return a list of said instructions
// So that we can later deduce which number is each instruction.
func Validate(validations []func([4]int, [4]int) [4]int, funcNames map[int]string, before, inst, after [4]int) (int, []string) {
	passed := 0
	valid := []string{}
	for i, exec := range validations {
		// Execute the operation with the values
		result := exec(before, inst)
		// If it evaluates to be equivalent, it could be this function
		if slices.Equal(after[:], result[:]) {
			passed++
			valid = append(valid, funcNames[i])
		}
	}
	return passed, valid
}

func runProgram(program []string, registers [4]int, opcodes map[int]func([4]int, [4]int) [4]int) [4]int {
	for _, inst := range program {
		fields := [4]int{}
		fmt.Sscanf(inst, "%2d %2d %2d %2d", &fields[0], &fields[1], &fields[2], &fields[3])
		// Get the operation we are supposed to do.
		op := opcodes[fields[0]]
		registers = op(registers, fields)
	}
	return registers
}

func deduceOpcodes(opcodesPossible map[int][]string) map[int]func([4]int, [4]int) [4]int {
	opcodes := map[int]func([4]int, [4]int) [4]int{}
	removed := []int{}
	done := false

	for !done {
		done = true
		operationToRemove := ""
		for k, v := range opcodesPossible {
			// If this opcode has only one operation and hasnt its operation
			// has not been removed from the other operations, do it.
			if len(v) == 1 && !slices.Contains(removed, k) {
				operationToRemove = v[0]
				removed = append(removed, k)
				removeOperation(opcodesPossible, operationToRemove)
			}
		}
		// Check if there are still deductions to be made
		for _, v := range opcodesPossible {
			if len(v) != 1 {
				done = false
			}
		}
	}
	// Return a map of each opcode mapped to a single function.
	for k, v := range opcodesPossible {
		switch v[0] {
		case "addr":
			opcodes[k] = addr
		case "addi":
			opcodes[k] = addi
		case "mulr":
			opcodes[k] = mulr
		case "muli":
			opcodes[k] = muli
		case "banr":
			opcodes[k] = banr
		case "bani":
			opcodes[k] = bani
		case "borr":
			opcodes[k] = borr
		case "bori":
			opcodes[k] = bori
		case "setr":
			opcodes[k] = setr
		case "seti":
			opcodes[k] = seti
		case "gtir":
			opcodes[k] = gtir
		case "gtri":
			opcodes[k] = gtri
		case "gtrr":
			opcodes[k] = gtrr
		case "eqir":
			opcodes[k] = eqir
		case "eqri":
			opcodes[k] = eqri
		case "eqrr":
			opcodes[k] = eqrr
		}
	}
	return opcodes
}

func removeOperation(opcodes map[int][]string, operation string) {
	for k, v := range opcodes {
		if len(v) > 1 {
			removed := slices.DeleteFunc(v, func(v string) bool {
				return v == operation
			})
			opcodes[k] = removed
		}
	}
}

func appendUnique(opcodes map[int][]string, op int, possibleFuncs []string) {
	if p, ok := opcodes[op]; ok {
		for _, f := range possibleFuncs {
			if !slices.Contains(p, f) { // If it doesnt have this valid function, add it
				p = append(p, f)
			}
		}
		opcodes[op] = p
	} else { // There is no opcode in the map
		opcodes[op] = possibleFuncs
	}
}

func addr(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] + registers[instructions[2]]
	return registers
}

func addi(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] + instructions[2]
	return registers
}

func mulr(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] * registers[instructions[2]]
	return registers
}

func muli(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] * instructions[2]
	return registers
}

func banr(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] & registers[instructions[2]]
	return registers
}

func bani(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] & instructions[2]
	return registers
}

func borr(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] | registers[instructions[2]]
	return registers
}

func bori(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]] | instructions[2]
	return registers
}

func setr(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = registers[instructions[1]]
	return registers
}

func seti(registers [4]int, instructions [4]int) [4]int {
	registers[instructions[3]] = instructions[1]
	return registers
}

func gtir(registers [4]int, instructions [4]int) [4]int {
	if instructions[1] > registers[instructions[2]] {
		registers[instructions[3]] = 1
	} else {
		registers[instructions[3]] = 0
	}
	return registers
}

func gtri(registers [4]int, instructions [4]int) [4]int {
	if registers[instructions[1]] > instructions[2] {
		registers[instructions[3]] = 1
	} else {
		registers[instructions[3]] = 0
	}
	return registers
}

func gtrr(registers [4]int, instructions [4]int) [4]int {
	if registers[instructions[1]] > registers[instructions[2]] {
		registers[instructions[3]] = 1
	} else {
		registers[instructions[3]] = 0
	}
	return registers
}

func eqir(registers [4]int, instructions [4]int) [4]int {
	if instructions[1] == registers[instructions[2]] {
		registers[instructions[3]] = 1
	} else {
		registers[instructions[3]] = 0
	}
	return registers
}

func eqri(registers [4]int, instructions [4]int) [4]int {
	if registers[instructions[1]] == instructions[2] {
		registers[instructions[3]] = 1
	} else {
		registers[instructions[3]] = 0
	}
	return registers
}

func eqrr(registers [4]int, instructions [4]int) [4]int {
	if registers[instructions[1]] == registers[instructions[2]] {
		registers[instructions[3]] = 1
	} else {
		registers[instructions[3]] = 0
	}
	return registers
}
