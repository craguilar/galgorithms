package math

//The Collatz conjecture is a conjecture in mathematics that concerns sequences defined as follows:
// start with any positive integer n. Then each term is obtained from the previous term as follows:
// - If the previous term is even, the next term is one half of the previous term.
// - If the previous term is odd, the next term is 3 times the previous term plus 1.
func CollatzConjecture(n int) []int {
	sequence := make([]int, 0)
	current := n
	for current > 1 {
		sequence = append(sequence, current)
		if current%2 == 0 {
			current = current / 2
		} else {
			current = 3*current + 1
		}
	}
	sequence = append(sequence, current)
	return sequence
}
