package math

import "testing"

func TestCollatzConjecture(t *testing.T) {

	result := CollatzConjecture(3)
	if !equals(result, []int{3, 10, 5, 16, 8, 4, 2, 1}) {
		t.Fatalf(`Array %v is not sorted`, result)
	}
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
