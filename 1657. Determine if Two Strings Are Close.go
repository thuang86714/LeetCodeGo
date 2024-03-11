func closeStrings(word1 string, word2 string) bool {
    charCount1, charCount2, charCount3, charCount4 := make([]int, 26),  make([]int, 26),  make([]int, 26),  make([]int, 26)
    for _, ch := range word1{
        charCount1[ch - 'a']++
        charCount2[ch - 'a'] = 1
    }

    for _, ch := range word2{
        charCount3[ch - 'a']++
        charCount4[ch - 'a'] = 1
    }

    sort.Ints(charCount1)
    sort.Ints(charCount3)

    for i := 0; i < 26; i++{
        if charCount1[i] != charCount3[i]{
            return false
        }
        if charCount2[i] != charCount4[i]{
            return false
        }
    }

    return true
}