package leetcodego

func lengthOfLongestSubstring3(s string) int {
	//TC O(s.length), SC O(s.length)
    maxLen, left, right := 0, 0, 0
    charCountMap := make(map[byte]int)
    for ;right < len(s); right++{
        charCountMap[s[right]]++
        
        for charCountMap[s[right]] > 1{
            charCountMap[s[left]]--
            left++
        }

        maxLen = max(maxLen, right - left + 1)
    }
    return maxLen
}