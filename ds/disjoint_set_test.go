package ds

import (
	"testing"
)

func TestConnect(t *testing.T) {
	ds := NewDisjointSet()
	ds.Make(1)
	ds.Make(2)
	ds.Make(3)
	ds.Make(4)
	ds.Make(5)
	ds.Make(6)
	ds.Make(7)
	// Connected set S1
	ds.Union(1, 2)
	ds.Union(1, 3)
	ds.Union(2, 3)
	ds.Union(2, 4)
	// Connected set S2
	ds.Union(5, 6)
	// The third disjoint set contains only one element {7}
	if ds.Size() != 3 {
		t.Fatalf(`Incorrect result %d`, ds.Size())
	}
}
