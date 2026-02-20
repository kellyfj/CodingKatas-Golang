package hash

type ConsistentHashRing struct {
	nodes       []string
	keysToNodes map[string]string
}

func NewConsistentHashRing() *ConsistentHashRing {
	return &ConsistentHashRing{
		nodes:       []string{},
		keysToNodes: make(map[string]string),
	}
}

func (r *ConsistentHashRing) AddKey(key string) {
	node := r.GetNode(key)
	r.keysToNodes[key] = node
}

func (r *ConsistentHashRing) AddNode(node string) {
	r.nodes = append(r.nodes, node)
}

func (r *ConsistentHashRing) GetNode(key string) string {
	if len(r.nodes) == 0 {
		return ""
	}
	hash := hashKey(key)
	index := hash % uint32(len(r.nodes))
	return r.nodes[index]
}

func hashKey(key string) uint32 {
	var hash uint32 = 2166136261
	for i := 0; i < len(key); i++ {
		hash = (hash * 16777619) ^ uint32(key[i])
	}
	return hash
}

func (r *ConsistentHashRing) RemoveNode(node string) {
	for i, n := range r.nodes {
		if n == node {
			r.nodes = append(r.nodes[:i], r.nodes[i+1:]...)
			return
		}
	}
}
