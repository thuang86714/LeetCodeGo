package leetcodego

// credit to archit91, Dynamic Programming - Tabulation, TC O(n), SC O(n)
func numTilings(n int) int {
    MOD := 1000000007
    memoSlice := make([][]int, 1001)
    for idx := range memoSlice{
        memoSlice[idx] = make([]int, 2)
    }
    memoSlice[1] = []int{1, 1}
    memoSlice[2] = []int{2, 2}
    for i := 3; i <= n; i++{
        memoSlice[i][0] = (memoSlice[i - 1][0] + memoSlice[i - 2][0] + 2*memoSlice[i - 2][1])%MOD
        memoSlice[i][1] = (memoSlice[i - 1][0] + memoSlice[i - 1][1])%MOD
    }
    return memoSlice[n][0]
}