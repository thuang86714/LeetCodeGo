/*
Given an integer array nums of unique elements, return all possible 
subsets
 (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

 

Example 1:

Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
Example 2:

Input: nums = [0]
Output: [[],[0]]
*/
func subsets(nums []int) [][]int {
    result := [][]int{}
    helper(&result, nums, []int{}, 0)
    return result
}

func helper(result *[][]int, nums []int, path []int, start int) {
    // this is a new subset, all the time (even empty array at beginning)
    // so we should copy path, and add to result
    pathCopy := make([]int, len(path))
    copy(pathCopy, path)
    *result = append(*result, pathCopy)
    
    // go from start to end (because we want no repeaters)
    for i := start; i < len(nums); i++ {
        // take this number, recurse that path, and then pop cause done searching
        path = append(path, nums[i])
        helper(result, nums, path, i+1)
        path = path[:len(path)-1]
    }
}