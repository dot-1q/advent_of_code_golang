package main

import (
	"fmt"
	"maps"
	"math"
	"os"
	"strconv"
	"strings"

	pq "github.com/emirpasic/gods/queues/priorityqueue"
	utils "github.com/emirpasic/gods/utils"
)

type Player struct {
	Name      string
	Hitpoints int
	Damage    int
	Mana      int
	Armor     int
}

type Spell struct {
	Name         string
	Damage       int
	Cost         int
	Turns        int
	TickDamage   int
	Heal         int
	Armor        int
	ManaRecharge int
}

// This will be the State of the simulation at each point,
// Which will allow us to perform Dijkstra's shortest path to
// Victory, ie, less mana spent
type GameState struct {
	Boss      Player
	Player    Player
	ManaSpent int
	// Keep track of the turns. True = us, False means boss
	Turn       bool
	isFinished bool
	// Keep track of the effects currently active, and how many turns
	onGoingStatusEffects map[string]int
}

func main() {
	f, _ := os.ReadFile("input.txt")

	l := string(f)

	lines := strings.Split(l, "\n")

	// Dont mind this nonsense, just getting the values from the file in a compact way
	// Not very readable
	hp, _ := strconv.Atoi(strings.TrimSpace(strings.Split(lines[0], ":")[1]))
	damage, _ := strconv.Atoi(strings.TrimSpace(strings.Split(lines[1], ":")[1]))

	boss := Player{Name: "Boss", Hitpoints: hp, Damage: damage, Armor: 0}
	// boss = Player{Name: "Boss", Hitpoints: 14, Damage: 8, Armor: 0}
	// Create myself as a player to fight the boss
	myself := Player{Name: "Player", Hitpoints: 50, Mana: 500, Armor: 0}
	spells := createSpells()
	partOne(myself, boss, spells)
	partTwo(myself, boss, spells)

}

func partOne(myself, boss Player, spells map[string]Spell) {
	// Part 1
	winningGames := simulateGame(myself, boss, spells, 1)
	minCost := math.MaxInt32
	minGame := GameState{}
	for game, cost := range winningGames {
		if cost < minCost {
			minGame = deepCopy(*game)
			minCost = cost
		}
	}

	fmt.Printf("Part1: Gamestate: %+v, with cost: %d\n", minGame, minCost)

}

func partTwo(myself, boss Player, spells map[string]Spell) {
	// Part 1
	winningGames := simulateGame(myself, boss, spells, 2)
	minCost := math.MaxInt32
	minGame := GameState{}
	for game, cost := range winningGames {
		if cost < minCost {
			minGame = deepCopy(*game)
			minCost = cost
		}
	}

	fmt.Printf("Part2: Gamestate: %+v, with cost: %d\n", minGame, minCost)

}

// Simulate game and return all the games the player wins
func simulateGame(player, boss Player, spells map[string]Spell, part int) map[*GameState]int {
	// Turn = true, means its our turn
	gameState := GameState{Boss: boss, Player: player, ManaSpent: 0, Turn: true, isFinished: false, onGoingStatusEffects: map[string]int{}}
	pqueue := pq.NewWith(byPriority)
	fmt.Println("---")
	// Save the gamestates where the player wins
	playerWins := map[*GameState]int{}
	minManaCost := math.MaxInt32

	pqueue.Enqueue(gameState)
	fmt.Printf("Start: %+v\n", gameState)

	for !pqueue.Empty() {
		i, _ := pqueue.Dequeue()
		lowestGameState := i.(GameState)
		// fmt.Printf("Original: %+v\n", lowestGameState)

		listPossibleGameState := calculateNextGameState(lowestGameState, spells, part)
		for _, gs := range listPossibleGameState {
			// fmt.Printf("PossibleGS: %+v\n", gs)
			// Enqueue not finished games
			if !gs.isFinished {
				// The games can continue on til infinity,
				// so only continue enqueueing if they have lower cost than the min
				if gs.ManaSpent < minManaCost {
					// fmt.Printf("Enqueued: %+v\n", gs)
					pqueue.Enqueue(gs)
				}
			} else {
				// Game is finished, and player won, save the mana cost
				if gs.Boss.Hitpoints <= 0 {
					playerWins[&gs] = gs.ManaSpent
					// Save this minimum mana cost
					minManaCost = min(minManaCost, gs.ManaSpent)
				}
			}
		}
	}
	return playerWins
}

