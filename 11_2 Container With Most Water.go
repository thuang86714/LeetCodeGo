package leetcodego

func maxArea3(height []int) int {
	left, right, curWater := 0, len(height)-1, 0
	for left < right {
		curHeight := min(height[left], height[right])
		curWidth := right - left
		curWater = max(curWater, curWidth*curHeight)
		for left <= right && height[left] <= curHeight {
			left++
		}
		for left <= right && height[right] <= curHeight {
			right--
		}
	}

	return curWater
}