package leetcodego

func minKBitFlips(nums []int, k int) int {
    //credit to votrubac, solution2, TC O(n), SC O(n)
    flipTrackSlice := []int{}
    res := 0
    for idx, num := range nums{
        var curFlipTimes int
        //The size of the queue will indicate number of flips
        if len(flipTrackSlice)%2 == 1{
            curFlipTimes = 0
        }else{
            curFlipTimes = 1
        }
        if num != curFlipTimes{
            res++
            flipTrackSlice = append(flipTrackSlice, idx + k - 1)
        }
        //meaning the sliding window is moving out of this range, the old bit would not be fliped anymore
        if len(flipTrackSlice) != 0 && flipTrackSlice[0] <= idx{
            flipTrackSlice = flipTrackSlice[1:]
        }
    }
    if len(flipTrackSlice) == 0{
        return res
    }
    //meaning there's still 0 in the array
    return -1
}

/*
func minKBitFlips(nums []int, k int) int {
    //credit to votrubac, TC O(n), SC O(1)
    //we can track the total number of flips, and use the source array to mark flips with negative values.
    res, flipRemainedCount := 0, 0
    for idx, num := range nums{
        if num == flipRemainedCount%2{
            if idx > len(nums) - k{
                return - 1
            }
            res++
            flipRemainedCount++
            nums[idx] -= 2
        }
        if idx >= k - 1 && nums[idx - k + 1] < 0{
            flipRemainedCount--
            nums[idx - k + 1] += 2
        }
    }
    return res
}
*/