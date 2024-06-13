package leetcodego

var ans [][]int
func findCombo2(candidates, curCombo []int, target, curSum, idx int){
    //both TC SC areO(2^n*k), where k is avg len of all combo, 
    if curSum > target {
        return
    }
    if curSum == target {
        temp := make([]int, len(curCombo))
        copy(temp, curCombo)
        ans = append(ans, temp)
        return
    }

    for i := idx; i < len(candidates); i++{
        curSum += candidates[i]
        curCombo = append(curCombo, candidates[i])
        findCombo2(candidates, curCombo, target, curSum, i)
        curSum -= candidates[i]
        curCombo = curCombo[:len(curCombo) - 1]
    }
}
func combinationSum2(candidates []int, target int) [][]int {
    ans = [][]int{}
    curSum, curCombo := 0, []int{}
    findCombo2(candidates, curCombo, target, curSum, 0)
    return ans
}