package leetcodego

func maxProfit2(prices []int) int {
    curMin, curProfit := 100001, 0
    for _, price := range prices{
        curMin = min(curMin, price)
        curProfit = max(curProfit, price - curMin)
    }
    return curProfit
}