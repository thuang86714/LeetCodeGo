package leetcodego
func removeDuplicates(nums []int) int {
    //TC O(n), SC O(1)
    count := 0;
    for i:= 2; i < len(nums); i++{
        if nums[i] == nums[i - 2- count]{
            count++
        }else{
            nums[i - count] = nums[i]
        }
    }
    return len(nums) - count
}