package leetcodego

import(
	"sort"
)

func majorityElement(nums []int) int {
    //TC O(nlogn), SC O(1)
    sort.Ints(nums)
    return nums[len(nums)/2]
}
/*
func majorityElement(nums []int) int {
    //TC O(n), SC O(1)
    count, majority := 0, 0
    for _,n := range nums{
        if count == 0{
            majority = n
        }
        if n == majority{
            count++
        }else{
            count--
        }
    }
    return majority
}
*/