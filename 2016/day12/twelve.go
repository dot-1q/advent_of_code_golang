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
	registers["a"] = 0
	registers["b"] = 0
	if part == 2 {
		registers["c"] = 1
	} else {
		registers["c"] = 0
	}
	registers["d"] = 0

	// PC will be seen as the program counter. it jumps to whatever instructions it needs to jump
	pc := 0
	for pc < len(instructions) {
		pc += decodeInst(instructions[pc], registers)
	}

	fmt.Println(registers)
}

// From the instruction, decode the next jump address in the array of instructions
// By default return 1, meaning, go to the intnext immediate instruction
func decodeInst(inst string, registers map[string]int) int {
	i := strings.Split(inst, " ")
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
		jmp, _ := strconv.Atoi(i[2])
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
	}
	return 1
}
