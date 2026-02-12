package graphs

import (
	"fmt"
	"strings"
)

type Node struct {
	value    int
	adjacent []*Node
	previous *Node
}

func (n *Node) addAdjacent(node *Node) {
	n.adjacent = append(n.adjacent, node)
}

func (n *Node) removeAdjacent(node *Node) {
	for i, adj := range n.adjacent {
		if adj == node {
			n.adjacent = append(n.adjacent[:i], n.adjacent[i+1:]...)
			return
		}
	}
}

type Graph struct {
	nodes []*Node
}

func (g *Graph) addNode(value int) *Node {
	node := &Node{value: value}
	g.nodes = append(g.nodes, node)
	return node
}

func (g *Graph) removeNode(node *Node) {
	// Remove the node from the graph's node list
	for i, n := range g.nodes {
		if n == node {
			g.nodes = append(g.nodes[:i], g.nodes[i+1:]...)
			break
		}
	}
	// Remove the node from all adjacency lists
	for _, n := range g.nodes {
		n.removeAdjacent(node)
	}
}

// For directed graphs, only add edge from node1 to node2
func (g *Graph) addEdge(node1, node2 *Node) {
	node1.addAdjacent(node2)
}

// For directed graphs, only remove edge from node1 to node2
func (g *Graph) removeEdge(node1, node2 *Node) {
	node1.removeAdjacent(node2)
}

func depthFirstSearch(start, target *Node) bool {
	visited := make(map[*Node]bool)
	var dfs func(node *Node) bool
	dfs = func(node *Node) bool {
		if node == target {
			return true
		}
		visited[node] = true
		for _, adj := range node.adjacent {
			if !visited[adj] {
				if dfs(adj) {
					return true
				}
			}
		}
		return false
	}
	return dfs(start)
}

func breadthFirstSearch(start, target *Node) bool {
	visited := make(map[*Node]bool)
	queue := []*Node{start}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == target {
			return true
		}
		visited[node] = true
		for _, adj := range node.adjacent {
			if !visited[adj] {
				queue = append(queue, adj)
				visited[adj] = true
			}
		}
	}
	return false
}

func (g *Graph) prettyPrint() string {
	result := ""
	for _, node := range g.nodes {
		result += "Node " + fmt.Sprintf("%d", node.value) + ": ["
		adjVals := []string{}
		for _, adj := range node.adjacent {
			adjVals = append(adjVals, fmt.Sprintf("%d", adj.value))
		}
		result += strings.Join(adjVals, " ") + "]\n"
	}
	return result
}
