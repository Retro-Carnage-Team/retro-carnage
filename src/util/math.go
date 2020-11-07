package util

type MathUtil struct{}

func (mu MathUtil) Max(a float32, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func (mu MathUtil) Min(a float32, b float32) float32 {
	if a < b {
		return a
	}
	return b
}
