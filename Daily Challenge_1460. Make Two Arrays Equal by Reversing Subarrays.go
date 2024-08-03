package leetcodego

import (
	"sort"
)

func canBeEqual(target []int, arr []int) bool {
	//Unlimited swap-->sorting, do sorting then deep compare to find out if each element is same
	//TC O(nlogn), SC O(1)
	//Use hashmap to count-->TC O(n), SC O(n)
	sort.Ints(target)
	sort.Ints(arr)
	for idx := range target {
		if arr[idx] != target[idx] {
			return false
		}
	}
	return true
}
