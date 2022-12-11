package ds

const ROOT_REFERENCE = -1

// Maintains a collection S of disjoint dynamic sets
type DisjointSet struct {
	nodes map[int]*forestNode
}

// For a given node keep a reference of its parent and its
// associated rank
type forestNode struct {
	rank   int
	parent int
}

// Initialize an empty graph, parents indicate
func NewDisjointSet() *DisjointSet {
	return &DisjointSet{
		nodes: make(map[int]*forestNode),
	}
}

// Creates a new set whose only member is x. It asserts the fact that
// x is not already in other set.
func (s *DisjointSet) Make(x int) {
	if s.nodes == nil {
		s.nodes = make(map[int]*forestNode)
	}
	// Don't override if already exists
	if _, exists := s.nodes[x]; !exists {
		s.nodes[x] = &forestNode{rank: 0, parent: ROOT_REFERENCE}
	}
}

// Unites the dynamic sets that contain x, y into a new set
func (s *DisjointSet) Union(x, y int) {
	s.link(s.Find(x), s.Find(y))
}

// Number of DisjointSets
func (s DisjointSet) Size() int {
	count := 0
	for _, node := range s.nodes {
		if node.parent == -1 {
			count++
		}
	}
	return count
}

func (s *DisjointSet) link(x, y int) {
	if x == y {
		return // Already linked
	}
	if s.nodes[x].rank > s.nodes[y].rank {
		s.nodes[y].parent = x
	} else {
		s.nodes[x].parent = y
		if s.nodes[x].rank == s.nodes[y].rank {
			s.nodes[y].rank++
		}
	}
}

// Returns a pointer to the representative of the set containing x
func (s *DisjointSet) Find(x int) int {
	if ROOT_REFERENCE == s.nodes[x].parent {
		return x
	}
	s.nodes[x].parent = s.Find(s.nodes[x].parent)
	return s.nodes[x].parent
}
