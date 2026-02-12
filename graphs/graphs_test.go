package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph(t *testing.T) {
	g := &Graph{}

	// Add nodes
	node1 := g.addNode(1)
	node2 := g.addNode(2)
	node3 := g.addNode(3)

	assert.Equal(t, 3, len(g.nodes))

	// Add edges
	g.addEdge(node1, node2)
	g.addEdge(node2, node3)

	assert.Contains(t, node1.adjacent, node2)
	assert.NotContains(t, node2.adjacent, node1) // Directed: node2 should not have node1 as adjacent
	assert.Contains(t, node2.adjacent, node3)
	assert.NotContains(t, node3.adjacent, node2) // Directed: node3 should not have node2 as adjacent

	// Remove edge
	g.removeEdge(node1, node2)
	assert.NotContains(t, node1.adjacent, node2)
	assert.NotContains(t, node2.adjacent, node1)

	// Remove node
	g.removeNode(node2)
	assert.NotContains(t, g.nodes, node2)
	assert.NotContains(t, node1.adjacent, node2)
	assert.NotContains(t, node3.adjacent, node2)
}

func TestDFS(t *testing.T) {
	g := &Graph{}

	// Create nodes
	node1 := g.addNode(1)
	node2 := g.addNode(2)
	node3 := g.addNode(3)
	node4 := g.addNode(4)

	// Create edges
	g.addEdge(node1, node2)
	g.addEdge(node1, node3)
	g.addEdge(node2, node4)

	// Test DFS
	assert.True(t, depthFirstSearch(node1, node4))
	assert.False(t, depthFirstSearch(node3, node4))
}

func TestBFS(t *testing.T) {
	g := &Graph{}

	// Create nodes
	node1 := g.addNode(1)
	node2 := g.addNode(2)
	node3 := g.addNode(3)
	node4 := g.addNode(4)

	// Create edges
	g.addEdge(node1, node2)
	g.addEdge(node1, node3)
	g.addEdge(node2, node4)

	// Test BFS
	assert.True(t, breadthFirstSearch(node1, node4))
	assert.False(t, breadthFirstSearch(node3, node4))
}

func TestDFSWithCycle(t *testing.T) {
	g := &Graph{}

	// Create nodes
	node1 := g.addNode(1)
	node2 := g.addNode(2)
	node3 := g.addNode(3)

	// Create edges with a cycle
	g.addEdge(node1, node2)
	g.addEdge(node2, node3)
	g.addEdge(node3, node1) // Cycle

	// Test DFS
	assert.True(t, depthFirstSearch(node1, node3))
	assert.True(t, depthFirstSearch(node2, node1))
	assert.True(t, depthFirstSearch(node3, node2))
}

func TestBFSWithCycle(t *testing.T) {

	g := &Graph{}

	// Create nodes
	node1 := g.addNode(1)
	node2 := g.addNode(2)
	node3 := g.addNode(3)

	// Create edges with a cycle
	g.addEdge(node1, node2)
	g.addEdge(node2, node3)
	g.addEdge(node3, node1) // Cycle

	// Test BFS
	assert.True(t, breadthFirstSearch(node1, node3))
	assert.True(t, breadthFirstSearch(node2, node1))
	assert.True(t, breadthFirstSearch(node3, node2))

	// Test BFS with a node that is not connected
	node4 := g.addNode(4)
	assert.False(t, breadthFirstSearch(node1, node4))
	assert.False(t, breadthFirstSearch(node2, node4))
	assert.False(t, breadthFirstSearch(node3, node4))
	assert.False(t, breadthFirstSearch(node4, node1))
	assert.False(t, breadthFirstSearch(node4, node2))
	assert.False(t, breadthFirstSearch(node4, node3))

}

func TestPrettyPrint(t *testing.T) {

	g := &Graph{}

	// Create nodes
	node1 := g.addNode(1)
	node2 := g.addNode(2)
	node3 := g.addNode(3)
	// Create edges
	g.addEdge(node1, node2)
	g.addEdge(node1, node3)

	expected := "Node 1: [2 3]\nNode 2: []\nNode 3: []\n"
	assert.Equal(t, expected, g.prettyPrint())
}

func TestPrettyPrintEmptyGraph(t *testing.T) {

	g := &Graph{}
	expected := ""
	assert.Equal(t, expected, g.prettyPrint())
}

func TestPrettyPrintSingleNode(t *testing.T) {
	g := &Graph{}
	_ = g.addNode(1)
	expected := "Node 1: []\n"
	assert.Equal(t, expected, g.prettyPrint())
}
