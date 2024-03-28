package leetcodego

func firstMissingPositive(nums []int) int {
    //credit to makuiyu and jianchao-li, V0idk. TC O(n), SC O(1)
    i, n := 0, len(nums)

    for i < n{
        if nums[i] > 0 && nums[i] <= n && nums[nums[i] - 1] != nums[i]{
            nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
        }else{
            i++
        }
    }

    for i := 0; i < n; i++{
        if nums[i] != i + 1{
            return i + 1
        }
    }
    return n + 1
}