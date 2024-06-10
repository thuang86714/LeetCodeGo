package leetcodego

func reverseString(s []byte)  {
    //TC O(n), SC O(1)
    left, right := 0, len(s) - 1
    for left < right{
        s[left], s[right] = s[right], s[left]
        left++
        right--
    }
}