// From a given game state, calculate all the next possible
// game moves, meaning, spells.
func calculateNextGameState(gamestate GameState, spells map[string]Spell, part int) []GameState {
	possibleGameState := []GameState{}
	// At the start of each of my turns i take 1hp damage
	if part == 2 && gamestate.Turn {
		gamestate.Player.Hitpoints--
		// Check if game ends. The previous damage can kill me
		if gamestate.Player.Hitpoints <= 0 {
			gamestate.isFinished = true
			possibleGameState = append(possibleGameState, gamestate)
			return possibleGameState
		}
	}
	tickStatusEffects(&gamestate, spells)
	// Check if game ends. The effect tick can kill the boss
	if gamestate.Boss.Hitpoints <= 0 {
		gamestate.isFinished = true
		possibleGameState = append(possibleGameState, gamestate)
		return possibleGameState
	}
	checkEndEffects(&gamestate)

	// If its our turn
	if gamestate.Turn {
		// For each spell available, check if its a valid move, i.e, if enough mana
		for name, spell := range spells {
			// Create a new gamestate for each move
			newGS := deepCopy(gamestate)
			newGS.Turn = !newGS.Turn
			switch name {
			case "Magic Missile":
				// Can cast
				if newGS.Player.Mana-spell.Cost >= 0 {
					newGS.Player.Mana -= spell.Cost
					newGS.Boss.Hitpoints -= spell.Damage
					newGS.ManaSpent += spell.Cost
					// Check if game ends
					if newGS.Boss.Hitpoints <= 0 {
						newGS.isFinished = true
					}
					possibleGameState = append(possibleGameState, newGS)
				} else {
					// The player didn't have mana to cast anything. lost
					newGS.isFinished = true
					possibleGameState = append(possibleGameState, newGS)
				}

			case "Drain":
				// Can cast
				if newGS.Player.Mana-spell.Cost >= 0 {
					newGS.Player.Mana -= spell.Cost
					newGS.Player.Hitpoints += spell.Heal
					newGS.Boss.Hitpoints -= spell.Damage
					newGS.ManaSpent += spell.Cost
					// Check if game ends
					if newGS.Boss.Hitpoints <= 0 {
						newGS.isFinished = true
					}
					possibleGameState = append(possibleGameState, newGS)
				} else {
					// The player didn't have mana to cast anything. lost
					newGS.isFinished = true
					possibleGameState = append(possibleGameState, newGS)
				}

			case "Shield":
				// Can cast
				if newGS.Player.Mana-spell.Cost >= 0 {
					// Check if doesnt exist, meaning it can cast
					if _, ok := newGS.onGoingStatusEffects[name]; !ok {
						// Start the Effect and increase armor
						newGS.onGoingStatusEffects[name] = 6
						newGS.Player.Armor = spell.Armor
						newGS.Player.Mana -= spell.Cost
						newGS.ManaSpent += spell.Cost
						possibleGameState = append(possibleGameState, newGS)
					}
				} else {
					// The player didn't have mana to cast anything. lost
					newGS.isFinished = true
					possibleGameState = append(possibleGameState, newGS)
				}

			case "Poison":
				// Can cast
				if newGS.Player.Mana-spell.Cost >= 0 {
					// Check if doesnt exist, meaning it can cast
					if _, ok := newGS.onGoingStatusEffects[name]; !ok {
						// Start the Effect and increase armor
						newGS.onGoingStatusEffects[name] = 6
						newGS.Player.Mana -= spell.Cost
						newGS.ManaSpent += spell.Cost
						possibleGameState = append(possibleGameState, newGS)
					}

				} else {
					// The player didn't have mana to cast anything. lost
					newGS.isFinished = true
					possibleGameState = append(possibleGameState, newGS)
				}

			case "Recharge":
				// Check if doesnt exist, meaning it can cast
				if newGS.Player.Mana-spell.Cost >= 0 {
					// Check if exists and is bigger than 0
					if _, ok := newGS.onGoingStatusEffects[name]; !ok {
						// Start the Effect and increase armor
						newGS.onGoingStatusEffects[name] = 5
						newGS.Player.Mana -= spell.Cost
						newGS.ManaSpent += spell.Cost
						possibleGameState = append(possibleGameState, newGS)
					}
				} else {
					// The player didn't have mana to cast anything. lost
					newGS.isFinished = true
					possibleGameState = append(possibleGameState, newGS)
				}
			}
		}
	} else {
		// Boss's turn
		newGS := deepCopy(gamestate)
		newGS.Turn = !newGS.Turn
		newGS.Player.Hitpoints -= (newGS.Boss.Damage - newGS.Player.Armor)
		// Check if game ends. The effect tick can kill the boss
		if newGS.Player.Hitpoints <= 0 {
			newGS.isFinished = true
		}
		possibleGameState = append(possibleGameState, newGS)
	}
	return possibleGameState
}

