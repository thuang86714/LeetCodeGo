package leetcodego

func trap(height []int) int {
    //credit to hiepit and jacklee20499, TC O(n), SC O(1)
    water, left, right := 0, 0, len(height) - 1
    maxLeft, maxRight := height[left], height[right]
    for left < right{
        if maxLeft < maxRight{
            left++
            maxLeft = max(maxLeft, height[left])
            water += maxLeft - height[left]
        }else{
            right--
            maxRight = max(maxRight, height[right])
            water += maxRight - height[right]
        }
    }
    return water
}