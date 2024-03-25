package leetcodego

func removeStars(s string) string {
    //TC O(n), SC O(n)
    runeSlice := []rune{}
    for _, r := range s{
        if r == '*'{
            runeSlice = runeSlice[:len(runeSlice) - 1]
        }else{
            runeSlice = append(runeSlice, r)
        }
    }

    ans := string(runeSlice)
    return ans;
}