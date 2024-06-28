package leetcodego
func findMaxAverage(nums []int, k int) float64 {
    sum := 0
    for i:= 0; i < k; i++{
        sum += nums[i]
    }
    temp := sum
    for i:=0; i < len(nums) - k; i++{
        temp -= nums[i]
        temp += nums[i + k]
        sum = max(sum, temp)
    }
    return float64(sum)/float64(k)
}