package CountingElements

func PermutationDetectorSolution(A []int) int {
	sum := 0
	seen := make(map[int]bool)
	arrayLen := len(A)

	for _, value := range A {
		if seen[value] {
			return 0
		}

		seen[value] = true
		sum += value
	}

	if (arrayLen * (arrayLen + 1) / 2) == sum {
		return 1
	} else {
		return 0
	}
}
