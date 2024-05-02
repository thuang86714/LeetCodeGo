package leetcodego

func rotate(nums []int, k int)  {
    //TC O(n), SC O(1)
    k %= len(nums)
    //reverse the whole slice first
    for left, right :=0, len(nums) - 1; left < right; left, right = left+1, right-1{
        nums[left], nums[right] = nums[right], nums[left]
    }
    //reverse the first k elements
    for left, right :=0, k - 1; left < right; left, right = left+1, right-1{
        nums[left], nums[right] = nums[right], nums[left]
    }
    //reverse the rest of elements
    for left, right :=k, len(nums) - 1; left < right; left, right = left+1, right-1{
        nums[left], nums[right] = nums[right], nums[left]
    }
}
/*
func rotate(nums []int, k int) {
    // Get the length of the array
    n := len(nums)
    
    // If k is larger than the length of the array,
    // take its modulo with the length of the array
    k = k % n
    
    // Reverse the entire array
    reverse(nums)
    
    // Reverse the first k elements
    reverse(nums[:k])
    
    // Reverse the remaining elements after k
    reverse(nums[k:])
}

func reverse(nums []int) {
    for left, right := 0, len(nums)-1; left < right; left, right = left+1, right-1 {
        nums[left], nums[right] = nums[right], nums[left]
    }
}
*/