package leetcodego

func threeConsecutiveOdds(arr []int) bool {
	curLen := 0
	for idx := 0; idx < len(arr); idx++ {
		if arr[idx]%2 == 0 {
			curLen = 0
			continue
		}
		curLen++
		if curLen == 3 {
			return true
		}
	}
	return false
}