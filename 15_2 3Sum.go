package leetcodego
import(
	"sort"
)
func threeSum3(nums []int) [][]int {
	//TC O(n^2), SC O(1)
	ansSlice := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			// Skip duplicate elements
			continue
		}
		target := nums[i] * -1
		left, right := i+1, len(nums)-1
		for left < right {
			curSum := nums[left] + nums[right]
			if curSum == target {
				ansSlice = append(ansSlice, []int{nums[i], nums[left], nums[right]})

				// Skip duplicates for left pointer
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// Skip duplicates for right pointer
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// Move pointers
				left++
				right--
			} else if curSum < target {
				left++
			} else if curSum > target {
				right--
			}
		}
	}
	return ansSlice
}