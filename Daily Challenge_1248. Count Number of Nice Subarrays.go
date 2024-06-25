package leetcodego

func atMostCount(nums []int, k int) int {
    //credit to lee215, TC O(n), SC O(1)
    res, left := 0, 0
    for right, num := range nums{
        k -= (num%2)
        for k < 0 && left < len(nums){
            k += nums[left]%2
            left++
        }
        res += right - left + 1
    }
    return res
}
func numberOfSubarrays(nums []int, k int) int {
    return atMostCount(nums, k) - atMostCount(nums, k - 1)
}