package leetcodego

func firstUniqChar(s string) int {
    //credit to maxyzli, TC O(n), SC O(1)
    //<byte, int>
    charCounter := map[byte]int{}
    for i := 0; i < len(s); i++ {
        charCounter[s[i]]++
    }
    for i := 0; i < len(s); i++{
        if charCounter[s[i]] == 1{
            return i
        }
    }
    return -1
}