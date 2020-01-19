package TimeComplexity

import "sort"

func PermMissingElemSolution(A []int) int {
	sort.Ints(A)
	var ce int

	for _, e := range A {
		if ce+1 != e && ce != e {
			return ce + 1
		}
		ce = e
	}

	return ce + 1
}
