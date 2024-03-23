func search(nums []int, target int) int {
    left, right := 0, len(nums) - 1
    for left <= right{
        mid := left + (right - left)/2
        if nums[mid] == target{
            return mid
        }
        //find out how the slice is rotated
        if nums[mid] >= nums[left]{
            if nums[mid] < target || nums[left] > target{
                left = mid + 1
            }else{
                right = mid - 1
            }
        }else{
            if nums[mid] > target || target > nums[right]{
                right = mid - 1
            }else{
                left = mid + 1
            }
        }
    }
    return -1
}