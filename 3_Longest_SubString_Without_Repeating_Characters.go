/*
Given a string s, find the length of the longest 
substring
 without repeating characters.

 

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.*/
package leetcodego
func lengthOfLongestSubstring(s string) int {
    dict := [128]bool{}//s consists of English letters, digits, symbols and spaces.
    length, maxstr := 0, 0
    for i, j :=0, 0; i < len(s); i++{
        index := s[i]
        if dict[index]{
            for ;dict[index] ;j++{
                length--
                dict[s[j]] = false
            } 
        }

        dict[index] = true
        length++
        if length > maxstr{
            maxstr = length
        }
    }
    return maxstr
}