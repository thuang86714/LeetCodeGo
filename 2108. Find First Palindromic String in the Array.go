package leetcodego
func firstPalindrome(words []string) string {
    for _, word:= range words{
        i, j := 0, len(word) - 1
        for word[i] == word[j]{
            if i >= j{
                return word
            }
            i++
            j--
        }
    }
    return ""
}