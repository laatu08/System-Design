package demo

import (
	"fmt"
	"consistent-hashing/ring"
)

func WeightedNodeDemo() {
	r := ring.NewRing(100)

	// nodeA is 3x stronger than nodeC
	r.AddNodeWithWeight("nodeA", 3)
	r.AddNodeWithWeight("nodeB", 2)
	r.AddNodeWithWeight("nodeC", 1)

	totalKeys := 10000
	counts := make(map[string]int)

	for i := 0; i < totalKeys; i++ {
		key := fmt.Sprintf("user%d", i)
		node := r.GetNode(key)
		counts[node]++
	}

	fmt.Println("---- Weighted Key Distribution ----")
	for node, count := range counts {
		fmt.Printf("%s -> %d\n", node, count)
	}
}
