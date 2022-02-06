package algorithm

// MaximumProfit
func MaximumProfit(n int, R []int) int {
	maxv := -2000000000
	minv := R[0]

	for i := 1; i < n; i++ {
		if maxv < R[i] - minv {
			maxv = R[i] - minv
		}
		if minv > R[i] {
			minv = R[i]
		}
	}

	return maxv
}