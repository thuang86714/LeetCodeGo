package leetcodego

//credit to linhduong
func diff(a, b byte) int {
    if a < b {
        return int(b-a)
    }
    return int(a - b)
}

func equalSubstring(s string, t string, maxCost int) int {
    left, curDiff, n, maxLen := 0, 0, len(s), 0
    for right := range s{
        curDiff += diff(s[right], t[right])
        for left < n && curDiff > maxCost{
            curDiff -= diff(s[left], t[left])
            left++
        }

        maxLen = max(maxLen, right - left + 1)
    }
    return maxLen
}