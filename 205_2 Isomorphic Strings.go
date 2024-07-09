package leetcodego

func isIsomorphic2(s string, t string) bool {
    //TC O(n), SC O(n)
    sMap, tMap := make(map[byte]byte), make(map[byte]byte)

    for i := 0; i < len(s); i++{
        _, sOk := sMap[s[i]]
        _, tOk := tMap[t[i]]
        if !sOk && !tOk{
            sMap[s[i]] = t[i]
            tMap[t[i]] = s[i]
        }else{
            if sMap[s[i]] == t[i] && tMap[t[i]] == s[i]{
                continue
            }
            return false
        }
    }
    return true
}