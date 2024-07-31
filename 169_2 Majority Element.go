package leetcodego

func majorityElement2(nums []int) int {
	//counting sort, TC O(n), SC O(n)
	ans, curCount := 0, 0
	numCountMap := make(map[int]int)
	for _, num := range nums {
		numCountMap[num]++
	}

	for _, num := range nums {
		if curCount < numCountMap[num] {
			ans = num
			curCount = numCountMap[num]
		}
	}
	return ans
}