package leetcodego

func strStr(haystack string, needle string) int {
    //credit to shikuan, TC O(n), SC O(1)
    if needle == ""{
        return 0
    }
    hLen := len(haystack)
    nLen := len(needle)
    for i:= 0; i < hLen - nLen + 1;i++{
        if haystack[i:i + nLen] == needle{
            return i
        }
    }
    return -1
}