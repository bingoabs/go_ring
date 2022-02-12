package ring

import (
	"fmt"
	"hash/crc32"
	"log"
)

type HashRing struct {
	node2hashs map[string][]uint32
	hash2nodes map[uint32]string
	tree       *BalanceTree
}

func (ring *HashRing) HasNode(node string) bool {
	_, ok := ring.node2hashs[node]
	return ok
}

func (ring *HashRing) InsertNode(node string, weight int) {
	_, ok := ring.node2hashs[node]
	if ok {
		log.Panic("Hash ring insert the repeated node")
	}
	insert_number := 0
	for i := 0; i < weight; i++ {
		full_node := fmt.Sprintf("%s_%d", node, i)
		node_hash := get_int(full_node)
		// check whether it is duplicate
		_, ok = ring.hash2nodes[node_hash]
		if !ok {
			ring.tree.Add(node_hash)
			ring.node2hashs[node] = append(ring.node2hashs[node], node_hash)
			ring.hash2nodes[node_hash] = node
			insert_number++
		}
	}
	log.Println("Hash ring insert the virtual node for the node, the inserted number is ", insert_number)
}

func (ring *HashRing) RemoveNode(node string) {
	hash_nodes, ok := ring.node2hashs[node]
	if !ok {
		log.Panic("Hash ring remove unexisting nodes")
	}
	delete(ring.node2hashs, node)
	for i := 0; i < len(hash_nodes); i++ {
		hash_node := hash_nodes[i]
		ring.tree.Remove(hash_node)
		delete(ring.hash2nodes, hash_node)
	}
}

func (ring *HashRing) GetMathNode(name string) string {
	val := get_int(name)
	match_hash_value := ring.tree.Find(val)
	return ring.hash2nodes[match_hash_value]
}

func get_int(val string) uint32 {
	return crc32.ChecksumIEEE([]byte(val))
}
