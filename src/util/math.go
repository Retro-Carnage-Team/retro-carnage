package util

type MathUtil struct{}

func (mu MathUtil) Max(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func (mu MathUtil) Min(a float64, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
