package leetcodego

func minSubArrayLen(target int, nums []int) int {
    curSum, ans, left := 0, len(nums) + 1, 0
    for right := 0; right < len(nums); right++{
        curSum += nums[right]
        for left <= right && curSum >= target{
            ans = min(ans, right - left + 1)
            curSum -= nums[left]
            left++
        }
    }
    if ans == len(nums) + 1{
        return 0
    }
    return ans
}