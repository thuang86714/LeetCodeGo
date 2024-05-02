package leetcodego

func findMaxK(nums []int) int {
    //TC O(n), SC O(n)
    negativeNumMap := make(map[int]bool)
    ans := -1
    for _, num := range nums{
        if num < 0{
            negativeNumMap[num] = true
        }
    }

    for _, num := range nums{
        if num < 0{
            continue
        }
        if _, found := negativeNumMap[-1*num]; found{
            ans = max(ans, num)
        }
    }
    return ans
}

/*
func findMaxK(nums []int) int {
    //TC O(nlogn), SC O(1), sorting + two ptr
    sort.Ints(nums)
    left, right := 0, len(nums) - 1
    
    for left < right{
        if nums[left] + nums[right] == 0{
            return nums[right]
        }else if nums[left] + nums[right] < 0{
            left++
        }else{
            right--
        }
    }
    return -1
}

*/