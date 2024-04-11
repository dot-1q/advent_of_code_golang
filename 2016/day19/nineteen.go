package main

import (
	"fmt"
)

type Seat struct {
	number int
	next   *Seat
}

func main() {
	puzzle := 3018458
	partOne(puzzle)
	partTwo(puzzle)
}

func partOne(puzzle int) {
	start := &Seat{number: 1}
	iterator := start
	// Array of the sitting positions. Will be easier to remove the elf sitting across for part 2
	seats := []*Seat{}
	seats = append(seats, iterator)
	// Sit the elves in a ring list
	for i := 2; i <= puzzle; i++ {
		iterator.next = &Seat{number: i}
		iterator = iterator.next
		seats = append(seats, iterator)
	}
	// Link the last seat to the first
	iterator.next = start

	// Given that each time we take a present from one elf, that elf gets removed...
	// The next seat form this one, is the next seat OF THE NEXT ONE (next.next)
	// Then update the stealing from the next position, which was updated
	for start.next != start {
		start.next = start.next.next
		start = start.next
	}
	fmt.Printf("Part1 : Iterator %d is last\n", start.number)
}

func partTwo(puzzle int) {
	// LOL
	// https://www.reddit.com/r/adventofcode/comments/5j4lp1/comment/dbdf50n/
	// I could not figure this out.
	// My previous solution involved removing the middle elements of an array, which is
	// insanely costly, but it worked for low N, given it was O(n2). For the puzzle input, there was no way
	// i was going to be clever enough.
	i := 1
	for i*3 < puzzle {
		i *= 3
	}
	fmt.Println(puzzle - i)
}
