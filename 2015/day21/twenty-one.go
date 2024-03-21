package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Weapon struct {
	Name   string
	Damage int
	Cost   int
	id     int
}

type Armor struct {
	Name  string
	Armor int
	Cost  int
	id    int
}

type Ring struct {
	Name   string
	Damage int
	Armor  int
	Cost   int
	id     int
}

type Player struct {
	Name      string
	Hitpoints int
	Damage    int
	Armor     int
}

func main() {
	f, _ := os.ReadFile("input.txt")

	l := string(f)

	lines := strings.Split(l, "\n")

	// Dont mind this nonsense, just getting the values from the file in a compact way
	// Not very readable
	hp, _ := strconv.Atoi(strings.TrimSpace(strings.Split(lines[0], ":")[1]))
	damage, _ := strconv.Atoi(strings.TrimSpace(strings.Split(lines[1], ":")[1]))
	armor, _ := strconv.Atoi(strings.TrimSpace(strings.Split(lines[2], ":")[1]))

	boss := Player{Name: "Boss", Hitpoints: hp, Damage: damage, Armor: armor}
	simulateGame(boss)
}

func simulateGame(boss Player) {
	// Create myself as a player to fight the boss
	myself := Player{Name: "me", Hitpoints: 100}
	fmt.Println(myself)
	weapons, armors, rings := createItemShop()
	cost_map, attack_map, defense_map := createMap(armors, rings)

	// We always have to choose a weapon.
	// Armor and rings are optional, and rings can be between 0-2
	all_layouts := layoutsArmorRings(armors, rings)
	// fmt.Println(all_layouts)
	// Uncomment the line above too see the combinations, it paints a clearer picture

	min_cost := math.MaxInt32
	max_cost := 0
	for _, w := range weapons {
		// Calculate the cost for each layout
		for _, layout := range all_layouts {
			// Add this weapon cost
			cost := layoutCost(layout, cost_map) + w.Cost
			attack := layoutAttack(layout, attack_map) + w.Damage
			defense := layoutDefense(layout, defense_map)

			// Damage we receive
			damage_rec := boss.Damage - defense
			// The minimum is 1 hp
			if damage_rec <= 0 {
				damage_rec = 1
			}
			damage_give := attack - boss.Armor
			// Number of hits to kill boss | Round up
			kill_boss := (boss.Hitpoints / (damage_give)) + 1
			// Number of hits to kill me | Round up
			kill_me := (myself.Hitpoints / (damage_rec)) + 1

			// Since i go first, I win
			if kill_boss <= kill_me {
				min_cost = min(min_cost, cost)
			} else {
				//  I loose
				max_cost = max(max_cost, cost)
			}
		}
	}
	fmt.Println("Part1: Min cost to kill the boss:", min_cost)
	fmt.Println("Part2: Max cost to kill the boss and still loose:", max_cost)
}

func layoutDefense(layout []int, id_map map[int]int) int {
	defense := 0

	// From each item id, get the cost associated with the item
	for _, item_id := range layout {
		defense += id_map[item_id]
	}

	return defense
}

func layoutAttack(layout []int, id_map map[int]int) int {
	attack := 0

	// From each item id, get the cost associated with the item
	for _, item_id := range layout {
		attack += id_map[item_id]
	}

	return attack
}

func layoutCost(layout []int, id_map map[int]int) int {
	cost := 0

	// From each item id, get the cost associated with the item
	for _, item_id := range layout {
		cost += id_map[item_id]
	}

	return cost
}

// Chose a layout of one armor and rings, and return it and its cost
func layoutsArmorRings(armors []Armor, rings []Ring) [][]int {
	r := []int{}
	for _, ring := range rings {
		r = append(r, ring.id)
	}

	// Basically we just need to calculate all the ring combinations, which by definition do not repeat.
	// The array of rings already contain the EMPTY ring, which is essentially no Ring.
	layouts := [][]int{}
	for _, armor := range armors {
		items := []int{}
		items = append(items, r...)
		c := generateCombinations(items, 2)
		// Manually add the 2 empty rings to every armor layout
		c = append(c, []int{17, 17})

		// For each combinations of rings, add this armor, so it creates a layout
		for _, d := range c {
			d = append(d, armor.id)
			layouts = append(layouts, d)
		}
	}

	// return the layout combinations
	return layouts
}

