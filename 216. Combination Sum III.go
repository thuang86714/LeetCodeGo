package leetcodego

func combinationSum3(k int, n int) [][]int {
    //TC O(9^k), SC O(k)
    result := [][]int{}
    var findCombo func(k, n, curNum, curSum int, curCombo []int)
    findCombo = func(k, n, curNum, curSum int, curCombo []int){
        if len(curCombo) == k{
            if curSum == n{
                comboCopy := make([]int, len(curCombo))
                copy(comboCopy, curCombo)
                result = append(result, comboCopy)
                return
            }
        }

        for i := curNum; i <= 9; i++{
            findCombo(k, n, i + 1, curSum + i, append(curCombo, i))
        }
    }
    cur := []int{}
    findCombo(k, n, 1, 0, cur)
    return result
}