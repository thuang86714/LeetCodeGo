package leetcodego
import(
    "sort"
)
func bagOfTokensScore(tokens []int, power int) int {
    //TC O(nlogn), SC O(1), greedy
    ans, score, left, right := 0, 0, 0, len(tokens) - 1
    sort.Ints(tokens)
    for left <= right{
        if power >= tokens[left]{
            score++
            ans = max(ans, score)
            power -= tokens[left]
            left++
        }else if score > 0{
            score--
            power += tokens[right]
            right--
        }else{
            break
        }
    }
    return ans
}