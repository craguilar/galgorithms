package ds

import "testing"

func TestSingleConnected(t *testing.T) {
	g := NewGraph()
	for _, game := range [][]int{{2, 3}, {1, 2}, {1, 4}, {1, 3}} {
		winner := game[0]
		looser := game[1]
		g.AddNode(winner)
		g.AddNode(looser)
		g.AddEdge(winner, looser)
	}
	if g.ConnectedComponents() != 1 {
		t.Fatalf(`Incorrect result %d`, g.ConnectedComponents())
	}
}

func TestMultupleConnected(t *testing.T) {
	g := NewGraph()
	for _, game := range [][]int{{1, 2}, {2, 3}, {1, 4}, {5, 6}, {6, 7}, {9, 9}} {
		winner := game[0]
		looser := game[1]
		g.AddNode(winner)
		g.AddNode(looser)
		g.AddEdge(winner, looser)
	}
	if g.ConnectedComponents() != 3 {
		t.Fatalf(`Incorrect result %d`, g.ConnectedComponents())
	}
}
