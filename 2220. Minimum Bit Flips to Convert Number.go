package leetcodego
func minBitFlips(start int, goal int) int {
    count := 0
    for start > 0 || goal > 0{
        if start%2 != goal%2{
            count++
        }
        start /= 2
        goal /= 2
    }
    return count
}