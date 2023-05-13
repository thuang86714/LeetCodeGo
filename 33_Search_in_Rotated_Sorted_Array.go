/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

 

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
*/
func search(nums []int, target int) int {
	//credit to neetcode
    //require solution in time complexity O(log n), possibly hinting using binary search
    l, r := 0, len(nums)-1
    for l <= r{
        mid := (l+r)/2
        if nums[mid]==target{
            return mid
        }
        //left sorted portion
        if nums[l] <= nums[mid]{
            //find out where is target
            if target < nums[l] || target > nums[mid]{
                l = mid + 1
            }else{
                r = mid - 1
            }
        }else{//right sorted portion
            if target < nums[mid] || nums[r] < target{
                r = mid -1
            }else{
                l = mid+1
            }
        }
        
    }
    return -1
}

func search(nums []int, target int) int {
    n := len(nums)
    //credit to casd82
    // Find the pivot.
    left, right := 0, n-1
    for left < right {
        mid := left+(right-left)/2
        if nums[mid] > nums[right] {
            left = mid+1
        } else {
            right = mid
        }
    }
    
    pivot := left
    
	// Regular binary search
    left, right = pivot, pivot-1+n
    for left <= right {
        mid := left+(right-left)/2
        midVal := nums[mid % n]
        
        if midVal > target {
            right = mid-1
        } else if midVal < target {
            left = mid+1
        } else {
            return mid % n
        }
    }
    
    return -1
}