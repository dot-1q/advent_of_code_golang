package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	registers := map[string]int{}
	maxHeld := 0
	for _, line := range lines {
		fields := strings.Fields(line)

		// Operation
		switch fields[5] {
		case "<":
			regValue := getRegisterValue(fields[4], registers)
			val, _ := strconv.Atoi(fields[6])
			if regValue < val {
				amount, _ := strconv.Atoi(fields[2])
				if fields[1] == "inc" {
					val := getRegisterValue(fields[0], registers)
					val += amount
					registers[fields[0]] = val
				} else {
					val := getRegisterValue(fields[0], registers)
					val -= amount
					registers[fields[0]] = val
				}
			}
		case ">":
			regValue := getRegisterValue(fields[4], registers)
			val, _ := strconv.Atoi(fields[6])
			if regValue > val {
				amount, _ := strconv.Atoi(fields[2])
				if fields[1] == "inc" {
					val := getRegisterValue(fields[0], registers)
					val += amount
					registers[fields[0]] = val
				} else {
					val := getRegisterValue(fields[0], registers)
					val -= amount
					registers[fields[0]] = val
				}
			}
		case ">=":
			regValue := getRegisterValue(fields[4], registers)
			val, _ := strconv.Atoi(fields[6])
			if regValue >= val {
				amount, _ := strconv.Atoi(fields[2])
				if fields[1] == "inc" {
					val := getRegisterValue(fields[0], registers)
					val += amount
					registers[fields[0]] = val
				} else {
					val := getRegisterValue(fields[0], registers)
					val -= amount
					registers[fields[0]] = val
				}
			}
		case "<=":
			regValue := getRegisterValue(fields[4], registers)
			val, _ := strconv.Atoi(fields[6])
			if regValue <= val {
				amount, _ := strconv.Atoi(fields[2])
				if fields[1] == "inc" {
					val := getRegisterValue(fields[0], registers)
					val += amount
					registers[fields[0]] = val
				} else {
					val := getRegisterValue(fields[0], registers)
					val -= amount
					registers[fields[0]] = val
				}
			}
		case "==":
			regValue := getRegisterValue(fields[4], registers)
			val, _ := strconv.Atoi(fields[6])
			if regValue == val {
				amount, _ := strconv.Atoi(fields[2])
				if fields[1] == "inc" {
					val := getRegisterValue(fields[0], registers)
					val += amount
					registers[fields[0]] = val
				} else {
					val := getRegisterValue(fields[0], registers)
					val -= amount
					registers[fields[0]] = val
				}
			}
		case "!=":
			regValue := getRegisterValue(fields[4], registers)
			val, _ := strconv.Atoi(fields[6])
			if regValue != val {
				amount, _ := strconv.Atoi(fields[2])
				if fields[1] == "inc" {
					val := getRegisterValue(fields[0], registers)
					val += amount
					registers[fields[0]] = val
				} else {
					val := getRegisterValue(fields[0], registers)
					val -= amount
					registers[fields[0]] = val
				}
			}
		}
		currentMax := getMaxRegister(registers)
		if currentMax > maxHeld {
			maxHeld = currentMax
		}
	}

	fmt.Println(registers)
	fmt.Printf("Part1 | Max Value: %d\n", getMaxRegister(registers))
	fmt.Printf("Part2 | Max Value held: %d\n", maxHeld)
}

func getRegisterValue(reg string, registers map[string]int) int {
	if val, ok := registers[reg]; ok {
		return val
	}
	return 0
}

func getMaxRegister(registers map[string]int) int {
	max := 0
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	return max
}
