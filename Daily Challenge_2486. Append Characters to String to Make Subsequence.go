package leetcodego

func appendCharacters(s string, t string) int {
    //TC O(n), SC O(1)
    sIdx, tIdx := 0, 0
    for sIdx < len(s) && tIdx < len(t){
        if s[sIdx] == t[tIdx]{
            tIdx++
        }
        sIdx++
    }
    return len(t) - tIdx
}