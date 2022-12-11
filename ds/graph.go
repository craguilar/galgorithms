package ds

type integerSet map[int]bool

// Int based value Graph
// TODO: Can we make it generic to support any node?
// TODO: Can we make this Thread safe?
// TODO: Transitive reduction https://dominikbraun.io/blog/graphs/reducing-graph-complexity-using-go-and-transitive-reduction/
type Graph struct {
	nodes map[int]integerSet
}

// Initialize an empty graph
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[int]integerSet),
	}
}

func (g *Graph) AddNode(n int) {
	if g.nodes == nil {
		g.nodes = make(map[int]integerSet)
	}
	_, exists := g.nodes[n]
	if !exists {
		g.nodes[n] = integerSet{}
	}
}

// AddEdge adds an edge to the graph, if the nodes don't exist  will be created
func (g *Graph) AddEdge(n1, n2 int) {
	g.AddNode(n1)
	g.AddNode(n2)
	g.nodes[n1][n2] = true
}

// Topological sort of the graph
func (g Graph) TopSort() []int {
	if len(g.nodes) == 0 {
		return make([]int, 0)
	}

	visited := make(map[int]bool)
	var rank []int
	for key := range g.nodes {
		if !visited[key] {
			rank = g.dfs(key, visited, rank)
		}
	}
	return reverse(rank)
}

// Count the number of connected components in the graph
func (g Graph) ConnectedComponents() int {

	ds := NewDisjointSet()
	for node := range g.nodes {
		ds.Make(node)
	}
	for node, neighbors := range g.nodes {
		for neighbor := range neighbors {
			if ds.Find(node) != ds.Find(neighbor) {
				ds.Union(node, neighbor)
			}
		}
	}

	return ds.Size()
}

// Get a Spanning Tree of the Graph
func (g Graph) SpanningTree() {
	/* TODO: Kruska algorithm
	1. Initialize a Tk wich is a Spanning tree algorithm
	2. Put all edges in a min Priority Queue ordered by weight.
	3. Create a disjoint ds set ds where each node exists on its own set.
	4. For each edge e {v,w}
		 	if ds.Find(v) != ds.Find(w)
				add edge v,w to Tk
				ds.Union(v, w)
	*/
}

func (g Graph) dfs(n int, visited map[int]bool, rank []int) []int {
	visited[n] = true
	for neighbor := range g.nodes[n] {
		if !visited[neighbor] {
			rank = g.dfs(neighbor, visited, rank)
		}
	}
	rank = append(rank, n)
	return rank
}

func reverse(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
