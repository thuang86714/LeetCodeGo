/*
Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

If target is not found in the array, return [-1, -1].

You must write an algorithm with O(log n) runtime complexity.

 

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
*/
package leetcodego

func searchRange(nums []int, target int) []int {
	//credit to needcode and johnson_xie(global_cache)
    l, r := 0, len(nums)-1
    result := []int{-1, -1}
    for l <=r {
        mid := (l+r)/2
        if target > nums[mid]{
            l = mid+1
        }else{
            if target == nums[mid]{
                result[0]= mid
                r = mid-1
            }else{
                r = mid -1
            }
        }
    }
    l, r = 0, len(nums)-1
    for l <= r{
        mid := (l+r)/2
        if target < nums[mid]{
            r = mid -1
        }else{
            if target == nums[mid]{
                result[1] = mid
                l = mid +1
            }else{
                l = mid +1
            }
        }
    }
    
    return result
}