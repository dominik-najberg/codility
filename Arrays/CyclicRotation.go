package Arrays

func CyclicRotationSolution(A []int, K int) []int {
	for i := 0; i < K; i++ {
		A = rotate(A)
	}

	return A
}

func rotate(A []int) []int {
	if len(A) == 0 {
		return A
	}

	le := A[len(A)-1]
	tr := A[0 : len(A)-1]

	return append([]int{le}, tr...)
}
