package CountingElements

func MissingIntegerSolution(A []int) int {
	p := 1
	seen := make(map[int]bool)

	for _, v := range A {
		if v > 0 && !seen[v] {
			seen[v] = true
		}
	}

	for i := p; i <= len(seen); i += 1 {
		if !seen[p] {
			return p
		}
		p++
	}

	return p
}
