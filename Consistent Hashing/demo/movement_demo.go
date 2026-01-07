package demo

import (
	"fmt"
	"consistent-hashing/ring"
)

// Measure how many keys move when a node is removed
func KeyMovementDemo() {
	r := ring.NewRing(100)

	nodes := []string{"nodeA", "nodeB", "nodeC"}
	for _, n := range nodes {
		r.AddNode(n)
	}

	totalKeys := 10000

	// Step 1: initial assignment
	before := make(map[string]string)

	for i := 0; i < totalKeys; i++ {
		key := fmt.Sprintf("user%d", i)
		before[key] = r.GetNode(key)
	}

	// Step 2: remove a node
	fmt.Println("Removing nodeB...\n")
	r.RemoveNode("nodeB")

	// Step 3: reassignment
	moved := 0

	for i := 0; i < totalKeys; i++ {
		key := fmt.Sprintf("user%d", i)
		after := r.GetNode(key)

		if before[key] != after {
			moved++
		}
	}

	percentage := float64(moved) / float64(totalKeys) * 100

	fmt.Println("---- Key Movement Report ----")
	fmt.Printf("Total keys      : %d\n", totalKeys)
	fmt.Printf("Keys moved      : %d\n", moved)
	fmt.Printf("Movement percent: %.2f%%\n", percentage)
}
