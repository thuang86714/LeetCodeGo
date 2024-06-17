package leetcodego
import(
	"math"
)
func judgeSquareSum(c int) bool {
	//TC O(c^(1/2)), SC O(1)
    left, right := 0, int(math.Sqrt(float64(c)))
    for left <= right{
        curSum := left*left + right*right
        if curSum == c{
            return true
        }else if curSum > c{
            right--
        }else{
            left++
        }
    }
    return false
}