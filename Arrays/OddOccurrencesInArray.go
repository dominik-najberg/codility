package Arrays

func OddOccurrencesInArraySolution(A []int) int {
	var e int

	for i := 0; i < len(A); i++ {
		e ^= A[i]
	}

	return e
}
