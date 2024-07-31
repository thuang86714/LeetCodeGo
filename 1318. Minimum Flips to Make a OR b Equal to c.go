package leetcodego
import "math/bits"
func minFlips(a int, b int, c int) int {
    c ^= a | b // Perform the operation c = c ^ (a | b)
    return bits.OnesCount(uint(c)) + bits.OnesCount(uint(a&b&c))

}