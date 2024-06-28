package leetcodego
import(
    "slices"
)
func minCostClimbingStairs(cost []int) int {
    if(len(cost) < 3){
        return slices.Min(cost)
    }
    cur, oneStep, twoStep := 0, cost[1], cost[0]
    for i := 2; i < len(cost); i++{
        cur = min(oneStep, twoStep) + cost[i]
        twoStep = oneStep
        oneStep = cur
    }
    return min(oneStep, twoStep) 
}