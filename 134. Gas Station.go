package leetcodego

func canCompleteCircuit(gas []int, cost []int) int {
    //TC O(n), SC O(1)
    gasSum, costSum, diff, index := 0, 0, 0, 0
    for i := 0; i < len(gas); i++{
        gasSum += gas[i]
        costSum += cost[i]
        diff += gas[i] - cost[i]
        if diff < 0{
            index = i + 1
            diff = 0
        }
    }
    if gasSum < costSum{
        return -1
    }
    return index
}