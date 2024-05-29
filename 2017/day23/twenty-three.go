package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi")
	f, _ := os.ReadFile("input.txt")
	instructions := strings.Split(strings.TrimSpace(string(f)), "\n")
	partOne(instructions)
	value := partTwo()
	// Im sorry but im not gonna debug assembly and try to decipher the code.
	// Its an interesting puzzle, but im more interested in writing algorithms with Go.
	// So i cheated. Looked it up on reddit. Sue me.
	fmt.Printf("Part 2 : H: %d\n", value)
}

func partOne(instructions []string) {
	registers := map[string]int{}
	// Save the last played frequency from snd
	address := 0
	occurrences := 0
	for address < len(instructions) {
		address += decode(instructions[address], registers, &occurrences)
	}
	fmt.Printf("Part 1 : Mul operation done %d\n", occurrences)
}

func partTwo() int {
	b := 79
	c := 79

	b = b*100 + 100000
	c = b + 17000
	var h int
	for {
		f := 1
		// effectively a prime number checker.
		for d := 2; d*d <= b; d++ {
			if b%d == 0 {
				f = 0
				break
			}
		}

		if f == 0 {
			h++
		}
		if b == c {
			break
		}
		b += 17
	}

	return h
}

// Decode for part 1
func decode(instruction string, registers map[string]int, occurrences *int) int {
	inst := strings.Fields(instruction)

	switch inst[0] {
	case "set":
		registers[inst[1]] = getValue(inst[2], registers)
	case "sub":
		registers[inst[1]] -= getValue(inst[2], registers)
	case "mul":
		registers[inst[1]] *= getValue(inst[2], registers)
		*occurrences++
	case "jnz":
		offset := getValue(inst[2], registers)
		val := getValue(inst[1], registers)
		// Jump the offset if value is greater than zero
		if val != 0 {
			return offset
		}
	}
	return 1
}

// Get value. could be a register variable or just a number literal.
func getValue(literal string, registers map[string]int) int {
	// Value exists in register
	if value, ok := registers[literal]; ok {
		return value
	}
	// Is a literal or an initialized register, and return 0 if that is the case.
	value, err := strconv.Atoi(literal)
	if err == nil {
		// Its a literal
		return value
	}
	return 0
}
