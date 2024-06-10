package leetcodego

func longestPalindrome2(s string) int {
    //TC O(n), SC O(1)
    length := 0
    hasOddCount := false
    runeCountMap := make(map[rune]int)
    for _, char := range s{
        runeCountMap[char]++
    }
    for _, count := range runeCountMap{
        if count%2 == 0{
            length += count
        }else{
            hasOddCount = true
            length += count - 1
        }
    }

    if hasOddCount{
        length += 1
    }

    return length
}