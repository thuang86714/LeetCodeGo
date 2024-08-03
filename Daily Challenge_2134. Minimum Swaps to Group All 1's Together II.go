package leetcodego

func minSwaps(nums []int) int {
	//Sliding Window, TC O(n), SC O(1), credit to votrubac
	ones, n := 0, len(nums)
	res, j, curOne := n, 0, 0
	for _, num := range nums {
		ones += num
	}

	for i := 0; i < n; i++ {
		for j-i < ones {
			curOne += nums[(j % n)]
			j++
		}
		res = min(res, ones-curOne)
		curOne -= nums[i]
	}
	return res
}