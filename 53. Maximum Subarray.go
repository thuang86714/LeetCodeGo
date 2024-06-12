package leetcodego

func maxSubArray(nums []int) int {
    //TC O(n), SC O(1)
    maxi, curMax := -10001, -10001
    for _, num := range nums{
        curMax = max(curMax + num, num)
        maxi = max(maxi, curMax)
    }
    return maxi
}

/*
func maxSubArray(nums []int) int {
    //TC O(n), SC O(n)
    memoSlice := make([]int, len(nums))
    maxi := -10001
    for idx := 0; idx < len(nums); idx++{
        if(idx == 0){
            memoSlice[idx] = nums[0]
        }else{
            memoSlice[idx] = max(memoSlice[idx - 1] + nums[idx], nums[idx])
        }
        maxi = max(maxi, memoSlice[idx])
    }
    return maxi
}
*/