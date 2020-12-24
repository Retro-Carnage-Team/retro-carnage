package util

import (
	"math"
	"retro-carnage/logging"
)

func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(input []float64) float64 {
	if 0 == len(input) {
		logging.Error.Fatalf("Cannot get max of empty slice")
	}

	if 1 == len(input) {
		return input[0]
	}

	var result = input[0]
	for _, value := range input {
		result = math.Max(result, value)
	}
	return result
}
