package TimeComplexity

func FrogJmpSolution(X int, Y int, D int) int {
	jumps := (Y - X) / D

	if (Y-X)%D > 0 {
		jumps++
	}

	return jumps
}
