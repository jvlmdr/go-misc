package split

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func ceilDiv(a, b int) int {
	if b <= 0 {
		panic("non-positive denominator")
	}
	if a < 0 {
		panic("negative numerator")
	}
	return (a + b - 1) / b
}
