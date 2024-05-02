package leetcodego

import (
    "slices"
)

func rob(nums []int) int {
    if(len(nums) <= 2){
        return slices.Max(nums)
    }

    oneStep, twoStep, ans := max(nums[0], nums[1]), nums[0], 0
    for i:= 2; i < len(nums); i++{
        ans = max(twoStep + nums[i], oneStep)
        twoStep = oneStep
        oneStep = ans
    }
    return ans
}