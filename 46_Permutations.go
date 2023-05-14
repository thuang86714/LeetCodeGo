/*
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.

 

Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]
Example 3:

Input: nums = [1]
Output: [[1]] 
*/

func permute(nums []int) [][]int {
    var result [][]int
    backtrack(&result, nums, 0)
    return result
}

func backtrack(res *[][]int, nums []int, idx int){
    if idx >= len(nums){
        cpy := make([]int, len(nums))
        copy(cpy, nums)
        *res = append(*res, cpy)
        return
    }

    for i:=idx; i < len(nums); i++{
        nums[i], nums[idx] = nums[idx], nums[i]
        backtrack(res, nums, idx+1)
        nums[i], nums[idx] = nums[idx], nums[i]
    }
}
