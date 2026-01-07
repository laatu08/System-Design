package main

import (
	"consistent-hashing/demo"
)

func main() {
	keys := []string{
		"user1", "user2", "user3",
		"user4", "user5",
	}

	nodes := []string{
		"nodeA", "nodeB", "nodeC",
	}

	before := demo.NaiveAssign(keys, nodes)
	demo.PrintAssignments("Before removing node", before)

	// Remove one node
	nodes = []string{"nodeA", "nodeB"}

	after := demo.NaiveAssign(keys, nodes)
	demo.PrintAssignments("After removing node", after)
}
