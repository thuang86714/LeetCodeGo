package leetcodego

func searchInsert(nums []int, target int) int {
	//binary search, TC O(logn)-->since it's a sorted slice, SC O(1)
    left, right := 0, len(nums) - 1
    for left <= right{
        mid := right + (left - right)/2
        if nums[mid] == target{
            return mid
        }else if nums[mid] < target{
            left = mid + 1
        }else{
            right = mid - 1
        }
    }
    return left
}