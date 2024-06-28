package leetcodego

func maxProfit3(prices []int, fee int) int {
	n := len(prices)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, 2)
	}

	for i := range memo {
		if i == 0 {
			memo[i][0] = 0
			memo[i][1] = -prices[i] - fee
			continue
		}
		memo[i][0] = max(memo[i-1][0], memo[i-1][1]+prices[i])
		memo[i][1] = max(memo[i-1][1], memo[i-1][0]-prices[i]-fee)
	}
	return memo[n-1][0]
}