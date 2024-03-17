package main

import (
	"bufio"
	"fmt"
	"os"
	c "strconv"
	s "strings"
)

type Aunt struct {
	Name  int
	Items map[string]int
}

func (aunt *Aunt) addItem(item string, number int) {
	aunt.Items[item] = number
}

func main() {
	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	aunts := [500]Aunt{}
	i := 0
	for scanner.Scan() {
		lineB, lineA, _ := s.Cut(scanner.Text(), ":")
		name, _ := c.Atoi(s.Split(lineB, " ")[1])
		aunts[i] = Aunt{Name: name, Items: map[string]int{}}
		things := s.Split(s.TrimSpace(lineA), ",")

		for _, items := range things {
			info := s.Split(s.TrimSpace(items), ":")
			name := info[0]
			number, _ := c.Atoi(s.TrimSpace(info[1]))
			aunts[i].addItem(name, number)
		}
		i++
	}

	partOne(aunts)
	partTwo(aunts)

}

// Find which of the aunts has items that match the description
// Every aunt only has 3 items, so we need three hits.
// Given that values that are not present, aren't 0, just means we don't remember
// So we can't escape early
func partOne(aunts [500]Aunt) {

	numberOfHits := 0
	for _, aunt := range aunts {
		n, ok := aunt.Items["children"]
		if ok {
			if n == 3 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["cats"]
		if ok {
			if n == 7 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["samoyeds"]
		if ok {
			if n == 2 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["pomeranians"]
		if ok {
			if n == 3 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["akitas"]
		if ok {
			if n == 0 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["vizslas"]
		if ok {
			if n == 0 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["goldfish"]
		if ok {
			if n == 5 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["trees"]
		if ok {
			if n == 3 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["cars"]
		if ok {
			if n == 2 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["perfumes"]
		if ok {
			if n == 1 {
				numberOfHits++
			}
		}
		if numberOfHits >= 3 {
			fmt.Printf("Part1: Aunt %d\n", aunt.Name)
		}
		numberOfHits = 0
	}
}

// Same as partOne, but repalce exact matches with ranges
func partTwo(aunts [500]Aunt) {
	numberOfHits := 0
	for _, aunt := range aunts {
		n, ok := aunt.Items["children"]
		if ok {
			if n == 3 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["cats"]
		if ok {
			if n > 7 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["samoyeds"]
		if ok {
			if n == 2 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["pomeranians"]
		if ok {
			if n < 3 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["akitas"]
		if ok {
			if n == 0 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["vizslas"]
		if ok {
			if n == 0 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["goldfish"]
		if ok {
			if n < 5 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["trees"]
		if ok {
			if n > 3 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["cars"]
		if ok {
			if n == 2 {
				numberOfHits++
			}
		}
		n, ok = aunt.Items["perfumes"]
		if ok {
			if n == 1 {
				numberOfHits++
			}
		}
		if numberOfHits >= 3 {
			fmt.Printf("Part2: Aunt %d\n", aunt.Name)
		}
		numberOfHits = 0
	}

}
