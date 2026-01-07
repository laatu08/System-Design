package demo

import (
	"fmt"
	"consistent-hashing/hash"
	"consistent-hashing/ring"
)

// ---------- Naive hashing helpers ----------

func naiveAssign(keys []string, nodes []string) map[string]string {
	result := make(map[string]string)

	for _, key := range keys {
		h := hash.Hash(key)
		idx := int(h) % len(nodes)
		result[key] = nodes[idx]
	}

	return result
}

// ---------- Comparison demo ----------

func CompareNaiveVsConsistent() {
	totalKeys := 10000

	// Generate keys
	keys := make([]string, 0, totalKeys)
	for i := 0; i < totalKeys; i++ {
		keys = append(keys, fmt.Sprintf("user%d", i))
	}

	// ---------------- Naive hashing ----------------

	nodesNaiveBefore := []string{"nodeA", "nodeB", "nodeC"}
	nodesNaiveAfter := []string{"nodeA", "nodeC"} // remove nodeB

	beforeNaive := naiveAssign(keys, nodesNaiveBefore)
	afterNaive := naiveAssign(keys, nodesNaiveAfter)

	movedNaive := 0
	for _, k := range keys {
		if beforeNaive[k] != afterNaive[k] {
			movedNaive++
		}
	}

	naivePercent := float64(movedNaive) / float64(totalKeys) * 100

	// ---------------- Consistent hashing ----------------

	r := ring.NewRing(100)
	r.AddNode("nodeA")
	r.AddNode("nodeB")
	r.AddNode("nodeC")

	beforeConsistent := make(map[string]string)
	for _, k := range keys {
		beforeConsistent[k] = r.GetNode(k)
	}

	r.RemoveNode("nodeB")

	movedConsistent := 0
	for _, k := range keys {
		if beforeConsistent[k] != r.GetNode(k) {
			movedConsistent++
		}
	}

	consistentPercent := float64(movedConsistent) / float64(totalKeys) * 100

	// ---------------- Result ----------------

	fmt.Println("========== Rebalancing Comparison ==========")
	fmt.Printf("Total keys: %d\n\n", totalKeys)

	fmt.Println("Naive hashing:")
	fmt.Printf("  Keys moved      : %d\n", movedNaive)
	fmt.Printf("  Movement percent: %.2f%%\n\n", naivePercent)

	fmt.Println("Consistent hashing (with virtual nodes):")
	fmt.Printf("  Keys moved      : %d\n", movedConsistent)
	fmt.Printf("  Movement percent: %.2f%%\n", consistentPercent)
}
