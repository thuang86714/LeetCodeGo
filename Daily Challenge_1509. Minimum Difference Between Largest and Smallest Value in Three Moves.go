package leetcodego
import(
	"sort"
)
func minDifference(nums []int) int {
    n := len(nums)
    if n <= 4{
        return 0
    }
    ans := 10000000000
    sort.Ints(nums)
    for i := 4; i > 0; i--{
        ans = min(ans, nums[n - i] - nums[4 - i])
    }
    return ans
}