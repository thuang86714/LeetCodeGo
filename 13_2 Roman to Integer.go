package leetcodego

var charToIntMap = map[byte]int{
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}
func romanToInt2(s string) int {
    ans := 0
    for idx := 0; idx < len(s); idx++{
        if(idx < len(s) - 1 && charToIntMap[s[idx]] < charToIntMap[s[idx + 1]]){
            ans -= charToIntMap[s[idx]]
        }else{
            ans += charToIntMap[s[idx]]
        }
    }
    return ans
}