package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
	Name          string
	Weight        int
	Children      []*Node
	Parent        *Node
	SubtreeWeight int
}

func main() {
	f, _ := os.ReadFile("input.txt")
	lines := strings.Split(strings.TrimSpace(string(f)), "\n")

	nodes := map[string]*Node{}
	for _, l := range lines {
		line := strings.Fields(l)
		name := line[0]
		weight, _ := strconv.Atoi(strings.TrimFunc(line[1], func(r rune) bool {
			return (r == '(' || r == ')')
		}))
		node := Node{Name: name, Weight: weight}
		nodes[name] = &node
	}
	makeTree(lines, nodes)
	calculateWeights(nodes)
	root := findBottom(nodes)
	fmt.Println("Part1: ", root)

	neededWeight := 0
	siblingWeight := 0
	for n := nodes[root]; n != nil && len(n.Children) > 1; {
		fmt.Printf("Root: %s \n", n.Name)
		// For all of this node's children weight+subtree weight, compact the array.
		// From this, all equal weighted elements will be discarded.
		// If this compact array has 1 element, all the children have the same weight,
		// If it has 2 elements, one is the unbalanced one.
		c := slices.CompactFunc(n.Children, func(a, b *Node) bool {
			return (a.Weight + a.SubtreeWeight) == (b.Weight + b.SubtreeWeight)
		})
		if len(c) > 1 {
			// Find the max node in the previous compacted array
			unbalanceNode := slices.MaxFunc(c, func(a, b *Node) int {
				return (a.SubtreeWeight + a.Weight) - (b.SubtreeWeight + b.Weight)
			})
			// Im assuming the Max will be c[0]. Which might not be the case. But works
			// We need this value for when we find the root of the unbalanced node.
			// We need to compare it to one of its siblings that is balanced.
			siblingWeight = c[1].SubtreeWeight + c[1].Weight
			fmt.Printf("Unbalanced is: %s \n", unbalanceNode.Name)
			n = unbalanceNode
		} else {
			// Get the diff of this unbalanced tree to a siblinc balanced tree
			weightDiff := (c[0].Parent.SubtreeWeight + c[0].Parent.Weight) - siblingWeight
			// Find the needed weight to be balanced
			neededWeight = c[0].Parent.Weight - weightDiff
			fmt.Println("No more unbalanced")
			break
		}
	}
	fmt.Println("Part 2: ", neededWeight)
}

func getChildren(line []string) []string {
	children := []string{}
	for _, child := range line {
		c, _ := strings.CutSuffix(child, ",")
		children = append(children, c)

	}
	return children
}

// Make the necessary children and parent connections.
func makeTree(lines []string, nodes map[string]*Node) {
	for _, l := range lines {
		line := strings.Fields(l)
		name := line[0]
		// info on a node that has children
		if len(line) > 3 {
			children := getChildren(line[3:])
			for _, child := range children {
				nodes[name].Children = append(nodes[name].Children, nodes[child])
				nodes[child].Parent = nodes[name]
			}
		}
	}
}

// Simply iterate over the nodes and find the one that does not have a parent
func findBottom(nodes map[string]*Node) string {
	for _, n := range nodes {
		if n.Parent == nil {
			return n.Name
		}
	}
	return ""
}

func calculateWeights(nodes map[string]*Node) {
	for _, n := range nodes {
		for k := n.Parent; k != nil; k = k.Parent {
			k.SubtreeWeight += n.Weight
		}
	}
}
