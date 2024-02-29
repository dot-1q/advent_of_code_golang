package main

import (
	// "bufio"
	"fmt"
	// "io"
	"os"
)

func main() {

	// In this case, data is just a huge string in the text file
	data, err := os.ReadFile("input.txt")
	floor := 0

	if err != nil {
		fmt.Print(err)
	}

	parenthesis := string(data)

	for i := 0; i < len(data); i++ {
		if string(parenthesis[i]) == "(" {
			floor += 1
		}
		if string(parenthesis[i]) == ")" {
			floor -= 1
		}
	}

	fmt.Printf("Final floor: %d\n", floor)

}
