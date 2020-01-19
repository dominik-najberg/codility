package Iterations

import (
	"strconv"
)

func BinaryGapSolution(N int) int {
	var (
		longestGap = 0
		zeroCount  = 0

		foundFirstOne = false
		//foundClosingOne = false
	)

	bin := strconv.FormatInt(int64(N), 2)
	for _, b := range bin {
		bs := string(b)

		if foundFirstOne == false && bs == "1" {
			foundFirstOne = true
			continue
		}

		if foundFirstOne == false && bs == "0" {
			continue
		}

		if bs == "0" {
			zeroCount++
		}

		if bs == "1" {
			if longestGap < zeroCount {
				longestGap = zeroCount
			}

			zeroCount = 0
		}
	}

	return longestGap
}
