package main

import (
	"bufio"
	"fmt"
	"os"
	c "strconv"
	s "strings"
)

type Ingredient struct {
	Name       string
	Capacity   int
	Durability int
	Flavor     int
	Texture    int
	Calories   int
}

func main() {

	f, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(f)

	ingredients := []Ingredient{}
	for scanner.Scan() {
		line := s.Split(scanner.Text(), ":")
		name := line[0]
		weights := s.Split(s.TrimSpace(line[1]), ",")
		capacity, _ := c.Atoi(s.Split(s.TrimSpace(weights[0]), " ")[1])
		durability, _ := c.Atoi(s.Split(s.TrimSpace(weights[1]), " ")[1])
		flavor, _ := c.Atoi(s.Split(s.TrimSpace(weights[2]), " ")[1])
		texture, _ := c.Atoi(s.Split(s.TrimSpace(weights[3]), " ")[1])
		calories, _ := c.Atoi(s.Split(s.TrimSpace(weights[4]), " ")[1])

		ingredients = append(ingredients, Ingredient{Name: name, Capacity: capacity, Durability: durability, Flavor: flavor, Texture: texture, Calories: calories})
	}

	fmt.Println(ingredients)

	partOne(ingredients)
	partTwo(ingredients)
}

func partOne(ingredients []Ingredient) {
	// Number of table spoons
	tbsp := 100
	maxScore := 0
	tSprinkes := 100
	tPeanut := 0
	tFrosting := 0
	tSugar := 0

	finalSp := 0
	finalPb := 0
	finalFr := 0
	finalSu := 0

	// Brute force the 100 tbps
	// how many tbps? 0,1,2,....99,100 | 100 has to be calculated as well
	for s := tSprinkes; s >= 0; s-- {
		tPeanut = (tbsp - tSprinkes)
		for p := tPeanut; p >= 0; p-- {
			tFrosting = (tbsp - tSprinkes - tPeanut)
			for p := tFrosting; p >= 0; p-- {
				tSugar = (tbsp - tSprinkes - tPeanut - tFrosting)

				s := CalculateScore([]int{tSprinkes, tPeanut, tFrosting, tSugar}, ingredients)

				if s > maxScore {
					maxScore = s
					finalSp = tSprinkes
					finalPb = tPeanut
					finalFr = tFrosting
					finalSu = tSugar
				}
				tFrosting--
			}
			tPeanut--
		}
		tSprinkes--

	}
	fmt.Printf("Part1: With %d of Sprinkes, %d of PeanutButter, %d of Frosting and %d of Sugar, Max Score is: %d\n", finalSp, finalPb, finalFr, finalSu, maxScore)
}

func partTwo(ingredients []Ingredient) {
	// Number of table spoons
	tbsp := 100
	maxScore := 0
	tSprinkes := 100
	tPeanut := 0
	tFrosting := 0
	tSugar := 0

	finalSp := 0
	finalPb := 0
	finalFr := 0
	finalSu := 0

	// Brute force the 100 tbps
	// how many tbps? 0,1,2,....99,100 | 100 has to be calculated as well
	for s := tSprinkes; s >= 0; s-- {
		tPeanut = (tbsp - tSprinkes)
		for p := tPeanut; p >= 0; p-- {
			tFrosting = (tbsp - tSprinkes - tPeanut)
			for p := tFrosting; p >= 0; p-- {
				tSugar = (tbsp - tSprinkes - tPeanut - tFrosting)

				calories := CalculateCalories([]int{tSprinkes, tPeanut, tFrosting, tSugar}, ingredients)

				score := 0
				// Only calculate Score if calories are exactly 500
				if calories == 500 {
					score = CalculateScore([]int{tSprinkes, tPeanut, tFrosting, tSugar}, ingredients)
				}

				if score > maxScore {
					maxScore = score
					finalSp = tSprinkes
					finalPb = tPeanut
					finalFr = tFrosting
					finalSu = tSugar
				}
				tFrosting--
			}
			tPeanut--
		}
		tSprinkes--

	}
	fmt.Printf("Part2: With %d of Sprinkes, %d of PeanutButter, %d of Frosting and %d of Sugar, Max Score is: %d\n", finalSp, finalPb, finalFr, finalSu, maxScore)

}

// Calculate Score of ingredients, they're weights are passed in accordance to their order in the array
func CalculateScore(weights []int, ingredients []Ingredient) int {

	capacity := CalculateCapacity(weights, ingredients)
	durability := CalculateDurability(weights, ingredients)
	flavor := CalculateFlavor(weights, ingredients)
	texture := CalculateTexture(weights, ingredients)

	score := capacity * durability * flavor * texture
	return score
}

func CalculateCapacity(weights []int, ingredients []Ingredient) int {

	capacity := 0
	for i := 0; i < len(ingredients); i++ {
		capacity += weights[i] * ingredients[i].Capacity
	}

	if capacity > 0 {
		return capacity
	} else {
		return 0
	}
}

func CalculateDurability(weights []int, ingredients []Ingredient) int {

	durability := 0
	for i := 0; i < len(ingredients); i++ {
		durability += weights[i] * ingredients[i].Durability
	}

	if durability > 0 {
		return durability
	} else {
		return 0
	}
}
func CalculateFlavor(weights []int, ingredients []Ingredient) int {

	flavor := 0
	for i := 0; i < len(ingredients); i++ {
		flavor += weights[i] * ingredients[i].Flavor
	}
	if flavor > 0 {
		return flavor
	} else {
		return 0
	}
}
func CalculateTexture(weights []int, ingredients []Ingredient) int {

	texture := 0
	for i := 0; i < len(ingredients); i++ {
		texture += weights[i] * ingredients[i].Texture
	}
	if texture > 0 {
		return texture
	} else {
		return 0
	}
}
func CalculateCalories(weights []int, ingredients []Ingredient) int {

	calories := 0
	for i := 0; i < len(ingredients); i++ {
		calories += weights[i] * ingredients[i].Calories
	}
	if calories > 0 {
		return calories
	} else {
		return 0
	}
}
