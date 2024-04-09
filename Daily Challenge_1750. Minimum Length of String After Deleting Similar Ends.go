package leetcodego

func minimumLength(s string) int {
    //two pointer, TC O(n), SC O(1)
    left, right := 0, len(s) - 1
    for left < right && s[left] == s[right]{
        cur := s[left]
        for left <= right && s[left] == cur{
            left++
        }
        for left <= right && s[right] == cur{
            right--
        }
    }
    return right - left + 1
}