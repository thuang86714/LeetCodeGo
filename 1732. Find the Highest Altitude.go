package leetcodego

func largestAltitude(gain []int) int {
    ans, cur := 0, 0
    for _, g := range(gain){
        cur += g
        ans = max(ans, cur)
    }
    return ans
}