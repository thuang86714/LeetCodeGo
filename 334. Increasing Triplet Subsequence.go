package leetcodego

import "math"
func increasingTriplet(nums []int) bool {
    if len(nums) < 3{
        return false
    }
    nums1, nums2 := math.MaxInt32, math.MaxInt32
    for _, n := range nums{
        if n <= nums1{
            nums1 = n
        }else if n <= nums2{
            nums2 = n
        }else{
            return true
        }
    }
    return false
}