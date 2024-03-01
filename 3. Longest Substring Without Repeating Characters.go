func lengthOfLongestSubstring(s string) int {
    //credit to andik, TC O(n^2), SC O(1)
    dict := [128]bool{}
    length, ans:= 0, 0
    for i, j := 0, 0; i < len(s); i++{
        index := s[i]
        if dict[index]{
            for ;dict[index]; j++{
                length--
                dict[s[j]] = false
            }
        }

        dict[index] = true
        length++
        ans = max(ans, length)
    }
    return ans
}