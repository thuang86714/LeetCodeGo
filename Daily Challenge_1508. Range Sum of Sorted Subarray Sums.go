package leetcodego

import (
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	//TC O(mlogm), where m is N^2, SC O(N^2)
	numSum := []int{}

	for idx := range nums {
		cur := 0
		for r := idx; r < len(nums); r++ {
			cur += nums[r]
			numSum = append(numSum, cur)
		}
	}

	sort.Ints(numSum)
	ans := 0
	for idx := left - 1; idx <= right-1; idx++ {
		ans = (ans + numSum[idx]) % 1000000007
	}
	return ans
}
