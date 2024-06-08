package main

import (
	"fmt"
)

// Node represents a node in the circular doubly linked list
type Node struct {
	value int
	right *Node
	left  *Node
}

// CircularDoublyLinkedList represents the circular doubly linked list
type CircularDoublyLinkedList struct {
	head *Node
}

// Append adds a new node with the given value next to the specified node
func (list *CircularDoublyLinkedList) Append(current *Node, value int) {
	newNode := &Node{value: value}
	if list.head == nil {
		newNode.right = newNode
		newNode.left = newNode
		list.head = newNode
	} else if current == nil {
		// If current is nil, append at the end of the list
		tail := list.head.left
		tail.right = newNode
		newNode.left = tail
		newNode.right = list.head
		list.head.left = newNode
	} else {
		newNode.right = current.right
		newNode.left = current
		current.right.left = newNode
		current.right = newNode
	}
}

func (list *CircularDoublyLinkedList) RemoveCurrent(node *Node) {
	node.left.right = node.right
	node.right.left = node.left
	if node == list.head {
		list.head = node.right
	}
	// fmt.Printf("Node with value %d removed\n", node.value) // Debug purposes
}

func main() {
	list := CircularDoublyLinkedList{}
	list.Append(nil, 0) // Start of the list
	part := 2

	// Im not going to parse a one line text file.
	const players = 425
	lastMarble := 70848
	if part == 2 {
		lastMarble = 70848 * 100
	}
	elves := [players]int{}

	// Start of the list
	current := list.head
	for i := 1; i <= lastMarble; i++ {
		if i%23 == 0 {
			player := (i % players) // list starts at 0
			elves[player] += i      // Add to score
			for range 7 {
				current = current.left
			}
			elves[player] += current.value // Add to score again
			list.RemoveCurrent(current)    // Remove
		} else {
			current = current.right // Go 1 to the right
			list.Append(current, i) // Insert
		}
		// Go clockwise
		current = current.right
	}
	fmt.Printf("Part %d | Highest Score: %d\n", part, getHighestScore(elves[:]))
}

func getHighestScore(elves []int) int {
	highestScore := 0
	for _, value := range elves {
		if value > highestScore {
			highestScore = value
		}
	}
	return highestScore
}

// PrintValues prints all the values in the circular doubly linked list
func (list *CircularDoublyLinkedList) PrintValues() {
	if list.head == nil {
		fmt.Println("The list is empty")
		return
	}

	current := list.head
	for {
		fmt.Printf("%d ", current.value)
		current = current.right
		if current == list.head {
			break
		}
	}
	fmt.Println()
}
