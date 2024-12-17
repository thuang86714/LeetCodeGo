package leetcodego

func countConsistentStrings(allowed string, words []string) int {
    charMap := make(map[byte]bool)

    for i := 0; i < len(allowed); i++{
        charMap[allowed[i]] = true
    }

    count := 0
    for _, word := range words{
        j := 0
        for ; j < len(word) ; j++{
            if _, exist := charMap[word[j]]; !exist{
                break
            }
        }
        if j == len(word){
            count++
        }
    }
    return count
}