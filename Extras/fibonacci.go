package Extras

func fibonacci() []int {
	// Fibonacci series:
	// the sum of two elements defines the next
	var output []int

	a, b := 0, 1
	for a < 10 {
		output = append(output, a)
		a, b = b, a+b
	}

	return output
}
