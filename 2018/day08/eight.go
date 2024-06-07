package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	id       int
	children []*Node
	metadata []int
	value    int
}

func main() {
	f, _ := os.ReadFile("input.txt")
	numbers := toInt(strings.Fields(strings.TrimSpace(string(f))))
	root := Node{0, []*Node{}, []int{}, 0}
	createTree(&root, 0, 0, numbers)
	// printTree(&root) // Debug purposes.
	fmt.Printf("Part 1 | Sum of the metadata: %d\n", addMetadata(&root))
	fmt.Printf("Part 2 | Value of the root node: %d\n", addMetadataWithChildren(&root))
}

func addMetadata(root *Node) int {
	s := 0
	if root != nil {
		s += root.value
		for _, children := range root.children {
			s += addMetadata(children)
		}
	}
	return s
}

func addMetadataWithChildren(root *Node) int {
	s := 0
	if root != nil {
		if len(root.children) > 0 {
			for _, childNumber := range root.metadata {
				if childNumber >= 1 && childNumber <= len(root.children) {
					s += addMetadataWithChildren(root.children[childNumber-1])
				}
			}
		} else {
			s = root.value
		}
	}
	return s
}

func createTree(root *Node, index int, id int, numbers []int) (*Node, int) {
	if index >= len(numbers) {
		return nil, 0
	}
	root.id = id
	children := numbers[index]
	metadata := numbers[index+1]
	index += 2
	// Index will always point to the number of children.
	for i := range children {
		child, j := createTree(&Node{}, index, id+1+i, numbers)
		root.children = append(root.children, child)
		index = j
	}
	sum := 0
	for range metadata {
		sum += numbers[index]
		root.metadata = append(root.metadata, numbers[index])
		index++
	}
	root.value = sum // This helps in readability to sum the metadata.
	return root, index
}

func printTree(root *Node) {
	if root != nil {
		fmt.Printf("Node: +%v\n", root)
		for _, node := range root.children {
			printTree(node)
		}
	}
}

func toInt(numbers []string) []int {
	n := make([]int, len(numbers))
	for i, number := range numbers {
		val, _ := strconv.Atoi(number)
		n[i] = val
	}
	return n
}
