package CountingElements

func MaxCountersSolution(N int, A []int) []int {
	var (
		counters = make([]int, N)
		base     int
		max      int
	)

	for _, val := range A {
		if val > N {
			base = max
			continue
		}

		if counters[val-1] < base {
			counters[val-1] = base
		}

		counters[val-1]++

		if max < counters[val-1] {
			max = counters[val-1]
		}
	}

	for i, value := range counters {
		if value < base {
			counters[i] = base
		}
	}

	return counters
}
