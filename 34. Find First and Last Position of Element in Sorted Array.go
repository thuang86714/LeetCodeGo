package leetcodego

func searchRange(nums []int, target int) []int {
    //binary search to find the left and right bound. TC O(logn), SC O(1)
    left, right := 0, len(nums) - 1
    ans := []int{-1, -1}
    for left <= right{//find the left bound
        mid := left + (right - left)/2
        if nums[mid] >= target{
            if nums[mid] == target{
                ans[0] = mid
            }
            right = mid - 1
        }else{
            left = mid + 1
        }
    }

    left, right = 0, len(nums) - 1
    for left <= right{//find the right bound
        mid := left + (right - left)/2
        if nums[mid] <= target{
            if nums[mid] == target{
                ans[1] = mid
            }
            left = mid + 1
        }else{
            right = mid - 1
        }
    }
    return ans
}