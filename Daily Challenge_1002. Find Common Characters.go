package leetcodego

func commonChars(words []string) []string {
    //TC O(n*m) where m is average len of all words, SC O(1)
    charCountSlice := make([]int, 26)
    ans := []string{}
    for _, char := range words[0]{
        charCountSlice[char - 'a']++
    }

    for _, word := range words{
        tempCharCountSlice := make([]int, 26)
        for _, char := range word{
            tempCharCountSlice[char - 'a']++
        }
        for i := 0; i < 26; i++{
            charCountSlice[i] = min(charCountSlice[i], tempCharCountSlice[i])
        }
    }

    for charCode, count := range charCountSlice{
        for i:= 0; i < count; i++{
            ans = append(ans, string(charCode + 'a'))
        }
    }
    return ans
}