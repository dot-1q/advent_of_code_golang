package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Claim struct {
	id   string
	row  int
	cell int
	wide int
	tall int
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")
	claims := createClaims(lines)

	cloth := [1000][1000]string{}
	overlaps := applyClaims(&cloth, claims)
	fmt.Printf("Part 1 | Overlap areas: %d\n", findOverlap(&cloth))
	fmt.Printf("Part 2 | Claim not overlapped: %s\n", findNoOverlap(overlaps))
}

func applyClaims(cloth *[1000][1000]string, claims []Claim) map[string]bool {
	overlap := map[string]bool{}
	for _, claim := range claims {
		row, cell := claim.row, claim.cell
		for range claim.tall {
			for range claim.wide {
				if cloth[row][cell] == "" {
					cloth[row][cell] = claim.id
					// If it hasn't been registered as not overlapped
					if _, ok := overlap[claim.id]; !ok {
						overlap[claim.id] = false
					}
				} else {
					// Overlap areas
					// Save that this ID and the one prior are overlapped
					overlap[cloth[row][cell]] = true
					overlap[claim.id] = true
					cloth[row][cell] = "X"
				}
				cell++
			}
			cell = claim.cell
			row++
		}
	}
	return overlap
}

func findNoOverlap(overlaps map[string]bool) string {
	for key, value := range overlaps {
		if !value {
			return key
		}
	}
	return ""
}
func findOverlap(cloth *[1000][1000]string) int {
	s := 0
	for _, row := range *cloth {
		for _, cell := range row {
			if cell == "X" {
				s++
			}
		}
	}
	return s
}

func createClaims(lines []string) []Claim {
	claims := []Claim{}
	for _, line := range lines {
		sep := strings.Fields(line)
		id, _ := strings.CutPrefix(sep[0], "#")
		p, _ := strings.CutSuffix(sep[2], ":")
		position := strings.Split(p, ",")
		size := strings.Split(sep[3], "x")
		cell, _ := strconv.Atoi(position[0])
		row, _ := strconv.Atoi(position[1])
		wide, _ := strconv.Atoi(size[0])
		tall, _ := strconv.Atoi(size[1])
		claim := Claim{id, row, cell, wide, tall}
		claims = append(claims, claim)
	}
	return claims
}
