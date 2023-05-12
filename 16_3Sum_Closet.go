/*
Given an integer array nums of length n and an integer target, find three integers in nums such that the sum is closest to target.

Return the sum of the three integers.

You may assume that each input would have exactly one solution.

 

Example 1:

Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).*/
func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
    n := len(nums)
    sum := nums[0]+nums[1]+nums[2]
    for i := 0; i < n-1; i++{
        j := i +1
        k := n -1
        for ;j < k;{
            temp := nums[i] + nums[j] + nums[k]
            if abs(temp - target) < abs(sum - target){
                sum = temp
            }
            if temp < target{
                j++
            }else if temp > target{
                k--
            }else{
                return target
            }
        }
    }
    return sum
}

func abs(x int) int{
    if x < 0{
        return -x
    }
    return x
} 