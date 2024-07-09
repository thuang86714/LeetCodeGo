package leetcodego

func isAnagram(s string, t string) bool {
    runeCountMap := make(map[rune]int)
    for _, b := range s{
        runeCountMap[b]++
    }
    for _, b := range t{
        runeCountMap[b]--
    }

    for _, val := range runeCountMap{
        if val != 0{
            return false
        }
    }
    return true
}