package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Bot struct {
	Name string
	// Values that this bot olds
	Values []int
	// Associations for other bots
	GiveLow  *Bot
	GiveHigh *Bot
}

func newBot(name string) *Bot {
	return &Bot{Name: name,
		Values: []int{}}
}

func (b *Bot) giveHigh(bot *Bot) {
	b.GiveHigh = bot
}

func (b *Bot) giveLow(bot *Bot) {
	b.GiveLow = bot
}

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	insts := []string{}
	for scanner.Scan() {
		insts = append(insts, scanner.Text())
	}
	finalState := partOne(insts)
	partTwo(finalState)
}

func partOne(inst []string) map[string]*Bot {
	a := createAssociations(inst)
	// Depth first search walk in the input and find the values required
	firstNode := getNextCompleteNode(a, []string{})
	fmt.Println("First Node:", firstNode)
	updateRun(a, firstNode, []string{})
	//After
	// for _, v := range a {
	// 	fmt.Printf("%+v\n", v)
	// }
	fmt.Printf("Bot %s, compared 61 and 17\n", getNodeThatCompares(a, 61, 17))
	return a
}

func partTwo(associtiations map[string]*Bot) {
	// Multiply outputs of 0,1,2
	outputs := []string{"output0", "output1", "output2"}
	fmt.Printf("Outputs 0,1,2 have value: %d\n", multiplyOutputs(associtiations, outputs...))
}

// Looks and is absolutely disgusting, but created the association map given the input.
// Don't know any better way to do it
func createAssociations(insts []string) map[string]*Bot {
	// Create a map of all the bots created, representing the associations between them
	associations := map[string]*Bot{}
	for _, instruction := range insts {
		inst := strings.Split(instruction, " ")

		switch inst[0] {
		// Association instruction
		case "bot":
			// Check if the three bots comteplated in this isntruction exist. If not, create
			if _, ok := associations[inst[1]]; !ok {
				associations[inst[1]] = newBot(inst[1])
			}
			//Giving to output and not bot
			if inst[5] == "output" {
				if _, ok := associations["output"+inst[6]]; !ok {
					associations["output"+inst[6]] = newBot(("output" + inst[6]))
				}
				associations[inst[1]].giveLow(associations["output"+inst[6]])
			} else {
				if _, ok := associations[inst[6]]; !ok {
					associations[inst[6]] = newBot(inst[6])
				}
				associations[inst[1]].giveLow(associations[inst[6]])
			}
			//Giving to output and not bot
			if inst[10] == "output" {
				if _, ok := associations["output"+inst[11]]; !ok {
					associations["output"+inst[11]] = newBot(("output" + inst[11]))
				}
				associations[inst[1]].giveHigh(associations["output"+inst[11]])
			} else {
				if _, ok := associations[inst[11]]; !ok {
					associations[inst[11]] = newBot(inst[11])
				}
				associations[inst[1]].giveHigh(associations[inst[11]])
			}
		case "value":
			// Doesnt exist
			if _, ok := associations[inst[5]]; !ok {
				// Create
				associations[inst[5]] = newBot(inst[5])
				v, _ := strconv.Atoi(inst[1])
				associations[inst[5]].Values = append(associations[inst[5]].Values, v)

			} else {
				v, _ := strconv.Atoi(inst[1])
				associations[inst[5]].Values = append(associations[inst[5]].Values, v)
			}
		}
	}
	return associations
}

// Perform Breadth First Search on the bots associations
func updateRun(associations map[string]*Bot, current string, seen []string) {
	if slices.Contains(seen, current) {
		return
	}

	if current == "" {
		return
	}

	if len(associations[current].Values) == 2 {
		seen = append(seen, current)
	}
	minV := slices.Min(associations[current].Values)
	maxV := slices.Max(associations[current].Values)
	associations[current].GiveLow.Values = append(associations[current].GiveLow.Values, minV)
	associations[current].GiveHigh.Values = append(associations[current].GiveHigh.Values, maxV)

	//Recurse only if its not an output node, which is terminal
	nextNode := getNextCompleteNode(associations, seen)
	updateRun(associations, nextNode, seen)
}

// Geth the nodes with can move, ie, hold two values, and have not yet been visited
func getNextCompleteNode(associations map[string]*Bot, seen []string) string {
	for name, bot := range associations {
		// Not visited and is not an output node, meaning a terminal one
		if !slices.Contains(seen, name) && !strings.Contains(name, "output") {
			// Has 2 values, ready to move. Return
			if len(bot.Values) == 2 {
				return name
			}
		}
	}
	return ""
}

func getNodeThatCompares(associations map[string]*Bot, a, b int) string {
	for name, bot := range associations {
		// This bot compared the two values we wanted
		if slices.Contains(bot.Values, a) && slices.Contains(bot.Values, b) {
			return name
		}
	}
	return ""
}

func multiplyOutputs(associations map[string]*Bot, outputs ...string) int {
	value := 1
	for _, output := range outputs {
		// Outputs only have one value
		value *= associations[output].Values[0]
	}
	return value
}
