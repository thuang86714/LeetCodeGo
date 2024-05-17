package leetcodego
import(
    "math"
)
// func sortedSquares(nums []int) []int {
//     //TC O(nlogn), SC O(1)
//     ans := []int{}
//     for _, num := range nums{
//         ans = append(ans, num*num)
//     }
//     sort.Ints(ans)
//     return ans
// }
func sortedSquares(nums []int) []int {
    //can we do TC O(n)-->yes, credit to andnik and yuluairoy
    idx, left, right := len(nums) - 1, 0, len(nums) -1
    ans := make([]int, right + 1)
    for ;idx >= 0; idx--{
        if(math.Abs(float64(nums[left])) > math.Abs(float64(nums[right]))){
            ans[idx] = nums[left] * nums[left]
            left++
        }else{
            ans[idx] = nums[right] * nums[right]
            right--
        }
    }
    return ans
}