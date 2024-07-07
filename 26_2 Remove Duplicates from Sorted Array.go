package leetcodego

func removeDuplicates3(nums []int) int {
    count := 0
    for idx := 1; idx < len(nums); idx++{
        if nums[idx] == nums[idx - 1]{
            count++
        }else{
            nums[idx - count] = nums[idx]
        }
    }
    return len(nums) - count
}