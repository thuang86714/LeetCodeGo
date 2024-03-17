func productExceptSelf(nums []int) []int {
    //TC O(n), SC O(1)
    res := make([]int, len(nums))
    prefix, postfix := 1, 1
    for idx := 0; idx < len(nums); idx++{
        res[idx] = prefix
        prefix *= nums[idx]
    }
    for idx := len(nums) - 1; idx >= 0; idx--{
        res[idx] *= postfix
        postfix *= nums[idx]
    }
    return res
}