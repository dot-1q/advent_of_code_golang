package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	id         int            // Program ID
	registers  map[string]int // Registers
	pc         int            // Program Counter
	waiting    bool           // Waiting to receive a value
	ended      bool           // End flag, to check if this program has exited
	valuesSent int            // Keep track of how many values each program sent
}

type SendQueue struct {
	queue0 []int // Send queue of the program 0
	queue1 []int // Send queue of the program 1
}

func main() {
	f, _ := os.ReadFile("input.txt")
	instructions := strings.Split(strings.TrimSpace(string(f)), "\n")
	partOne(instructions)
	partTwo(instructions)
}

func partOne(instructions []string) {
	registers := map[string]int{}
	// Save the last played frequency from snd
	registers["lp"] = 0
	address := 0
	for address < len(instructions) {
		next, stop := decode(instructions[address], registers)
		if stop {
			lastPlayed := registers["lp"]
			fmt.Println("Part 1 | Last Played: ", lastPlayed)
			break
		}
		address += next
	}
}

func partTwo(instructions []string) {
	pc0 := Program{0, map[string]int{}, 0, false, false, 0}
	pc1 := Program{1, map[string]int{}, 0, false, false, 0}
	sq := SendQueue{[]int{}, []int{}} // Their respective send queue
	// Start their 'p' register.
	pc0.registers["p"] = 0
	pc1.registers["p"] = 1

	for !pc0.ended || !pc1.ended {
		if pc0.pc < len(instructions) {
			pc0.pc += pc0.decode(instructions[pc0.pc], &sq)
		} else {
			pc0.ended = true
		}
		if pc1.pc < len(instructions) {
			pc1.pc += pc1.decode(instructions[pc1.pc], &sq)
		} else {
			pc1.ended = true
		}
		// If they're both waiting, it's a deadlock
		if pc0.waiting && pc1.waiting {
			break
		}
	}
	// fmt.Printf("%+v\n", pc0)
	// fmt.Printf("%+v\n", pc1)
	fmt.Printf("Part 2 | Program 1 sent %d values\n", pc1.valuesSent)
}

func (program *Program) decode(instruction string, sq *SendQueue) int {
	inst := strings.Fields(instruction)

	switch inst[0] {
	case "snd":
		val := getValue(inst[1], program.registers)
		// Get this program ID. We want to send this value to the other programs send queue.
		q := program.id
		if q == 0 {
			sq.queue1 = append(sq.queue1, val)
		} else {
			sq.queue0 = append(sq.queue0, val)
			// Values sent by program1
			program.valuesSent++
		}
	case "set":
		program.registers[inst[1]] = getValue(inst[2], program.registers)
	case "add":
		program.registers[inst[1]] += getValue(inst[2], program.registers)
	case "mul":
		program.registers[inst[1]] *= getValue(inst[2], program.registers)
	case "mod":
		program.registers[inst[1]] %= getValue(inst[2], program.registers)
	case "rcv":
		// Get this program ID. We have to receive from the opposite Send Queue.
		q := program.id
		if q == 0 {
			// If there are no values, we wait
			if len(sq.queue0) == 0 {
				program.waiting = true
				// Don't jump to the next instruction. Wait here
				return 0
			} else {
				// Remove from the front of the queue, and move
				val := sq.queue0[0]
				sq.queue0 = sq.queue0[1:]
				program.registers[inst[1]] = val
			}
		} else {
			// If there are no values, we wait
			if len(sq.queue1) == 0 {
				program.waiting = true
				// Don't jump to the next instruction. Wait here
				return 0
			} else {
				// Remove from the front of the queue, and move
				val := sq.queue1[0]
				sq.queue1 = sq.queue1[1:]
				program.registers[inst[1]] = val
			}
		}
	case "jgz":
		offset := getValue(inst[2], program.registers)
		val := getValue(inst[1], program.registers)
		// Jump the offset if value is greater than zero
		if val > 0 {
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

// Decode for part 1
func decode(instruction string, registers map[string]int) (int, bool) {
	inst := strings.Fields(instruction)

	switch inst[0] {
	case "snd":
		registers["lp"] = getValue(inst[1], registers)
	case "set":
		registers[inst[1]] = getValue(inst[2], registers)
	case "add":
		registers[inst[1]] += getValue(inst[2], registers)
	case "mul":
		registers[inst[1]] *= getValue(inst[2], registers)
	case "mod":
		registers[inst[1]] %= getValue(inst[2], registers)
	case "rcv":
		val := getValue(inst[1], registers)
		if val != 0 {
			// Get the last played sound stored in the "lp" register.
			return 0, true
		}
	case "jgz":
		offset := getValue(inst[2], registers)
		val := getValue(inst[1], registers)
		// Jump the offset if value is greater than zero
		if val > 0 {
			return offset, false
		}
	}
	return 1, false
}
