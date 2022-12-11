package ds

import "testing"

func TestFindWinner(t *testing.T) {
	result := FindWinner([][]int{{2, 3}, {1, 2}, {1, 4}, {1, 3}})
	if result != 1 {
		t.Fatalf(`Incorrect result %d`, result)
	}
}

func TestListRankr(t *testing.T) {

	rank := ListRank([][]int{{3, 1}, {2, 3}})
	if !arrayEquals(rank, []int{2, 3, 1}) {
		t.Fatalf(`Incorrect result %v`, rank)
	}

	rank = ListRank([][]int{{2, 3}, {1, 2}})
	if !arrayEquals(rank, []int{1, 2, 3}) {
		t.Fatalf(`Incorrect result %v`, rank)
	}
}

func arrayEquals(a, b []int) bool {
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
