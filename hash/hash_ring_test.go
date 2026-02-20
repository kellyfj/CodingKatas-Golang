package hash

import "testing"

func TestConsistentHashRing(t *testing.T) {
	ring := NewConsistentHashRing()
	ring.AddNode("Node1")
	ring.AddNode("Node2")
	ring.AddNode("Node3")

	keys := []string{"Key1", "Key2", "Key3", "Key4", "Key5"}

	for _, key := range keys {
		ring.AddKey(key)
	}

	for _, key := range keys {
		node := ring.GetNode(key)
		if node == "" {
			t.Errorf("Expected a node for key %s, got empty string", key)
		}
	}

	ring.RemoveNode("Node2")

	for _, key := range keys {
		node := ring.GetNode(key)
		if node == "Node2" {
			t.Errorf("Expected Node2 to be removed, but got it for key %s", key)
		}
	}
}

func TestConsistentHashRingEmpty(t *testing.T) {
	ring := NewConsistentHashRing()
	node := ring.GetNode("Key1")
	if node != "" {
		t.Errorf("Expected empty string for key when no nodes are present, got %s", node)
	}
}

func TestConsistentHashRingSingleNode(t *testing.T) {
	ring := NewConsistentHashRing()
	ring.AddNode("Node1")
	node := ring.GetNode("Key1")
	if node != "Node1" {
		t.Errorf("Expected Node1 for key, got %s", node)
	}
}
