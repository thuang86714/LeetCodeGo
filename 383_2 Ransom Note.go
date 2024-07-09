package leetcodego

func canConstruct2(ransomNote string, magazine string) bool {
    runeCountMap := make(map[rune]int)
    for _, r := range magazine{
        runeCountMap[r]++
    }

    for _, r := range ransomNote{
        runeCountMap[r]--
        if runeCountMap[r] < 0{
            return false
        }
    }
    return true
}