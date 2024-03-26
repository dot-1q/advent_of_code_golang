package main

import (
	"fmt"
	"math"
)

func main() {
	// This is given to us in the problem statement

	// So, given that [1,1] = 1 iteration
	// [1,3] = 1+2+3 = 6 iterations
	// [1,4] = 1+2+3+4 = 10 iterations
	// Our input column is 3083, so that means that
	// [1,3083] = 1+2+3+4...+3083 = ((n+1)*n)/2 == 4753986 // Gauss formula
	// Now, this is going in column order, left to right
	// To go down, its the similar. Take notice that from each row,
	// the iteration number is [1,x] = V(calculated above) and [2,x]= V+x=D, [3,x] = (D+(x+1))
	// Example: [1,3] = 1+2+3 = 6
	// [2,3] = 6+3 == 9
	// [3,3] = 9+4 = 11
	// [4,3] = 11+5 = 16

	a := 20151125
	x := 1
	y := 1
	for {
		// Go diagonally in the iterations
		if y == 1 {
			y = x + 1
			x = 1
		} else {
			y -= 1
			x += 1
		}

		a = int(math.Mod(float64(a*252533), 33554393))

		if x == 3083 && y == 2978 {
			fmt.Println(a)
			break
		}
	}

}
