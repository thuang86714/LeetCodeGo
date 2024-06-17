package leetcodego
import(
	"sort"
)
func minIncrementForUnique(nums []int) int {
    //TC O(nlogn), SC O(1)
    sort.Ints(nums) 
    count := 0
    for i:= 1; i < len(nums); i++{
        diff := nums[i - 1] + 1 - nums[i]
        if diff > 0{
            count += diff
            nums[i] += diff
        }
    }
    return count
}