func tickStatusEffects(gamestate *GameState, spells map[string]Spell) {
	for name := range gamestate.onGoingStatusEffects {
		switch name {
		case "Shield":
			gamestate.onGoingStatusEffects[name]--
		case "Poison":
			gamestate.onGoingStatusEffects[name]--
			gamestate.Boss.Hitpoints -= spells[name].TickDamage
		case "Recharge":
			gamestate.onGoingStatusEffects[name]--
			gamestate.Player.Mana += spells[name].ManaRecharge
		}
	}
}

func checkEndEffects(gamestate *GameState) {
	for name, turns := range gamestate.onGoingStatusEffects {
		if turns == 0 {
			delete(gamestate.onGoingStatusEffects, name)
			if name == "Shield" {
				gamestate.Player.Armor = 0
			}
		}
	}
}

func createSpells() map[string]Spell {
	var spellsMap = map[string]Spell{
		"Magic Missile": {
			Name:   "Magic Missile",
			Cost:   53,
			Damage: 4,
		},
		"Drain": {
			Name:   "Drain",
			Cost:   73,
			Damage: 2,
			Heal:   2,
		},
		"Shield": {
			Name:  "Shield",
			Cost:  113,
			Turns: 6,
			Armor: 7,
		},
		"Poison": {
			Name:       "Poison",
			Cost:       173,
			Turns:      6,
			TickDamage: 3,
		},
		"Recharge": {
			Name:         "Recharge",
			Cost:         229,
			Turns:        5,
			ManaRecharge: 101,
		},
	}
	return spellsMap
}

// The priority of each game state will be given by its mana spent
func byPriority(a, b interface{}) int {
	priorityA := a.(GameState).ManaSpent
	priorityB := b.(GameState).ManaSpent
	return utils.IntComparator(priorityA, priorityB) // "-" descending order
}

func deepCopy(gs GameState) GameState {
	return GameState{
		Boss:                 deepCopyPlayer(gs.Boss),
		Player:               deepCopyPlayer(gs.Player),
		ManaSpent:            gs.ManaSpent,
		Turn:                 gs.Turn,
		isFinished:           gs.isFinished,
		onGoingStatusEffects: maps.Clone(gs.onGoingStatusEffects),
	}
}

func deepCopyPlayer(p Player) Player {
	return Player{
		Name:      p.Name,
		Hitpoints: p.Hitpoints,
		Damage:    p.Damage,
		Mana:      p.Mana,
		Armor:     p.Armor,
	}
}
