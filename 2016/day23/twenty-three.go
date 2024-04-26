package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := string(f)
	instructions := strings.Split(lines, "\n")
	part := 2

	// Registers
	registers := map[string]int{}
	if part == 1 {
		registers["a"] = 7
	} else {
		registers["a"] = 12
	}
	registers["b"] = 0
	registers["c"] = 0
	registers["d"] = 0

	// PC will be seen as the program counter. it jumps to whatever instructions it needs to jump
	pc := 0
	for pc < len(instructions) {
		pc += decodeInst(instructions, pc, registers)
	}

	fmt.Println(registers)
}

// From the instruction, decode the next jump address in the array of instructions
// By default return 1, meaning, go to the intnext immediate instruction
func decodeInst(instructions []string, address int, registers map[string]int) int {
	i := strings.Split(instructions[address], " ")
	oper := i[0]
	switch oper {
	case "cpy":
		// Check if the second operator is a register. If not, means its an int
		value, err := strconv.Atoi(i[1])
		if err == nil {
			// Its a literal
			registers[i[2]] = value
		} else {
			// Its a register
			registers[i[2]] = registers[i[1]]
		}
	case "inc":
		registers[i[1]]++
	case "dec":
		registers[i[1]]--
	case "jnz":
		jmp, err := strconv.Atoi(i[2])
		if err != nil {
			// Its a register
			jmp = registers[i[2]]
		}
		// Check if the value of the register is *NOT* zero
		if val, ok := registers[i[1]]; ok && val != 0 {
			return jmp
		} else {
			// Previous failed, means first operand is a number
			val, _ := strconv.Atoi(i[1])
			if val != 0 {
				return jmp
			}
		}
	case "tgl":
		jump := registers[i[1]]
		// The jump is outside the program.
		if (address + jump) > len(instructions) {
			return 1
		}
		toggleInst := strings.Split(instructions[address+jump], " ")
		oper := toggleInst[0]
		// Alter the toggle Instruction operation
		if len(toggleInst) == 2 {
			if oper == "inc" {
				toggleInst[0] = "dec"
			} else {
				toggleInst[0] = "inc"
			}
		} else if len(toggleInst) == 3 {
			if oper == "jnz" {
				toggleInst[0] = "cpy"
			} else {
				toggleInst[0] = "jnz"
			}
		}
		// Place the toggled instruction in the original instruction array
		instructions[address+jump] = strings.Join(toggleInst, " ")
	}
	return 1
}
