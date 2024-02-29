func maxArea(height []int) int {
    //TC O(n), SC O(1)
    water, left, right := 0, 0, len(height) - 1
    for left <= right{
        w, h := right - left, min(height[left], height[right])
        water = max(water, w*h)
        for left <= right && height[left] <= h{
            left++
        }
        for left <= right && height[right] <= h{
            right--
        }
    }
    return water
}