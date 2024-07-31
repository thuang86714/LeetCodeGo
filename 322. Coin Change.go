package leetcodego
func coinChange(coins []int, amount int) int {
    //TC O(n), SC O(n)
    memo := make([]int, amount + 1)
    for i := range memo{
        memo[i] = amount + 1
    }
    memo[0] = 0
    for i := 1; i <= amount; i++{
        for _, c := range coins{
            if i >= c{
                memo[i] = min(memo[i], memo[i - c] + 1)
            }
        }
    }
    if memo[amount] == amount + 1{
        return -1
    }else{
        return memo[amount]
    }
}