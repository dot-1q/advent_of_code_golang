package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Layer struct {
	Depth      int  // Depth number
	Range      int  // Range number
	Scanner    int  // Position of the Security, that  ranges from 0-[Range]
	ScannerDir bool // Direction: UP or down
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	firewall := map[int]Layer{}
	for _, l := range lines {
		info := strings.Split(l, ":")
		depth, _ := strconv.Atoi(strings.TrimSpace(info[0]))
		r, _ := strconv.Atoi(strings.TrimSpace(info[1]))
		newL := Layer{Depth: depth, Range: r, Scanner: 0, ScannerDir: false}
		firewall[depth] = newL
	}

	partOne(firewall)
	partTwo(firewall)
}

func partOne(firewall map[int]Layer) {
	// Packet position
	packet := -1
	// Save the layers on where we were caught
	caught := []Layer{}
	maxLevel := lastLevel(firewall)

	for range maxLevel + 1 {
		packet++
		// If theres a security checkpoint at this depth
		if r, ok := firewall[packet]; ok {
			// Check the multiple of the number of steps it takes for the scanner to be at level 0
			index := (r.Range - 1) * 2
			if packet%index == 0 {
				caught = append(caught, r)
			}
		}
	}
	fmt.Println("Severity: ", findSeverity(caught))
}

func partTwo(firewall map[int]Layer) {
	maxLevel := lastLevel(firewall)
	delay := 0
	// I have to deep copy the firewall
	for !tryPassing(firewall, maxLevel, delay) {
		delay++
	}
	fmt.Println("Delay to not get caught: ", delay)
}

func tryPassing(firewall map[int]Layer, maxLevel, delay int) bool {
	// Packet position
	packet := -1
	for range maxLevel + 1 {
		packet++
		// If theres a security checkpoint at this depth
		if r, ok := firewall[packet]; ok {
			// Check the multiple of the number of steps it takes for the scanner to be at level 0
			// If we get caught, this delay is not enough
			index := (r.Range - 1) * 2
			if (packet+delay)%index == 0 {
				return false
			}
		}
	}
	// We've escaped the maze with this delay level
	return true
}

func lastLevel(firewall map[int]Layer) int {
	maxLevel := 0
	for k := range firewall {
		if k > maxLevel {
			maxLevel = k
		}
	}
	return maxLevel
}

func findSeverity(caught []Layer) int {
	severity := 0
	for _, c := range caught {
		severity += (c.Depth * c.Range)
	}
	return severity
}
