package leetcodego

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums2) > len(nums1){
		return intersect(nums2, nums1)
	}
	// TC O(n + m), SC O(min(m, n))
	numCountMap := make(map[int]int)

	for _, num := range nums1 {
		numCountMap[num]++
	}
	ansSlice := []int{}
	for _, num := range nums2 {
		if numCountMap[num] > 0 {
			ansSlice = append(ansSlice, num)
			numCountMap[num]--
		}
	}
	return ansSlice
}

/*
func intersect(nums1 []int, nums2 []int) []int {
	//TC O(nlogn + mlogm), SC O(1)
	sort.Ints(nums1)
	sort.Ints(nums2)

	idx1, idx2 := 0, 0
	ansSlice := []int{}
	for idx1 < len(nums1) && idx2 < len(nums2) {
		if nums1[idx1] < nums2[idx2] {
			idx1++
		} else if nums1[idx1] > nums2[idx2] {
			idx2++
		} else {
			ansSlice = append(ansSlice, nums2[idx2])
			idx1++
			idx2++
		}
	}
	return ansSlice
}
*/