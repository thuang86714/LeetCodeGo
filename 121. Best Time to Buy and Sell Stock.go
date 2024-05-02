package leetcodego

func maxProfit(prices []int) int {
    curMin, ans:= 100000, 0
    for i := 0; i < len(prices); i++{
        if curMin > prices[i]{
            curMin = prices[i]
        }
        if prices[i] - curMin > ans{
            ans = prices[i] - curMin
        }
    }
    return ans
}