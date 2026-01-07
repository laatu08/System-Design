package hash

import "hash/fnv"

func Hash(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}


// Fowler–Noll–Vo (FNV) hash algorithm.