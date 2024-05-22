package leetcodego

func isPalindrome4(s string, left, right int) bool{
    for left < right{
        if s[left] != s[right]{
            return false
        }
        left++
        right--
    }
    return true
}

var res [][]string

func findPartition(s string, idx int, temp []string){
    if idx == len(s){
        // Create a copy of temp and append it to the results
        res = append(res, append([]string(nil), temp...))
        return
    }

    for i := idx; i < len(s); i++{
        if isPalindrome4(s, idx, i){
            temp = append(temp, s[idx:i + 1])
            findPartition(s, i + 1, temp)
            temp = temp[:len(temp) - 1]
        }
    }
    return
}
func partition(s string) [][]string {
    //TC O(n*2^n), SC O(n*2^n)
    res = [][]string{}
    temp := []string{}
    findPartition(s, 0, temp)
    return res
}