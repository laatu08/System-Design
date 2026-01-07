package demo

import (
	"fmt"
	"consistent-hashing/ring"
)

func ConsistentDemo() {
	r := ring.NewRing(100) // 100 virtual nodes per physical node

	r.AddNode("nodeA")
	r.AddNode("nodeB")
	r.AddNode("nodeC")
	r.AddNode("nodeD")
	r.AddNode("nodeE")

	keys := []string{
		"user1", "user2", "user3",
		"user4", "user5", "user6",
		"user7", "user8", "user9",
	}

	fmt.Println("---- Before removing node ----")
	for _, k := range keys {
		fmt.Printf("%s -> %s\n", k, r.GetNode(k))
	}

	r.RemoveNode("nodeB")

	fmt.Println("\n---- After removing nodeB ----")
	for _, k := range keys {
		fmt.Printf("%s -> %s\n", k, r.GetNode(k))
	}

	
}
