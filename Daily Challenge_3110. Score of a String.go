package leetcodego
import (
	"math"
)
func scoreOfString(s string) int {
	//TC O(n), SC O(1)
    score := 0
    for i := 1; i < len(s); i++{
        score += int(math.Abs(float64(s[i]) - float64(s[i - 1]))) 
    }
    return score
}