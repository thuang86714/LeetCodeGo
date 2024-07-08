package leetcodego

func removeElement2(nums []int, val int) int {
	count := 0
	for idx := 0; idx < len(nums); idx++ {
		if nums[idx] == val {
			count++
		} else {
			nums[idx-count] = nums[idx]
		}
	}
	return len(nums) - count
}