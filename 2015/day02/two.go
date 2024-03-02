package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var l int64 = 0
	var w int64 = 0
	var h int64 = 0
	var total int64 = 0
	var ribbon int64 = 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "x")

		l, _ = strconv.ParseInt(line[0], 10, 64)
		w, _ = strconv.ParseInt(line[1], 10, 64)
		h, _ = strconv.ParseInt(line[2], 10, 64)
		wrap := calculateAreaAndSlack(l, w, h)
		r := calculateRibbonAndBow(l, w, h)

		total += wrap
		ribbon += r
	}

	fmt.Printf("Total wrap %d\n", total)
	fmt.Printf("Total ribbon %d\n", ribbon)
}

// Calculate area and the smallest side and return their addition
func calculateAreaAndSlack(l, w, h int64) int64 {

	side1 := l * w
	side2 := w * h
	side3 := h * l
	min_side := min(side1, side2, side3)
	return (2 * (side1 + side2 + side3)) + min_side
}

func calculateRibbonAndBow(l, w, h int64) int64 {

	volume := l * h * w
	// Calculate all perimeters
	perimeter1 := 2*l + 2*w
	perimeter2 := 2*w + 2*h
	perimeter3 := 2*h + 2*l
	// The bow has the same length has the smallest perimeter
	min_perimeter := min(perimeter1, perimeter2, perimeter3)

	return volume + min_perimeter

}
