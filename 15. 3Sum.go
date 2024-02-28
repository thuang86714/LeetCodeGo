func threeSum(nums []int) [][]int {
    //TC O(nlogn), SC O(1)
    resultSlice := [][]int{}
    sort.Ints(nums)
    for i := 0; i < len(nums) - 2; i++{
        if i > 0 && nums[i] == nums[i - 1]{
            continue
        }
        target, left, right := -nums[i], i + 1, len(nums) - 1
        for left < right{//binary search
            sum := nums[left] + nums[right]
            if sum == target{
                resultSlice = append(resultSlice, []int{nums[i], nums[left], nums[right]})
                left++
                right--
                for left < right && nums[left] == nums[left - 1]{
                    left++
                }
                for left < right && nums[right] == nums[right + 1]{
                    right--
                }
            }else if sum > target{
                right--
            }else{
                left++
            }
        }
    }
    return resultSlice
}