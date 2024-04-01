package leetcodego

func countSubarrays(nums []int, k int) int64 {
    curMax, maxCount := 0, 0
    var left int64
    var subarrayCount int64
    for _, num := range nums{
        curMax = max(curMax, num)
    }
    for right := 0; right < len(nums); right++{
        if nums[right] == curMax{
            maxCount++
        }

        for maxCount >= k{
            if nums[left] == curMax{
                maxCount--
            }
            left++
        }
        subarrayCount += left
    }
    return subarrayCount
}