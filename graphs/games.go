package graphs

// In a tennis tournament of N players every player plays with every other player.
// The following condition always hold:
// If player P1 has won the match with P2 and player P2 has won from P3, then Player P1 has also defeated P3.
// Plus: Find winner of tournament in O(N) time and O(1) space. Find rank of players in O(NlogN) time.
func FindWinner(games [][]int) int {
	g := NewGraph()
	for _, game := range games {
		winner := Node{game[0]}
		looser := Node{game[1]}
		g.AddNode(&winner)
		g.AddNode(&looser)
		g.AddEdge(&winner, &looser)
	}
	// Find player with max degree
	winner := -1
	min_degree := 0
	for _, node := range g.nodes {
		if len(g.edges[*node]) > min_degree {
			winner = node.value
			min_degree = len(g.edges[*node])
		}
	}

	return winner
}

// Score
func ListRank(games [][]int) []int {
	g := NewGraph()
	// Build graph
	for _, game := range games {
		winner := Node{game[0]}
		looser := Node{game[1]}
		g.AddNode(&winner)
		g.AddNode(&looser)
		g.AddEdge(&winner, &looser)
	}
	return g.TopSort()
}
