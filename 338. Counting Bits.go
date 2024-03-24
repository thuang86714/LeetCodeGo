package leetcodego

func countBits(n int) []int {
    //credit to fengcc, TC O(n), SC O(1)
    ans := make([]int, n + 1)
    for i := 1; i <= n; i++{
        ans[i] = ans[i & (i - 1)] + 1
    }
    return ans
}