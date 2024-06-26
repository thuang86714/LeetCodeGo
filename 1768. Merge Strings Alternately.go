package leetcodego

func mergeAlternately(word1 string, word2 string) string {
    i, j := 0, 0
    ans := ""
    for i < len(word1) || j < len(word2){
        if i < len(word1){
            ans += string(word1[i])
            i++
        }
        if j < len(word2){
            ans += string(word2[j])
            j++
        }
    }
    return ans
}