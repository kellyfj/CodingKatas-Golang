package graphs

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

type Node struct {
	value    int
	adjacent []*Node
	edges    map[*Node]int // weighted edges: node -> distance
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
	node := &Node{value: value, edges: make(map[*Node]int)}
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

// For directed graphs, only add edge from node1 to node2 with weight
func (g *Graph) addEdge(node1, node2 *Node) {
	node1.addAdjacent(node2)
	node1.edges[node2] = 1 // default weight
}

// Add weighted edge from node1 to node2
func (g *Graph) addWeightedEdge(node1, node2 *Node, weight int) {
	node1.addAdjacent(node2)
	node1.edges[node2] = weight
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

// Priority queue item for Dijkstra's algorithm
type Item struct {
	node     *Node
	distance int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// Dijkstra's algorithm - returns path and shortest distance from start to target
func dijkstra(start, target *Node) ([]*Node, int) {
	distances := make(map[*Node]int)
	previous := make(map[*Node]*Node)
	visited := make(map[*Node]bool)

	// Initialize distances
	var allNodes []*Node
	if start != nil {
		// Collect all nodes reachable from start
		var collect func(*Node)
		collect = func(n *Node) {
			if visited[n] {
				return
			}
			visited[n] = true
			allNodes = append(allNodes, n)
			for _, adj := range n.adjacent {
				collect(adj)
			}
		}
		collect(start)
	}

	for _, node := range allNodes {
		distances[node] = math.MaxInt
	}
	distances[start] = 0
	visited = make(map[*Node]bool) // Reset for algorithm

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Item{node: start, distance: 0})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)
		currentNode := current.node
		currentDist := current.distance

		if visited[currentNode] {
			continue
		}
		visited[currentNode] = true

		if currentNode == target {
			// Reconstruct path
			var path []*Node
			node := target
			for node != nil {
				path = append([]*Node{node}, path...)
				node = previous[node]
			}
			return path, distances[target]
		}

		// Check all adjacent nodes
		for _, adj := range currentNode.adjacent {
			if visited[adj] {
				continue
			}

			weight := currentNode.edges[adj]
			newDist := currentDist + weight

			if newDist < distances[adj] {
				distances[adj] = newDist
				previous[adj] = currentNode
				heap.Push(&pq, &Item{node: adj, distance: newDist})
			}
		}
	}

	// Target not reachable
	return nil, -1
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
