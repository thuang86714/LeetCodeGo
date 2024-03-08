func maxFrequencyElements(nums []int) int {
    //TC O(n), SC O(1)
    countSlice := make([]int, 101)
    maxCount, res := 0, 0
    for _, num := range nums{
        countSlice[num]++
    }
    for _, count := range countSlice{
        maxCount = max(maxCount, count)
    }
    for _, count := range countSlice{
        if count == maxCount{
            res += maxCount
        }
    }
    return res
}