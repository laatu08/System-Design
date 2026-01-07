package demo

import (
	"fmt"
	"consistent-hashing/hash"
)

func NaiveAssign(keys []string, nodes []string) map[string]string {
	result := make(map[string]string)

	for _, key := range keys {
		h := hash.Hash(key)
		idx := int(h) % len(nodes)
		result[key] = nodes[idx]
	}
	return result
}

func PrintAssignments(title string, m map[string]string) {
	fmt.Println("----", title, "----")
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}
	fmt.Println()
}
