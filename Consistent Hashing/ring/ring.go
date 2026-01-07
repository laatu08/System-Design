package ring

import (
	"sort"
	"consistent-hashing/hash"
)

type Ring struct {
	nodes map[uint32]string
	keys  []uint32 // sorted hashes
}

func NewRing() *Ring {
	return &Ring{
		nodes: make(map[uint32]string),
		keys:  []uint32{},
	}
}


func (r *Ring) AddNode(node string) {
	h := hash.Hash(node)

	r.nodes[h] = node
	r.keys = append(r.keys, h)

	sort.Slice(r.keys, func(i, j int) bool {
		return r.keys[i] < r.keys[j]
	})
}

func (r *Ring) RemoveNode(node string) {
	h := hash.Hash(node)

	delete(r.nodes, h)

	for i, key := range r.keys {
		if key == h {
			r.keys = append(r.keys[:i], r.keys[i+1:]...)
			break
		}
	}
}

func (r *Ring) GetNode(key string) string {
	if len(r.keys) == 0 {
		return ""
	}

	h := hash.Hash(key)

	// Binary search for first node >= key hash
	idx := sort.Search(len(r.keys), func(i int) bool {
		return r.keys[i] >= h
	})

	// Wrap around
	if idx == len(r.keys) {
		idx = 0
	}

	return r.nodes[r.keys[idx]]
}
