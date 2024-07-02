package leetcodego
func minSubArrayLen2(target int, nums []int) int {
    //TC O(n), SC O(1)
    curSum, left, right, curMin:= 0, 0, 0, len(nums) + 1
    for right < len(nums){
        curSum += nums[right]
        for left <= right && curSum >= target{
            curMin = min(curMin, right - left + 1)
            curSum -= nums[left]
            left++
        }
        right++
    }
    if curMin == len(nums) + 1{
        return 0
    }
    return curMin
}