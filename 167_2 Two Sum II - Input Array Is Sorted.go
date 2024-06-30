package leetcodego

func twoSum3(numbers []int, target int) []int {
	//two pointer, TC O(n), SC O(1)
	left, right := 0, len(numbers)-1
	for left < right {
		curSum := numbers[left] + numbers[right]
		if curSum == target {
			return []int{left + 1, right + 1}
		} else if curSum > target {
			right--
		} else if curSum < target {
			left++
		}
	}
	return []int{-1, -1}
}