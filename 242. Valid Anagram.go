func isAnagram(s string, t string) bool {
    //TC O(s.len + t.len), SC O(1)
    charMap := make(map[rune]int)
    for _,ch := range s{
        charMap[ch]++
    }
    for _,ch := range t{
        charMap[ch]--
    }
    for _, val := range charMap{
        if val != 0{
            return false
        }
    }
    return true
}