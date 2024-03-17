func numSubarraysWithSum(nums []int, goal int) int {
    //credit to lee215 and Ruolinbi, TC O(n), SC O(1)
    var atMostSubArray func([]int, int) int
    atMostSubArray = func(nums []int, goal int)int{//atMost(A, S) counts the number of the subarrays of A the sum of which is at most(less than or equal to) S.
        if goal < 0{
            return 0
        }
        left, right, ans := 0, 0, 0
        for ; right < len(nums); right++{
            goal -= nums[right]
            for goal < 0{
                goal += nums[left]//Through this process, we can count the number of the subarrays.
                left++
            }
            ans += right - left + 1
        }
        return ans
    }
    return atMostSubArray(nums, goal) - atMostSubArray(nums, goal - 1)
}