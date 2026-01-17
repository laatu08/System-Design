package demo

import (
	"fmt"
	"consistent-hashing/ring"
)

// Test consistent hashing behavior with many keys
func LargeKeyTest() {
	r := ring.NewRing(100) // 100 virtual nodes per physical node

	// Add nodes
	nodes := []string{"nodeA", "nodeB", "nodeC"	}
	for _, n := range nodes {
		r.AddNode(n)
	}

	// Generate many keys
	totalKeys := 10000
	counts := make(map[string]int)

	for i := 0; i < totalKeys; i++ {
		key := fmt.Sprintf("user%d", i)
		node := r.GetNode(key)
		counts[node]++
	}

	fmt.Println("---- Key distribution BEFORE removing node ----")
	for node, count := range counts {
		fmt.Printf("%s -> %d\n", node, count)
	}

	// Remove one node
	fmt.Println("\nRemoving nodeB...")
	r.RemoveNode("nodeB")

	// Recount after removal
	counts = make(map[string]int)
	for i := 0; i < totalKeys; i++ {
		key := fmt.Sprintf("user%d", i)
		node := r.GetNode(key)
		counts[node]++
	}

	fmt.Println("---- Key distribution AFTER removing nodeB ----")
	for node, count := range counts {
		fmt.Printf("%s -> %d\n", node, count)
	}
}
