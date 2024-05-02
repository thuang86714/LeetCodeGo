/*
Given a string s, return the longest 
palindromic
 
substring
 in s.

 

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"*/

package leetcodego
func longestPalindrome(s string) string {
    //TC = O(N^2), SC = O(1)
    length := len(s)
    if length == 0{
        return ""
    }
    var left, right, prev_left, prev_right int
    for right < length {
        // gobble up dup chars
        for right + 1 < length && s[left] == s[right+1]{
            right++
        }
        // find size of this palindrome
        for left - 1 >= 0 && right +1 < length && s[left -1] == s[right +1]{
            left--
            right++
        }
        if right -left > prev_right -prev_left{
            prev_right, prev_left = right, left
        }
        // reset to next mid point
		left = (left + right)/2 + 1
		right = left
    }
    return s[prev_left:prev_right+1]
}