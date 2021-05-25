package sorting

import (
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {

	sorted := MergeSort([]int{4, 3, 2, 1})

	if !sort.IntsAreSorted(sorted) {
		t.Fatalf(`Array %v is not sorted`, sorted)
	}

	sorted = MergeSort([]int{2, 3, 1})

	if !sort.IntsAreSorted(sorted) {
		t.Fatalf(`Array %v is not sorted`, sorted)
	}

	sorted = MergeSort([]int{1})

	if !sort.IntsAreSorted(sorted) {
		t.Fatalf(`Array %v is not sorted`, sorted)
	}

	sorted = MergeSort([]int{3, 1})

	if !sort.IntsAreSorted(sorted) {
		t.Fatalf(`Array %v is not sorted`, sorted)
	}

	sorted = MergeSort([]int{2, 4, 5, 6, 8, 9, 4, 6, 8, 9, 4, 6, 5, 6, 8, 9, 4, 6, 8, 9, 4})
	if !sort.IntsAreSorted(sorted) {
		t.Fatalf(`Array %v is not sorted`, sorted)
	}
}
