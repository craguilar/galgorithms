package graphs

import (
	"fmt"
	"sync"
)

//Node class
type Node struct {
	value int
}

func (n Node) String() string {
	return fmt.Sprintf("%d", n.value)
}

// ItemGraph the Items graph
type IntegerGraph struct {
	nodes map[int]*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

// Initialize an empty graph
func NewGraph() *IntegerGraph {
	return &IntegerGraph{
		nodes: make(map[int]*Node),
		edges: make(map[Node][]*Node),
	}
}

func (g *IntegerGraph) AddNode(n *Node) {
	g.lock.Lock()
	if g.nodes == nil {
		g.nodes = make(map[int]*Node)
	}
	_, exists := g.nodes[n.value]
	if !exists {
		g.nodes[n.value] = n
	}
	g.lock.Unlock()
}

// AddEdge adds an edge to the graph
func (g *IntegerGraph) AddEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.lock.Unlock()
}

// Topological sort of the graph
func (g *IntegerGraph) TopSort() []int {

	if len(g.nodes) == 0 {
		return make([]int, 0)
	}

	visited := make(map[int]bool)
	var rank []int
	for _, node := range g.nodes {
		if !visited[node.value] {
			rank = g.dfs(node, visited, rank)
		}
	}
	return reverse(rank)
}

func (g *IntegerGraph) dfs(n *Node, visited map[int]bool, rank []int) []int {
	visited[n.value] = true
	for _, edge := range g.edges[*n] {
		if !visited[edge.value] {
			rank = g.dfs(edge, visited, rank)
		}
	}
	rank = append(rank, n.value)
	return rank
}

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// AddEdge adds an edge to the graph
func (g *IntegerGraph) String() string {
	g.lock.RLock()
	s := ""
	for _, node := range g.nodes {
		s += node.String() + " -> "
		near := g.edges[*node]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}

	g.lock.RUnlock()
	return s
}
