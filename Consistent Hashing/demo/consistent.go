package demo

import (
	"fmt"
	"consistent-hashing/ring"
)

func ConsistentDemo() {
	r := ring.NewRing()

	r.AddNode("nodeA")
	r.AddNode("nodeB")
	r.AddNode("nodeC")

	keys := []string{
		"user1", "user2", "user3",
		"user4", "user5",
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
