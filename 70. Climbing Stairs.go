package leetcodego
func climbStairs(n int) int {
    if n <= 2{
        return n
    }

    oneStep, twoStep, ans := 2, 1, 0
    for i := 3; i <= n; i++{
        ans = oneStep + twoStep
        twoStep = oneStep
        oneStep = ans
    }
    return ans
}