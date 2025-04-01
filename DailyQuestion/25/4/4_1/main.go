package __1

func mostPoints(questions [][]int) (ret int64) {
	n := len(questions)
	dp := make([]int64, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			dp[i] = max(dp[i], dp[i-1])
		}
		q := questions[i]
		nextIdx := i + q[1] + 1
		if nextIdx < n {
			dp[nextIdx] = max(dp[i]+int64(q[0]), dp[nextIdx])
		} else {
			ret = max(dp[i]+int64(q[0]), ret)
		}
	}
	return ret
}
