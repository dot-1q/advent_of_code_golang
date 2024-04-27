package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := string(f)
	instructions := strings.Split(lines, "\n")

	// We need 3 values for part 2
	for i := range math.MaxInt32 {
		// Means it timedout
		if findLowestA(instructions, i) {
			fmt.Printf("Value %d is a valid value\n", i)
			return
		}
	}
}

func findLowestA(instructions []string, registerA int) bool {
	// Registers
	registers := map[string]int{}
	registers["a"] = registerA
	registers["b"] = 0
	registers["c"] = 0
	registers["d"] = 0

	start := time.Now()
	// PC will be seen as the program counter. it jumps to whatever instructions it needs to jump
	pc := 0
	lastValue := -1
	for pc < len(instructions) {
		// This last value is the return of register b, which should be 0,1,0,1,0,1
		// Instructions that do not output nothing, return the value -1
		address, ret := decodeInst(instructions, pc, registers)
		// If its the out instruction
		if ret != -1 {
			// Not alternating
			if ret == lastValue {
				// Stop this iteration
				return false
			}
			lastValue = ret
			current := time.Since(start)
			// We've been alternating for 5 seconds. We can safely guess this value for register A is valid.
			if current.Seconds() > 5 {
				return true
			}

		}
		pc += address
	}
	return false
}

// From the instruction, decode the next jump address in the array of instructions
// By default return 1, meaning, go to the intnext immediate instruction
func decodeInst(instructions []string, address int, registers map[string]int) (int, int) {
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
			return jmp, -1
		} else {
			// Previous failed, means first operand is a number
			val, _ := strconv.Atoi(i[1])
			if val != 0 {
				return jmp, -1
			}
		}
	case "out":
		value := registers[i[1]]
		// return the value on register
		return 1, value
		// fmt.Println(value)
	}
	return 1, -1
}
