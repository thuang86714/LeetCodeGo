package leetcodedynamicprogramming


func removeDuplicates(nums []int) int {
    //TC O(n), SC O(1)
    count := 0
    for i := 1; i < len(nums); i++{
        if nums[i] == nums[i - 1 - count]{
            count++
        }else{
            nums[i - count] = nums[i]
        }
    }
    return len(nums) - count
}