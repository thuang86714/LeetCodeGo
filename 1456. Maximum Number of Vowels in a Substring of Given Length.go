package leetcodego

func maxVowels(s string, k int) int {
    //TC O(n), SC O(1)
    charMap := map[byte]bool{
        'a': true,
        'e': true,
        'i':true,
        'o':true,
        'u':true,
    }
    left, right, maxCount, curCount := 0, 0, 0, 0
    for left < len(s) && right < len(s){//sliding window with fixed size
        for right < len(s) && right - left + 1 <= k{
            if _, ok := charMap[s[right]]; ok{
                curCount++
            }
            right++
        }
        maxCount = max(maxCount, curCount)
        
        if _, ok := charMap[s[left]]; ok{
            curCount--
        }
        left++
    }
    return maxCount
}