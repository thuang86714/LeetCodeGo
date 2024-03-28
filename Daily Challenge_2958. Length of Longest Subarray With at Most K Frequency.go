package leetcodego

func maxSubarrayLength(nums []int, k int) int {
    //TC O(n), SC O(n)
    numCountMap := map[int]int{}//<nums[i], count>
    maxLen, left := 0, 0
    for right := 0; right < len(nums); right++{
        numCountMap[nums[right]]++
        for numCountMap[nums[right]] > k{
            numCountMap[nums[left]]--
            left++
        }
        maxLen = max(maxLen, right - left + 1)
    }
    return maxLen
}