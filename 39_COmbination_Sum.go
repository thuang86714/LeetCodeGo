/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the 
frequency
 of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

 

Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
*/
func combinationSum(candidates []int, target int) [][]int {
	//credit to ganesh
    result := [][]int{}
    if len(candidates) == 0{
        return result
    }
    sort.Ints(candidates)
    backtrack(&result, make([]int, 0), candidates, target, 0)
    return result
}

func backtrack(res *[][]int, temp, nums []int, target int, i int) {
    if target == 0 {
        cpyTemp := make([]int, len(temp))
        copy(cpyTemp, temp)
        *res = append(*res, cpyTemp)
        return
    }

    for i < len(nums) && target - nums[i]>=0{
        temp = append(temp, nums[i])
        backtrack(res, temp, nums, target - nums[i], i)
        i++
        temp = temp[:len(temp)-1]
    }
}