package geometry

func max(a float32, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

func min(a float32, b float32) float32 {
	if a < b {
		return a
	}
	return b
}
