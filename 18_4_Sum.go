/*
Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:

0 <= a, b, c, d < n
a, b, c, and d are distinct.
nums[a] + nums[b] + nums[c] + nums[d] == target
You may return the answer in any order.

 

Example 1:

Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]*/
func fourSum(nums []int, target int) [][]int {
    result := [][]int{}
    sort.Ints(nums)
    if len(nums) < 4{
        return result
    }
    for i := 0; i < len(nums); i++{
        var target_of_3 int
        target_of_3 = target - nums[i]
        for j := i +1; j < len(nums); j++{
            var target_of_2 int
            target_of_2 = target_of_3 - nums[j]
            front := j+1
            back := len(nums)-1
            for front < back{
                two_sum := nums[front]+ nums[back]
                if two_sum < target_of_2{
                    front++
                }else if two_sum > target_of_2 {
                    back--
                }else {
                    temp := []int{nums[i], nums[j], nums[front], nums[back]}
                    result = append(result, temp)
                    for front < back && nums[front]==temp[2]{front++}
                    for front < back&& nums[back]==temp[3]{back--}
                }
                
            }
            for j +1 <len(nums) && nums[j+1]==nums[j]{j++}
        }
        for i +1 < len(nums)&& nums[i+1]==nums[i]{i++}
    }

    return result
}