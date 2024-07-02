package leetcodego

func lengthOfLongestSubstring5(s string) int {
	//TC O(n), SC O(n)
	byteCountMap, left, ans := make(map[byte]int), 0, 0

	for right := 0; right < len(s); right++ {
		byteCountMap[s[right]]++

		for left < right && byteCountMap[s[right]] > 1 {
			byteCountMap[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}