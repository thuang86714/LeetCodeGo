package leetcodego

func atMostK(A []int, K int) int {
	i, res := 0, 0
	count := make(map[int]int)
	for j, val := range A {
		if _, ok := count[val]; !ok {
			count[val] = 0
		}
		count[val]++
		if count[val] == 1 {
			K--
		}
		for K < 0 {
			count[A[i]]--
			if count[A[i]] == 0 {
				K++
			}
			i++
		}
		res += j - i + 1
	}
	return res
}

func subarraysWithKDistinct(A []int, K int) int {
	return atMostK(A, K) - atMostK(A, K-1)
}