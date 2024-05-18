package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi")
	partOne()
	partTwo()
}

func partOne() {
	a := 634
	b := 301

	judge := 0
	for range 40000000 {
		a = generateNextValue(a, 16807)
		b = generateNextValue(b, 48271)
		// xOr the two binary values and shift 48 positions.
		// The result will be 0, if the last 16 bits are zero.
		xor := (a ^ b) << 48
		if (xor) == 0 {
			judge++
		}
	}
	fmt.Printf("Part 1 | Judge: %d\n", judge)
}

func partTwo() {
	a := 634
	b := 301

	judge := 0
	for range 5000000 {
		a = generateNextValueCriteria(a, 16807, 4)
		b = generateNextValueCriteria(b, 48271, 8)
		// xOr the two binary values and shift 48 positions.
		// The result will be 0, if the last 16 bits are zero.
		xor := (a ^ b) << 48
		if (xor) == 0 {
			judge++
		}
	}
	fmt.Printf("Part 2 | Judge: %d\n", judge)
}

func generateNextValue(value, factor int) int {
	return (value * factor) % 2147483647
}

func generateNextValueCriteria(value, factor, divisor int) int {
	v := (value * factor) % 2147483647
	for v%divisor != 0 {
		v = (v * factor) % 2147483647
	}
	return v
}
