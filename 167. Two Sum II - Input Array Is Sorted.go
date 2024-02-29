func twoSum(numbers []int, target int) []int {
    //binary search, TC O(logn), SC O(1)
    left, right := 0, len(numbers) -1
    for left < right{
        curSum := numbers[left] + numbers[right]
        if curSum == target{
            return []int{left + 1, right + 1}
        }else if curSum < target{
            left++
        }else{
            right--
        }
    }
    return []int{}
}