package ring

import (
	"consistent-hashing/hash"
	"sort"
	"strconv"
)

type Ring struct {
	nodes    map[uint32]string // vnode hash -> physical node
	keys     []uint32          // sorted vnode hashes
	replicas int               // virtual nodes per physical node
}


func NewRing(replicas int) *Ring {
	return &Ring{
		nodes:    make(map[uint32]string),
		keys:     []uint32{},
		replicas: replicas,
	}
}


func (r *Ring) AddNode(node string) {
	for i := 0; i < r.replicas; i++ {
		vnodeKey := node + "#" + strconv.Itoa(i)
		h := hash.Hash(vnodeKey)

		r.nodes[h] = node
		r.keys = append(r.keys, h)
	}

	sort.Slice(r.keys, func(i, j int) bool {
		return r.keys[i] < r.keys[j]
	})
}


func (r *Ring) RemoveNode(node string) {
	for i := 0; i < r.replicas; i++ {
		vnodeKey := node + "#" + strconv.Itoa(i)
		h := hash.Hash(vnodeKey)

		delete(r.nodes, h)

		for j, key := range r.keys {
			if key == h {
				r.keys = append(r.keys[:j], r.keys[j+1:]...)
				break
			}
		}
	}
}


func (r *Ring) GetNode(key string) string {
	if len(r.keys) == 0 {
		return ""
	}

	h := hash.Hash(key)

	idx := sort.Search(len(r.keys), func(i int) bool {
		return r.keys[i] >= h
	})

	if idx == len(r.keys) {
		idx = 0
	}

	return r.nodes[r.keys[idx]]
}
