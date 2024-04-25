package leetcodego

func longestIdealString(s string, k int) int {
	//TC O(n*k), SC O(1)
    charMemoSlice := [26]int{}
    ans := 0

    for _, char := range s {
        maxLen := 0
        // Directly work with indices derived from character ranges
        start := int(max(0, int(char-'a') - k))
        end := int(min(25, int(char-'a') + k))

        for j := start; j <= end; j++ {
            maxLen = max(maxLen, charMemoSlice[j])
        }

        charMemoSlice[char-'a'] = maxLen + 1
        ans = max(ans, charMemoSlice[char-'a'])
    }

    return ans
}