func createMap(armors []Armor, rings []Ring) (map[int]int, map[int]int, map[int]int) {
	cost_map := map[int]int{}
	attack_map := map[int]int{}
	defense_map := map[int]int{}

	for _, a := range armors {
		cost_map[a.id] = a.Cost
		defense_map[a.id] = a.Armor
	}
	for _, r := range rings {
		cost_map[r.id] = r.Cost
		attack_map[r.id] = r.Damage
		defense_map[r.id] = r.Armor
	}

	return cost_map, attack_map, defense_map
}

// return the list of Ids as well
func createItemShop() ([]Weapon, []Armor, []Ring) {
	// Create Item Shop
	/* Weapons:    Cost  Damage  Armor
	   Dagger        8     4       0
	   Shortsword   10     5       0
	   Warhammer    25     6       0
	   Longsword    40     7       0
	   Greataxe     74     8       0

	   Armor:      Cost  Damage  Armor
	   Leather      13     0       1
	   Chainmail    31     0       2
	   Splintmail   53     0       3
	   Bandedmail   75     0       4
	   Platemail   102     0       5

	   Rings:      Cost  Damage  Armor
	   Damage +1    25     1       0
	   Damage +2    50     2       0
	   Damage +3   100     3       0
	   Defense +1   20     0       1
	   Defense +2   40     0       2
	   Defense +3   80     0       3 */

	// Since Armor and Rings can be optional, meaning we have NOT buy them, add an empty item to their list

	w := []Weapon{{Name: "Dagger", Cost: 8, Damage: 4, id: 0}, {Name: "Shortsword", Cost: 10, Damage: 5, id: 1}, {Name: "Warhammer", Cost: 25, Damage: 6, id: 2}, {Name: "Longsword", Cost: 40, Damage: 7, id: 3}, {Name: "Greataxe", Cost: 74, Damage: 8, id: 4}}
	a := []Armor{{Name: "Leather", Cost: 13, Armor: 1, id: 5}, {Name: "Chainmail", Cost: 31, Armor: 2, id: 6}, {Name: "Splintmail", Cost: 53, Armor: 3, id: 7}, {Name: "Bandedmail", Cost: 75, Armor: 4, id: 8}, {Name: "Platemail", Cost: 102, Armor: 5, id: 9}, {Name: "Empty", Cost: 0, Armor: 0, id: 10}}
	r := []Ring{{Name: "Damage1", Cost: 25, Damage: 1, Armor: 0, id: 11}, {Name: "Damage2", Cost: 50, Damage: 2, Armor: 0, id: 12}, {Name: "Damage3", Cost: 100, Damage: 3, Armor: 0, id: 13}, {Name: "Defense1", Cost: 20, Damage: 0, Armor: 1, id: 14}, {Name: "Defense2", Cost: 40, Damage: 0, Armor: 2, id: 15}, {Name: "Defense3", Cost: 80, Damage: 0, Armor: 3, id: 16}, {Name: "Empty", Cost: 0, Armor: 0, id: 17}}

	return w, a, r
}

func generateCombinations(numbers []int, length int) [][]int {
	if length == 0 {
		return [][]int{{}}
	}

	combs := [][]int{}
	for i, n := range numbers {
		// Create an empty list with this first element
		l := append([]int{}, n)

		// Generate the lists taking out the first element, and with -1 length
		lists := generateCombinations(numbers[i+1:], length-1)

		// For all the different lists of length -1, append the elements, to the original
		// list 'l', which only had the first element.
		// After that, append those lists to the list of all combinations
		for _, list := range lists {

			combs = append(combs, append(l, list...))
		}
	}
	return combs
}
