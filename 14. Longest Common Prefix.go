package leetcodego

func longestCommonPrefix(strs []string) string {
    sort.Strings(strs)
    firstStr, lastStr := strs[0], strs[len(strs) - 1]
    for i:= range(firstStr){
        if firstStr[i] != lastStr[i]{
            return firstStr[:i]
        }
    }
    return firstStr
}