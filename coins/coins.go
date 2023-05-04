package coins

func Piles(n int) int {
	dp := make([][]int, n+1)
	dp[0] = make([]int, n+1)
	for i := range dp[0] {
		dp[0][i] = 1
	}
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			if i >= j {
				dp[i][j] = dp[i-j][j] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[n][n]
}
