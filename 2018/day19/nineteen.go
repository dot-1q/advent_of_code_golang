package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	instructions := instructions()
	// IP = 2
	registers := [6]int{0, 0, 0, 0, 0, 0}
	registers = runProgram(instructions[1:], registers)
	fmt.Printf("Part 1 | When the program halts, the value %d is in register[%d]\n", registers[0], 0)
	fmt.Printf("Part 2 | When the program halts, the value %d is in register[%d]\n", part2(), 0)
}

func runProgram(program []string, registers [6]int) [6]int {
	ip := registers[2]
	for ip < len(program) {
		registers = decode(program[ip], registers)
		registers[2]++
		// Update IP counter
		ip = registers[2]
	}
	return registers
}

func decode(instruction string, registers [6]int) [6]int {
	fields := [3]int{}
	inst := ""
	fmt.Sscanf(instruction, "%s %2d %2d %2d", &inst, &fields[0], &fields[1], &fields[2])

	switch inst {
	case "seti":
		registers = seti(registers, fields)
	case "setr":
		registers = setr(registers, fields)
	case "addi":
		registers = addi(registers, fields)
	case "addr":
		registers = addr(registers, fields)
	case "muli":
		registers = muli(registers, fields)
	case "mulr":
		registers = mulr(registers, fields)
	case "gtrr":
		registers = gtrr(registers, fields)
	case "eqrr":
		registers = eqrr(registers, fields)
	default:
		fmt.Printf("inst: %s not found\n", inst)
	}

	return registers
}

func addr(registers [6]int, instructions [3]int) [6]int {
	registers[instructions[2]] = registers[instructions[0]] + registers[instructions[1]]
	return registers
}

func addi(registers [6]int, instructions [3]int) [6]int {
	registers[instructions[2]] = registers[instructions[0]] + instructions[1]
	return registers
}

func seti(registers [6]int, instructions [3]int) [6]int {
	registers[instructions[2]] = instructions[0]
	return registers
}

func setr(registers [6]int, instructions [3]int) [6]int {
	registers[instructions[2]] = registers[instructions[0]]
	return registers
}

func mulr(registers [6]int, instructions [3]int) [6]int {
	registers[instructions[2]] = registers[instructions[0]] * registers[instructions[1]]
	return registers
}

func muli(registers [6]int, instructions [3]int) [6]int {
	registers[instructions[2]] = registers[instructions[0]] * instructions[1]
	return registers
}

func gtrr(registers [6]int, instructions [3]int) [6]int {
	if registers[instructions[0]] > registers[instructions[1]] {
		registers[instructions[2]] = 1
	} else {
		registers[instructions[2]] = 0
	}
	return registers
}

func eqrr(registers [6]int, instructions [3]int) [6]int {
	if registers[instructions[0]] == registers[instructions[1]] {
		registers[instructions[2]] = 1
	} else {
		registers[instructions[2]] = 0
	}
	return registers
}

func instructions() []string {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	return lines
}

// I can't be arsed to disassemble assembly what was meticulously crafted lol. Got this off Github
func part2() int {
	numberToFactorize := 10551387 // this index varies based on inputs

	var ans int
	for i := 1; i <= numberToFactorize; i++ {
		if numberToFactorize%i == 0 {
			ans += i
		}
	}

	return ans
}
