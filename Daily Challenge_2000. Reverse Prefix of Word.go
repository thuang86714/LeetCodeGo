package leetcodego

func reversePrefix(word string, ch byte) string {
    //TC O(n), SC O(n)
    idx := 0
    for ; idx < len(word); idx++{
        if word[idx] == ch{
            break
        }
    }
    if idx == len(word){
        return word
    }
    runes := []rune(word)
    for i, j := 0, idx; i <= j; i, j = i + 1, j - 1{
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}