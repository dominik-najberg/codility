package PrefixSums

import "math"

func calculatePrefixSums(A []int) []int {
	var (
		n = len(A)
		P = make([]int, n+1)
	)

	for i := 1; i < n+1; i++ {
		P[i] = P[i-1] + A[i-1]
	}

	return P
}

func calculateSuffixSums(A []int) []int {
	var (
		n = len(A)
		P = make([]int, n)
	)

	P[n-1] = A[n-1]

	for i := n - 2; i >= 0; i-- {
		P[i] = P[i+1] + A[i]
	}

	return P
}

func countTotalOfSliceOfSums(A []int, x int, y int) int {
	return int(math.Abs(float64(A[y+1] - A[x])))
}
