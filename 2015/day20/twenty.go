package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	puzzle := "36000000"

	p, _ := strconv.Atoi(puzzle)
	h := partOne(p)
	partTwo(p, h)
}

func partOne(puzzle int) int {
	// Essentially, part one needs us to calculate the sum of the divisions of a given house, and  check if the division
	// is the same as the puzzle number.
	// Now, instead of just brute forcing the calculation of the whole proper factors of a number, and then calculating its sum,
	// there's actually a formula to calculate that sum, relying on the *PRIME* factors, so we just have to calculate those

	for house := 1; house <= math.MaxInt32; house++ {
		presents := sumOfDivisors(house)
		// Divide by ten, because the elves are multiplying by 10 for no reason lol
		if presents >= (puzzle / 10) {
			fmt.Println("Part1, House:", house)
			return house
		}
	}

	return 0
}
func partTwo(puzzle int, startingHouse int) int {
	// Since the elves stop delivering after 50 presents,
	// We need to find the factors of a House Number, which (Factor*50>=House Number), because each elf behaves
	// as a factor of a given house number, and they stop after 50.
	// To find such factors which obey the rule above, its easier to start finding factors in a decreasing manner,
	// since the smallest ones are guaranteed to not follow the rule, i.e, those elves would have stopped giving presents
	// already.

	// Store the calculated factors in a lookup table, to be easier to find them
	for house := startingHouse; house <= math.MaxInt32; house++ {
		factors := factorsInverse(house)

		presents := calculatePresents(factors)
		if presents >= (puzzle / 11) {
			fmt.Println("Part2, House:", house)
			return house
		}
	}
	return 0
}

// Get the prime factors of a number
// https://siongui.github.io/2017/05/09/go-find-all-prime-factors-of-integer-number/
func primeFactors(n int) []int {
	// Get the number of 2s that divide n
	pfs := []int{}
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return pfs
}

// Sum of divisors has a formula thats anchored on the values of the prime factorization
// Taken from: https://math.stackexchange.com/questions/22721/is-there-a-formula-to-calculate-the-sum-of-all-proper-divisors-of-a-number
func sumOfDivisors(number int) int {
	factors := primeFactors(number)
	sod := 1.0

	i := 0
	for i < len(factors) {
		c := count(&factors, factors[i], i)
		p := (math.Pow(float64(factors[i]), float64(c)+1) - 1)
		sod *= (p / (float64(factors[i]) - 1))
		i += (c)
	}
	return int(sod)
}

func calculatePresents(factors []int) int {
	count := 0
	for _, factor := range factors {
		count += factor
	}

	return count
}

// Count, but since its an ordered array, start from a given position
// also, pass a reference to array, to save memory
func count(slice *[]int, n int, start int) int {
	count := 0
	for i := start; i < len(*slice); i++ {
		if (*slice)[i] == n {
			count++
		}
	}
	return count
}

// Find the factors in inverse order, starting from the biggest and stopping when
// they do not obey by the rule (factor*50>number)
func factorsInverse(num int) []int {
	var factors []int

	// We start immediately with the /2 of the number, because there are no factors between [n/2,n], only
	// below/2
	factors = append(factors, num)
	for i := num / 2; i >= 1; i-- {
		if math.Remainder(float64(num), float64(i)) == 0 {
			// Only append if they obey the rule, else stop.
			// The subsequent factors dont need to be calculated
			if (i * 50) >= num {
				factors = append(factors, i)
			} else {
				break
			}
		}
	}
	return factors
}
