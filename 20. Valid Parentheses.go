func isValid(s string) bool {
    //credit to vsir, TC O(n), SC O(n)
    if len(s) == 0 || len(s)%2 == 1{
        return false
    }
    pairs := map[rune]rune{
        '(' : ')',
        '{' : '}',
        '[' : ']',
    }
    stackSlice := []rune{}

    for _,val := range s{
        if _, ok := pairs[val]; ok{
            stackSlice = append(stackSlice, val)
        }else if len(stackSlice) == 0 || pairs[stackSlice[len(stackSlice) - 1]] != val{
            return false
        }else{
            stackSlice = stackSlice[:len(stackSlice) - 1]
        }
    }
    return len(stackSlice) == 0
}