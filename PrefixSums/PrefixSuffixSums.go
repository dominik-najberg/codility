package PrefixSums

/**
Based on the following lesson:
https://codility.com/media/train/3-PrefixSums.pdf
*/

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

func mushrooms(A []int, startPos int, steps int) int {
	n := len(A)
	result := 0
	pref := calculatePrefixSums(A)

	for p := 0; p < min(steps, startPos); p++ {
		leftPos := startPos - p
		rightPos := min(n-1, max(startPos, startPos+steps-2*p))
		result = max(result, countTotalOfSliceOfSums(pref, leftPos, rightPos))
	}

	for p := 0; p < min(steps+1, n-startPos); p++ {
		rightPos := startPos + p
		leftPos := max(0, min(startPos, startPos-(steps-2*p)))
		result = max(result, countTotalOfSliceOfSums(pref, leftPos, rightPos))
	}

	return result
}

func min(vs ...int) int {
	min := math.MaxInt32
	for _, v := range vs {
		if min > v {
			min = v
		}
	}

	return min
}

func max(vs ...int) int {
	max := math.MinInt32
	for _, v := range vs {
		if max < v {
			max = v
		}
	}

	return max
}
